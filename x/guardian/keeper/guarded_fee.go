package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bandprotocol/chain/v2/x/guardian/types"
)

// HasGuardedFee checks if the data source of this ID exists in the storage.
func (k Keeper) HasGuardedFee(ctx sdk.Context, id types.GuardedFeeID) bool {
	return ctx.KVStore(k.storeKey).Has(types.GuardedFeeStoreKey(id))
}

// GetGuardedFee returns the data source struct for the given ID or error if not exists.
func (k Keeper) GetGuardedFee(ctx sdk.Context, id types.GuardedFeeID) (types.GuardedFee, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.GuardedFeeStoreKey(id))
	if bz == nil {
		return types.GuardedFee{}, sdkerrors.Wrapf(types.ErrGuardedFeeNotFound, "id: %d", id)
	}
	var guardedFee types.GuardedFee
	k.cdc.MustUnmarshal(bz, &guardedFee)
	return guardedFee, nil
}

// GetGuardedFeeList returns the data source struct for the given ID or error if not exists.
func (k Keeper) GetGuardedFeeList(ctx sdk.Context, account string, status types.STATUS) ([]*types.GuardedFee, error) {
	allGuardedFees := k.GetAllGuardedFees(ctx)
	var guardedFees []*types.GuardedFee
	for _, fee := range allGuardedFees {
		if fee.Payer == account {
			guardedFees = append(guardedFees, &fee)
			continue
		}
		for _, payee := range fee.Payees {
			if payee.Payee == account {
				guardedFees = append(guardedFees, &fee)
				break
			}
		}
	}

	if status == -1 {
		return guardedFees, nil
	}

	var filteredFees []*types.GuardedFee
	for _, fee := range guardedFees {
		for _, payee := range fee.Payees {
			if payee.Status == status {
				filteredFees = append(filteredFees, fee)
				break
			}
		}
	}

	return filteredFees, nil
}

// MustGetGuardedFee returns the data source struct for the given ID. Panic if not exists.
func (k Keeper) MustGetGuardedFee(ctx sdk.Context, id types.GuardedFeeID) types.GuardedFee {
	guardedFee, err := k.GetGuardedFee(ctx, id)
	if err != nil {
		panic(err)
	}
	return guardedFee
}

// SetGuardedFee saves the given data source to the storage without performing validation.
func (k Keeper) SetGuardedFee(ctx sdk.Context, id types.GuardedFeeID, guardedFee types.GuardedFee) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GuardedFeeStoreKey(id), k.cdc.MustMarshal(&guardedFee))
}

// AddGuardedFee adds the given data source to the storage.
func (k Keeper) AddGuardedFee(ctx sdk.Context, guardedFee types.GuardedFee) types.GuardedFeeID {
	id := k.GetNextGuardedFeeID(ctx)
	k.SetGuardedFee(ctx, id, guardedFee)
	return id
}

func (k Keeper) UpdateGuardedFee(ctx sdk.Context, id types.GuardedFeeID, fee types.GuardedFee) types.GuardedFeeID {
	k.SetGuardedFee(ctx, id, fee)
	return id
}

// GetAllGuardedFees returns the list of all data sources in the store, or nil if there is none.
func (k Keeper) GetAllGuardedFees(ctx sdk.Context) (guardedFees []types.GuardedFee) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GuardedFeeStoreKeyPrefix)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var guardedFee types.GuardedFee
		k.cdc.MustUnmarshal(iterator.Value(), &guardedFee)
		guardedFees = append(guardedFees, guardedFee)
	}
	return guardedFees
}
