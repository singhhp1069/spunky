package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/singhp1069/spunky/x/spunky/types"
)

// GetScoresCount get the total number of scores
func (k Keeper) GetScoresCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ScoresCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetScoresCount set the total number of scores
func (k Keeper) SetScoresCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ScoresCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendScores appends a scores in the store with a new id and update the count
func (k Keeper) AppendScores(
	ctx sdk.Context,
	scores types.Scores,
) uint64 {
	// Create the scores
	count := k.GetScoresCount(ctx)

	// Set the ID of the appended value
	scores.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScoresKey))
	appendedValue := k.cdc.MustMarshal(&scores)
	store.Set(GetScoresIDBytes(scores.Id), appendedValue)

	// Update scores count
	k.SetScoresCount(ctx, count+1)

	return count
}

// SetScores set a specific scores in the store
func (k Keeper) SetScores(ctx sdk.Context, scores types.Scores) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScoresKey))
	b := k.cdc.MustMarshal(&scores)
	store.Set(GetScoresIDBytes(scores.Id), b)
}

// GetScores returns a scores from its id
func (k Keeper) GetScores(ctx sdk.Context, id uint64) (val types.Scores, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScoresKey))
	b := store.Get(GetScoresIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveScores removes a scores from the store
func (k Keeper) RemoveScores(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScoresKey))
	store.Delete(GetScoresIDBytes(id))
}

// GetAllScores returns all scores
func (k Keeper) GetAllScores(ctx sdk.Context) (list []types.Scores) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ScoresKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Scores
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetScoresIDBytes returns the byte representation of the ID
func GetScoresIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetScoresIDFromBytes returns ID in uint64 format from a byte array
func GetScoresIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
