package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/singhp1069/spunky/x/spunky/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NFTAll(c context.Context, req *types.QueryAllNFTRequest) (*types.QueryAllNFTResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nFTs []types.NFT
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	nFTStore := prefix.NewStore(store, types.KeyPrefix(types.NFTKey))

	pageRes, err := query.Paginate(nFTStore, req.Pagination, func(key []byte, value []byte) error {
		var nFT types.NFT
		if err := k.cdc.Unmarshal(value, &nFT); err != nil {
			return err
		}

		nFTs = append(nFTs, nFT)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNFTResponse{NFT: nFTs, Pagination: pageRes}, nil
}

func (k Keeper) NFT(c context.Context, req *types.QueryGetNFTRequest) (*types.QueryGetNFTResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	nFT, found := k.GetNFT(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetNFTResponse{NFT: nFT}, nil
}
