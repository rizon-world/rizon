package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BankKeeper defines the expected bank keeper
type BankKeeper interface {
	// MintCoins is used for minting new stable coins
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	// SendCoinsFromModuleToAccount used for sending minted coins to proper receiver
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error

	// BurnCoins is used for burning coins
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	// SendCoinsFromAccountToModule is used for moving coins from owner address to module account to burn
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}
