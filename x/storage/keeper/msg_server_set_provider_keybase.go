package keeper

import (
	"context"

	"github.com/MonikaCat/canine-chain/v2/x/storage/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetProviderKeybase(goCtx context.Context, msg *types.MsgSetProviderKeybase) (*types.MsgSetProviderKeybaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider, found := k.GetProviders(ctx, msg.Creator)

	if !found {
		return nil, types.ErrProviderNotFound
	}

	provider.KeybaseIdentity = msg.Keybase

	k.SetProviders(ctx, provider)

	return &types.MsgSetProviderKeybaseResponse{}, nil
}
