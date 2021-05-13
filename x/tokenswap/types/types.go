package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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
