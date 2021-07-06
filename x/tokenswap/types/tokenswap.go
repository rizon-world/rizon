package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	MaxSupply      = 2_851_982_500_000_000
	SwapLimitation = 1_500_000_000_000_000 // estimation of maximum swappable amount
)

// NewTokenswap creates a new swap state object
func NewTokenswap(txHash string, receiver string, signer string, amount sdk.Coins) Tokenswap {
	return Tokenswap{
		TxHash:   txHash,
		Receiver: receiver,
		Signer:   signer,
		Amount:   amount,
	}
}

// NewSwappedAmount creates a new swapped amount
func NewSwappedAmount(amount int64) SwappedAmount {
	return SwappedAmount{
		Amount: amount,
	}
}
