package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BankKeeper defines the expected bank keeper
type BankKeeper interface {
	// MintCoins is used for minting new coins to swap
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	// SendCoinsFromModuleToAccount is used for sending minted coins to proper receiver
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}
