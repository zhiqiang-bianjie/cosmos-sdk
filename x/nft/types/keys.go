package types

const (
	// module name
	ModuleName = "nft"

	// routerKey
	RouterKey = ModuleName

	// StoreKey is the default store key for nft
	StoreKey = ModuleName
)

var (
	TypeKey       = []byte{0x01}
	TypeIssuerKey = []byte{0x02}
	NFTKey        = []byte{0x03}
)

// GetTypeKey returns the byte representation of the nft type key
func GetTypeKey(typ string) []byte {
	return append(TypeKey, []byte(typ)...)
}

// GetTypeIssuerKey returns the byte representation of the nft type owner key
func GetTypeIssuerKey(typ string) []byte {
	return append(TypeIssuerKey, []byte(typ)...)
}

// GetNFTKey returns the byte representation of the nft
func GetNFTKey(typ string) []byte {
	return append(NFTKey, []byte(typ)...)
}

func GetNFTIdKey(id string) []byte {
	return []byte(id)
}
