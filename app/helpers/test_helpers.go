package helpers

import (
	"encoding/json"
	"time"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/CosmWasm/wasmd/x/wasm"

	rizonapp "github.com/rizon-world/rizon/app"
	rizontypes "github.com/rizon-world/rizon/types"
)

var SetConfigFlag = false

func isConfigSet() bool {
	return SetConfigFlag
}
func toggleConfigFlag() {
	SetConfigFlag = !SetConfigFlag
}

// DefaultConsensusParams defines the default Tendermint consensus params used in
// SimApp testing.
var DefaultConsensusParams = &abci.ConsensusParams{
	Block: &abci.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}

func SetRizonConfig() {
	if !isConfigSet() {
		rizontypes.SetConfig()
		toggleConfigFlag()
	}
}

func setup(withGenesis bool, invCheckPeriod uint) (*rizonapp.RizonApp, rizonapp.GenesisState) {
	SetRizonConfig()

	db := dbm.NewMemDB()
	encCdc := rizonapp.MakeTestEncodingConfig()
	var emptyWasmOpts []wasm.Option

	// A Nop logger is set in RizonApp.
	app := rizonapp.NewRizonApp(log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		rizonapp.DefaultNodeHome,
		invCheckPeriod,
		encCdc,
		rizonapp.GetEnabledProposals(),
		EmptyAppOptions{},
		emptyWasmOpts,
	)

	if withGenesis {
		return app, rizonapp.NewDefaultGenesisState()
	}
	return app, rizonapp.GenesisState{}
}

// Setup initializes a new RizonApp.
func Setup(isCheckTx bool) *rizonapp.RizonApp {
	app, genesisState := setup(!isCheckTx, 5)
	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		app.InitChain(abci.RequestInitChain{
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: DefaultConsensusParams,
			AppStateBytes:   stateBytes,
		})
	}

	return app
}

// EmptyAppOptions is a stub implementing AppOptions
type EmptyAppOptions struct{}

// Get implements AppOptions
func (ao EmptyAppOptions) Get(o string) interface{} {
	return nil
}
