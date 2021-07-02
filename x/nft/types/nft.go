package types

import (
	"encoding/hex"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Coin returns a coin with a amount of 1
func (n NFT) Coin() sdk.Coin {
	return sdk.NewCoin(CreateDenom(n.Type, n.ID), sdk.OneInt())
}

// Validate is responsible for verifying the legality of nft content
func (n *NFT) Validate() error {
	if err := ValidateType(n.Type); err != nil {
		return err
	}
	return ValidateID(n.ID)
}

// Validate is responsible for verifying the legality of metadata content
func (m *Metadata) Validate() error {
	if m == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty metadata")
	}
	return ValidateType(m.Type)
}

// CreateDenom return a coin denom from the type and id of nft
func CreateDenom(typ, id string) string {
	nm := fmt.Sprintf("%s-%s", typ, id)
	//nmHex := hex.EncodeToString([]byte(nm))
	denom := fmt.Sprintf("%s/%s", ModuleName, nm)
	return denom
}

// Parse the type and id from nft's coin denom
func ParseTypeAndIDFrom(coinDenom string) (typ, id string, err error) {
	prefix := fmt.Sprintf("%s/", ModuleName)
	if !strings.HasPrefix(coinDenom, prefix) {
		return typ, id, fmt.Errorf("invalid ntf denom: %s", coinDenom)
	}

	nmHex, err := hex.DecodeString(strings.TrimPrefix(coinDenom, prefix))
	if err != nil {
		return typ, id, fmt.Errorf("invalid ntf denom: %s", coinDenom)
	}

	result := strings.Split(string(nmHex), "/")
	if len(result) != 2 {
		return typ, id, fmt.Errorf("invalid ntf denom: %s", coinDenom)
	}
	return result[0], result[1], nil
}

// ValidateType is responsible for verifying the legality of nft type
// TODO should there be length and character restrictions ?
func ValidateType(typ string) error {
	if len(typ) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty type")
	}
	return nil
}

// ValidateID is responsible for verifying the legality of nft id
// TODO should there be length and character restrictions ?
func ValidateID(id string) error {
	if len(id) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty id")
	}
	return nil
}
