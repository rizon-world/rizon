package keeper

import (
	"github.com/rizon-world/rizon/x/tokenswap/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	abci "github.com/tendermint/tendermint/abci/types"
)

// creates a querier for legacy rest endpoints
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		var (
			res []byte
			err error
		)

		switch path[0] {
		case types.QueryTokenswap:
			res, err = queryTokenswap(ctx, req, k, legacyQuerierCdc)

		case types.QuerySwappedAmount:
			res, err = querySwappedAmount(ctx, k, legacyQuerierCdc)

		case types.QueryParams:
			res, err = queryParams(ctx, k, legacyQuerierCdc)

		default:
			err = sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}

		return res, err
	}
}

// query a tokenswap item by tx hash
func queryTokenswap(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var param types.QueryTokenswapParam

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &param)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	if types.ValidTxHash(param.TxHash) == false {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid tx hash (%s)", param.TxHash)
	}
	res, err := k.GetSwap(ctx, param.TxHash)
	if err != nil {
		panic(err)
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, res)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

// query current swapped amount of tokenswap
func querySwappedAmount(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	amt := k.GetSwappedAmount(ctx)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, types.NewSwappedAmount(amt))
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

// query the tokenswap parameters
func queryParams(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	params := k.GetParams(ctx)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
