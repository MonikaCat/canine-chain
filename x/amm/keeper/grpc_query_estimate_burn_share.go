package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/amm/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EstimateBurnShare(
	goCtx context.Context,
	req *types.QueryEstimateBurnShareRequest,
) (*types.QueryEstimateBurnShareResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	pool, found := k.GetPool(ctx, req.PoolName)

	if !found {
		return nil, types.ErrLiquidityPoolNotFound
	}

	poolCoins := sdk.NewCoins(pool.Coins...)

	burnAmount, err := sdk.ParseCoinNormalized(req.Amount)

	if err != nil || burnAmount.IsNegative() {
		return nil, status.Error(codes.InvalidArgument, "invalid burn amount")
	}

	AMM, _ := types.GetAMM(pool.AMM_Id)

	returnedCoins, err := AMM.EstimateBurnShare(poolCoins, burnAmount)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	return &types.QueryEstimateBurnShareResponse{Coins: returnedCoins}, nil
}
