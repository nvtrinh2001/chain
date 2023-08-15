package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/chain/v2/x/globalfee/types"
)

var _ types.QueryServer = &Querier{}

type Querier struct {
	Keeper
}

// MinimumGasPrices return minimum gas prices
func (q Querier) MinimumGasPrices(
	stdCtx context.Context,
	_ *types.QueryMinimumGasPricesRequest,
) (*types.QueryMinimumGasPricesResponse, error) {
	ctx := sdk.UnwrapSDKContext(stdCtx)
	return &types.QueryMinimumGasPricesResponse{
		MinimumGasPrices: q.GetParams(ctx).MinimumGasPrices,
	}, nil
}
