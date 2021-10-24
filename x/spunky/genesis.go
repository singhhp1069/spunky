package spunky

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/singhp1069/spunky/x/spunky/keeper"
	"github.com/singhp1069/spunky/x/spunky/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the scores
	for _, elem := range genState.ScoresList {
		k.SetScores(ctx, elem)
	}

	// Set scores count
	k.SetScoresCount(ctx, genState.ScoresCount)
	// Set all the nFT
	for _, elem := range genState.NFTList {
		k.SetNFT(ctx, elem)
	}

	// Set nFT count
	k.SetNFTCount(ctx, genState.NFTCount)
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.ScoresList = k.GetAllScores(ctx)
	genesis.ScoresCount = k.GetScoresCount(ctx)
	genesis.NFTList = k.GetAllNFT(ctx)
	genesis.NFTCount = k.GetNFTCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
