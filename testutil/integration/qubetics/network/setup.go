package network

import (
	"fmt"
	"time"

	"github.com/qubetics/qubetics-blockchain/v2/app"
	"github.com/qubetics/qubetics-blockchain/v2/encoding"

	"cosmossdk.io/simapp"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/testutil/mock"
	"github.com/cosmos/gogoproto/proto"

	sdkmath "cosmossdk.io/math"
	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmtypes "github.com/cometbft/cometbft/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	simutils "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/crypto"
	qubeticstypes "github.com/qubetics/qubetics-blockchain/v2/types"
	qubeticsutil "github.com/qubetics/qubetics-blockchain/v2/utils"
	erc20types "github.com/qubetics/qubetics-blockchain/v2/x/erc20/types"
	evmtypes "github.com/qubetics/qubetics-blockchain/v2/x/evm/types"
)

// createValidatorSetAndSigners creates validator set with the amount of validators specified
// with the default power of 1.
func createValidatorSetAndSigners(numberOfValidators int) (*tmtypes.ValidatorSet, map[string]tmtypes.PrivValidator) {
	// Create validator set
	tmValidators := make([]*tmtypes.Validator, 0, numberOfValidators)
	signers := make(map[string]tmtypes.PrivValidator, numberOfValidators)

	for i := 0; i < numberOfValidators; i++ {
		privVal := mock.NewPV()
		pubKey, _ := privVal.GetPubKey()
		validator := tmtypes.NewValidator(pubKey, 1)
		tmValidators = append(tmValidators, validator)
		signers[pubKey.Address().String()] = privVal
	}

	return tmtypes.NewValidatorSet(tmValidators), signers
}

// createGenesisAccounts returns a slice of genesis accounts from the given
// account addresses.
func createGenesisAccounts(accounts []sdktypes.AccAddress) []authtypes.GenesisAccount {
	numberOfAccounts := len(accounts)
	genAccounts := make([]authtypes.GenesisAccount, 0, numberOfAccounts)
	emptyCodeHash := crypto.Keccak256Hash(nil).String()
	for _, acc := range accounts {
		baseAcc := authtypes.NewBaseAccount(acc, nil, 0, 0)
		ethAcc := &qubeticstypes.EthAccount{
			BaseAccount: baseAcc,
			CodeHash:    emptyCodeHash,
		}
		genAccounts = append(genAccounts, ethAcc)
	}
	return genAccounts
}

// getAccAddrsFromBalances returns a slice of genesis accounts from the
// given balances.
func getAccAddrsFromBalances(balances []banktypes.Balance) []sdktypes.AccAddress {
	numberOfBalances := len(balances)
	genAccounts := make([]sdktypes.AccAddress, 0, numberOfBalances)
	for _, balance := range balances {
		genAccounts = append(genAccounts, balance.GetAddress())
	}
	return genAccounts
}

// createBalances creates balances for the given accounts and coin
func createBalances(accounts []sdktypes.AccAddress, coin sdktypes.Coin) []banktypes.Balance {
	numberOfAccounts := len(accounts)
	fundedAccountBalances := make([]banktypes.Balance, 0, numberOfAccounts)
	for _, acc := range accounts {
		balance := banktypes.Balance{
			Address: acc.String(),
			Coins:   sdktypes.NewCoins(coin),
		}

		fundedAccountBalances = append(fundedAccountBalances, balance)
	}
	return fundedAccountBalances
}

// createQubeticsApp creates an qubetics app
func createQubeticsApp(chainID string) *app.Qubetics {
	// Create qubetics app
	db := dbm.NewMemDB()
	logger := log.NewNopLogger()
	loadLatest := true
	skipUpgradeHeights := map[int64]bool{}
	homePath := app.DefaultNodeHome
	invCheckPeriod := uint(5)
	encodingConfig := encoding.MakeConfig(app.ModuleBasics)
	appOptions := simutils.NewAppOptionsWithFlagHome(app.DefaultNodeHome)
	baseAppOptions := []func(*baseapp.BaseApp){baseapp.SetChainID(chainID)}

	return app.NewQubetics(
		logger,
		db,
		nil,
		loadLatest,
		skipUpgradeHeights,
		homePath,
		invCheckPeriod,
		encodingConfig,
		appOptions,
		baseAppOptions...,
	)
}

