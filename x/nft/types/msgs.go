package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// TypeMsgSend nft message types
	TypeMsgSend = "send"
)

var (
	_ sdk.Msg = &MsgIssue{}
	_ sdk.Msg = &MsgMint{}
	_ sdk.Msg = &MsgEdit{}
	_ sdk.Msg = &MsgSend{}
	_ sdk.Msg = &MsgBurn{}
)

// ValidateBasic implements the sdk.Msg interface
func (m *MsgIssue) ValidateBasic() error {
	_, err := sdk.ValAddressFromBech32(m.Issuer)
	if err != nil {
		return err
	}
	return m.Metadata.Validate()
}

// GetSigners implements the sdk.Msg interface.
func (m *MsgIssue) GetSigners() []sdk.AccAddress {
	issuer, err := sdk.ValAddressFromBech32(m.Issuer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{issuer.Bytes()}
}

// ValidateBasic implements the sdk.Msg interface.
func (m *MsgMint) ValidateBasic() error {
	_, err := sdk.ValAddressFromBech32(m.Minter)
	if err != nil {
		panic(err)
	}
	return m.NFT.Validate()
}

// GetSigners implements the sdk.Msg interface.
func (m *MsgMint) GetSigners() []sdk.AccAddress {
	minter, err := sdk.ValAddressFromBech32(m.Minter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{minter.Bytes()}
}

// ValidateBasic implements the sdk.Msg interface.
func (m *MsgEdit) ValidateBasic() error {
	_, err := sdk.ValAddressFromBech32(m.Editor)
	if err != nil {
		panic(err)
	}
	return m.NFT.Validate()
}

// GetSigners implements the sdk.Msg interface.
func (m *MsgEdit) GetSigners() []sdk.AccAddress {
	editor, err := sdk.ValAddressFromBech32(m.Editor)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{editor.Bytes()}
}

// ValidateBasic implements the sdk.Msg interface.
func (m *MsgSend) ValidateBasic() error {
	_, err := sdk.ValAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}

	_, err = sdk.ValAddressFromBech32(m.Receiver)
	if err != nil {
		panic(err)
	}
	if len(m.Type) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty type")
	}
	if len(m.ID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty id")
	}
	return nil
}

// GetSigners implements the sdk.Msg interface.
func (m *MsgSend) GetSigners() []sdk.AccAddress {
	sender, err := sdk.ValAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender.Bytes()}
}

// ValidateBasic implements the sdk.Msg interface.
func (m *MsgBurn) ValidateBasic() error {
	_, err := sdk.ValAddressFromBech32(m.Destroyer)
	if err != nil {
		panic(err)
	}
	if err := ValidateType(m.Type); err != nil {
		return err
	}
	return ValidateType(m.ID)
}

// GetSigners implements the sdk.Msg interface.
func (m *MsgBurn) GetSigners() []sdk.AccAddress {
	destroyer, err := sdk.ValAddressFromBech32(m.Destroyer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{destroyer.Bytes()}
}
