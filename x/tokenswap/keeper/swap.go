package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/rizon-world/rizon/x/tokenswap/types"
)

// Swap processes a tokenswap by mint tokens and send them to proper receiver
func (k Keeper) Swap(ctx sdk.Context, swap types.Tokenswap) error {
	// mint some coins to swap
	err := k.bankKeeper.MintCoins(
		ctx,
		types.ModuleName,
		swap.Amount,
	)
	if err != nil {
		return err
	}

	// send minted coins to receiver
	receiver, err := sdk.AccAddressFromBech32(swap.Receiver)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		receiver,
		swap.Amount,
	)
	if err != nil {
		return err
	}

	// store swap item
	k.SetSwap(ctx, swap)

	return nil
}

// GetSwap returns a tokenswap item by tx hash
func (k Keeper) GetSwap(ctx sdk.Context, txHash string) (swap types.Tokenswap, err error) {
	store := k.Store(ctx)

	value := store.Get([]byte(txHash))
	err = k.cdc.UnmarshalBinaryBare(value, &swap)

	return swap, err
}

// AlreadySwapped returns true if tx hash has existed
func (k *Keeper) AlreadySwapped(ctx sdk.Context, txHash string) bool {
	store := k.Store(ctx)
	return store.Has([]byte(txHash))
}

// SetSwap stores a tokenswap item with tx hash as a key
func (k Keeper) SetSwap(ctx sdk.Context, swap types.Tokenswap) {
	store := k.Store(ctx)

	store.Set([]byte(swap.TxHash), k.cdc.MustMarshalBinaryBare(&swap))
}
