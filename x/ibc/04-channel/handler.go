package channel

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/04-channel/exported"
	"github.com/cosmos/cosmos-sdk/x/ibc/04-channel/keeper"
	"github.com/cosmos/cosmos-sdk/x/ibc/04-channel/types"
	commitment "github.com/cosmos/cosmos-sdk/x/ibc/23-commitment"
	abci "github.com/tendermint/tendermint/abci/types"
)

//TODO
var portCapability sdk.CapabilityKey

type Handler struct {
	k keeper.Keeper
}

func NewHandler(k Keeper) Handler {
	return Handler{k}
}

// ChanOpenInit defines the sdk.Handler for MsgChannelOpenInit
func (h Handler) ChanOpenInit(ctx sdk.Context, msg types.MsgChannelOpenInit) sdk.Result {
	err := h.k.ChanOpenInit(
		ctx, msg.Channel.Ordering, msg.Channel.ConnectionHops, msg.PortID, msg.ChannelID,
		msg.Channel.Counterparty, msg.Channel.Version, portCapability)
	if err != nil {
		return sdk.ResultFromError(err)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeChannelOpenInit,
			sdk.NewAttribute(types.AttributeKeySenderPort, msg.PortID),
			sdk.NewAttribute(types.AttributeKeyChannelID, msg.ChannelID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer.String()),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}

// ChanOpenTry defines the sdk.Handler for MsgChannelOpenTry
func (h Handler) ChanOpenTry(ctx sdk.Context, msg types.MsgChannelOpenTry) sdk.Result {
	err := h.k.ChanOpenTry(ctx, msg.Channel.Ordering, msg.Channel.ConnectionHops, msg.PortID, msg.ChannelID,
		msg.Channel.Counterparty, msg.Channel.Version, msg.CounterpartyVersion, msg.ProofInit, msg.ProofHeight, portCapability)
	if err != nil {
		return sdk.ResultFromError(err)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeChannelOpenTry,
			sdk.NewAttribute(types.AttributeKeyChannelID, msg.ChannelID),
			sdk.NewAttribute(types.AttributeKeySenderPort, msg.PortID), // TODO: double check sender and receiver
			sdk.NewAttribute(types.AttributeKeyReceiverPort, msg.Channel.Counterparty.PortID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer.String()),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}

// ChanOpenAck defines the sdk.Handler for MsgChannelOpenAck
func (h Handler) ChanOpenAck(ctx sdk.Context, msg types.MsgChannelOpenAck) sdk.Result {
	err := h.k.ChanOpenAck(
		ctx, msg.PortID, msg.ChannelID, msg.CounterpartyVersion, msg.ProofTry, msg.ProofHeight, portCapability)
	if err != nil {
		return sdk.ResultFromError(err)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeChannelOpenAck,
			sdk.NewAttribute(types.AttributeKeySenderPort, msg.PortID),
			sdk.NewAttribute(types.AttributeKeyChannelID, msg.ChannelID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer.String()),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}

// ChanOpenConfirm defines the sdk.Handler for MsgChannelOpenConfirm
func (h Handler) ChanOpenConfirm(ctx sdk.Context, msg types.MsgChannelOpenConfirm) sdk.Result {
	err := h.k.ChanOpenConfirm(ctx, msg.PortID, msg.ChannelID, msg.ProofAck, msg.ProofHeight, portCapability)
	if err != nil {
		return sdk.ResultFromError(err)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeChannelOpenConfirm,
			sdk.NewAttribute(types.AttributeKeySenderPort, msg.PortID),
			sdk.NewAttribute(types.AttributeKeyChannelID, msg.ChannelID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer.String()),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}

// ChanCloseInit defines the sdk.Handler for MsgChannelCloseInit
func (h Handler) ChanCloseInit(ctx sdk.Context, msg types.MsgChannelCloseInit) sdk.Result {
	err := h.k.ChanCloseInit(ctx, msg.PortID, msg.ChannelID, portCapability)
	if err != nil {
		return sdk.ResultFromError(err)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeChannelCloseInit,
			sdk.NewAttribute(types.AttributeKeySenderPort, msg.PortID),
			sdk.NewAttribute(types.AttributeKeyChannelID, msg.ChannelID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer.String()),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}

// ChanCloseConfirm defines the sdk.Handler for MsgChannelCloseConfirm
func (h Handler) ChanCloseConfirm(ctx sdk.Context, msg types.MsgChannelCloseConfirm) sdk.Result {
	err := h.k.ChanCloseConfirm(ctx, msg.PortID, msg.ChannelID, msg.ProofInit, msg.ProofHeight, portCapability)
	if err != nil {
		return sdk.ResultFromError(err)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeChannelCloseConfirm,
			sdk.NewAttribute(types.AttributeKeySenderPort, msg.PortID),
			sdk.NewAttribute(types.AttributeKeyChannelID, msg.ChannelID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer.String()),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func (h Handler) QueryChannel(ctx sdk.Context, req abci.RequestQuery) ([]byte, sdk.Error) {
	return QuerierChannel(ctx, req, h.k)
}

func (h Handler) SendPacket(
	ctx sdk.Context,
	packet exported.PacketI,
	portCapability sdk.CapabilityKey,
) error {
	return h.k.SendPacket(ctx, packet, portCapability)
}

func (h Handler) RecvPacket(
	ctx sdk.Context,
	packet exported.PacketI,
	proof commitment.ProofI,
	proofHeight uint64,
	acknowledgement []byte,
	portCapability sdk.CapabilityKey,
) (exported.PacketI, error) {
	return h.k.RecvPacket(ctx, packet, proof, proofHeight, acknowledgement, portCapability)
}

func (h Handler) AcknowledgePacket(
	ctx sdk.Context,
	packet exported.PacketI,
	acknowledgement []byte,
	proof commitment.ProofI,
	proofHeight uint64,
	portCapability sdk.CapabilityKey,
) (exported.PacketI, error) {
	return h.k.AcknowledgePacket(ctx, packet, acknowledgement, proof, proofHeight, portCapability)
}

func (h Handler) TimeoutOnClose(
	ctx sdk.Context,
	packet exported.PacketI,
	proofNonMembership,
	proofClosed commitment.ProofI,
	proofHeight uint64,
	portCapability sdk.CapabilityKey,
) (exported.PacketI, error) {
	return h.k.TimeoutOnClose(ctx, packet, proofNonMembership, proofClosed, proofHeight, portCapability)
}

func (h Handler) TimeoutPacket(
	ctx sdk.Context,
	packet exported.PacketI,
	proof commitment.ProofI,
	proofHeight uint64,
	nextSequenceRecv uint64,
	portCapability sdk.CapabilityKey,
) (exported.PacketI, error) {
	return h.k.TimeoutPacket(ctx, packet, proof, proofHeight, nextSequenceRecv, portCapability)
}

func (h Handler) CleanupPacket(
	ctx sdk.Context,
	packet exported.PacketI,
	proof commitment.ProofI,
	proofHeight,
	nextSequenceRecv uint64,
	acknowledgement []byte,
	portCapability sdk.CapabilityKey,
) (exported.PacketI, error) {
	return h.k.CleanupPacket(ctx, packet, proof, proofHeight, nextSequenceRecv, acknowledgement, portCapability)
}
