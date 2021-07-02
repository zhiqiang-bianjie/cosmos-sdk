package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft/types"
	gogotypes "github.com/gogo/protobuf/types"
)

// Issue defines a method for create a new nft type
func (k Keeper) IssueType(ctx sdk.Context,
	typ string,
	name string,
	symbol string,
	description string,
	mintRestricted bool,
	editRestricted bool,
	issuer sdk.AccAddress) error {
	if k.HasType(ctx, typ) {
		return sdkerrors.Wrap(types.ErrTypeExists, typ)
	}
	k.SetMetadata(ctx, types.Metadata{
		Type:           typ,
		Name:           name,
		Symbol:         symbol,
		Description:    description,
		MintRestricted: mintRestricted,
		EditRestricted: editRestricted,
	}, issuer)
	return nil
}

func (k Keeper) SetMetadata(ctx sdk.Context, metadata types.Metadata, issuer sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&metadata)
	store.Set(types.GetTypeKey(metadata.Type), bz)

	issuerWrap := gogotypes.StringValue{Value: issuer.String()}
	bz = k.cdc.MustMarshal(&issuerWrap)
	store.Set(types.GetTypeIssuerKey(metadata.Type), bz)
}

func (k Keeper) GetMetadata(ctx sdk.Context, typ string) (types.Metadata, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetTypeKey(typ))
	if len(bz) == 0 {
		return types.Metadata{}, false
	}

	var metadata types.Metadata
	k.cdc.MustUnmarshal(bz, &metadata)
	return metadata, true
}

func (k Keeper) GetTypeIssuer(ctx sdk.Context, typ string) string {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetTypeIssuerKey(typ))

	var issuerWrap gogotypes.StringValue
	k.cdc.MustUnmarshal(bz, &issuerWrap)
	return issuerWrap.Value
}

func (k Keeper) HasType(ctx sdk.Context, typ string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetTypeKey(typ))
}

func (k Keeper) IterateAllTypes(ctx sdk.Context, cb func(types.Metadata)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.TypeKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var metadata types.Metadata
		k.cdc.MustUnmarshal(iterator.Value(), &metadata)
		cb(metadata)
	}
}
