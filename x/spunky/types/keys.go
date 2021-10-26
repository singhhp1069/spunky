package types

const (
	// ModuleName defines the module name
	ModuleName = "spunky"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_spunky"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ScoresKey      = "Scores-value-"
	ScoresCountKey = "Scores-count-"
)

const (
	NFTKey      = "NFT-value-"
	NFTCountKey = "NFT-count-"
)

const (
	RewardsKey      = "Rewards-value-"
	RewardsCountKey = "Rewards-count-"
)
