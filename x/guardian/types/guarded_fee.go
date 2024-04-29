package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewGuardedFee(
	payer sdk.AccAddress, payees []sdk.AccAddress, fee sdk.Coins,
) GuardedFee {
	var payeeList []*Payee
	for _, payee := range payees {
		payeeObj := Payee{Status: STATUS_CLAIMABLE, Payee: payee.String()}
		payeeList = append(payeeList, &payeeObj)
	}
	return GuardedFee{payer.String(), payeeList, fee}
}
