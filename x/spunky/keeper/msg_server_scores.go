package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/singhp1069/spunky/x/spunky/types"
)

func (k msgServer) CreateScores(goCtx context.Context, msg *types.MsgCreateScores) (*types.MsgCreateScoresResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var scores = types.Scores{
		Creator:   msg.Creator,
		Highscore: msg.Highscore,
	}

	id := k.AppendScores(
		ctx,
		scores,
	)

	return &types.MsgCreateScoresResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateScores(goCtx context.Context, msg *types.MsgUpdateScores) (*types.MsgUpdateScoresResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var scores = types.Scores{
		Creator:   msg.Creator,
		Id:        msg.Id,
		Highscore: msg.Highscore,
	}

	// Checks that the element exists
	val, found := k.GetScores(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}



	k.SetScores(ctx, scores)

	return &types.MsgUpdateScoresResponse{}, nil
}
