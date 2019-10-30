package client

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
	abci "github.com/tendermint/tendermint/abci/types"
)

type Handler struct {
	k Keeper
}

func NewHandler(k Keeper) Handler {
	return Handler{k}
}

func (h Handler) CreateClient(ctx sdk.Context, msg MsgCreateClient) sdk.Result {
	clientType, err := exported.ClientTypeFromString(msg.ClientType)
	if err != nil {
		return sdk.ResultFromError(ErrInvalidClientType(DefaultCodespace, err.Error()))
	}

	// TODO: should we create an event with the new client state id ?
	_, err = h.k.CreateClient(ctx, msg.ClientID, clientType, msg.ConsensusState)
	if err != nil {
		return sdk.ResultFromError(err)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeCreateClient,
			sdk.NewAttribute(AttributeKeyClientID, msg.ClientID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer.String()),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func (h Handler) UpdateClient(ctx sdk.Context, msg MsgUpdateClient) sdk.Result {
	err := h.k.UpdateClient(ctx, msg.ClientID, msg.Header)
	if err != nil {
		return sdk.ResultFromError(err)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeUpdateClient,
			sdk.NewAttribute(AttributeKeyClientID, msg.ClientID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer.String()),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func (h Handler) SubmitMisbehaviour(ctx sdk.Context, msg MsgSubmitMisbehaviour) sdk.Result {
	err := h.k.CheckMisbehaviourAndUpdateState(ctx, msg.ClientID, msg.Evidence)
	if err != nil {
		return sdk.ResultFromError(err)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeSubmitMisbehaviour,
			sdk.NewAttribute(AttributeKeyClientID, msg.ClientID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer.String()),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func (h Handler) QueryConsensusState(ctx sdk.Context, req abci.RequestQuery) ([]byte, sdk.Error) {
	return QuerierConsensusState(ctx, req, h.k)
}

func (h Handler) QueryClientState(ctx sdk.Context, req abci.RequestQuery) ([]byte, sdk.Error) {
	return QuerierClientState(ctx, req, h.k)
}
