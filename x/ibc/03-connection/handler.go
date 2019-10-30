package connection

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/03-connection/keeper"
	"github.com/cosmos/cosmos-sdk/x/ibc/03-connection/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type Handler struct {
	k keeper.Keeper
}

func NewHandler(k Keeper) Handler {
	return Handler{k}
}

// ConnOpenInit defines the sdk.Handler for MsgConnectionOpenInit
func (h Handler) ConnOpenInit(ctx sdk.Context, msg types.MsgConnectionOpenInit) sdk.Result {
	err := h.k.ConnOpenInit(ctx, msg.ConnectionID, msg.ClientID, msg.Counterparty)
	if err != nil {
		return sdk.ResultFromError(err)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeConnectionOpenInit,
			sdk.NewAttribute(types.AttributeKeyConnectionID, msg.ConnectionID),
			sdk.NewAttribute(types.AttributeKeyCounterpartyClientID, msg.Counterparty.ClientID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer.String()),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}

// ConnOpenTry defines the sdk.Handler for MsgConnectionOpenTry
func (h Handler) ConnOpenTry(ctx sdk.Context, msg types.MsgConnectionOpenTry) sdk.Result {
	err := h.k.ConnOpenTry(
		ctx, msg.ConnectionID, msg.Counterparty, msg.ClientID,
		msg.CounterpartyVersions, msg.ProofInit, msg.ProofHeight, msg.ConsensusHeight)
	if err != nil {
		return sdk.ResultFromError(err)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeConnectionOpenTry,
			sdk.NewAttribute(types.AttributeKeyConnectionID, msg.ConnectionID),
			sdk.NewAttribute(types.AttributeKeyCounterpartyClientID, msg.Counterparty.ClientID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer.String()),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}

// ConnOpenAck defines the sdk.Handler for MsgConnectionOpenAck
func (h Handler) ConnOpenAck(ctx sdk.Context, msg types.MsgConnectionOpenAck) sdk.Result {
	err := h.k.ConnOpenAck(
		ctx, msg.ConnectionID, msg.Version, msg.ProofTry,
		msg.ProofHeight, msg.ConsensusHeight,
	)
	if err != nil {
		return sdk.ResultFromError(err)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeConnectionOpenAck,
			sdk.NewAttribute(types.AttributeKeyConnectionID, msg.ConnectionID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer.String()),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}

// ConnOpenConfirm defines the sdk.Handler for MsgConnectionOpenConfirm
func (h Handler) ConnOpenConfirm(ctx sdk.Context, msg types.MsgConnectionOpenConfirm) sdk.Result {
	err := h.k.ConnOpenConfirm(ctx, msg.ConnectionID, msg.ProofAck, msg.ProofHeight)
	if err != nil {
		return sdk.ResultFromError(err)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeConnectionOpenConfirm,
			sdk.NewAttribute(types.AttributeKeyConnectionID, msg.ConnectionID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer.String()),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func (h Handler) QueryConnection(ctx sdk.Context, req abci.RequestQuery) ([]byte, sdk.Error) {
	return QuerierConnection(ctx, req, h.k)
}
