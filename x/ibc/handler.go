package ibc

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	client "github.com/cosmos/cosmos-sdk/x/ibc/02-client"
	connection "github.com/cosmos/cosmos-sdk/x/ibc/03-connection"
	transfer "github.com/cosmos/cosmos-sdk/x/ibc/20-transfer"
)

// NewHandler defines the IBC handler
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// IBC client msgs
		case client.MsgCreateClient:
			return k.CreateClient(ctx, msg)

		case client.MsgUpdateClient:
			return k.UpdateClient(ctx, msg)

		case client.MsgSubmitMisbehaviour:
			return k.SubmitMisbehaviour(ctx, msg)

		// IBC connection  msgs
		case connection.MsgConnectionOpenInit:
			return k.ConnOpenInit(ctx, msg)

		case connection.MsgConnectionOpenTry:
			return k.ConnOpenTry(ctx, msg)

		case connection.MsgConnectionOpenAck:
			return k.ConnOpenAck(ctx, msg)

		case connection.MsgConnectionOpenConfirm:
			return k.ConnOpenConfirm(ctx, msg)

			// // IBC channel msgs
			// case channel.MsgChannelOpenInit:
			// 	return channel.HandleMsgChannelOpenInit(ctx, k.ChannelKeeper, msg)

			// case channel.MsgChannelOpenTry:
			// 	return channel.HandleMsgChannelOpenTry(ctx, k.ChannelKeeper, msg)

			// case channel.MsgChannelOpenAck:
			// 	return channel.HandleMsgChannelOpenAck(ctx, k.ChannelKeeper, msg)

			// case channel.MsgChannelOpenConfirm:
			// 	return channel.HandleMsgChannelOpenConfirm(ctx, k.ChannelKeeper, msg)

			// case channel.MsgChannelCloseInit:
			// 	return channel.HandleMsgChannelCloseInit(ctx, k.ChannelKeeper, msg)

			// case channel.MsgChannelCloseConfirm:
			// 	return channel.HandleMsgChannelCloseConfirm(ctx, k.ChannelKeeper, msg)

		// IBC transfer msgs
		case transfer.MsgTransfer:
			return k.Transfer(ctx, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized IBC message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}
