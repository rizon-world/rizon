package keeper_test

import (
	"github.com/rizon-world/rizon/x/tokenswap/keeper"
	"github.com/rizon-world/rizon/x/tokenswap/types"
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
	querier := keeper.NewQuerier(suite.keeper, suite.cdc)
	res, err := querier(suite.ctx, []string{"other"}, req)
	suite.Error(err)
	suite.Nil(res)
}

func (suite *KeeperTestSuite) TestqueryTokenswap() {
	querier := keeper.NewQuerier(suite.keeper, suite.cdc)
	res, err := querier(suite.ctx, []string{types.QueryTokenswap}, abci.RequestQuery{})
	suite.Error(err)

	var swap types.Tokenswap
	e := suite.cdc.UnmarshalJSON(res, swap)
	suite.Error(e)
}

func (suite *KeeperTestSuite) TestqueryParams() {
	querier := keeper.NewQuerier(suite.keeper, suite.cdc)
	res, err := querier(suite.ctx, []string{types.QueryParams}, abci.RequestQuery{})
	suite.NoError(err)

	var param types.Params
	e := suite.cdc.UnmarshalJSON(res, param)
	suite.Error(e)
}
