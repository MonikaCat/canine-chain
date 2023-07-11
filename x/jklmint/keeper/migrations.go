package keeper

// DONTCOVER

import (
	v210 "github.com/MonikaCat/canine-chain/v2/x/jklmint/legacy/v210"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	k Keeper
}

// NewMigrator returns a new Migrator
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{
		k: keeper,
	}
}

// Migrate1to2 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v210.MigrateStore(ctx, &m.k.paramSpace)
}
