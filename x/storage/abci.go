package storage

import (
	"time"

	"github.com/MonikaCat/canine-chain/v2/x/storage/keeper"
	"github.com/MonikaCat/canine-chain/v2/x/storage/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	err := k.HandleRewardBlock(ctx)
	if err != nil {
		ctx.Logger().Error(err.Error())
	}

	k.KillOldContracts(ctx)

	var week int64 = (7 * 24 * 60 * 60) / 6

	if ctx.BlockHeight()%week == 0 { // clear out files once a week
		k.ClearDeadFiles(ctx)
	}
}
