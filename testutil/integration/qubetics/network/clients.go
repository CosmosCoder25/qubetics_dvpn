package network

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/qubetics/qubetics-blockchain/v2/app"
	"github.com/qubetics/qubetics-blockchain/v2/encoding"
	erc20types "github.com/qubetics/qubetics-blockchain/v2/x/erc20/types"
	evmtypes "github.com/qubetics/qubetics-blockchain/v2/x/evm/types"
	feemarkettypes "github.com/qubetics/qubetics-blockchain/v2/x/feemarket/types"
)

func getQueryHelper(ctx sdktypes.Context) *baseapp.QueryServiceTestHelper {
	encCfg := encoding.MakeConfig(app.ModuleBasics)
	interfaceRegistry := encCfg.InterfaceRegistry
	// This is needed so that state changes are not committed in precompiles
	// simulations.
	cacheCtx, _ := ctx.CacheContext()
	return baseapp.NewQueryServerTestHelper(cacheCtx, interfaceRegistry)
}

func (n *IntegrationNetwork) GetERC20Client() erc20types.QueryClient {
	queryHelper := getQueryHelper(n.GetContext())
	erc20types.RegisterQueryServer(queryHelper, n.app.Erc20Keeper)
	return erc20types.NewQueryClient(queryHelper)
}

func (n *IntegrationNetwork) GetEvmClient() evmtypes.QueryClient {
	queryHelper := getQueryHelper(n.GetContext())
	evmtypes.RegisterQueryServer(queryHelper, n.app.EvmKeeper)
	return evmtypes.NewQueryClient(queryHelper)
}

func (n *IntegrationNetwork) GetGovClient() govtypes.QueryClient {
	queryHelper := getQueryHelper(n.GetContext())
	govtypes.RegisterQueryServer(queryHelper, n.app.GovKeeper)
	return govtypes.NewQueryClient(queryHelper)
}

func (n *IntegrationNetwork) GetBankClient() banktypes.QueryClient {
	queryHelper := getQueryHelper(n.GetContext())
	banktypes.RegisterQueryServer(queryHelper, n.app.BankKeeper)
	return banktypes.NewQueryClient(queryHelper)
}

func (n *IntegrationNetwork) GetFeeMarketClient() feemarkettypes.QueryClient {
	queryHelper := getQueryHelper(n.GetContext())
	feemarkettypes.RegisterQueryServer(queryHelper, n.app.FeeMarketKeeper)
	return feemarkettypes.NewQueryClient(queryHelper)
}

func (n *IntegrationNetwork) GetMintClient() minttypes.QueryClient {
	queryHelper := getQueryHelper(n.GetContext())
	minttypes.RegisterQueryServer(queryHelper, n.app.MintKeeper)
	return minttypes.NewQueryClient(queryHelper)
}

func (n *IntegrationNetwork) GetAuthClient() authtypes.QueryClient {
	queryHelper := getQueryHelper(n.GetContext())
	authtypes.RegisterQueryServer(queryHelper, n.app.AccountKeeper)
	return authtypes.NewQueryClient(queryHelper)
}

func (n *IntegrationNetwork) GetAuthzClient() authz.QueryClient {
	queryHelper := getQueryHelper(n.GetContext())
	authz.RegisterQueryServer(queryHelper, n.app.AuthzKeeper)
	return authz.NewQueryClient(queryHelper)
}

func (n *IntegrationNetwork) GetStakingClient() stakingtypes.QueryClient {
	queryHelper := getQueryHelper(n.GetContext())
	stakingtypes.RegisterQueryServer(queryHelper, stakingkeeper.Querier{Keeper: n.app.StakingKeeper.Keeper})
	return stakingtypes.NewQueryClient(queryHelper)
}
