package utils_test

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	grpchandler "github.com/qubetics/qubetics-blockchain/v2/testutil/integration/qubetics/grpc"
	testkeyring "github.com/qubetics/qubetics-blockchain/v2/testutil/integration/qubetics/keyring"
	"github.com/qubetics/qubetics-blockchain/v2/testutil/integration/qubetics/network"
	"github.com/qubetics/qubetics-blockchain/v2/testutil/integration/qubetics/utils"
	"github.com/stretchr/testify/require"
)

func TestCheckBalances(t *testing.T) {
	testDenom := "atest"
	keyring := testkeyring.New(1)
	nw := network.New(
		network.WithDenom(testDenom),
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
	)
	handler := grpchandler.NewIntegrationHandler(nw)

	testcases := []struct {
		name        string
		address     string
		expAmount   math.Int
		expPass     bool
		errContains string
	}{
		{
			name:      "pass",
			address:   keyring.GetAccAddr(0).String(),
			expAmount: network.PrefundedAccountInitialBalance,
			expPass:   true,
		},
		{
			name:        "fail - wrong amount",
			address:     keyring.GetAccAddr(0).String(),
			expAmount:   sdk.NewInt(1),
			errContains: "expected balance",
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			balances := []banktypes.Balance{{
				Address: tc.address,
				Coins: sdk.NewCoins(
					sdk.NewCoin(testDenom, tc.expAmount),
				),
			}}

			err := utils.CheckBalances(handler, balances)
			if tc.expPass {
				require.NoError(t, err, "unexpected error checking balances")
			} else {
				require.Error(t, err, "expected error checking balances")
				require.ErrorContains(t, err, tc.errContains, "expected different error checking balances")
			}
		})
	}
}
