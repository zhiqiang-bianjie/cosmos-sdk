package transfer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/20-transfer/types"
)

type Handler struct {
	k Keeper
}

func NewHandler(k Keeper) Handler {
	return Handler{k}
}

// HandleMsgTransfer defines the sdk.Handler for MsgTransfer
func (h Handler) Transfer(ctx sdk.Context, msg MsgTransfer) (res sdk.Result) {
	err := h.k.SendTransfer(ctx, msg.SourcePort, msg.SourceChannel, msg.Amount, msg.Sender, msg.Receiver, msg.Source)
	if err != nil {
		return sdk.ResultFromError(err)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver.String()),
		))

	return sdk.Result{Events: ctx.EventManager().Events()}
}
