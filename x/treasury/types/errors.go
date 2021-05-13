package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrModuleUnmintable   = sdkerrors.Register(ModuleName, 2, "treasury module is unable to mint / burn")
	ErrCurrencyUnmintable = sdkerrors.Register(ModuleName, 3, "this currency is unable to mint / burn")
	ErrUnauthorizedSigner = sdkerrors.Register(ModuleName, 4, "unauthorized signer")
)
