package keeper_test

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	rizon "github.com/rizon-world/rizon/app"
	"github.com/rizon-world/rizon/x/treasury/keeper"
	"github.com/rizon-world/rizon/x/treasury/types"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc 		codec.JSONMarshaler
	legacycdc 	*codec.LegacyAmino
	ctx 		sdk.Context
	keeper 		keeper.Keeper
	app 		*rizon.RizonApp
}

func (suite *KeeperTestSuite) SetupTest() {
	app := rizon.Setup(false)
	suite.cdc = codec.NewAminoCodec(app.LegacyAmino())
	suite.legacycdc = app.LegacyAmino()
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

func TestMintBurnSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestMint() {
	coin, errCoin := sdk.ParseCoinNormalized("40rizon")
	suite.Nil(errCoin)
	msg := types.NewMsgMintRequest(
		"rizon19hydu78c6wrwalr0g4gkdgv4glpt4j0hflc29s",
		"rizon148xe53d4438sr084ehf8xz30x4xm65x3yxpy6t",
		coin,
		)
	err := suite.keeper.Mint(suite.ctx, msg.Receiver, msg.Amount)
	suite.Nil(err)
}

func (suite *KeeperTestSuite) TestBurn() {
	coin, errCoin := sdk.ParseCoinNormalized("40rizon")
	suite.Nil(errCoin)
	msgMint := types.NewMsgMintRequest(
		"rizon19hydu78c6wrwalr0g4gkdgv4glpt4j0hflc29s",
		"rizon148xe53d4438sr084ehf8xz30x4xm65x3yxpy6t",
		coin,
	)
	suite.keeper.Mint(suite.ctx, msgMint.Receiver, msgMint.Amount)
	msg := types.NewMsgBurnRequest(
		"rizon19hydu78c6wrwalr0g4gkdgv4glpt4j0hflc29s",
		coin,
		)
	err := suite.keeper.Burn(suite.ctx, msg.Signer, msg.Amount)
	suite.Nil(err)
}