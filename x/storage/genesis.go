package storage

import (
	"github.com/MonikaCat/canine-chain/v2/x/storage/keeper"
	"github.com/MonikaCat/canine-chain/v2/x/storage/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the contracts
	for _, elem := range genState.ContractsList {
		k.SetContracts(ctx, elem)
	}
	// Set all the activeDeals
	for _, elem := range genState.ActiveDealsList {
		k.SetActiveDeals(ctx, elem)
	}
	// Set all the Providers
	for _, elem := range genState.ProvidersList {
		k.SetProviders(ctx, elem)
	}

	// Set all the strays
	for _, elem := range genState.StraysList {
		k.SetStrays(ctx, elem)
	}
	// Set all the fidCid
	for _, elem := range genState.FidCidList {
		k.SetFidCid(ctx, elem)
	}

	// Set all the paymentinfo
	for _, elem := range genState.PaymentInfoList {
		k.SetStoragePaymentInfo(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ContractsList = k.GetAllContracts(ctx)
	genesis.ActiveDealsList = k.GetAllActiveDeals(ctx)
	genesis.ProvidersList = k.GetAllProviders(ctx)
	genesis.StraysList = k.GetAllStrays(ctx)
	genesis.FidCidList = k.GetAllFidCid(ctx)
	genesis.PaymentInfoList = k.GetAllStoragePaymentInfo(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
