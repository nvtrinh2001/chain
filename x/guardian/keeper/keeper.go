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

func (k Keeper) Lock(ctx sdk.Context, fromAddr string, toAddrs []string, amt sdk.Coins) error {
	payer, err := sdk.AccAddressFromBech32(fromAddr)
	if err != nil {
		return err
	}

	var payeeList []sdk.AccAddress
	for _, payee := range toAddrs {
		payeeAccount, err := sdk.AccAddressFromBech32(payee)
		if err != nil {
			return err
		}
		payeeList = append(payeeList, payeeAccount)
	}

	var paidFee sdk.Coins
	for _, coin := range amt {
		temp := sdk.Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount.Mul(sdk.NewInt(int64(len(payeeList)))),
		}
		paidFee = append(paidFee, temp)
	}

	// integrate bank module here
	err = k.bankKeeper.SendCoinsFromAccountToModule(
		ctx, payer, k.feeCollectorName, paidFee,
	)

	if err != nil {
		return err
	}

	id := k.AddGuardedFee(ctx, types.NewGuardedFee(
		payer, payeeList, amt,
	))

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeLock,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", id)),
	))

	return nil
}

func (k Keeper) Claim(ctx sdk.Context, toAddr string, id uint64, amt sdk.Coins) error {
	payee, err := sdk.AccAddressFromBech32(toAddr)
	if err != nil {
		return err
	}

	guardedFee, err := k.GetGuardedFee(ctx, types.GuardedFeeID(id))
	if err != nil {
		return err
	}

	if amt.IsAnyGT(guardedFee.Fee) {
		return types.ErrInvalidAmount
	}

	check := 0
	for _, feePayee := range guardedFee.Payees {
		if feePayee.Payee != toAddr {
			continue
		}

		if feePayee.Status != types.STATUS_CLAIMABLE {
			return types.ErrFeeHasBeenClaimed
		}

		check = 1
	}
	if check == 0 {
		return types.ErrUnAuthorizedAccount
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, k.feeCollectorName, payee, amt,
	)

	if err != nil {
		return err
	}

	remain, isNeg := guardedFee.Fee.SafeSub(amt)
	if isNeg {
		return types.ErrInvalidAmount
	}

	payer, err := sdk.AccAddressFromBech32(guardedFee.Payer)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, k.feeCollectorName, payer, remain,
	)
	if err != nil {
		return err
	}

	for i, payeeInFee := range guardedFee.Payees {
		if payeeInFee.Payee == toAddr {
			guardedFee.Payees[i].Status = types.STATUS_CLAIMED
			break
		}
	}

	_ = k.UpdateGuardedFee(ctx, types.GuardedFeeID(id), guardedFee)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeClaim,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", id)),
		sdk.NewAttribute(types.AttributeClaimedPayee, fmt.Sprintf("%v", toAddr)),
	))

	return nil
}
