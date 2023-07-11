package filetree

import (
	"github.com/MonikaCat/canine-chain/v2/x/filetree/keeper"
	types "github.com/MonikaCat/canine-chain/v2/x/filetree/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the files
	for _, elem := range genState.FilesList {
		k.SetFiles(ctx, elem)
	}
	// Set all the pubkey
	for _, elem := range genState.PubkeyList {
		k.SetPubkey(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.FilesList = k.GetAllFiles(ctx)
	genesis.PubkeyList = k.GetAllPubkey(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
