package beta6

import (
	"github.com/MonikaCat/canine-chain/v2/app/upgrades"
	storagemodulekeeper "github.com/MonikaCat/canine-chain/v2/x/storage/keeper"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

var _ upgrades.Upgrade = &Upgrade{}

// Upgrade represents the v4 upgrade
type Upgrade struct {
	mm             *module.Manager
	configurator   module.Configurator
	storeageKeeper storagemodulekeeper.Keeper
}

// NewUpgrade returns a new Upgrade instance
func NewUpgrade(mm *module.Manager, configurator module.Configurator, storeageKeeper storagemodulekeeper.Keeper) *Upgrade {
	return &Upgrade{
		mm:             mm,
		configurator:   configurator,
		storeageKeeper: storeageKeeper,
	}
}

// Name implements upgrades.Upgrade
func (u *Upgrade) Name() string {
	return "beta6"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		return fromVM, nil
	}
}

// StoreUpgrades implements upgrades.Upgrade
func (u *Upgrade) StoreUpgrades() *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{}
}
