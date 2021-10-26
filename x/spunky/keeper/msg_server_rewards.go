package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/singhp1069/spunky/x/spunky/types"
	"github.com/tendermint/tendermint/crypto"
)

func (k msgServer) CreateRewards(goCtx context.Context, msg *types.MsgCreateRewards) (*types.MsgCreateRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var rewards = types.Rewards{
		Creator:     msg.Creator,
		Milestone:   msg.Milestone,
		Description: msg.Description,
		Reward:      msg.Reward,
		Spunker:     msg.Spunker,
	}

	// get address of the spunky reward module account
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	// convert the message creator address from a string into sdk.AccAddress
	spunkyReward, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	// convert tokens from string into sdk.Coins
	reward, err := sdk.ParseCoinsNormalized(rewards.Reward)
	if err != nil {
		panic(err)
	}

	// send tokens from the scavenge creator to the module account
	sdkError := k.bankKeeper.SendCoins(ctx, spunkyReward, moduleAcct, reward)
	if sdkError != nil {
		return nil, sdkError
	}

	id := k.AppendRewards(
		ctx,
		rewards,
	)

	return &types.MsgCreateRewardsResponse{
		Id: id,
	}, nil
}

func (k msgServer) DeleteRewards(goCtx context.Context, msg *types.MsgDeleteRewards) (*types.MsgDeleteRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetRewards(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// reward is already claim
	if val.Spunker != "nil" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Reward is already claimed")
	}

	// get address of the spunky reward module account
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	// convert the message creator address from a string into sdk.AccAddress
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	// convert tokens from string into sdk.Coins
	reward, err := sdk.ParseCoinsNormalized(val.Reward)
	if err != nil {
		panic(err)
	}

	// send tokens from the scavenge creator to the module account
	sdkError := k.bankKeeper.SendCoins(ctx, moduleAcct, creator, reward)
	if sdkError != nil {
		return nil, sdkError
	}

	k.RemoveRewards(ctx, msg.Id)

	return &types.MsgDeleteRewardsResponse{}, nil
}
