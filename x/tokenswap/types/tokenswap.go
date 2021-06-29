package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	SwapLimitation = 2_800_000_000
)

// Newtokenswap creates a new swap state object
func NewTokenswap(txHash string, receiver string, signer string, amount sdk.Coins) Tokenswap {
	return Tokenswap{
		TxHash:   txHash,
		Receiver: receiver,
		Signer:   signer,
		Amount:   amount,
	}
}
