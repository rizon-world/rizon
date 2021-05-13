package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/rizon-world/rizon/x/tokenswap/types"
)

// Swappable returns the tokenswap is enabled or not
func (k *Keeper) Swappable(ctx sdk.Context) (swappable bool) {
	k.params.Get(ctx, types.KeySwappable, &swappable)
	return
}

// Signer returns authorized signer
func (k *Keeper) Signer(ctx sdk.Context) (signer string) {
	k.params.Get(ctx, types.KeySigner, &signer)
	return
}

func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.params.SetParamSet(ctx, &params)
}

func (k *Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.Swappable(ctx),
		k.Signer(ctx),
	)
}
