package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/singhp1069/spunky/testutil/keeper"
	"github.com/singhp1069/spunky/x/spunky/keeper"
	"github.com/singhp1069/spunky/x/spunky/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SpunkyKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
