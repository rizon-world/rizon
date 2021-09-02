package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	rizon "github.com/rizon-world/rizon/app"
	"github.com/rizon-world/rizon/x/treasury/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

func TestMint(t *testing.T) {
	app := rizon.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	treasuryKeeper := app.TreasuryKeeper
	bankKeeper := app.BankKeeper

	coin, errCoin := sdk.ParseCoinNormalized("40rizon")
	require.Nil(t, errCoin)

	msg := types.NewMsgMintRequest(
		"rizon19hydu78c6wrwalr0g4gkdgv4glpt4j0hflc29s",
		"rizon148xe53d4438sr084ehf8xz30x4xm65x3yxpy6t",
		coin,
	)
	err := treasuryKeeper.Mint(ctx, msg.Receiver, msg.Amount)
	require.Nil(t, err)

	// If mint is success, bankKeeper give a "40rizon" as a reuslt of getbalance.
	mintedAddr, err := sdk.AccAddressFromBech32("rizon19hydu78c6wrwalr0g4gkdgv4glpt4j0hflc29s")
	require.Nil(t, err, "String type bech32 address should be changed AccAddress type.")
	accBalance := bankKeeper.GetBalance(ctx, mintedAddr, coin.Denom)
	require.Equal(t, "40rizon", accBalance.String())
}

func TestBurn(t *testing.T) {
	app := rizon.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	treasuryKeeper := app.TreasuryKeeper
	bankKeeper := app.BankKeeper

	coin, errCoin := sdk.ParseCoinNormalized("40rizon")
	require.Nil(t, errCoin)

	msgMint := types.NewMsgMintRequest(
		"rizon19hydu78c6wrwalr0g4gkdgv4glpt4j0hflc29s",
		"rizon148xe53d4438sr084ehf8xz30x4xm65x3yxpy6t",
		coin,
	)

	treasuryKeeper.Mint(ctx, msgMint.Receiver, msgMint.Amount)
	msg := types.NewMsgBurnRequest(
		"rizon19hydu78c6wrwalr0g4gkdgv4glpt4j0hflc29s",
		coin,
	)
	err := treasuryKeeper.Burn(ctx, msg.Signer, msg.Amount)
	require.Nil(t, err)

	// If burn is success, bankKeeper give a "0rizon" as a reuslt of getbalance.
	burnedAddr, err := sdk.AccAddressFromBech32("rizon19hydu78c6wrwalr0g4gkdgv4glpt4j0hflc29s")
	require.Nil(t, err, "String type bech32 address should be changed AccAddress type.")
	accBalance := bankKeeper.GetBalance(ctx, burnedAddr, coin.Denom)
	require.Equal(t, "0rizon", accBalance.String())
}
