package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/bandprotocol/chain/v2/x/feelocker/types"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

func (k Querier) QueryLockedFeeList(c context.Context, req *types.QueryLockedFeeListRequest) (*types.QueryLockedFeeListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	lockedFeeList, err := k.GetLockedFeeList(ctx, req.AccountAddress, req.Status)
	if err != nil {
		return nil, err
	}

	return &types.QueryLockedFeeListResponse{LockedFees: lockedFeeList}, nil
}

func (k Querier) QueryLockedFee(c context.Context, req *types.QueryLockedFeeRequest) (*types.QueryLockedFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	lockedFee, err := k.GetLockedFee(ctx, types.LockedFeeID(req.LockedFeeId))
	if err != nil {
		return nil, err
	}

	return &types.QueryLockedFeeResponse{LockedFee: &lockedFee}, nil
}
