package types_test

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	rizon "github.com/rizon-world/rizon/app"
	tkswaptypes "github.com/rizon-world/rizon/x/tokenswap/types"
	"github.com/rizon-world/rizon/x/treasury/keeper"
	"github.com/rizon-world/rizon/x/treasury/types"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

var (
	receiver = "rizon1alhgwh8cex84aacc62xxh3yslq4wgprgnhtsvc"
	signer = tkswaptypes.DefaultSigner
	coin, coinErr = sdk.ParseCoinNormalized("10uatolo")
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

func TestMsgSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) TestMsgMintRequest() {
	msg := types.NewMsgMintRequest(receiver, signer, coin)
	suite.Equal(receiver, msg.Receiver)
	suite.Equal(signer, msg.Signer)
	suite.Equal(coin, msg.Amount)
}

func (suite *TestSuite) TestMsgMintRequest_Route() {
	msg := types.NewMsgMintRequest(receiver, signer, coin)
	suite.Equal(types.RouterKey, msg.Route())
}

func (suite *TestSuite) TestMsgMintRequest_Type() {
	msg := types.NewMsgMintRequest(receiver, signer, coin)
	suite.Equal("MintRequest", msg.Type())
}

func (suite *TestSuite) TestMsgMintRequest_GetSignBytes() {
	msg := types.NewMsgMintRequest(receiver, signer, coin)
	res := msg.GetSignBytes()
	expected := `{"type":"treasury/MsgMintRequest","value":{"amount":{"amount":"10","denom":"uatolo"},"receiver":"rizon1alhgwh8cex84aacc62xxh3yslq4wgprgnhtsvc","signer":"rizon1cafygu0kppg46tq9kkzz9z9nrf5v8zwwnf5t9l"}}`
	suite.Equal(expected, string(res))
}

func (suite *TestSuite) TestMsgMintRequest_GetSigners() {
	msg := types.NewMsgMintRequest(receiver, signer, coin)
	res := msg.GetSigners()
	for _, addr := range res {
		suite.Equal(signer, addr.String())
	}
}

func (suite *TestSuite) TestMsgBurnRequest() {
	msg := types.NewMsgBurnRequest(signer, coin)
	suite.Equal(signer, msg.Signer)
	suite.Equal(coin, msg.Amount)
}

func (suite *TestSuite) TestMsgBurnRequest_Route() {
	msg := types.NewMsgBurnRequest(signer, coin)
	suite.Equal(types.RouterKey, msg.Route())
}

func (suite *TestSuite) TestMsgBurnRequest_Type() {
	msg := types.NewMsgBurnRequest(signer, coin)
	suite.Equal("BurnRequest", msg.Type())
}

func (suite *TestSuite) TestMsgBurnRequest_GetSignBytes() {
	msg := types.NewMsgBurnRequest(signer, coin)
	res := msg.GetSignBytes()
	expected := `{"type":"treasury/MsgBurnRequest","value":{"amount":{"amount":"10","denom":"uatolo"},"signer":"rizon1cafygu0kppg46tq9kkzz9z9nrf5v8zwwnf5t9l"}}`
	suite.Equal(expected, string(res))
}

func (suite *TestSuite) TestMsgBurnRequest_GetSigners() {
	msg := types.NewMsgBurnRequest(signer, coin)
	res := msg.GetSigners()
	for _, addr := range res {
		suite.Equal(signer, addr.String())
	}
}