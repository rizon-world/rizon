package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/rizon-world/rizon/x/treasury/types"
)

func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.params.SetParamSet(ctx, &params)
}

func (k *Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.params.GetParamSet(ctx, &params)
	return params
}

// Mintable returns the treasury module's minting/burning is enabled or not
func (k *Keeper) Mintable(ctx sdk.Context) (mintable bool) {
	k.params.Get(ctx, types.KeyMintable, &mintable)
	return
}

// GetParamSequence returns the parameter sequence
func (k *Keeper) GetParamSequence(ctx sdk.Context) (seq int64) {
	k.params.Get(ctx, types.KeyParamSequence, &seq)
	return
}