// createStakingValidator creates a staking validator from the given tm validator and bonded
func createStakingValidator(val *tmtypes.Validator, bondedAmt sdkmath.Int) (stakingtypes.Validator, error) {
	pk, err := cryptocodec.FromTmPubKeyInterface(val.PubKey)
	if err != nil {
		return stakingtypes.Validator{}, err
	}

	pkAny, err := codectypes.NewAnyWithValue(pk)
	if err != nil {
		return stakingtypes.Validator{}, err
	}

	commission := stakingtypes.NewCommission(sdktypes.ZeroDec(), sdktypes.ZeroDec(), sdktypes.ZeroDec())
	validator := stakingtypes.Validator{
		OperatorAddress:   sdktypes.ValAddress(val.Address).String(),
		ConsensusPubkey:   pkAny,
		Jailed:            false,
		Status:            stakingtypes.Bonded,
		Tokens:            bondedAmt,
		DelegatorShares:   sdktypes.OneDec(),
		Description:       stakingtypes.Description{},
		UnbondingHeight:   int64(0),
		UnbondingTime:     time.Unix(0, 0).UTC(),
		Commission:        commission,
		MinSelfDelegation: sdktypes.ZeroInt(),
	}
	return validator, nil
}

// createStakingValidators creates staking validators from the given tm validators and bonded
// amounts
func createStakingValidators(tmValidators []*tmtypes.Validator, bondedAmt sdkmath.Int) ([]stakingtypes.Validator, error) {
	amountOfValidators := len(tmValidators)
	stakingValidators := make([]stakingtypes.Validator, 0, amountOfValidators)
	for _, val := range tmValidators {
		validator, err := createStakingValidator(val, bondedAmt)
		if err != nil {
			return nil, err
		}
		stakingValidators = append(stakingValidators, validator)
	}
	return stakingValidators, nil
}

// createDelegations creates delegations for the given validators and account
func createDelegations(tmValidators []*tmtypes.Validator, fromAccount sdktypes.AccAddress) []stakingtypes.Delegation {
	amountOfValidators := len(tmValidators)
	delegations := make([]stakingtypes.Delegation, 0, amountOfValidators)
	for _, val := range tmValidators {
		delegation := stakingtypes.NewDelegation(fromAccount, val.Address.Bytes(), sdktypes.OneDec())
		delegations = append(delegations, delegation)
	}
	return delegations
}

// StakingCustomGenesisState defines the staking genesis state
type StakingCustomGenesisState struct {
	denom string

	validators  []stakingtypes.Validator
	delegations []stakingtypes.Delegation
}

// setDefaultStakingGenesisState sets the staking genesis state
func setDefaultStakingGenesisState(qubeticsApp *app.Qubetics, genesisState simapp.GenesisState, overwriteParams StakingCustomGenesisState) simapp.GenesisState {
	// Set staking params
	stakingParams := stakingtypes.DefaultParams()
	stakingParams.BondDenom = overwriteParams.denom

	stakingGenesis := stakingtypes.NewGenesisState(
		stakingParams,
		overwriteParams.validators,
		overwriteParams.delegations,
	)
	genesisState[stakingtypes.ModuleName] = qubeticsApp.AppCodec().MustMarshalJSON(stakingGenesis)
	return genesisState
}

func setDefaultMintGenesisState(qubeticsApp *app.Qubetics, genesisState simapp.GenesisState) simapp.GenesisState {
	minter := minttypes.DefaultInitialMinter()
	mintParams := minttypes.DefaultParams()
	mintGenesis := minttypes.NewGenesisState(minter, mintParams)
	genesisState[minttypes.ModuleName] = qubeticsApp.AppCodec().MustMarshalJSON(mintGenesis)
	return genesisState
}

type BankCustomGenesisState struct {
	totalSupply sdktypes.Coins
	balances    []banktypes.Balance
}

// setDefaultBankGenesisState sets the bank genesis state
func setDefaultBankGenesisState(
	qubeticsApp *app.Qubetics,
	genesisState simapp.GenesisState,
	overwriteParams BankCustomGenesisState,
) simapp.GenesisState {
	bankGenesis := banktypes.NewGenesisState(
		banktypes.DefaultGenesisState().Params,
		overwriteParams.balances,
		overwriteParams.totalSupply,
		[]banktypes.Metadata{},
		[]banktypes.SendEnabled{},
	)
	genesisState[banktypes.ModuleName] = qubeticsApp.AppCodec().MustMarshalJSON(bankGenesis)
	return genesisState
}

// genSetupFn is the type for the module genesis setup functions
type genSetupFn func(
	qubeticsApp *app.Qubetics,
	genesisState simapp.GenesisState,
	customGenesis interface{},
) (simapp.GenesisState, error)

// defaultGenesisParams contains the params that are needed to
// setup the default genesis for the testing setup
type defaultGenesisParams struct {
	genAccounts []authtypes.GenesisAccount
	staking     StakingCustomGenesisState
	bank        BankCustomGenesisState
}

// genStateSetter is a generic function to set module-specific genesis state
func genStateSetter[T proto.Message](moduleName string) genSetupFn {
	return func(
		qubeticsApp *app.Qubetics,
		genesisState simapp.GenesisState,
		customGenesis interface{},
	) (simapp.GenesisState, error) {
		moduleGenesis, ok := customGenesis.(T)
		if !ok {
			return nil, fmt.Errorf("invalid type %T for %s module genesis state", customGenesis, moduleName)
		}

		genesisState[moduleName] = qubeticsApp.AppCodec().MustMarshalJSON(moduleGenesis)
		return genesisState, nil
	}
}

