package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

// SetReportForm sets a specific report in the store from its index
func (k Keeper) SetReportForm(ctx sdk.Context, report types.ReportForm) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReportKeyPrefix))
	b := k.cdc.MustMarshal(&report)
	store.Set(types.ReportKey(
		report.Cid,
	), b)
}

// GetReportForm returns a report from its index
func (k Keeper) GetReportForm(
	ctx sdk.Context,
	cid string,
) (val types.ReportForm, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReportKeyPrefix))

	b := store.Get(types.ReportKey(
		cid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveReport removes an attestation from the store
func (k Keeper) RemoveReport(
	ctx sdk.Context,
	cid string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReportKeyPrefix))
	store.Delete(types.ReportKey(
		cid,
	))
}

// GetAllReport returns all reports
func (k Keeper) GetAllReport(ctx sdk.Context) (list []types.ReportForm) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReportKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ReportForm
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
