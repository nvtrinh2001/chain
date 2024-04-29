package keeper

import (
	"context"
	"fmt"
	"github.com/bandprotocol/chain/v2/x/guardian/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) Lock(goCtx context.Context, msg *types.MsgLockRequest) (*types.MsgLockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	payer, err := sdk.AccAddressFromBech32(msg.Payer)
	if err != nil {
		return nil, err
	}

	var payeeList []sdk.AccAddress
	for _, payee := range msg.Payees {
		payeeAccount, err := sdk.AccAddressFromBech32(payee)
		if err != nil {
			return nil, err
		}
		payeeList = append(payeeList, payeeAccount)
	}

	var paidFee sdk.Coins
	for _, coin := range msg.Fee {
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
		return nil, err
	}

	id := k.AddGuardedFee(ctx, types.NewGuardedFee(
		payer, payeeList, msg.Fee,
	))

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeLock,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", id)),
	))

	return &types.MsgLockResponse{}, nil
}

func (k msgServer) Claim(goCtx context.Context, msg *types.MsgClaimRequest) (*types.MsgClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	payee, err := sdk.AccAddressFromBech32(msg.Payee)
	if err != nil {
		return nil, err
	}

	guardedFee, err := k.GetGuardedFee(ctx, msg.GuardedFeeID)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, k.feeCollectorName, payee, guardedFee.Fee,
	)

	if err != nil {
		return nil, err
	}

	for i, payeeInFee := range guardedFee.Payees {
		if payeeInFee.Payee == msg.Payee {
			guardedFee.Payees[i].Status = types.STATUS_CLAIMED
			break
		}
	}

	_ = k.UpdateGuardedFee(ctx, msg.GuardedFeeID, guardedFee)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeClaim,
		sdk.NewAttribute(types.AttributeKeyID, fmt.Sprintf("%d", msg.GuardedFeeID)),
		sdk.NewAttribute(types.AttributeClaimedPayee, fmt.Sprintf("%v", msg.Payee)),
	))

	return &types.MsgClaimResponse{}, nil
}
