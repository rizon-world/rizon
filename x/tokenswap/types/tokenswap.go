package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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
