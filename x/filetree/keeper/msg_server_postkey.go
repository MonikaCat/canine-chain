package keeper

import (
	"context"

	"github.com/MonikaCat/canine-chain/v2/x/filetree/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Postkey(goCtx context.Context, msg *types.MsgPostkey) (*types.MsgPostkeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pubKey := types.Pubkey{
		Address: msg.Creator,
		Key:     msg.Key,
	}
	k.SetPubkey(ctx, pubKey)

	return &types.MsgPostkeyResponse{}, nil
}
