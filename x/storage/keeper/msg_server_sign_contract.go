package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

const SigningProviderLength = 3

func (k Keeper) PostNewContract(ctx sdk.Context, fid string, creator string, merkle string, duration int64, fileSize uint64) (*types.MsgSignContractResponse, error) {
	providers := k.GetActiveProviders(ctx)
	providerList := make([]string, 0)

	if len(providers) < 3 {
		provs := k.GetAllProviders(ctx)
		providers = make([]types.ActiveProviders, 0)
		for i := 0; i < len(provs); i++ {
			ap := types.ActiveProviders{
				Address: provs[i].Address,
			}
			providers = append(providers, ap)
			if i > 20 {
				break
			}
		}
	}

	w := sha256.New() // creating new cid
	w.Write([]byte(fid))
	w.Write([]byte(creator))
	for i := 0; i < len(providerList); i++ {
		providerList = append(providerList, providers[i].Address)
		_, err := w.Write([]byte(providerList[i]))
		if err != nil {
			return nil, err
		}
	}
	cid := w.Sum(nil)
	cidString, err := MakeCid(cid)
	if err != nil {
		return nil, err
	}

	contract := types.ContractV2{
		Cid:       cidString,
		Creator:   creator,
		Providers: providerList,
		Merkle:    merkle,
		Duration:  duration,
		Filesize:  fileSize,
		Fid:       fid,
		Age:       0,
	}

	_, foundD := k.GetActiveDeals(ctx, cidString)
	_, foundS := k.GetStrays(ctx, cidString)
	if foundD || foundS {
		return nil, types.ErrContractExists
	}

	k.SetContracts(ctx, contract)

	cids := []string{contract.Cid}

	cidArr, err := json.Marshal(cids)
	if err != nil {
		return nil, err
	}

	newFtoC := types.FidCid{
		Fid:  contract.Fid,
		Cids: string(cidArr),
	}

	k.SetFidCid(ctx, newFtoC)

	if duration == 0 {
		payInfo, found := k.GetStoragePaymentInfo(ctx, creator)
		if !found {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "payment info not found, please purchase storage space")
		}

		fSize := int64(fileSize)

		// check if user has any free space
		if (payInfo.SpaceUsed + (fSize * 3)) > payInfo.SpaceAvailable {
			return nil, fmt.Errorf("not enough storage space")
		}
		// check if storage subscription still active
		if payInfo.End.Before(ctx.BlockTime()) {
			return nil, fmt.Errorf("storage subscription has expired")
		}

		payInfo.SpaceUsed += fSize * 3

		k.SetStoragePaymentInfo(ctx, payInfo)
	} else {
		size := fileSize / 1024
		if size < 1 {
			size = 1
		}
		var hour int64 = 60 * 60 // 60 seconds * 60 minutes
		timing := (duration * 6) / hour
		if timing < 720 {
			return nil, types.ErrDurationTooShort
		}
		cost := k.GetStorageCostKbs(ctx, int64(size), timing)
		depositor, err := sdk.AccAddressFromBech32(creator)
		if err != nil {
			return nil, err
		}
		deposit, err := sdk.AccAddressFromBech32(k.GetParams(ctx).DepositAccount)
		if err != nil {
			return nil, err
		}
		coins := sdk.NewCoins(sdk.NewCoin("ujkl", cost))
		err = k.bankkeeper.SendCoinsFromAccountToModule(ctx, depositor, types.ModuleName, coins)
		if err != nil {
			return nil, err
		}
		err = k.bankkeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, deposit, coins)
		if err != nil {
			return nil, err
		}
	}

	return &types.MsgSignContractResponse{}, nil
}

