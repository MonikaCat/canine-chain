package keeper

import (
	"context"

	"github.com/MonikaCat/canine-chain/v2/x/storage/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetProviderIP(goCtx context.Context, msg *types.MsgSetProviderIP) (*types.MsgSetProviderIPResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider, found := k.GetProviders(ctx, msg.Creator)

	if !found {
		return nil, types.ErrProviderNotFound
	}

	provider.Ip = msg.Ip

	k.SetProviders(ctx, provider)

	return &types.MsgSetProviderIPResponse{}, nil
}
