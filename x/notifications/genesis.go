package notifications

import (
	"github.com/MonikaCat/canine-chain/v2/x/notifications/keeper"
	"github.com/MonikaCat/canine-chain/v2/x/notifications/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the notifications
	for _, elem := range genState.NotificationsList {
		k.SetNotifications(ctx, elem, elem.Address)
	}
	// Set all the notiCounter
	for _, elem := range genState.NotiCounterList {
		k.SetNotiCounter(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// How to export every single custom prefixKey for each user?...

	genesis.NotificationsList = k.GetAllNotifications(ctx)
	genesis.NotiCounterList = k.GetAllNotiCounter(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
