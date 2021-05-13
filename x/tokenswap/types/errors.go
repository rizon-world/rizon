package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrUnswappable        = sdkerrors.Register(ModuleName, 2, "tokenswap is disabled")
	ErrUnauthorizedSigner = sdkerrors.Register(ModuleName, 3, "unauthorized signer")
	ErrInvalidTxHash      = sdkerrors.Register(ModuleName, 4, "invalid tx hash")
	ErrAlreadySwapped     = sdkerrors.Register(ModuleName, 5, "already swapped tx hash")
)
