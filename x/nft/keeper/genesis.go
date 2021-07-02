package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft/types"
)

// InitGenesis initializes the nft module's state from a given genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState *types.GenesisState) {
	if err := genState.Validate(); err != nil {
		panic(err)
	}
	for _, collection := range genState.Collections {
		issuer, err := sdk.AccAddressFromBech32(collection.Issuer)
		if err != nil {
			panic(err)
		}

		k.SetMetadata(ctx, collection.Metadata, issuer)
		for _, nft := range collection.NFTs {
			k.SetNFT(ctx, nft)
		}
	}
}

// ExportGenesis returns the nft module's genesis state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	var collections []*types.Collection
	k.IterateAllTypes(ctx, func(metadata types.Metadata) {
		collections = append(collections, &types.Collection{
			Metadata: metadata,
			Issuer:   k.GetTypeIssuer(ctx, metadata.Type),
			NFTs:     k.GetNFTs(ctx, metadata.Type),
		})
	})
	return &types.GenesisState{Collections: collections}
}
