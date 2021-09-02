package tokenswap_test

import (
	"github.com/rizon-world/rizon/x/tokenswap"
	"testing"

	rizon "github.com/rizon-world/rizon/app"
	"github.com/rizon-world/rizon/x/tokenswap/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

// After InitGenesis, Default Limit, Signer, Swappable config is correctly set in param
func TestExportGenesis(t *testing.T) {
	app := rizon.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	keeper := app.TokenswapKeeper

	exportedGenesis := tokenswap.ExportGenesis(ctx, keeper)
	defaultGenesis := types.DefaultGenesisState()

	// In param struct, limit uses int64 type
	var expectedLimit int64 = types.DefaultLimit
	require.Equal(t, exportedGenesis, defaultGenesis)
	require.Equal(t, expectedLimit, exportedGenesis.GetParams().Limit)
	require.Equal(t, types.DefaultSigner, exportedGenesis.GetParams().Signer)
	require.Equal(t, types.DefaultSwappable, exportedGenesis.GetParams().Swappable)
}
