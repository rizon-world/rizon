package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMintRequest{}

// NewMsgMintRequest creates a new MsgMintRequest object
func NewMsgMintRequest(receiver string, signer string, amount sdk.Coin) *MsgMintRequest {
	return &MsgMintRequest{
		Receiver: receiver,
		Signer:   signer,
		Amount:   amount,
	}
}

// implements of the sdk.Msg interface //
// Route returns the RouterKey
func (msg MsgMintRequest) Route() string {
	return RouterKey
}

// Type returns the name of message type
func (msg MsgMintRequest) Type() string {
	return "MintRequest"
}

// GetSigners returns the AccAddress of signer
func (msg MsgMintRequest) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// GetSignBytes returns the message bytes to sign over
func (msg MsgMintRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validates the MsgMintRequest
func (msg MsgMintRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	return nil
}

var _ sdk.Msg = &MsgBurnRequest{}

// NewMsgBurnRequest creates a new MsgBurnRequest object
func NewMsgBurnRequest(signer string, amount sdk.Coin) *MsgBurnRequest {
	return &MsgBurnRequest{
		Signer: signer,
		Amount: amount,
	}
}

// implements of the sdk.Msg interface //
// Route returns the RouterKey
func (msg MsgBurnRequest) Route() string {
	return RouterKey
}

// Type returns the name of message type
func (msg MsgBurnRequest) Type() string {
	return "BurnRequest"
}

// GetSigners returns the AccAddress of signer
func (msg MsgBurnRequest) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// GetSignBytes returns the message bytes to sign over
func (msg MsgBurnRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validates the MsgBurnRequest
func (msg MsgBurnRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	return nil
}
