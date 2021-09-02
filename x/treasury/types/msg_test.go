package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	rizon "github.com/rizon-world/rizon/app"
	tkswaptypes "github.com/rizon-world/rizon/x/tokenswap/types"
	"github.com/rizon-world/rizon/x/treasury/types"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	receiver      = "rizon1alhgwh8cex84aacc62xxh3yslq4wgprgnhtsvc"
	signer        = tkswaptypes.DefaultSigner
	coin, coinErr = sdk.ParseCoinNormalized("10uatolo")
)

func init() {
	rizon.SetRizonConfig()
}

// MsgMintRequest Test
func TestMsgMintRequest(t *testing.T) {
	msg := types.NewMsgMintRequest(receiver, signer, coin)
	require.Equal(t, receiver, msg.Receiver)
	require.Equal(t, signer, msg.Signer)
	require.Equal(t, coin, msg.Amount)
}

func TestMsgMintRequest_Route(t *testing.T) {
	msg := types.NewMsgMintRequest(receiver, signer, coin)
	require.Equal(t, types.RouterKey, msg.Route())
}

func TestMsgMintRequest_Type(t *testing.T) {
	msg := types.NewMsgMintRequest(receiver, signer, coin)
	require.Equal(t, "MintRequest", msg.Type())
}

func TestMsgMintRequest_GetSignBytes(t *testing.T) {
	msg := types.NewMsgMintRequest(receiver, signer, coin)
	res := msg.GetSignBytes()
	expected := `{"type":"treasury/MsgMintRequest","value":{"amount":{"amount":"10","denom":"uatolo"},"receiver":"rizon1alhgwh8cex84aacc62xxh3yslq4wgprgnhtsvc","signer":"rizon1cafygu0kppg46tq9kkzz9z9nrf5v8zwwnf5t9l"}}`
	require.Equal(t, expected, string(res))
}

func TestMsgMintRequest_GetSigners(t *testing.T) {
	msg := types.NewMsgMintRequest(receiver, signer, coin)
	res := msg.GetSigners()
	for _, addr := range res {
		require.Equal(t, signer, addr.String())
	}
}

// MsgBurnRequest Test
func TestMsgBurnRequest(t *testing.T) {
	msg := types.NewMsgBurnRequest(signer, coin)
	require.Equal(t, signer, msg.Signer)
	require.Equal(t, coin, msg.Amount)
}

func TestMsgBurnRequest_Route(t *testing.T) {
	msg := types.NewMsgBurnRequest(signer, coin)
	require.Equal(t, types.RouterKey, msg.Route())
}

func TestMsgBurnRequest_Type(t *testing.T) {
	msg := types.NewMsgBurnRequest(signer, coin)
	require.Equal(t, "BurnRequest", msg.Type())
}

func TestMsgBurnRequest_GetSignBytes(t *testing.T) {
	msg := types.NewMsgBurnRequest(signer, coin)
	res := msg.GetSignBytes()
	expected := `{"type":"treasury/MsgBurnRequest","value":{"amount":{"amount":"10","denom":"uatolo"},"signer":"rizon1cafygu0kppg46tq9kkzz9z9nrf5v8zwwnf5t9l"}}`
	require.Equal(t, expected, string(res))
}

func TestMsgBurnRequest_GetSigners(t *testing.T) {
	msg := types.NewMsgBurnRequest(signer, coin)
	res := msg.GetSigners()
	for _, addr := range res {
		require.Equal(t, signer, addr.String())
	}
}
