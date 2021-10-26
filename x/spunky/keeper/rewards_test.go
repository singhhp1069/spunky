package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/singhp1069/spunky/testutil/keeper"
	"github.com/singhp1069/spunky/x/spunky/keeper"
	"github.com/singhp1069/spunky/x/spunky/types"
	"github.com/stretchr/testify/require"
)

func createNRewards(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Rewards {
	items := make([]types.Rewards, n)
	for i := range items {
		items[i].Id = keeper.AppendRewards(ctx, items[i])
	}
	return items
}

func TestRewardsGet(t *testing.T) {
	keeper, ctx := keepertest.SpunkyKeeper(t)
	items := createNRewards(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetRewards(ctx, item.Id)
		require.True(t, found)
		require.Equal(t, item, got)
	}
}

func TestRewardsRemove(t *testing.T) {
	keeper, ctx := keepertest.SpunkyKeeper(t)
	items := createNRewards(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRewards(ctx, item.Id)
		_, found := keeper.GetRewards(ctx, item.Id)
		require.False(t, found)
	}
}

func TestRewardsGetAll(t *testing.T) {
	keeper, ctx := keepertest.SpunkyKeeper(t)
	items := createNRewards(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllRewards(ctx))
}

func TestRewardsCount(t *testing.T) {
	keeper, ctx := keepertest.SpunkyKeeper(t)
	items := createNRewards(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetRewardsCount(ctx))
}
