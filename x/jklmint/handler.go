package jklmint

import (
	"fmt"

	"github.com/MonikaCat/canine-chain/v2/x/jklmint/keeper"
	"github.com/MonikaCat/canine-chain/v2/x/jklmint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(_ keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager()) //nolint:staticcheck

		errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
	}
}
