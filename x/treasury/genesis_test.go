package treasury_test

import (
	"encoding/json"
	"github.com/stretchr/testify/require"

	rizon "github.com/rizon-world/rizon/app/helpers"
	"github.com/rizon-world/rizon/x/treasury"
	"github.com/rizon-world/rizon/x/treasury/types"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

func TestExportGenesis(t *testing.T) {
	app := rizon.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	treasuryKeeper := app.TreasuryKeeper

	genesisState := types.DefaultGenesisState()
	stateBytes, err := json.MarshalIndent(genesisState, "", "  ")
	require.NoError(t, err)

	// Initialize the chain
	app.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	app.Commit()

	exportedGenesis := treasury.ExportGenesis(ctx, treasuryKeeper)
	defaultGenesis := types.DefaultGenesisState()
	require.Equal(t, exportedGenesis, defaultGenesis)
}

// After InitGenesis, "skrw" denom is set as Currency into KVStore
func TestInitGenesis(t *testing.T) {
	app := rizon.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	treasuryKeeper := app.TreasuryKeeper

	treasury.InitGenesis(ctx, treasuryKeeper, types.DefaultGenesisState())

	acc := treasuryKeeper.GetCurrency(ctx, "skrw")
	require.Equal(t, "skrw", acc.Denom)
}
