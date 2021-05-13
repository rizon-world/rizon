package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/rizon-world/rizon/x/treasury/types"
)

// Mint
func (k Keeper) Mint(ctx sdk.Context, recv string, amount sdk.Coin) error {
	// mint some coins to send
	err := k.bankKeeper.MintCoins(
		ctx,
		types.ModuleName,
		sdk.NewCoins(amount),
	)
	if err != nil {
		return err
	}

	// send minted coins to receiver
	receiver, err := sdk.AccAddressFromBech32(recv)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		receiver,
		sdk.NewCoins(amount),
	)
	if err != nil {
		return err
	}

	return nil
}

// Burn
func (k Keeper) Burn(ctx sdk.Context, owner string, amount sdk.Coin) error {
	// send coins to module account for burning
	sender, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		sender,
		types.ModuleName,
		sdk.NewCoins(amount),
	)
	if err != nil {
		return err
	}

	// burn coins
	err = k.bankKeeper.BurnCoins(
		ctx,
		types.ModuleName,
		sdk.NewCoins(amount),
	)
	if err != nil {
		return err
	}

	return nil
}
