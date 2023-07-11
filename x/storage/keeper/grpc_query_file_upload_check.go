package keeper

import (
	"context"

	"github.com/MonikaCat/canine-chain/v2/x/storage/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FileUploadCheck(c context.Context, req *types.QueryFileUploadCheckRequest) (*types.QueryFileUploadCheckResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	spi, found := k.GetStoragePaymentInfo(ctx, req.Address)

	if !found {
		return nil, status.Error(codes.NotFound, "address not found")
	}

	if req.Bytes < 0 {
		return nil, status.Error(codes.InvalidArgument, "bytes cannot be negative")
	}

	return &types.QueryFileUploadCheckResponse{Valid: (spi.SpaceAvailable - spi.SpaceUsed) < req.Bytes}, nil
}