func (k msgServer) SignContract(goCtx context.Context, msg *types.MsgSignContract) (*types.MsgSignContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	return k.Keeper.PostNewContract(ctx, msg.Fid, msg.Creator, msg.Merkle, msg.Duration, msg.FileSize)

	// contract, found := k.GetContracts(ctx, msg.Cid)
	// if !found {
	//	return nil, fmt.Errorf("contract not found")
	// }
	//
	// _, found = k.GetActiveDeals(ctx, msg.Cid)
	// if found {
	//	return nil, fmt.Errorf("contract already exists")
	// }
	//
	// _, found = k.GetStrays(ctx, msg.Cid)
	// if found {
	//	return nil, fmt.Errorf("contract already exists")
	// }
	//
	// if contract.Signee != msg.Creator {
	//	return nil, fmt.Errorf("you do not have permission to approve this contract")
	// }
	//
	// size, ok := sdk.NewIntFromString(contract.Filesize)
	// if !ok {
	//	return nil, fmt.Errorf("cannot parse size")
	// }
	//
	// pieces := size.Quo(sdk.NewInt(k.GetParams(ctx).ChunkSize))
	//
	// var pieceToStart int64
	//
	// if !pieces.IsZero() {
	//	pieceToStart = ctx.BlockHeight() % pieces.Int64()
	// }
	//
	// var end int64
	// if msg.PayOnce {
	//	s := size.Quo(sdk.NewInt(1_000_000_000)).Int64()
	//	if s <= 0 {
	//		s = 1
	//	}
	//	cost := k.GetStorageCost(ctx, s, 720*12*200) // pay for 200 years
	//	deposit, err := sdk.AccAddressFromBech32(k.GetParams(ctx).DepositAccount)
	//	if err != nil {
	//		return nil, err
	//	}
	//	err = k.bankkeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, deposit, sdk.NewCoins(sdk.NewCoin("ujkl", cost)))
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	end = (200*31_536_000)/6 + ctx.BlockHeight()
	// }
	//
	// deal := types.ActiveDeals{
	//	Cid:           contract.Cid,
	//	Signee:        contract.Signee,
	//	Provider:      contract.Creator,
	//	Startblock:    fmt.Sprintf("%d", ctx.BlockHeight()),
	//	Endblock:      fmt.Sprintf("%d", end),
	//	Filesize:      contract.Filesize,
	//	Proofverified: "false",
	//	Blocktoprove:  fmt.Sprintf("%d", pieceToStart),
	//	Creator:       msg.Creator,
	//	Proofsmissed:  "0",
	//	Merkle:        contract.Merkle,
	//	Fid:           contract.Fid,
	// }
	//
	// if end == 0 {
	//	fsize, ok := sdk.NewIntFromString(contract.Filesize)
	//	if !ok {
	//		return nil, fmt.Errorf("cannot parse file size")
	//	}
	//	payInfo, found := k.GetStoragePaymentInfo(ctx, msg.Creator)
	//	if !found {
	//		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "payment info not found, please purchase storage space")
	//	}
	//
	//	// check if user has any free space
	//	if (payInfo.SpaceUsed + (fsize.Int64() * 3)) > payInfo.SpaceAvailable {
	//		return nil, fmt.Errorf("not enough storage space")
	//	}
	//	// check if storage subscription still active
	//	if payInfo.End.Before(ctx.BlockTime()) {
	//		return nil, fmt.Errorf("storage subscription has expired")
	//	}
	//
	//	payInfo.SpaceUsed += fsize.Int64() * 3
	//
	//	k.SetStoragePaymentInfo(ctx, payInfo)
	// }
	//
	// k.SetActiveDeals(ctx, deal)
	// k.RemoveContracts(ctx, contract.Cid)
	//
	// ftc, found := k.GetFidCid(ctx, contract.Fid)
	//
	// cids := []string{contract.Cid}
	//
	// if found {
	//	var ncids []string
	//	err := json.Unmarshal([]byte(ftc.Cids), &ncids)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	cids = append(cids, ncids...)
	// }
	//
	// for i := 0; i < 2; i++ {
	//	h := sha256.New()
	//	_, err := io.WriteString(h, fmt.Sprintf("%s%d", contract.Cid, i))
	//	if err != nil {
	//		return nil, err
	//	}
	//	hashName := h.Sum(nil)
	//
	//	scid, err := MakeCid(hashName)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	newContract := types.Strays{
	//		Cid:      scid,
	//		Signee:   contract.Signee,
	//		Fid:      contract.Fid,
	//		Filesize: contract.Filesize,
	//		Merkle:   contract.Merkle,
	//		End:      end,
	//	}
	//
	//	cids = append(cids, scid)
	//
	//	k.SetStrays(ctx, newContract)
	//
	// }
	//
	// cidarr, err := json.Marshal(cids)
	// if err != nil {
	//	return nil, err
	// }
	//
	// nftc := types.FidCid{
	//	Fid:  contract.Fid,
	//	Cids: string(cidarr),
	// }
	//
	// k.SetFidCid(ctx, nftc)
	//
	// return &types.MsgSignContractResponse{}, nil
}
