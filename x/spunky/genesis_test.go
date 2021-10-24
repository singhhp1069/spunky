package spunky_test

import (
	"testing"

	keepertest "github.com/singhp1069/spunky/testutil/keeper"
	"github.com/singhp1069/spunky/x/spunky"
	"github.com/singhp1069/spunky/x/spunky/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		ScoresList: []types.Scores{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ScoresCount: 2,
		NFTList: []types.NFT{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		NFTCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SpunkyKeeper(t)
	spunky.InitGenesis(ctx, *k, genesisState)
	got := spunky.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.ScoresList, len(genesisState.ScoresList))
	require.Subset(t, genesisState.ScoresList, got.ScoresList)
	require.Equal(t, genesisState.ScoresCount, got.ScoresCount)
	require.Len(t, got.NFTList, len(genesisState.NFTList))
	require.Subset(t, genesisState.NFTList, got.NFTList)
	require.Equal(t, genesisState.NFTCount, got.NFTCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
