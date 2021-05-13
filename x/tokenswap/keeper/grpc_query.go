package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/rizon-world/rizon/x/tokenswap/types"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// Tokenswap queries a tokenswap item by tx hash
func (k Querier) Tokenswap(c context.Context, req *types.QueryTokenswapRequest) (*types.QueryTokenswapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	if req.TxHash == "" {
		return nil, status.Error(codes.InvalidArgument, "tx hash cannot be empty")
	}
	// check whether tx hash is valid type
	// tx hash should be 64-length, lower case, hexadecimal type string
	if types.ValidTxHash(req.TxHash) == false {
		return nil, status.Error(codes.InvalidArgument, "invalid tx hash")
	}
	ctx := sdk.UnwrapSDKContext(c)

	res, err := k.Keeper.GetSwap(ctx, req.TxHash)
	if err != nil {
		panic(err)
	}

	return &types.QueryTokenswapResponse{Tokenswap: &res}, nil
}

// Params queries the parameters of tokenswap
func (k Querier) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.Keeper.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}
