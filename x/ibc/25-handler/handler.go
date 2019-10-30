package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	client "github.com/cosmos/cosmos-sdk/x/ibc/02-client"
	connection "github.com/cosmos/cosmos-sdk/x/ibc/03-connection"
	channel "github.com/cosmos/cosmos-sdk/x/ibc/04-channel"
	exp04 "github.com/cosmos/cosmos-sdk/x/ibc/04-channel/exported"
	transfer "github.com/cosmos/cosmos-sdk/x/ibc/20-transfer"
	commitment "github.com/cosmos/cosmos-sdk/x/ibc/23-commitment"
	abci "github.com/tendermint/tendermint/abci/types"
)

type ClientHandler interface {
	CreateClient(
		ctx sdk.Context,
		msg client.MsgCreateClient,
	) sdk.Result

	UpdateClient(
		ctx sdk.Context,
		msg client.MsgUpdateClient,
	) sdk.Result

	SubmitMisbehaviour(
		ctx sdk.Context,
		msg client.MsgSubmitMisbehaviour,
	) sdk.Result

	QueryConsensusState(
		ctx sdk.Context,
		req abci.RequestQuery,
	) ([]byte, sdk.Error)

	QueryClientState(
		ctx sdk.Context,
		req abci.RequestQuery,
	) ([]byte, sdk.Error)
}

type ConnectionHandler interface {
	ConnOpenInit(
		ctx sdk.Context,
		msg connection.MsgConnectionOpenInit,
	) sdk.Result

	ConnOpenTry(
		ctx sdk.Context,
		msg connection.MsgConnectionOpenTry,
	) sdk.Result

	ConnOpenAck(
		ctx sdk.Context,
		msg connection.MsgConnectionOpenAck,
	) sdk.Result

	ConnOpenConfirm(
		ctx sdk.Context,
		msg connection.MsgConnectionOpenConfirm,
	) sdk.Result

	QueryConnection(
		ctx sdk.Context,
		req abci.RequestQuery,
	) ([]byte, sdk.Error)
}

type ChannelHandler interface {
	ChanOpenInit(
		ctx sdk.Context,
		msg channel.MsgChannelOpenInit,
	) sdk.Result

	ChanOpenTry(
		ctx sdk.Context,
		msg channel.MsgChannelOpenTry,
	) sdk.Result

	ChanOpenAck(
		ctx sdk.Context,
		msg channel.MsgChannelOpenAck,
	) sdk.Result

	ChanOpenConfirm(
		ctx sdk.Context,
		msg channel.MsgChannelOpenConfirm,
	) sdk.Result

	ChanCloseInit(
		ctx sdk.Context,
		msg channel.MsgChannelCloseInit,
	) sdk.Result

	ChanCloseConfirm(
		ctx sdk.Context,
		msg channel.MsgChannelCloseConfirm,
	) sdk.Result

	QueryChannel(
		ctx sdk.Context,
		req abci.RequestQuery,
	) ([]byte, sdk.Error)
}

type PacketHandler interface {
	SendPacket(
		ctx sdk.Context,
		packet exp04.PacketI,
		portCapability sdk.CapabilityKey,
	) error //sdk.Error

	RecvPacket(
		ctx sdk.Context,
		packet exp04.PacketI,
		proof commitment.ProofI,
		proofHeight uint64,
		acknowledgement []byte,
		portCapability sdk.CapabilityKey,
	) (exp04.PacketI, error) //sdk.Error

	AcknowledgePacket(
		ctx sdk.Context,
		packet exp04.PacketI,
		acknowledgement []byte,
		proof commitment.ProofI,
		proofHeight uint64,
		portCapability sdk.CapabilityKey,
	) (exp04.PacketI, error) //sdk.Error

	TimeoutOnClose(
		ctx sdk.Context,
		packet exp04.PacketI,
		proofNonMembership,
		proofClosed commitment.ProofI,
		proofHeight uint64,
		portCapability sdk.CapabilityKey,
	) (exp04.PacketI, error) //sdk.Error

	TimeoutPacket(
		ctx sdk.Context,
		packet exp04.PacketI,
		proof commitment.ProofI,
		proofHeight uint64,
		nextSequenceRecv uint64,
		portCapability sdk.CapabilityKey,
	) (exp04.PacketI, error) //sdk.Error

	CleanupPacket(
		ctx sdk.Context,
		packet exp04.PacketI,
		proof commitment.ProofI,
		proofHeight,
		nextSequenceRecv uint64,
		acknowledgement []byte,
		portCapability sdk.CapabilityKey,
	) (exp04.PacketI, error) //sdk.Error
}

type TransferHandler interface {
	Transfer(ctx sdk.Context, msg transfer.MsgTransfer) (res sdk.Result)
}
type Handler interface {
	ClientHandler
	ChannelHandler
	ConnectionHandler
	PacketHandler
	TransferHandler
}

type HManager struct {
	ClientHandler
	ChannelHandler
	ConnectionHandler
	PacketHandler
	TransferHandler
}

func NewHManager() *HManager {
	return &HManager{}
}

func (hm *HManager) WithClientHandler(client ClientHandler) *HManager {
	hm.ClientHandler = client
	return hm
}

func (hm *HManager) WithChannelHandler(channel ChannelHandler) *HManager {
	hm.ChannelHandler = channel
	return hm
}

func (hm *HManager) WithConnectionHandler(connection ConnectionHandler) *HManager {
	hm.ConnectionHandler = connection
	return hm
}

func (hm *HManager) WithPacketHandler(packet PacketHandler) *HManager {
	hm.PacketHandler = packet
	return hm
}

func (hm *HManager) WithTransferHandler(transfer TransferHandler) *HManager {
	hm.TransferHandler = transfer
	return hm
}
