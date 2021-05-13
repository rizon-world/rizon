package treasury

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/rizon-world/rizon/x/treasury/keeper"
	"github.com/rizon-world/rizon/x/treasury/types"
)

// EndBlocker updates currencies status from parameter
// Treasury module handles all supported currencies by parameter, and
// this parameter should be modified via governance proposals.
// So we should check whether the parameter is changed or not per every
// endblock and apply it if parameter has changed.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	// get sequence from params and state, compare
	paramseq := k.GetParamSequence(ctx)
	seq := k.GetSequence(ctx)
	// if there is no change, pass
	if paramseq == seq {
		return
	}
	// different sequence means there was a change at parameter

	// clear existing currencies first
	k.ClearCurrencies(ctx)
	// reset state sequence number with parameter sequence
	k.SetSequence(ctx, types.NewSequence(paramseq))

	// update parameter's currencies to state
	params := k.GetParams(ctx)
	var denoms []string
	for _, c := range params.CurrencyList {
		k.SetCurrency(ctx, c)
		denoms = append(denoms, c.Denom)
	}
	k.SetCurrencies(ctx, types.NewCurrencies(denoms))
}
