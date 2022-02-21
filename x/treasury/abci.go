package treasury

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	rizon "github.com/rizon-world/rizon/types"
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
		if rizon.DefaultDenom == c.Denom {
			continue
		}
		k.SetCurrency(ctx, c)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.EventTypeCurrencyUpdate,
				sdk.NewAttribute(types.AttributeKeyDenom, c.Denom),
				sdk.NewAttribute(types.AttributeKeyDesc, c.Desc),
				sdk.NewAttribute(types.AttributeKeyOwner, c.Owner),
				sdk.NewAttribute(types.AttributeKeyMintable, strconv.FormatBool(c.Mintable)),
			),
		)
		denoms = append(denoms, c.Denom)
	}
	k.SetCurrencies(ctx, types.NewCurrencies(denoms))
}
