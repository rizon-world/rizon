package rizon_test

import (
	"encoding/json"
	"os"
	"testing"

	rizonapp "github.com/rizon-world/rizon/app"
	"github.com/rizon-world/rizon/app/helpers"
	"github.com/rizon-world/rizon/types"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

func TestRizonAppExportAndBlockedAddress(t *testing.T) {
	types.SetConfig()

	encCfg := rizonapp.MakeTestEncodingConfig()
	db := dbm.NewMemDB()
	var emptyWasmOpts []wasm.Option
	app := rizonapp.NewRizonApp(
		log.NewTMLogger(log.NewSyncWriter(os.Stdout)),
		db,
		nil,
		true,
		map[int64]bool{},
		rizonapp.DefaultNodeHome,
		0,
		encCfg,
		rizonapp.GetEnabledProposals(),
		helpers.EmptyAppOptions{},
		emptyWasmOpts)

	genesisState := rizonapp.NewDefaultGenesisState()
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

	// Making a new app object with the db, so that initchain hasn't been called
	app2 := rizonapp.NewRizonApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, rizonapp.DefaultNodeHome, 0, encCfg, rizonapp.GetEnabledProposals(), helpers.EmptyAppOptions{}, emptyWasmOpts)
	_, err = app2.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}
