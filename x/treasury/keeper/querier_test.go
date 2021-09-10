package keeper_test

import (
	rizon "github.com/rizon-world/rizon/app"
	"github.com/rizon-world/rizon/x/treasury/keeper"
	"github.com/rizon-world/rizon/x/treasury/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

func TestNewQuerier(t *testing.T) {
	app := rizon.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	treasuryKeeper := app.TreasuryKeeper
	cdc := app.LegacyAmino()

	req := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}
	querier := keeper.NewQuerier(treasuryKeeper, cdc)
	res, err := querier(ctx, []string{"other"}, req)
	require.Error(t, err)
	require.Nil(t, res)
}

func TestQueryCurrencies(t *testing.T) {
	app := rizon.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	treasuryKeeper := app.TreasuryKeeper
	cdc := app.LegacyAmino()

	querier := keeper.NewQuerier(treasuryKeeper, cdc)
	res, err := querier(ctx, []string{types.QueryCurrencies}, abci.RequestQuery{})
	require.Nil(t, err)

	var currencies types.Currencies
	e := cdc.UnmarshalJSON(res, currencies)
	require.Error(t, e)
}

func TestQueryCurrency(t *testing.T) {
	app := rizon.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	treasuryKeeper := app.TreasuryKeeper
	cdc := app.LegacyAmino()

	querier := keeper.NewQuerier(treasuryKeeper, cdc)
	res, err := querier(ctx, []string{types.QueryCurrency}, abci.RequestQuery{})
	require.Error(t, err)

	var param types.QueryCurrencyParam
	e := cdc.UnmarshalJSON(res, param)
	require.Error(t, e)
}

func TestQueryParams(t *testing.T) {
	app := rizon.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	treasuryKeeper := app.TreasuryKeeper
	cdc := app.LegacyAmino()

	querier := keeper.NewQuerier(treasuryKeeper, cdc)
	res, err := querier(ctx, []string{types.QueryParams}, abci.RequestQuery{})
	require.NoError(t, err)

	var param types.Params
	e := cdc.UnmarshalJSON(res, param)
	require.Error(t, e)
}
