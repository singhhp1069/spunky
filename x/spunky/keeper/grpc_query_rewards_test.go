package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/singhp1069/spunky/testutil/keeper"
	"github.com/singhp1069/spunky/x/spunky/types"
)

func TestRewardsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.SpunkyKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRewards(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRewardsRequest
		response *types.QueryGetRewardsResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetRewardsRequest{Id: msgs[0].Id},
			response: &types.QueryGetRewardsResponse{Rewards: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetRewardsRequest{Id: msgs[1].Id},
			response: &types.QueryGetRewardsResponse{Rewards: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetRewardsRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Rewards(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestRewardsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.SpunkyKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRewards(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRewardsRequest {
		return &types.QueryAllRewardsRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RewardsAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Rewards), step)
			require.Subset(t, msgs, resp.Rewards)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RewardsAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Rewards), step)
			require.Subset(t, msgs, resp.Rewards)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RewardsAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RewardsAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
