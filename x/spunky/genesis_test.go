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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SpunkyKeeper(t)
	spunky.InitGenesis(ctx, *k, genesisState)
	got := spunky.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.ScoresList, len(genesisState.ScoresList))
	require.Subset(t, genesisState.ScoresList, got.ScoresList)
	require.Equal(t, genesisState.ScoresCount, got.ScoresCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
