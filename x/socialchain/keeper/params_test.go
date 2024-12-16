package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/emad-siddiq/socialchain/testutil/keeper"
	"github.com/emad-siddiq/socialchain/x/socialchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.SocialchainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
