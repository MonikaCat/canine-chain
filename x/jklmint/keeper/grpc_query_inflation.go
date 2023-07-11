package keeper

import (
	"context"

	"github.com/MonikaCat/canine-chain/v2/x/jklmint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Inflation(c context.Context, _ *types.QueryInflationRequest) (*types.QueryInflationResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	inflation, err := k.GetInflation(ctx)

	return &types.QueryInflationResponse{Inflation: inflation}, err
}
