package nebulae

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	icacontrollertypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/types"

	"github.com/rizon-world/rizon/app/keepers"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("Preparing module migrations...")

		// adding denom metadata
		keepers.BankKeeper.SetDenomMetaData(ctx, getDenomMetadata())
		ctx.Logger().Info("Adding denom metadata complete")

		// Change hostParams allow_messages = [*] instead of whitelisting individual messages
		hostParams := icahosttypes.Params{
			HostEnabled:   true,
			AllowMessages: []string{"*"},
		}

		// Update params for ica host keepers
		keepers.ICAHostKeeper.SetParams(ctx, hostParams)
		ctx.Logger().Info("Updated ICAHostKeeper params")

		ctx.Logger().Info("Starting module migrations...")

		vm, err := mm.RunMigrations(ctx, configurator, vm)
		if err != nil {
			return vm, err
		}

		ctx.Logger().Info("Setting module parameters...")

		// Update params for ica controller keepers
		keepers.ICAControllerKeeper.SetParams(ctx, icacontrollertypes.Params{ControllerEnabled: true})
		ctx.Logger().Info("Updated ICAControllerKeeper params")

		ctx.Logger().Info("Upgrade complete")
		return vm, err
	}
}

func getDenomMetadata() banktypes.Metadata {
	return banktypes.Metadata{
		Name:        "ATOLO",
		Symbol:      "ATOLO",
		Description: "The native staking token of the RIZON Blockchain",
		DenomUnits: []*banktypes.DenomUnit{
			{"uatolo", uint32(0), []string{"microatolo"}},
			{"matolo", uint32(3), []string{"milliatolo"}},
			{"atolo", uint32(6), nil},
		},
		Base:    "uatolo",
		Display: "atolo",
	}
}
