package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/rizon-world/rizon/x/treasury/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// Mint
func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMintRequest) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check module mintable
	if !k.Keeper.Mintable(ctx) {
		return nil, types.ErrModuleUnmintable
	}

	// check currency mintable
	if !k.Keeper.CurrencyMintable(ctx, msg.Amount.Denom) {
		return nil, sdkerrors.Wrapf(types.ErrCurrencyUnmintable, "%s", msg.Amount.Denom)
	}

	// check authorized signer
	// XXX should think about signer could be multi at future...
	if k.Keeper.Owner(ctx, msg.Amount.Denom) != msg.Signer {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorizedSigner, "%s", msg.Signer)
	}

	// mint it
	err := k.Keeper.Mint(ctx, msg.Receiver, msg.Amount)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Amount.Denom),
			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.Amount.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
	})

	return &types.MsgMintResponse{}, err
}

// Burn
func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurnRequest) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check module mintable
	if !k.Keeper.Mintable(ctx) {
		return nil, types.ErrModuleUnmintable
	}

	// check currency mintable
	if !k.Keeper.CurrencyMintable(ctx, msg.Amount.Denom) {
		return nil, sdkerrors.Wrapf(types.ErrCurrencyUnmintable, "%s", msg.Amount.Denom)
	}

	// check authorized signer
	// XXX should think about signer could be multi at future...
	if k.Keeper.Owner(ctx, msg.Amount.Denom) != msg.Signer {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorizedSigner, "%s", msg.Signer)
	}

	// burn it
	err := k.Keeper.Burn(ctx, msg.Signer, msg.Amount)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeBurn,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Amount.Denom),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.Amount.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
	})

	return &types.MsgBurnResponse{}, err
}
