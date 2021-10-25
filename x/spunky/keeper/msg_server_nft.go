package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/singhp1069/spunky/x/spunky/types"
)

func (k msgServer) CreateNFT(goCtx context.Context, msg *types.MsgCreateNFT) (*types.MsgCreateNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var nFT = types.NFT{
		Creator:     msg.Creator,
		Name:        msg.Name,
		Symbol:      msg.Symbol,
		Description: msg.Description,
		Uri:         msg.Uri,
		UriHash:     msg.UriHash,
		Owner:       msg.Creator,
	}

	id := k.AppendNFT(
		ctx,
		nFT,
	)

	return &types.MsgCreateNFTResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateNFT(goCtx context.Context, msg *types.MsgUpdateNFT) (*types.MsgUpdateNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var nFT = types.NFT{
		Creator:     msg.Creator,
		Id:          msg.Id,
		Name:        msg.Name,
		Symbol:      msg.Symbol,
		Description: msg.Description,
		Uri:         msg.Uri,
		UriHash:     msg.UriHash,
	}

	// Checks that the element exists
	val, found := k.GetNFT(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Cannot change the meta if you are not a owner
	if msg.Creator != val.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetNFT(ctx, nFT)

	return &types.MsgUpdateNFTResponse{}, nil
}
