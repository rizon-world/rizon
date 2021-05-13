package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/rizon-world/rizon/x/treasury/types"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// Currencies queries all supported currency denom list
func (k Querier) Currencies(c context.Context, req *types.QueryCurrenciesRequest) (*types.QueryCurrenciesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	store := k.Keeper.Store(ctx)
	currencyStore := prefix.NewStore(store, types.PrefixCurrency)
	var denoms []string

	pageRes, err := query.Paginate(currencyStore, req.Pagination, func(_, value []byte) error {
		var currency types.Currency
		err := k.cdc.UnmarshalBinaryBare(value, &currency)
		if err != nil {
			return err
		}
		denoms = append(denoms, currency.Denom)
		return nil
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}

	currencies := types.NewCurrencies(denoms)

	return &types.QueryCurrenciesResponse{
		Currencies: &currencies,
		Pagination: pageRes,
	}, nil
}

// Currency queries an information of a currency
func (k Querier) Currency(c context.Context, req *types.QueryCurrencyRequest) (*types.QueryCurrencyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	if req.Denom == "" {
		return nil, status.Error(codes.InvalidArgument, "denom cannot be empty")
	}
	// check whether denom is valid
	if err := sdk.ValidateDenom(req.Denom); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid denom")
	}
	ctx := sdk.UnwrapSDKContext(c)

	res := k.Keeper.GetCurrency(ctx, req.Denom)

	return &types.QueryCurrencyResponse{Currency: &res}, nil
}

// Params queries the parameters of treasury
func (k Querier) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.Keeper.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}
