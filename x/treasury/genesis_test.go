package treasury_test

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	rizon "github.com/rizon-world/rizon/app"
	"github.com/rizon-world/rizon/x/treasury"
	"github.com/rizon-world/rizon/x/treasury/keeper"
	"github.com/rizon-world/rizon/x/treasury/types"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

type TestSuite struct {
	suite.Suite

	cdc codec.JSONMarshaler
	ctx sdk.Context
	keeper keeper.Keeper
	app *rizon.RizonApp
}

func (suite *TestSuite) SetupTest() {
	app := rizon.Setup(false)
	suite.cdc = codec.NewAminoCodec(app.LegacyAmino())
	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	suite.keeper = app.TreasuryKeeper

	genesisState := types.DefaultGenesisState()
	stateBytes, err := json.MarshalIndent(genesisState, "", "  ")
	suite.NoError(err)

	// Initialize the chain
	app.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	app.Commit()

	suite.app = app
}

func TestGenesisSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) TestExportGenesis() {
	exportedGenesis := treasury.ExportGenesis(suite.ctx, suite.keeper)
	defaultGenesis := types.DefaultGenesisState()
	suite.Equal(exportedGenesis, defaultGenesis)
}

func (suite *TestSuite) TestInitGenesis() {
	treasury.InitGenesis(suite.ctx, suite.keeper, types.DefaultGenesisState())

	acc := suite.keeper.GetCurrency(suite.ctx, "skrw")
	suite.Equal("skrw", acc.Denom)
}