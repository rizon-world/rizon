package keeper_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	rizon "github.com/rizon-world/rizon/app"
	rizontypes "github.com/rizon-world/rizon/types"
	"github.com/rizon-world/rizon/x/tokenswap/types"
)

func TestSwap(t *testing.T) {
	app := rizon.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	tokenswapKeeper := app.TokenswapKeeper

	amount := sdk.NewDec(10)
	msg := types.NewMsgCreateTokenswapRequest(
		"1accecb6c80859478bbcf8fbc1bb04c2625efca9e79c050916f9fd318ca3e5d7",
		"rizon19hydu78c6wrwalr0g4gkdgv4glpt4j0hflc29s",
		types.DefaultSigner,
		amount,
	)
	newCoin := sdk.NewCoins(sdk.NewCoin(rizontypes.DefaultDenom, msg.Amount.RoundInt()))
	newSwap := types.NewTokenswap(msg.TxHash, msg.Receiver, msg.Signer, newCoin)
	err := tokenswapKeeper.Swap(ctx, newSwap)
	// If success, return nil
	require.Nil(t, err)
}

// Receiver address should be an address of bech32 string
func TestSwapFail(t *testing.T) {
	app := rizon.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	tokenswapKeeper := app.TokenswapKeeper

	amount := sdk.NewDec(10)
	msg := types.NewMsgCreateTokenswapRequest(
		"1accecb6c80859478bbcf8fbc1bb04c2625efca9e79c050916f9fd318ca3e5d7",
		"rizon19hydu78c6wrwalr0g4gkdgv4glpt4j0hflc29s",
		types.DefaultSigner,
		amount,
	)
	newCoin := sdk.NewCoins(sdk.NewCoin(rizontypes.DefaultDenom, msg.Amount.RoundInt()))
	newSwapReceiver := types.NewTokenswap(msg.TxHash, "msg.Receiver", msg.Signer, newCoin)
	errReceiver := tokenswapKeeper.Swap(ctx, newSwapReceiver)
	require.NotNil(t, errReceiver)
}
