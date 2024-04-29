package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/bandprotocol/chain/v2/x/guardian/types"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

func (k Querier) QueryGuardedFeeList(c context.Context, req *types.QueryGuardedFeeListRequest) (*types.QueryGuardedFeeListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	guardedFeeList, err := k.GetGuardedFeeList(ctx, req.AccountAddress, req.Status)
	if err != nil {
		return nil, err
	}

	return &types.QueryGuardedFeeListResponse{GuardedFees: guardedFeeList}, nil
}

func (k Querier) QueryGuardedFee(c context.Context, req *types.QueryGuardedFeeRequest) (*types.QueryGuardedFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	guardedFee, err := k.GetGuardedFee(ctx, types.GuardedFeeID(req.GuardedFeeId))
	if err != nil {
		return nil, err
	}

	return &types.QueryGuardedFeeResponse{GuardedFee: &guardedFee}, nil
}
