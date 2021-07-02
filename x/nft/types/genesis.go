package types

// Validate performs basic validation of supply genesis data returning an
// error for any failed validation criteria.
func (gs GenesisState) Validate() error {
	for _, collection := range gs.Collections {
		if err := collection.Metadata.Validate(); err != nil {
			return err
		}
		for _, nft := range collection.NFTs {
			if err := nft.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

// DefaultGenesisState returns a default nft module genesis state.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}
