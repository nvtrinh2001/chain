package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrGuardedFeeNotFound  = sdkerrors.Register(ModuleName, 1, "guarded fee not found")
	ErrNotEnoughFee        = sdkerrors.Register(ModuleName, 2, "not enough fee")
	ErrFeeHasBeenClaimed   = sdkerrors.Register(ModuleName, 3, "fee has been claimed")
	ErrUnAuthorizedAccount = sdkerrors.Register(ModuleName, 4, "account is not authorized to claim")
	ErrInvalidAmount       = sdkerrors.Register(ModuleName, 5, "invalid amount to claim")
)
