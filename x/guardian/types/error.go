package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrGuardedFeeNotFound = sdkerrors.Register(ModuleName, 1, "guarded fee not found")
	ErrNotEnoughFee       = sdkerrors.Register(ModuleName, 2, "not enough fee")
)
