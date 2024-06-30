package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bandprotocol/chain/v2/x/feelocker/types"
)

// HasLockedFee checks if the data source of this ID exists in the storage.
func (k Keeper) HasLockedFee(ctx sdk.Context, id types.LockedFeeID) bool {
	return ctx.KVStore(k.storeKey).Has(types.LockedFeeStoreKey(id))
}

// GetLockedFee returns the data source struct for the given ID or error if not exists.
func (k Keeper) GetLockedFee(ctx sdk.Context, id types.LockedFeeID) (types.LockedFee, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.LockedFeeStoreKey(id))
	if bz == nil {
		return types.LockedFee{}, sdkerrors.Wrapf(types.ErrLockedFeeNotFound, "id: %d", id)
	}
	var lockedFee types.LockedFee
	k.cdc.MustUnmarshal(bz, &lockedFee)
	return lockedFee, nil
}

// GetLockedFeeList returns the data source struct for the given ID or error if not exists.
func (k Keeper) GetLockedFeeList(ctx sdk.Context, account string, status types.STATUS) ([]*types.LockedFee, error) {
	allLockedFees := k.GetAllLockedFees(ctx)
	var lockedFees []*types.LockedFee
	for _, fee := range allLockedFees {
		if fee.Payer == account {
			lockedFees = append(lockedFees, &fee)
			continue
		}
		for _, payee := range fee.Payees {
			if payee.Payee == account {
				lockedFees = append(lockedFees, &fee)
				break
			}
		}
	}

	if status == -1 {
		return lockedFees, nil
	}

	var filteredFees []*types.LockedFee
	for _, fee := range lockedFees {
		for _, payee := range fee.Payees {
			if payee.Status == status {
				filteredFees = append(filteredFees, fee)
				break
			}
		}
	}

	return filteredFees, nil
}

// MustGetLockedFee returns the data source struct for the given ID. Panic if not exists.
func (k Keeper) MustGetLockedFee(ctx sdk.Context, id types.LockedFeeID) types.LockedFee {
	lockedFee, err := k.GetLockedFee(ctx, id)
	if err != nil {
		panic(err)
	}
	return lockedFee
}

// SetLockedFee saves the given data source to the storage without performing validation.
func (k Keeper) SetLockedFee(ctx sdk.Context, id types.LockedFeeID, lockedFee types.LockedFee) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LockedFeeStoreKey(id), k.cdc.MustMarshal(&lockedFee))
}

// AddLockedFee adds the given data source to the storage.
func (k Keeper) AddLockedFee(ctx sdk.Context, lockedFee types.LockedFee) types.LockedFeeID {
	id := k.GetNextLockedFeeID(ctx)
	k.SetLockedFee(ctx, id, lockedFee)
	return id
}

func (k Keeper) UpdateLockedFee(ctx sdk.Context, id types.LockedFeeID, fee types.LockedFee) types.LockedFeeID {
	k.SetLockedFee(ctx, id, fee)
	return id
}

// GetAllLockedFees returns the list of all data sources in the store, or nil if there is none.
func (k Keeper) GetAllLockedFees(ctx sdk.Context) (lockedFees []types.LockedFee) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.LockedFeeStoreKeyPrefix)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var lockedFee types.LockedFee
		k.cdc.MustUnmarshal(iterator.Value(), &lockedFee)
		lockedFees = append(lockedFees, lockedFee)
	}
	return lockedFees
}