// genesisSetupFunctions contains the available genesis setup functions
// that can be used to customize the network genesis
var genesisSetupFunctions = map[string]genSetupFn{
	authtypes.ModuleName:  genStateSetter[*authtypes.GenesisState](authtypes.ModuleName),
	evmtypes.ModuleName:   genStateSetter[*evmtypes.GenesisState](evmtypes.ModuleName),
	govtypes.ModuleName:   genStateSetter[*govtypesv1.GenesisState](govtypes.ModuleName),
	minttypes.ModuleName:  genStateSetter[*minttypes.GenesisState](minttypes.ModuleName),
	erc20types.ModuleName: genStateSetter[*erc20types.GenesisState](erc20types.ModuleName),
}

// setDefaultAuthGenesisState sets the default auth genesis state
func setDefaultAuthGenesisState(
	qubeticsApp *app.Qubetics,
	genesisState simapp.GenesisState,
	genAccs []authtypes.GenesisAccount,
) simapp.GenesisState {
	defaultAuthGen := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
	genesisState[authtypes.ModuleName] = qubeticsApp.AppCodec().MustMarshalJSON(defaultAuthGen)
	return genesisState
}

// setDefaultGovGenesisState sets the default gov genesis state
func setDefaultGovGenesisState(qubeticsApp *app.Qubetics, genesisState simapp.GenesisState) simapp.GenesisState {
	govGen := govtypesv1.DefaultGenesisState()
	updatedParams := govGen.Params
	// set 'tics' as deposit denom
	updatedParams.MinDeposit = sdktypes.NewCoins(sdktypes.NewCoin(qubeticsutil.BaseDenom, sdkmath.NewInt(1e18)))
	govGen.Params = updatedParams
	genesisState[govtypes.ModuleName] = qubeticsApp.AppCodec().MustMarshalJSON(govGen)
	return genesisState
}

func setDefaultErc20GenesisState(qubeticsApp *app.Qubetics, genesisState simapp.GenesisState) simapp.GenesisState {
	erc20Gen := erc20types.DefaultGenesisState()
	genesisState[erc20types.ModuleName] = qubeticsApp.AppCodec().MustMarshalJSON(erc20Gen)
	return genesisState
}

// defaultAuthGenesisState sets the default genesis state
// for the testing setup
func newDefaultGenesisState(qubeticsApp *app.Qubetics, params defaultGenesisParams) simapp.GenesisState {
	genesisState := app.NewDefaultGenesisState()

	genesisState = setDefaultAuthGenesisState(qubeticsApp, genesisState, params.genAccounts)
	genesisState = setDefaultStakingGenesisState(qubeticsApp, genesisState, params.staking)
	genesisState = setDefaultBankGenesisState(qubeticsApp, genesisState, params.bank)
	genesisState = setDefaultGovGenesisState(qubeticsApp, genesisState)
	genesisState = setDefaultMintGenesisState(qubeticsApp, genesisState)
	genesisState = setDefaultErc20GenesisState(qubeticsApp, genesisState)

	return genesisState
}

// customizeGenesis modifies genesis state if there're any custom genesis state
// for specific modules
func customizeGenesis(
	qubeticsApp *app.Qubetics,
	customGen CustomGenesisState,
	genesisState simapp.GenesisState,
) (simapp.GenesisState, error) {
	var err error
	for mod, modGenState := range customGen {
		if fn, found := genesisSetupFunctions[mod]; found {
			genesisState, err = fn(qubeticsApp, genesisState, modGenState)
			if err != nil {
				return genesisState, err
			}
		} else {
			panic(fmt.Sprintf("module %s not found in genesis setup functions", mod))
		}
	}
	return genesisState, err
}

// calculateTotalSupply calculates the total supply from the given balances
func calculateTotalSupply(fundedAccountsBalances []banktypes.Balance) sdktypes.Coins {
	totalSupply := sdktypes.NewCoins()
	for _, balance := range fundedAccountsBalances {
		totalSupply = totalSupply.Add(balance.Coins...)
	}
	return totalSupply
}

// addBondedModuleAccountToFundedBalances adds bonded amount to bonded pool module account and include it on funded accounts
func addBondedModuleAccountToFundedBalances(
	fundedAccountsBalances []banktypes.Balance,
	totalBonded sdktypes.Coin,
) []banktypes.Balance {
	return append(fundedAccountsBalances, banktypes.Balance{
		Address: authtypes.NewModuleAddress(stakingtypes.BondedPoolName).String(),
		Coins:   sdktypes.Coins{totalBonded},
	})
}
