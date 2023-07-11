package keeper

import (
	"context"
	"fmt"
	"strings"

	"github.com/MonikaCat/canine-chain/v2/x/rns/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) CancelOneBid(ctx sdk.Context, sender string, name string) error {
	name = strings.ToLower(name)

	bidder, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return err
	}

	bid, bidFound := k.GetBids(ctx, fmt.Sprintf("%s%s", sender, name))

	if !bidFound {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "Bid does not exist or has expired.")
	}

	price, err := sdk.ParseCoinsNormalized(bid.Price)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidder, price)
	if err != nil {
		return err
	}

	k.RemoveBids(ctx, fmt.Sprintf("%s%s", sender, name))

	return nil
}

func (k msgServer) CancelBid(goCtx context.Context, msg *types.MsgCancelBid) (*types.MsgCancelBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.CancelOneBid(ctx, msg.Creator, msg.Name)

	return &types.MsgCancelBidResponse{}, err
}
