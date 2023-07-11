package keeper

import (
	"strings"

	"github.com/MonikaCat/canine-chain/v2/x/rns/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetNames set a specific names in the store from its index
func (k Keeper) SetNames(ctx sdk.Context, names types.Names) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NamesKeyPrefix))
	b := k.cdc.MustMarshal(&names)
	store.Set(types.NamesKey(
		names.Name,
		names.Tld,
	), b)
}

// GetNames returns a names from its index
func (k Keeper) GetNames(
	ctx sdk.Context,
	name string,
	tld string,
) (val types.Names, found bool) {
	name = strings.ToLower(name)
	tld = strings.ToLower(tld)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NamesKeyPrefix))

	b := store.Get(types.NamesKey(
		name,
		tld,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNames removes a names from the store
func (k Keeper) RemoveNames(
	ctx sdk.Context,
	name string,
	tld string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NamesKeyPrefix))
	store.Delete(types.NamesKey(
		name,
		tld,
	))
}

// GetAllNames returns all names
func (k Keeper) GetAllNames(ctx sdk.Context) (list []types.Names) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NamesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Names
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// quickly checks if there are any domains registered
func (k Keeper) CheckExistence(ctx sdk.Context) bool {
	// intializing the iterator
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NamesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	// looping to see if at least 1 element exists
	i := 0
	for ; iterator.Valid(); iterator.Next() {
		if i > 0 {
			break
		}
		i++
	}
	exist := false
	if i > 0 {
		exist = true
	}
	return exist
}
