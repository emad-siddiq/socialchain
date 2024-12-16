package socialchain_test

import (
	"testing"

	keepertest "github.com/emad-siddiq/socialchain/testutil/keeper"
	"github.com/emad-siddiq/socialchain/testutil/nullify"
	socialchain "github.com/emad-siddiq/socialchain/x/socialchain/module"
	"github.com/emad-siddiq/socialchain/x/socialchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SocialchainKeeper(t)
	socialchain.InitGenesis(ctx, k, genesisState)
	got := socialchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
