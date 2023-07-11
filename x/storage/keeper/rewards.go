package keeper

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerror "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/MonikaCat/canine-chain/v2/x/storage/types"
)

func getTotalSize(allDeals []types.ActiveDeals) sdk.Dec {
	networkSize := sdk.NewDecFromInt(sdk.NewInt(0))
	for i := 0; i < len(allDeals); i++ {
		deal := allDeals[i]
		ss, err := sdk.NewDecFromStr(deal.Filesize)
		if err != nil {
			continue
		}
		networkSize = networkSize.Add(ss)
	}
	return networkSize
}

func (k Keeper) manageDealReward(ctx sdk.Context, deal types.ActiveDeals, networkSize sdk.Dec, balance sdk.Coin) error {
	toprove, ok := sdk.NewIntFromString(deal.Blocktoprove)
	if !ok {
		return sdkerror.Wrapf(sdkerror.ErrInvalidType, "int parse failed")
	}

	iprove := toprove.Int64()

	totalSize, err := sdk.NewDecFromStr(deal.Filesize)
	if err != nil {
		return err
	}

	var byteHash byte
	if len(ctx.HeaderHash().Bytes()) > 2 {
		byteHash = ctx.HeaderHash().Bytes()[0] + ctx.HeaderHash().Bytes()[1] + ctx.HeaderHash().Bytes()[2]
	} else {
		byteHash = byte(ctx.BlockHeight()) // support for running simulations
	}

	d := totalSize.TruncateInt().Int64() / k.GetParams(ctx).ChunkSize

	if d > 0 {
		iprove = (int64(byteHash) + int64(ctx.BlockGasMeter().GasConsumed())) % d
	}

	deal.Blocktoprove = fmt.Sprintf("%d", iprove)

	verified, errb := strconv.ParseBool(deal.Proofverified)

	if errb != nil {
		return errb
	}

	if !verified {
		ctx.Logger().Debug("%s\n", "Not verified!")
		intt, ok := sdk.NewIntFromString(deal.Proofsmissed)
		if !ok {
			return sdkerror.Wrapf(sdkerror.ErrInvalidType, "int parse failed")
		}

		sb, ok := sdk.NewIntFromString(deal.Startblock)
		if !ok {
			return sdkerror.Wrapf(sdkerror.ErrInvalidType, "int parse failed")
		}

		DayBlocks := k.GetParams(ctx).ProofWindow

		if sb.Int64() >= ctx.BlockHeight()-DayBlocks {
			return sdkerror.Wrapf(sdkerror.ErrUnauthorized, "ignore young deals")
		}

		misses := intt.Int64() + 1
		missesToBurn := k.GetParams(ctx).MissesToBurn

		if misses > missesToBurn {
			provider, ok := k.GetProviders(ctx, deal.Provider)
			if !ok {
				return sdkerror.Wrapf(sdkerror.ErrKeyNotFound, "provider not found")
			}

			curburn, ok := sdk.NewIntFromString(provider.BurnedContracts)
			if !ok {
				return sdkerror.Wrapf(sdkerror.ErrInvalidType, "int parse failed")
			}
			provider.BurnedContracts = fmt.Sprintf("%d", curburn.Int64()+1)
			k.SetProviders(ctx, provider)

			intBlock, ok := sdk.NewIntFromString(deal.Endblock)
			if !ok {
				return sdkerror.Wrapf(sdkerror.ErrInvalidType, "int parse failed for endblock")
			}
			// Creating new stray file from the burned active deal
			strayDeal := types.Strays{
				Cid:      deal.Cid,
				Fid:      deal.Fid,
				Signee:   deal.Signee,
				Filesize: deal.Filesize,
				Merkle:   deal.Merkle,
				End:      intBlock.Int64(),
			}
			k.SetStrays(ctx, strayDeal)
			k.RemoveActiveDeals(ctx, deal.Cid)
			return nil
		}

		deal.Proofsmissed = fmt.Sprintf("%d", misses)
		k.SetActiveDeals(ctx, deal)
		return nil
	}

	ctx.Logger().Debug(fmt.Sprintf("File size: %s\n", deal.Filesize))
	ctx.Logger().Debug(fmt.Sprintf("Total size: %d\n", networkSize))

	nom := totalSize

	den := networkSize

	res := nom.Quo(den)

	ctx.Logger().Debug("Percentage of network space * 1000: %f\n", res)

	coinfloat := res.Mul(balance.Amount.ToDec())

	ctx.Logger().Debug("%f\n", coinfloat)
	coin := sdk.NewCoin("ujkl", coinfloat.TruncateInt())
	coins := sdk.NewCoins(coin)

	provider, err := sdk.AccAddressFromBech32(deal.Provider)
	if err != nil {
		return err
	}
	ctx.Logger().Debug("Sending coins to %s\n", provider.String())
	errorr := k.bankkeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, provider, coins)
	if errorr != nil {
		ctx.Logger().Debug("ERR: %v\n", errorr)
		ctx.Logger().Error(errorr.Error())
		return errorr
	}

	ctx.Logger().Debug("%s\n", deal.Cid)

	misses, ok := sdk.NewIntFromString(deal.Proofsmissed)
	if !ok {
		e := errors.New("cannot parse string")
		ctx.Logger().Error(e.Error())
		return e
	}
	updatedMisses := misses.SubRaw(1)

	if updatedMisses.LT(sdk.NewInt(0)) {
		updatedMisses = sdk.NewInt(0)
	}

	deal.Proofsmissed = updatedMisses.String()
	deal.Proofverified = "false"
	k.SetActiveDeals(ctx, deal)

	return nil
}

