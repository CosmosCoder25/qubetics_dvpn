// This file contains all utility function that require the access to the unit
// test network and should only be used in unit tests.
package utils

import (
	"fmt"

	"cosmossdk.io/math"
	"github.com/qubetics/qubetics-blockchain/v2/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	"github.com/qubetics/qubetics-blockchain/v2/testutil/integration/qubetics/network"
	erc20types "github.com/qubetics/qubetics-blockchain/v2/x/erc20/types"
)

const (
	TokenToMint = 1e18
)

// RegisterQubeticsERC20Coins uses the UnitNetwork to register the qubetics token as an
// ERC20 token. The function performs all the required steps for the registration
// like registering the denom trace in the transfer keeper and minting the token
// with the bank. Returns the TokenPair or an error.
func RegisterQubeticsERC20Coins(
	network network.UnitTestNetwork,
	tokenReceiver sdk.AccAddress,
) (erc20types.TokenPair, error) {
	bondDenom := network.App.StakingKeeper.BondDenom(network.GetContext())

	coin := sdk.NewCoin(utils.BaseDenom, math.NewInt(TokenToMint))
	err := network.App.BankKeeper.MintCoins(
		network.GetContext(),
		minttypes.ModuleName,
		sdk.NewCoins(coin),
	)
	if err != nil {
		return erc20types.TokenPair{}, err
	}
	err = network.App.BankKeeper.SendCoinsFromModuleToAccount(
		network.GetContext(),
		minttypes.ModuleName,
		tokenReceiver,
		sdk.NewCoins(coin),
	)
	if err != nil {
		return erc20types.TokenPair{}, err
	}

	qubeticsMetadata, found := network.App.BankKeeper.GetDenomMetaData(network.GetContext(), utils.BaseDenom)
	if !found {
		return erc20types.TokenPair{}, fmt.Errorf("expected qubetics denom metadata")
	}

	_, err = network.App.Erc20Keeper.RegisterERC20Extension(network.GetContext(), qubeticsMetadata.Base)
	if err != nil {
		return erc20types.TokenPair{}, err
	}

	qubeticsdenomID := network.App.Erc20Keeper.GetDenomMap(network.GetContext(), bondDenom)
	tokenPair, ok := network.App.Erc20Keeper.GetTokenPair(network.GetContext(), qubeticsdenomID)
	if !ok {
		return erc20types.TokenPair{}, fmt.Errorf("expected qubetics erc20 token pair")
	}

	return tokenPair, nil
}

// RegisterIBCERC20Coins uses the UnitNetwork to register the denomTrace as an
// ERC20 token. The function performs all the required steps for the registration
// like registering the denom trace in the transfer keeper and minting the token
// with the bank. Returns the TokenPair or an error.
func RegisterIBCERC20Coins(
	network *network.UnitTestNetwork,
	tokenReceiver sdk.AccAddress,
	denomTrace transfertypes.DenomTrace,
) (erc20types.TokenPair, error) {
	ibcDenom := denomTrace.IBCDenom()
	network.App.TransferKeeper.SetDenomTrace(network.GetContext(), denomTrace)
	ibcMetadata := banktypes.Metadata{
		Name:        "Generic IBC name",
		Symbol:      "IBC",
		Description: "Generic IBC token description",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    ibcDenom,
				Exponent: 0,
				Aliases:  []string{ibcDenom},
			},
			{
				Denom:    ibcDenom,
				Exponent: 18,
			},
		},
		Display: ibcDenom,
		Base:    ibcDenom,
	}

	coin := sdk.NewCoin(ibcMetadata.Base, math.NewInt(TokenToMint))
	err := network.App.BankKeeper.MintCoins(
		network.GetContext(),
		minttypes.ModuleName,
		sdk.NewCoins(coin),
	)
	if err != nil {
		return erc20types.TokenPair{}, err
	}

	err = network.App.BankKeeper.SendCoinsFromModuleToAccount(
		network.GetContext(),
		minttypes.ModuleName,
		tokenReceiver,
		sdk.NewCoins(coin),
	)
	if err != nil {
		return erc20types.TokenPair{}, err
	}

	_, err = network.App.Erc20Keeper.RegisterERC20Extension(network.GetContext(), ibcMetadata.Base)
	if err != nil {
		return erc20types.TokenPair{}, err
	}

	ibcDenomID := network.App.Erc20Keeper.GetDenomMap(
		network.GetContext(),
		denomTrace.IBCDenom(),
	)
	tokenPair, ok := network.App.Erc20Keeper.GetTokenPair(network.GetContext(), ibcDenomID)
	if !ok {
		return erc20types.TokenPair{}, fmt.Errorf("expected %s erc20 token pair", ibcDenom)
	}

	return tokenPair, nil
}
