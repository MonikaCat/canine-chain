package v210

import (
	"github.com/MonikaCat/canine-chain/v2/x/jklmint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// MigrateStore performs in-place store migrations from v1 to v2
// The things done here are the following:
// 1. setting up the next reason id and report id keys for existing subspaces
// 2. setting up the module params
func MigrateStore(ctx sdk.Context, paramsSubspace *paramstypes.Subspace) error {
	ctx.Logger().Error("MIGRATING MINT STORE!")
	// Set the module params
	params := types.DefaultParams()

	params.ProviderRatio = 0

	paramsSubspace.SetParamSet(ctx, &params)

	return nil
}
