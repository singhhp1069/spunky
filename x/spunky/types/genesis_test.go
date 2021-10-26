package types_test

import (
	"testing"

	"github.com/singhp1069/spunky/x/spunky/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
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
				RewardsList: []types.Rewards{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				RewardsCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated scores",
			genState: &types.GenesisState{
				ScoresList: []types.Scores{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid scores count",
			genState: &types.GenesisState{
				ScoresList: []types.Scores{
					{
						Id: 1,
					},
				},
				ScoresCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated nFT",
			genState: &types.GenesisState{
				NFTList: []types.NFT{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid nFT count",
			genState: &types.GenesisState{
				NFTList: []types.NFT{
					{
						Id: 1,
					},
				},
				NFTCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated rewards",
			genState: &types.GenesisState{
				RewardsList: []types.Rewards{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid rewards count",
			genState: &types.GenesisState{
				RewardsList: []types.Rewards{
					{
						Id: 1,
					},
				},
				RewardsCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
