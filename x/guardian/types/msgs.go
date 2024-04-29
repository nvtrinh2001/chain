package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// oracle message types
const (
	TypeMsgLockRequest  = "lock"
	TypeMsgClaimRequest = "claim"
)

var (
	_ sdk.Msg = &MsgLockRequest{}
	_ sdk.Msg = &MsgClaimRequest{}
)

func NewMsgLockRequest(
	payer sdk.AccAddress,
	fee sdk.Coins,
	payees []sdk.AccAddress,
) *MsgLockRequest {
	var payeeList []string
	for _, payee := range payees {
		payeeList = append(payeeList, payee.String())
	}
	return &MsgLockRequest{
		payer.String(), payeeList, fee,
	}
}

// Route returns the route of MsgRequestData - "oracle" (sdk.Msg interface).
func (msg MsgLockRequest) Route() string { return RouterKey }

// Type returns the message type of MsgRequestData (sdk.Msg interface).
func (msg MsgLockRequest) Type() string { return TypeMsgLockRequest }

// ValidateBasic checks whether the given MsgRequestData instance (sdk.Msg interface).
func (msg MsgLockRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat([]byte(msg.Payer)); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "sender: %s", msg.Payer)
	}

	for _, payee := range msg.Payees {
		if err := sdk.VerifyAddressFormat([]byte(payee)); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "sender: %s", msg.Payer)
		}
	}

	if !msg.Fee.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Fee.String())
	}

	return nil
}

// GetSigners returns the required signers for the given MsgRequestData (sdk.Msg interface).
func (msg MsgLockRequest) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Payer)
	return []sdk.AccAddress{sender}
}

// GetSignBytes returns raw JSON bytes to be signed by the signers (sdk.Msg interface).
func (msg MsgLockRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

func NewMsgClaimRequest(
	account sdk.AccAddress,
	guardedFeeId GuardedFeeID,
) *MsgClaimRequest {

	return &MsgClaimRequest{
		account.String(), guardedFeeId,
	}
}

// Route returns the route of MsgRequestData - "oracle" (sdk.Msg interface).
func (msg MsgClaimRequest) Route() string { return RouterKey }

// Type returns the message type of MsgRequestData (sdk.Msg interface).
func (msg MsgClaimRequest) Type() string { return TypeMsgLockRequest }

// ValidateBasic checks whether the given MsgRequestData instance (sdk.Msg interface).
func (msg MsgClaimRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat([]byte(msg.Payee)); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "sender: %s", msg.Payee)
	}
	return nil
}

// GetSigners returns the required signers for the given MsgRequestData (sdk.Msg interface).
func (msg MsgClaimRequest) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Payee)
	return []sdk.AccAddress{sender}
}

// GetSignBytes returns raw JSON bytes to be signed by the signers (sdk.Msg interface).
func (msg MsgClaimRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}
