package keeper_test

import (
	rizon "github.com/rizon-world/rizon/app/helpers"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

func TestStore(t *testing.T) {
	app := rizon.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	keeper := app.TokenswapKeeper
	accStore := keeper.Store(ctx)
	accLogger := keeper.Logger(ctx)

	require.NotNil(t, accStore)
	// When rizon daemon is started, the logger is empty instance.
	require.NotNil(t, accLogger)
}
