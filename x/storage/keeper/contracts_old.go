package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// SetContracts set a specific contracts in the store from its index
func (k Keeper) SetContractsOld(ctx sdk.Context, contracts types.Contracts) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractsKeyPrefix))
	b := k.cdc.MustMarshal(&contracts)
	store.Set(types.ContractsKey(
		contracts.Cid,
	), b)
}

// GetContracts returns a contracts from its index
func (k Keeper) GetContractsOld(
	ctx sdk.Context,
	cid string,
) (val types.Contracts, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractsKeyPrefix))

	b := store.Get(types.ContractsKey(
		cid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveContracts removes a contracts from the store
func (k Keeper) RemoveContractsOld(
	ctx sdk.Context,
	cid string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractsKeyPrefix))
	store.Delete(types.ContractsKey(
		cid,
	))
}

// GetAllContracts returns all contracts
func (k Keeper) GetAllContractsOld(ctx sdk.Context) (list []types.Contracts) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Contracts
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
