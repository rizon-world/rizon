package keeper_test

import (
	rizon "github.com/rizon-world/rizon/app"
	"github.com/rizon-world/rizon/x/tokenswap/keeper"
	"github.com/rizon-world/rizon/x/tokenswap/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

func TestNewQuerier(t *testing.T) {
	app := rizon.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	tokenswapKeeper := app.TokenswapKeeper
	cdc := app.LegacyAmino()

	req := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}
	querier := keeper.NewQuerier(tokenswapKeeper, cdc)
	res, err := querier(ctx, []string{"other"}, req)
	require.Error(t, err)
	require.Nil(t, res)
}

func TestQueryTokenswap(t *testing.T) {
	app := rizon.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	tokenswapKeeper := app.TokenswapKeeper
	cdc := app.LegacyAmino()

	querier := keeper.NewQuerier(tokenswapKeeper, cdc)
	res, err := querier(ctx, []string{types.QueryTokenswap}, abci.RequestQuery{})
	require.Error(t, err)

	var swap types.Tokenswap
	e := cdc.UnmarshalJSON(res, swap)
	require.Error(t, e)
}

func TestQueryParams(t *testing.T) {
	app := rizon.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	tokenswapKeeper := app.TokenswapKeeper
	cdc := app.LegacyAmino()

	querier := keeper.NewQuerier(tokenswapKeeper, cdc)
	res, err := querier(ctx, []string{types.QueryParams}, abci.RequestQuery{})
	require.NoError(t, err)

	var param types.Params
	e := cdc.UnmarshalJSON(res, param)
	require.Error(t, e)
}
