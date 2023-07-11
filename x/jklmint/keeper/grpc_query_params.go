package keeper

import (
	"context"

	"github.com/MonikaCat/canine-chain/v2/x/jklmint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}
