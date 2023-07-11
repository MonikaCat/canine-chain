package keeper

import (
	"context"

	"github.com/MonikaCat/canine-chain/v2/x/storage/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) InitProvider(goCtx context.Context, msg *types.MsgInitProvider) (*types.MsgInitProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider := types.Providers{
		Address:         msg.Creator,
		Ip:              msg.Ip,
		Totalspace:      msg.Totalspace,
		Creator:         msg.Creator,
		BurnedContracts: "0",
		KeybaseIdentity: msg.Keybase,
		AuthClaimers:    []string{},
	}

	k.SetProviders(ctx, provider)

	return &types.MsgInitProviderResponse{}, nil
}
