package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	rizon "github.com/rizon-world/rizon/app"
	"github.com/rizon-world/rizon/x/tokenswap/keeper"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc		*codec.LegacyAmino
	ctx 	sdk.Context
	keeper 	keeper.Keeper
	app 	*rizon.RizonApp
}

func (suite *KeeperTestSuite) SetupTest() {
	app := rizon.Setup(false)

	suite.app = app
	suite.cdc = app.LegacyAmino()
	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	suite.keeper = app.TokenswapKeeper
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestStore() {
	acc := suite.keeper.Store(suite.ctx)
	suite.NotNil(acc)
}
