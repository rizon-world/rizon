package treasury

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/rizon-world/rizon/x/treasury/keeper"
	"github.com/rizon-world/rizon/x/treasury/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)
	k.SetSequence(ctx, state.Seq)
	// update parameter's currencies to state
	var denoms []string
	for _, c := range state.Params.CurrencyList {
		k.SetCurrency(ctx, c)
		denoms = append(denoms, c.Denom)
	}
	k.SetCurrencies(ctx, types.NewCurrencies(denoms))
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	seq := k.GetSequence(ctx)
	return types.NewGenesisState(
		k.GetParams(ctx),
		types.NewSequence(seq),
	)
}
