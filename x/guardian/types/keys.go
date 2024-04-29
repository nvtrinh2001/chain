package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the module.
	ModuleName = "guardian"

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

	GuardedFeeCountStoreKey = append(GlobalStoreKeyPrefix, []byte("GuardedFeeCount")...)

	GuardedFeeStoreKeyPrefix = []byte{0x01}
)

// GuardedFeeStoreKey returns the key to retrieve a specific guarded fee from the store.
func GuardedFeeStoreKey(id GuardedFeeID) []byte {
	return append(GuardedFeeStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(id))...)
}
