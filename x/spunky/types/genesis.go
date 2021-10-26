package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ScoresList:  []Scores{},
		NFTList:     []NFT{},
		RewardsList: []Rewards{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in scores
	scoresIdMap := make(map[uint64]bool)
	scoresCount := gs.GetScoresCount()
	for _, elem := range gs.ScoresList {
		if _, ok := scoresIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for scores")
		}
		if elem.Id >= scoresCount {
			return fmt.Errorf("scores id should be lower or equal than the last id")
		}
		scoresIdMap[elem.Id] = true
	}
	// Check for duplicated ID in nFT
	nFTIdMap := make(map[uint64]bool)
	nFTCount := gs.GetNFTCount()
	for _, elem := range gs.NFTList {
		if _, ok := nFTIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for nFT")
		}
		if elem.Id >= nFTCount {
			return fmt.Errorf("nFT id should be lower or equal than the last id")
		}
		nFTIdMap[elem.Id] = true
	}
	// Check for duplicated ID in rewards
	rewardsIdMap := make(map[uint64]bool)
	rewardsCount := gs.GetRewardsCount()
	for _, elem := range gs.RewardsList {
		if _, ok := rewardsIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for rewards")
		}
		if elem.Id >= rewardsCount {
			return fmt.Errorf("rewards id should be lower or equal than the last id")
		}
		rewardsIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
