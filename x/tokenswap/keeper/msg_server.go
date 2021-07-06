package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	rizon "github.com/rizon-world/rizon/types"
	"github.com/rizon-world/rizon/x/tokenswap/types"
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

// CreateTokenswap processes create tokenswap message
func (k msgServer) CreateTokenswap(goCtx context.Context, msg *types.MsgCreateTokenswapRequest) (*types.MsgCreateTokenswapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check swappable
	if !k.Keeper.Swappable(ctx) {
		return nil, types.ErrUnswappable
	}

	// check swap limit has exceeded
	swapTarget := k.Keeper.GetSwappedAmount(ctx) + msg.Amount.RoundInt64()
	if k.Keeper.Limit(ctx) < swapTarget {
		return nil, types.ErrSwapLimitExceed
	}

	// check authorized signer
	// XXX should think about signer could be multi at future...
	if k.Keeper.Signer(ctx) != msg.Signer {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorizedSigner, "%s", msg.Signer)
	}

	// check valid tx hash
	if !k.validTxHash(msg.TxHash) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidTxHash, "%s", msg.TxHash)
	}

	// check tx hash is already swapped
	if k.Keeper.AlreadySwapped(ctx, msg.TxHash) {
		return nil, sdkerrors.Wrapf(types.ErrAlreadySwapped, "%s", msg.TxHash)
	}

	// calculate the amount of coin
	newCoin := sdk.NewCoins(sdk.NewCoin(rizon.DefaultDenom, msg.Amount.RoundInt()))
	// prepare swap struct from message
	newSwap := types.NewTokenswap(msg.TxHash, msg.Receiver, msg.Signer, newCoin)

	// process swap
	err := k.Keeper.Swap(ctx, newSwap)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateTokenswap,
			sdk.NewAttribute(types.AttributeKeyTxHash, msg.TxHash),
			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
	})

	return &types.MsgCreateTokenswapResponse{}, err
}

/// utils ///

// validTxHash returns whether the tx hash is valid format or not
func (k msgServer) validTxHash(hash string) bool {
	return types.ValidTxHash(hash)
}
