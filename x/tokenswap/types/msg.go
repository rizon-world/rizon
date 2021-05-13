package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateTokenswapRequest{}

// NewMsgCreateTokenswapRequest creates a new MsgCreateTokenswapRequest object
func NewMsgCreateTokenswapRequest(tx_hash string, receiver string, signer string, amount sdk.Dec) *MsgCreateTokenswapRequest {
	return &MsgCreateTokenswapRequest{
		TxHash:   tx_hash,
		Receiver: receiver,
		Signer:   signer,
		Amount:   &amount,
	}
}

// implements of the sdk.Msg interface //
// Route returns the RouterKey
func (msg MsgCreateTokenswapRequest) Route() string {
	return RouterKey
}

// Type returns the name of message type
func (msg MsgCreateTokenswapRequest) Type() string {
	return "CreateTokenswapRequest"
}

// GetSigners returns the AccAddress of signer
func (msg MsgCreateTokenswapRequest) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// GetSignBytes returns the message bytes to sign over
func (msg MsgCreateTokenswapRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validates the MsgCreateTokenswapRequest
func (msg MsgCreateTokenswapRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if ValidTxHash(msg.TxHash) == false {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid tx hash (%s)", msg.TxHash)
	}
	return nil
}
