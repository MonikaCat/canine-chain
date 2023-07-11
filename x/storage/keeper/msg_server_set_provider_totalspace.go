package keeper

import (
	"context"
	"strconv"

	"github.com/MonikaCat/canine-chain/v2/x/storage/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetProviderTotalspace(goCtx context.Context, msg *types.MsgSetProviderTotalspace) (*types.MsgSetProviderTotalspaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider, found := k.GetProviders(ctx, msg.Creator)

	if !found {
		return nil, types.ErrProviderNotFound
	}

	validTotalSpace := isValidTotalSpace(msg.Space)

	if !validTotalSpace {
		return nil, types.ErrNotValidTotalSpace
	}

	provider.Totalspace = msg.Space

	k.SetProviders(ctx, provider)

	return &types.MsgSetProviderTotalspaceResponse{}, nil
}

func isValidTotalSpace(totalSpace string) bool {
	var isNumber bool

	if _, err := strconv.Atoi(totalSpace); err == nil {
		isNumber = true
	}
	return isNumber
}
