package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	rizon "github.com/rizon-world/rizon/types"
	"github.com/rizon-world/rizon/x/tokenswap/types"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestSwapSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestSwap() {
	amount := sdk.NewDec(10)
	msg := types.NewMsgCreateTokenswapRequest(
		"1accecb6c80859478bbcf8fbc1bb04c2625efca9e79c050916f9fd318ca3e5d7",
		"rizon19hydu78c6wrwalr0g4gkdgv4glpt4j0hflc29s",
		types.DefaultSigner,
		amount,
	)
	newCoin := sdk.NewCoins(sdk.NewCoin(rizon.DefaultDenom, msg.Amount.RoundInt()))
	newSwap := types.NewTokenswap(msg.TxHash, msg.Receiver, msg.Signer, newCoin)
	err := suite.keeper.Swap(suite.ctx, newSwap)
	suite.Nil(err)
}

func (suite *KeeperTestSuite) TestSwapFail() {
	amount := sdk.NewDec(10)
	msg := types.NewMsgCreateTokenswapRequest(
		"1accecb6c80859478bbcf8fbc1bb04c2625efca9e79c050916f9fd318ca3e5d7",
		"rizon19hydu78c6wrwalr0g4gkdgv4glpt4j0hflc29s",
		types.DefaultSigner,
		amount,
	)
	newCoin := sdk.NewCoins(sdk.NewCoin(rizon.DefaultDenom, msg.Amount.RoundInt()))
	newSwapReceiver := types.NewTokenswap(msg.TxHash, "msg.Receiver", msg.Signer, newCoin)
	errReceiver := suite.keeper.Swap(suite.ctx, newSwapReceiver)
	suite.NotNil(errReceiver)
}
