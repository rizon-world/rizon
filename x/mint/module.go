// fork from github.com/cosmos/cosmos-sdk v0.42.5
package mint

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/mint/keeper"
	"github.com/cosmos/cosmos-sdk/x/mint/types"
)

var (
	_ module.AppModule = AppModule{}
)

// AppModule implements an application module for the mint module.
type AppModule struct {
	mint.AppModule

	keeper     keeper.Keeper
	authKeeper types.AccountKeeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Marshaler, keeper keeper.Keeper, ak types.AccountKeeper) AppModule {
	return AppModule{
		AppModule:  mint.NewAppModule(cdc, keeper, ak),
		keeper:     keeper,
		authKeeper: ak,
	}
}

// BeginBlock returns the begin blocker for the mint module.
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	BeginBlocker(ctx, am.keeper)
}
