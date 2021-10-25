package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/singhp1069/spunky/x/spunky/types"
)

func (k msgServer) TransferNFT(goCtx context.Context, msg *types.MsgTransferNFT) (*types.MsgTransferNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetNFT(ctx, msg.Nft)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Nft))
	}

	// Checks if the owenership
	if msg.Creator != val.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// change the owner of NFT
	var nFT = types.NFT{
		Id:          val.Id,
		Name:        val.Name,
		Symbol:      val.Symbol,
		Description: val.Description,
		Uri:         val.Uri,
		UriHash:     val.UriHash,
		Creator:     val.Creator,
		Owner:       msg.Recipient,
	}

	k.SetNFT(ctx, nFT)
	return &types.MsgTransferNFTResponse{}, nil
}
