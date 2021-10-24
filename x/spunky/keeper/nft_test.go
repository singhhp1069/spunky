package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/singhp1069/spunky/testutil/keeper"
	"github.com/singhp1069/spunky/x/spunky/keeper"
	"github.com/singhp1069/spunky/x/spunky/types"
	"github.com/stretchr/testify/require"
)

func createNNFT(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NFT {
	items := make([]types.NFT, n)
	for i := range items {
		items[i].Id = keeper.AppendNFT(ctx, items[i])
	}
	return items
}

func TestNFTGet(t *testing.T) {
	keeper, ctx := keepertest.SpunkyKeeper(t)
	items := createNNFT(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetNFT(ctx, item.Id)
		require.True(t, found)
		require.Equal(t, item, got)
	}
}

func TestNFTRemove(t *testing.T) {
	keeper, ctx := keepertest.SpunkyKeeper(t)
	items := createNNFT(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNFT(ctx, item.Id)
		_, found := keeper.GetNFT(ctx, item.Id)
		require.False(t, found)
	}
}

func TestNFTGetAll(t *testing.T) {
	keeper, ctx := keepertest.SpunkyKeeper(t)
	items := createNNFT(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllNFT(ctx))
}

func TestNFTCount(t *testing.T) {
	keeper, ctx := keepertest.SpunkyKeeper(t)
	items := createNNFT(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetNFTCount(ctx))
}
