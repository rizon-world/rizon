package keeper_test

import (
	"github.com/rizon-world/rizon/x/treasury/keeper"
	"github.com/rizon-world/rizon/x/treasury/types"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestQuerierSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestNewQuerier() {
	req := abci.RequestQuery {
		Path: "",
		Data: []byte{},
	}
	querier := keeper.NewQuerier(suite.keeper, suite.legacycdc)
	res, err := querier(suite.ctx, []string{"other"}, req)
	suite.Error(err)
	suite.Nil(res)
}

func (suite *KeeperTestSuite) TestQueryCurrencies() {
	querier := keeper.NewQuerier(suite.keeper, suite.legacycdc)
	res, err := querier(suite.ctx, []string{types.QueryCurrencies}, abci.RequestQuery{})
	suite.Nil(err)

	var currencies types.Currencies
	e := suite.legacycdc.UnmarshalJSON(res, currencies)
	suite.Error(e)
}

func (suite *KeeperTestSuite) TestQueryCurrency() {
	querier := keeper.NewQuerier(suite.keeper, suite.legacycdc)
	res, err := querier(suite.ctx, []string{types.QueryCurrency}, abci.RequestQuery{})
	suite.Error(err)

	var param types.QueryCurrencyParam
	e := suite.legacycdc.UnmarshalJSON(res, param)
	suite.Error(e)
}

func (suite *KeeperTestSuite) TestQueryParams() {
	querier := keeper.NewQuerier(suite.keeper, suite.legacycdc)
	res, err := querier(suite.ctx, []string{types.QueryParams}, abci.RequestQuery{})
	suite.NoError(err)

	var param types.Params
	e := suite.legacycdc.UnmarshalJSON(res, param)
	suite.Error(e)
}
