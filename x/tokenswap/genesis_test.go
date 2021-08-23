package tokenswap_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	rizon "github.com/rizon-world/rizon/app"
	"github.com/rizon-world/rizon/x/tokenswap"
	"github.com/rizon-world/rizon/x/tokenswap/keeper"
	"github.com/rizon-world/rizon/x/tokenswap/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

type TestSuite struct {
	suite.Suite

	cdc codec.JSONMarshaler
	ctx sdk.Context
	keeper keeper.Keeper
}

func (suite *TestSuite) SetupTest() {
	app := rizon.Setup(false)
	suite.cdc = codec.NewAminoCodec(app.LegacyAmino())
	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	suite.keeper = app.TokenswapKeeper
}

func TestGenesisSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) TestExportGenesis() {
	exportedGenesis := tokenswap.ExportGenesis(suite.ctx, suite.keeper)
	defaultGenesis := types.DefaultGenesisState()
	suite.Equal(exportedGenesis, defaultGenesis)
}

