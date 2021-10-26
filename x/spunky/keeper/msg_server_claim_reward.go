package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/singhp1069/spunky/x/spunky/types"
	"github.com/tendermint/tendermint/crypto"
)

func (k msgServer) ClaimReward(goCtx context.Context, msg *types.MsgClaimReward) (*types.MsgClaimRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks the scorelist
	scoresList := k.GetAllScores(ctx)

	msgId, err := strconv.ParseUint(msg.Id, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msgId))
	}
	// check the claims
	reward, found := k.GetRewards(ctx, msgId)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msgId))
	}

	var highscore = ""

	for _, score := range scoresList {
		if score.Creator == msg.Creator {
			highscore = score.GetHighscore()
		}
	}

	if highscore == reward.Milestone {
		// reward is already claim
		if reward.Spunker != "nil" {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Reward is already claimed")
		}
		// get address of the spunky reward module account
		moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

		// convert the message creator address from a string into sdk.AccAddress
		spunker, err := sdk.AccAddressFromBech32(msg.Creator)
		if err != nil {
			panic(err)
		}

		// convert tokens from string into sdk.Coins
		coins, err := sdk.ParseCoinsNormalized(reward.Reward)
		if err != nil {
			panic(err)
		}

		// send tokens from the scavenge creator to the module account
		sdkError := k.bankKeeper.SendCoins(ctx, moduleAcct, spunker, coins)
		if sdkError != nil {
			return nil, sdkError
		}
	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Incorrect claim"))
	}

	return &types.MsgClaimRewardResponse{}, nil
}
