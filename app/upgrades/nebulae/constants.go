package nebulae

import (
	store "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/gaia/v10/x/globalfee"
	icacontrollertypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v4/modules/apps/29-fee/types"
	ibchookstypes "github.com/osmosis-labs/osmosis/x/ibc-hooks/types"
	routertypes "github.com/strangelove-ventures/packet-forward-middleware/v4/router/types"

	"github.com/rizon-world/rizon/app/upgrades"
)

const (
	// UpgradeName defines the on-chain upgrade name.
	UpgradeName = "Nebulae"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			globalfee.ModuleName,
			ibcfeetypes.ModuleName,
			ibchookstypes.StoreKey,
			icacontrollertypes.StoreKey,
			routertypes.StoreKey,
			wasm.ModuleName,
		},
	},
}
