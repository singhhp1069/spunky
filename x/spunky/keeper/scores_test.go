package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/singhp1069/spunky/testutil/keeper"
	"github.com/singhp1069/spunky/x/spunky/keeper"
	"github.com/singhp1069/spunky/x/spunky/types"
	"github.com/stretchr/testify/require"
)

func createNScores(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Scores {
	items := make([]types.Scores, n)
	for i := range items {
		items[i].Id = keeper.AppendScores(ctx, items[i])
	}
	return items
}

func TestScoresGet(t *testing.T) {
	keeper, ctx := keepertest.SpunkyKeeper(t)
	items := createNScores(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetScores(ctx, item.Id)
		require.True(t, found)
		require.Equal(t, item, got)
	}
}

func TestScoresRemove(t *testing.T) {
	keeper, ctx := keepertest.SpunkyKeeper(t)
	items := createNScores(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveScores(ctx, item.Id)
		_, found := keeper.GetScores(ctx, item.Id)
		require.False(t, found)
	}
}

func TestScoresGetAll(t *testing.T) {
	keeper, ctx := keepertest.SpunkyKeeper(t)
	items := createNScores(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllScores(ctx))
}

func TestScoresCount(t *testing.T) {
	keeper, ctx := keepertest.SpunkyKeeper(t)
	items := createNScores(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetScoresCount(ctx))
}
