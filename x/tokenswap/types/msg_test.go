package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	rizon "github.com/rizon-world/rizon/app/helpers"
	"github.com/rizon-world/rizon/x/tokenswap/types"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	tx_hash           = "84aa2e585705d40173c67e9ccbf4c71ab39984f7477e160a276252fd7d0b2f74"
	receiver          = "rizon1alhgwh8cex84aacc62xxh3yslq4wgprgnhtsvc"
	signer            = "rizon146a27r6mcx8qg8sl78ue49na3hk26un899yd4l"
	amount            = sdk.NewDec(1000)
	pAmount  *sdk.Dec = &amount

	invalid_tx_hash  = "84aa2e585705d40173c67e9ccbf4c71ab39984f7477e160a276252fd7d"
	invalid_receiver = "hdac1alhgwh8cex84aacc62xxh3yslq4wgprgnhtsvc"
	invalid_signer   = "hdac1alhgwh8cex84aacc62xxh3yslq4wgprgnhtsvc"
)

func init() {
	rizon.SetRizonConfig()
}

// MsgCreateTokenswapRequest Test
func TestNewMsgCreateTokenswapRequest(t *testing.T) {
	msg := types.NewMsgCreateTokenswapRequest(tx_hash, receiver, signer, amount)
	require.Equal(t, tx_hash, msg.TxHash)
	require.Equal(t, receiver, msg.Receiver)
	require.Equal(t, signer, msg.Signer)
	require.Equal(t, pAmount, msg.Amount)
}

func TestMsgCreateTokenswapRequest_Route(t *testing.T) {
	msg := types.NewMsgCreateTokenswapRequest(tx_hash, receiver, signer, amount)
	require.Equal(t, types.RouterKey, msg.Route())
}

func TestMsgCreateTokenswapRequest_Type(t *testing.T) {
	msg := types.NewMsgCreateTokenswapRequest(tx_hash, receiver, signer, amount)
	require.Equal(t, "CreateTokenswapRequest", msg.Type())
}

func TestMsgCreateTokenswapRequest_GetSignBytes(t *testing.T) {
	msg := types.NewMsgCreateTokenswapRequest(tx_hash, receiver, signer, amount)
	res := msg.GetSignBytes()
	expected := `{"type":"tokenswap/MsgCreateTokenswapRequest","value":{"amount":"1000.000000000000000000","receiver":"rizon1alhgwh8cex84aacc62xxh3yslq4wgprgnhtsvc","signer":"rizon146a27r6mcx8qg8sl78ue49na3hk26un899yd4l","tx_hash":"84aa2e585705d40173c67e9ccbf4c71ab39984f7477e160a276252fd7d0b2f74"}}`
	require.Equal(t, expected, string(res))
}

func TestMsgCreateTokenswapRequest_GetSigners(t *testing.T) {
	msg := types.NewMsgCreateTokenswapRequest(tx_hash, receiver, signer, amount)
	res := msg.GetSigners()
	for _, addr := range res {
		require.Equal(t, signer, addr.String())
	}
}

func TestMsgCreateTokenswapRequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name       string
		expectPass bool
		msg        *types.MsgCreateTokenswapRequest
	}{
		{"pass", true, types.NewMsgCreateTokenswapRequest(tx_hash, receiver, signer, amount)},
		{"invalid hash", false, types.NewMsgCreateTokenswapRequest(invalid_tx_hash, receiver, signer, amount)},
		{"invalid receiver address", false, types.NewMsgCreateTokenswapRequest(tx_hash, invalid_receiver, signer, amount)},
		{"invalid signer", false, types.NewMsgCreateTokenswapRequest(tx_hash, receiver, invalid_signer, amount)},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectPass {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
