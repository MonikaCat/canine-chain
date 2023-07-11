package jklmint

import (
	"time"

	"github.com/MonikaCat/canine-chain/v2/x/jklmint/keeper"
	"github.com/MonikaCat/canine-chain/v2/x/jklmint/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	k.BlockMint(ctx)
}
