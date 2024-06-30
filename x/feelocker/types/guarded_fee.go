package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewLockedFee(
	payer sdk.AccAddress, payees []sdk.AccAddress, fee sdk.Coins,
) LockedFee {
	var payeeList []*Payee
	for _, payee := range payees {
		payeeObj := Payee{Status: STATUS_CLAIMABLE, Payee: payee.String()}
		payeeList = append(payeeList, &payeeObj)
	}
	return LockedFee{payer.String(), payeeList, fee}
}
