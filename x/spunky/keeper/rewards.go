package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/singhp1069/spunky/x/spunky/types"
)

// GetRewardsCount get the total number of rewards
func (k Keeper) GetRewardsCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RewardsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetRewardsCount set the total number of rewards
func (k Keeper) SetRewardsCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RewardsCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendRewards appends a rewards in the store with a new id and update the count
func (k Keeper) AppendRewards(
	ctx sdk.Context,
	rewards types.Rewards,
) uint64 {
	// Create the rewards
	count := k.GetRewardsCount(ctx)

	// Set the ID of the appended value
	rewards.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardsKey))
	appendedValue := k.cdc.MustMarshal(&rewards)
	store.Set(GetRewardsIDBytes(rewards.Id), appendedValue)

	// Update rewards count
	k.SetRewardsCount(ctx, count+1)

	return count
}

// SetRewards set a specific rewards in the store
func (k Keeper) SetRewards(ctx sdk.Context, rewards types.Rewards) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardsKey))
	b := k.cdc.MustMarshal(&rewards)
	store.Set(GetRewardsIDBytes(rewards.Id), b)
}

// GetRewards returns a rewards from its id
func (k Keeper) GetRewards(ctx sdk.Context, id uint64) (val types.Rewards, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardsKey))
	b := store.Get(GetRewardsIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRewards removes a rewards from the store
func (k Keeper) RemoveRewards(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardsKey))
	store.Delete(GetRewardsIDBytes(id))
}

// GetAllRewards returns all rewards
func (k Keeper) GetAllRewards(ctx sdk.Context) (list []types.Rewards) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardsKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Rewards
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetRewardsIDBytes returns the byte representation of the ID
func GetRewardsIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetRewardsIDFromBytes returns ID in uint64 format from a byte array
func GetRewardsIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
