package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/singhp1069/spunky/x/spunky/types"
)

// GetNFTCount get the total number of nFT
func (k Keeper) GetNFTCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.NFTCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetNFTCount set the total number of nFT
func (k Keeper) SetNFTCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.NFTCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendNFT appends a nFT in the store with a new id and update the count
func (k Keeper) AppendNFT(
	ctx sdk.Context,
	nFT types.NFT,
) uint64 {
	// Create the nFT
	count := k.GetNFTCount(ctx)

	// Set the ID of the appended value
	nFT.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTKey))
	appendedValue := k.cdc.MustMarshal(&nFT)
	store.Set(GetNFTIDBytes(nFT.Id), appendedValue)

	// Update nFT count
	k.SetNFTCount(ctx, count+1)

	return count
}

// SetNFT set a specific nFT in the store
func (k Keeper) SetNFT(ctx sdk.Context, nFT types.NFT) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTKey))
	b := k.cdc.MustMarshal(&nFT)
	store.Set(GetNFTIDBytes(nFT.Id), b)
}

// GetNFT returns a nFT from its id
func (k Keeper) GetNFT(ctx sdk.Context, id uint64) (val types.NFT, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTKey))
	b := store.Get(GetNFTIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNFT removes a nFT from the store
func (k Keeper) RemoveNFT(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTKey))
	store.Delete(GetNFTIDBytes(id))
}

// GetAllNFT returns all nFT
func (k Keeper) GetAllNFT(ctx sdk.Context) (list []types.NFT) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NFT
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetNFTIDBytes returns the byte representation of the ID
func GetNFTIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetNFTIDFromBytes returns ID in uint64 format from a byte array
func GetNFTIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
