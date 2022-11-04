package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jackalLabs/canine-chain/x/rns/types"
)

func (k msgServer) List(goCtx context.Context, msg *types.MsgList) (*types.MsgListResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetForsale(ctx, msg.Name)

	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name already listed.")
	}

	n, tld, err := GetNameAndTLD(msg.Name)
	if err != nil {
		return nil, err
	}

	name, nfound := k.GetNames(ctx, n, tld)

	if !nfound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name does not exist or has expired.")
	}

	if name.Value != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "You do not own this name.")
	}

	blockHeight := ctx.BlockTime().Unix()

	if name.Locked > blockHeight {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "cannot transfer free name")
	}

	if blockHeight > name.Expires {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Name does not exist or has expired.")
	}

	newsale := types.Forsale{
		Name:  msg.Name,
		Price: msg.Price,
		Owner: msg.Creator,
	}

	k.SetForsale(ctx, newsale)

	return &types.MsgListResponse{}, nil
}