func (k Keeper) loopDeals(ctx sdk.Context, allDeals []types.ActiveDeals, networkSize sdk.Dec, balance sdk.Coin) {
	for _, deal := range allDeals {
		info, found := k.GetStoragePaymentInfo(ctx, deal.Signee)
		if !found {
			ctx.Logger().Debug(fmt.Sprintf("Removing %s due to no payment info", deal.Cid))
			cerr := CanContract(ctx, deal.Cid, deal.Signee, k)
			if cerr != nil {
				ctx.Logger().Error(cerr.Error())
			}
			continue
		}
		grace := info.End.Add(time.Hour * 24 * 30)
		if grace.Before(ctx.BlockTime()) {
			ctx.Logger().Debug(fmt.Sprintf("Removing %s after grace period", deal.Cid))
			cerr := CanContract(ctx, deal.Cid, deal.Signee, k)
			if cerr != nil {
				ctx.Logger().Error(cerr.Error())
			}
			continue
		}

		if info.SpaceUsed > info.SpaceAvailable { // remove file if the user doesn't have enough space
			ctx.Logger().Debug(fmt.Sprintf("Removing %s for space used", deal.Cid))
			err := CanContract(ctx, deal.Cid, deal.Signee, k)
			if err != nil {
				ctx.Logger().Error(err.Error())
			}
			continue
		}

		err := k.manageDealReward(ctx, deal, networkSize, balance)
		if err != nil {
			ctx.Logger().Error(err.Error())
			continue
		}

	}
}

func (k Keeper) InternalRewards(ctx sdk.Context, allDeals []types.ActiveDeals, address sdk.AccAddress) error {
	ctx.Logger().Debug("%s\n", "checking blocks")

	networkSize := getTotalSize(allDeals)

	balance := k.bankkeeper.GetBalance(ctx, address, "ujkl")

	k.loopDeals(ctx, allDeals, networkSize, balance)

	balance = k.bankkeeper.GetBalance(ctx, address, "ujkl")

	err := k.bankkeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(balance))
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) HandleRewardBlock(ctx sdk.Context) error {
	allDeals := k.GetAllActiveDeals(ctx)

	DayBlocks := k.GetParams(ctx).ProofWindow

	ctx.Logger().Debug("blockdiff : %d\n", ctx.BlockHeight()%DayBlocks)

	if ctx.BlockHeight()%DayBlocks > 0 {
		return sdkerror.Wrapf(sdkerror.ErrUnauthorized, "cannot check rewards before timer has been met")
	}

	address := k.accountkeeper.GetModuleAddress(types.ModuleName)

	err := k.InternalRewards(ctx, allDeals, address)
	if err != nil {
		return err
	}

	return nil
}
