package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bandprotocol/chain/v2/x/oracle/types"
)

// HasRequirementFile checks if the data source of this ID exists in the storage.
func (k Keeper) HasRequirementFile(ctx sdk.Context, id types.RequirementFileID) bool {
	return ctx.KVStore(k.storeKey).Has(types.RequirementFileStoreKey(id))
}

// GetRequirementFile returns the data source struct for the given ID or error if not exists.
func (k Keeper) GetRequirementFile(ctx sdk.Context, id types.RequirementFileID) (types.RequirementFile, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.RequirementFileStoreKey(id))
	if bz == nil {
		return types.RequirementFile{}, sdkerrors.Wrapf(types.ErrRequirementFileNotFound, "id: %d", id)
	}
	var requirementFile types.RequirementFile
	k.cdc.MustUnmarshal(bz, &requirementFile)
	return requirementFile, nil
}

// MustGetRequirementFile returns the data source struct for the given ID. Panic if not exists.
func (k Keeper) MustGetRequirementFile(ctx sdk.Context, id types.RequirementFileID) types.RequirementFile {
	requirementFile, err := k.GetRequirementFile(ctx, id)
	if err != nil {
		panic(err)
	}
	return requirementFile
}

// SetRequirementFile saves the given data source to the storage without performing validation.
func (k Keeper) SetRequirementFile(ctx sdk.Context, id types.RequirementFileID, requirementFile types.RequirementFile) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.RequirementFileStoreKey(id), k.cdc.MustMarshal(&requirementFile))
}

// AddRequirementFile adds the given data source to the storage.
func (k Keeper) AddRequirementFile(ctx sdk.Context, requirementFile types.RequirementFile) types.RequirementFileID {
	id := k.GetNextRequirementFileID(ctx)
	k.SetRequirementFile(ctx, id, requirementFile)
	return id
}

// MustEditRequirementFile edits the given data source by id and flushes it to the storage.
func (k Keeper) MustEditRequirementFile(ctx sdk.Context, id types.RequirementFileID, new types.RequirementFile) {
	requirementFile := k.MustGetRequirementFile(ctx, id)
	requirementFile.Owner = new.Owner
	requirementFile.Name = modify(requirementFile.Name, new.Name)
	requirementFile.Description = modify(requirementFile.Description, new.Description)
	requirementFile.Filename = modify(requirementFile.Filename, new.Filename)
	requirementFile.Treasury = new.Treasury
	requirementFile.Fee = new.Fee
	k.SetRequirementFile(ctx, id, requirementFile)
}

// GetAllRequirementFiles returns the list of all data sources in the store, or nil if there is none.
func (k Keeper) GetAllRequirementFiles(ctx sdk.Context) (requirementFiles []types.RequirementFile) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RequirementFileStoreKeyPrefix)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var requirementFile types.RequirementFile
		k.cdc.MustUnmarshal(iterator.Value(), &requirementFile)
		requirementFiles = append(requirementFiles, requirementFile)
	}
	return requirementFiles
}
