package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// nft module event types
const (
	AttributeValueCategory = ModuleName

	EventTypeIssue = "issue"
	EventTypeMint  = "mint"
	EventTypeEdit  = "edit"
	EventTypeSend  = "send"
	EventTypeBurn  = "burn"

	AttributeKeyIssuer    = "issuer"
	AttributeKeyMinter    = "minter"
	AttributeKeyEditor    = "editor"
	AttributeKeySender    = "sender"
	AttributeKeyReceiver  = "receiver"
	AttributeKeyDestroyer = "destroyer"
	AttributeKeyType      = "type"
	AttributeKeyID        = "id"
)

// NewTypeIssueEvent constructs a new nft type issued sdk.Event
func NewTypeIssueEvent(issuer sdk.AccAddress, typ string) sdk.Event {
	return sdk.NewEvent(
		EventTypeIssue,
		sdk.NewAttribute(AttributeKeyIssuer, issuer.String()),
		sdk.NewAttribute(AttributeKeyType, typ),
	)
}

// NewNFTMintEvent constructs a new nft minted sdk.Event
func NewNFTMintEvent(minter sdk.AccAddress, typ, id string) sdk.Event {
	return sdk.NewEvent(
		EventTypeMint,
		sdk.NewAttribute(AttributeKeyMinter, minter.String()),
		sdk.NewAttribute(AttributeKeyType, typ),
		sdk.NewAttribute(AttributeKeyID, id),
	)
}

// NewNFTEditEvent construct a new nft edited sdk.Event
func NewNFTEditEvent(editor sdk.AccAddress, typ, id string) sdk.Event {
	return sdk.NewEvent(
		EventTypeEdit,
		sdk.NewAttribute(AttributeKeyEditor, editor.String()),
		sdk.NewAttribute(AttributeKeyType, typ),
		sdk.NewAttribute(AttributeKeyID, id),
	)
}

// NewNFTSendEvent constructs a new nft sent sdk.Event
func NewNFTSendEvent(sender sdk.AccAddress, receiver sdk.AccAddress, typ, id string) sdk.Event {
	return sdk.NewEvent(
		EventTypeSend,
		sdk.NewAttribute(AttributeKeySender, sender.String()),
		sdk.NewAttribute(AttributeKeyReceiver, receiver.String()),
		sdk.NewAttribute(AttributeKeyType, typ),
		sdk.NewAttribute(AttributeKeyID, id),
	)
}

// NewNFTBurnEvent constructs a new nft burned sdk.Event
func NewNFTBurnEvent(destroyer sdk.AccAddress, typ, id string) sdk.Event {
	return sdk.NewEvent(
		EventTypeBurn,
		sdk.NewAttribute(AttributeKeyDestroyer, destroyer.String()),
		sdk.NewAttribute(AttributeKeyType, typ),
		sdk.NewAttribute(AttributeKeyID, id),
	)
}
