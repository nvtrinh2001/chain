package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the module.
	ModuleName = "feelocker"

	// StoreKey to be used when creating the KVStore.
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the oracle module
	QuerierRoute = ModuleName

	// RouterKey is the msg router key for the oracle module
	RouterKey = ModuleName
)

var (
	// GlobalStoreKeyPrefix is the prefix for global primitive state variables.
	GlobalStoreKeyPrefix = []byte{0x00}

	LockedFeeCountStoreKey = append(GlobalStoreKeyPrefix, []byte("GuardedFeeCount")...)

	LockedFeeStoreKeyPrefix = []byte{0x07}
)

// LockedFeeStoreKey returns the key to retrieve a specific guarded fee from the store.
func LockedFeeStoreKey(id LockedFeeID) []byte {
	return append(LockedFeeStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(id))...)
}
