package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ReportsAll(c context.Context, req *types.QueryAllReportRequest) (*types.QueryAllReportResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var reports []types.ReportForm
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	reportStore := prefix.NewStore(store, types.KeyPrefix(types.ReportKeyPrefix))

	pageRes, err := query.Paginate(reportStore, req.Pagination, func(key []byte, value []byte) error {
		var providers types.ReportForm
		if err := k.cdc.Unmarshal(value, &providers); err != nil {
			return err
		}

		reports = append(reports, providers)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllReportResponse{Reports: reports, Pagination: pageRes}, nil
}

func (k Keeper) Reports(c context.Context, req *types.QueryReportRequest) (*types.QueryReportResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetReportForm(
		ctx,
		req.Cid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryReportResponse{Report: val}, nil
}
