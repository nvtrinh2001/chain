package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/bandprotocol/chain/v2/x/guardian/types"
)

type Keeper struct {
	storeKey         sdk.StoreKey
	cdc              codec.BinaryCodec
	feeCollectorName string
	authKeeper       types.AccountKeeper
	bankKeeper       types.BankKeeper
	stakingKeeper    types.StakingKeeper
	distrKeeper      types.DistrKeeper
	authzKeeper      types.AuthzKeeper
}

// NewKeeper creates a new oracle Keeper instance.
func NewKeeper(
	cdc codec.BinaryCodec,
	key sdk.StoreKey,
	feeCollectorName string,
	authKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,
	distrKeeper types.DistrKeeper,
	authzKeeper types.AuthzKeeper,
) Keeper {
	return Keeper{
		storeKey:         key,
		cdc:              cdc,
		authKeeper:       authKeeper,
		feeCollectorName: feeCollectorName,
		bankKeeper:       bankKeeper,
		stakingKeeper:    stakingKeeper,
		distrKeeper:      distrKeeper,
		authzKeeper:      authzKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetGuardedFeeCount sets the number of guarded fee count to the given value.
func (k Keeper) SetGuardedFeeCount(ctx sdk.Context, count uint64) {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	ctx.KVStore(k.storeKey).Set(types.GuardedFeeCountStoreKey, bz)
}

// GetGuardedFeeCount returns the current number of all data sources ever exist.
func (k Keeper) GetGuardedFeeCount(ctx sdk.Context) uint64 {
	bz := ctx.KVStore(k.storeKey).Get(types.GuardedFeeCountStoreKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// GetNextGuardedFeeID increments and returns the current number of data sources.
func (k Keeper) GetNextGuardedFeeID(ctx sdk.Context) types.GuardedFeeID {
	guardedFeeCount := k.GetGuardedFeeCount(ctx)
	k.SetGuardedFeeCount(ctx, guardedFeeCount+1)
	return types.GuardedFeeID(guardedFeeCount + 1)
}
