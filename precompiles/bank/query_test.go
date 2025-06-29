package bank_test

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qubetics/qubetics-blockchain/v2/precompiles/bank"
	qubeticsutiltx "github.com/qubetics/qubetics-blockchain/v2/testutil/tx"
)

func (s *PrecompileTestSuite) TestBalances() {
	method := s.precompile.Methods[bank.BalancesMethod]

	testcases := []struct {
		name        string
		malleate    func() []interface{}
		expPass     bool
		errContains string
		expBalances []bank.Balance
	}{
		{
			"fail - invalid number of arguments",
			func() []interface{} {
				return []interface{}{
					"", "",
				}
			},
			false,
			"invalid number of arguments",
			nil,
		},
		{
			"fail - invalid account address",
			func() []interface{} {
				return []interface{}{
					"random text",
				}
			},
			false,
			"invalid type for account",
			nil,
		},
		{
			"pass - empty balances for new account",
			func() []interface{} {
				return []interface{}{
					qubeticsutiltx.GenerateAddress(),
				}
			},
			true,
			"",
			[]bank.Balance{},
		},
		{
			"pass - QUBETICS balance present",
			func() []interface{} {
				return []interface{}{
					s.keyring.GetAddr(0),
				}
			},
			true,
			"",
			[]bank.Balance{{
				ContractAddress: s.qubeticsAddr,
				Amount:          big.NewInt(4e18),
			}},
		},
		{
			"pass - QUBETICS and XMPL balances present",
			func() []interface{} {
				s.mintAndSendXMPLCoin(s.keyring.GetAccAddr(0), sdk.NewInt(1e18))
				return []interface{}{
					s.keyring.GetAddr(0),
				}
			},
			true,
			"",
			[]bank.Balance{{
				ContractAddress: s.qubeticsAddr,
				Amount:          big.NewInt(4e18),
			}, {
				ContractAddress: s.xmplAddr,
				Amount:          big.NewInt(1e18),
			}},
		},
	}

	for _, tc := range testcases {
		tc := tc

		s.Run(tc.name, func() {
			s.SetupTest()

			bz, err := s.precompile.Balances(
				s.network.GetContext(),
				nil,
				&method,
				tc.malleate(),
			)

			if tc.expPass {
				s.Require().NoError(err)
				var balances []bank.Balance
				err = s.precompile.UnpackIntoInterface(&balances, method.Name, bz)
				s.Require().NoError(err)
				s.Require().Equal(tc.expBalances, balances)
			} else {
				s.Require().Contains(err.Error(), tc.errContains)
			}
		})
	}
}

func (s *PrecompileTestSuite) TestTotalSupply() {
	method := s.precompile.Methods[bank.TotalSupplyMethod]

	qubeticsTotalSupply, ok := new(big.Int).SetString("11000000000000000000", 10)
	s.Require().True(ok)

	testcases := []struct {
		name      string
		malleate  func()
		expSupply []bank.Balance
	}{
		{
			"pass - QUBETICS and XMPL total supply",
			func() {
				s.mintAndSendXMPLCoin(s.keyring.GetAccAddr(0), sdk.NewInt(1e18))
			},
			[]bank.Balance{{
				ContractAddress: s.qubeticsAddr,
				Amount:          qubeticsTotalSupply,
			}, {
				ContractAddress: s.xmplAddr,
				Amount:          big.NewInt(1e18),
			}},
		},
	}

	for _, tc := range testcases {
		tc := tc

		s.Run(tc.name, func() {
			s.SetupTest()

			bz, err := s.precompile.TotalSupply(
				s.network.GetContext(),
				nil,
				&method,
				nil,
			)

			s.Require().NoError(err)
			var balances []bank.Balance
			err = s.precompile.UnpackIntoInterface(&balances, method.Name, bz)
			s.Require().NoError(err)
			s.Require().Equal(tc.expSupply, balances)
		})
	}
}

func (s *PrecompileTestSuite) TestSupplyOf() {
	method := s.precompile.Methods[bank.SupplyOfMethod]

	qubeticsTotalSupply, ok := new(big.Int).SetString("11000000000000000000", 10)
	s.Require().True(ok)

	testcases := []struct {
		name        string
		malleate    func() []interface{}
		expErr      bool
		errContains string
		expSupply   *big.Int
	}{
		{
			"fail - invalid number of arguments",
			func() []interface{} {
				return []interface{}{
					"", "", "",
				}
			},
			true,
			"invalid number of arguments",
			nil,
		},
		{
			"fail - invalid hex address",
			func() []interface{} {
				return []interface{}{
					"random text",
				}
			},
			true,
			"invalid type for erc20Address",
			nil,
		},
		{
			"pass - erc20 not registered return 0 supply",
			func() []interface{} {
				return []interface{}{
					qubeticsutiltx.GenerateAddress(),
				}
			},
			false,
			"",
			big.NewInt(0),
		},
		{
			"pass - XMPL total supply",
			func() []interface{} {
				return []interface{}{
					s.xmplAddr,
				}
			},
			false,
			"",
			big.NewInt(1e18),
		},

		{
			"pass - QUBETICS total supply",
			func() []interface{} {
				return []interface{}{
					s.qubeticsAddr,
				}
			},
			false,
			"",
			qubeticsTotalSupply,
		},
	}

	for _, tc := range testcases {
		tc := tc

		s.Run(tc.name, func() {
			s.SetupTest()

			bz, err := s.precompile.SupplyOf(
				s.network.GetContext(),
				nil,
				&method,
				tc.malleate(),
			)

			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.errContains)
			} else {
				out, err := method.Outputs.Unpack(bz)
				s.Require().NoError(err, "expected no error unpacking")
				supply, ok := out[0].(*big.Int)
				s.Require().True(ok, "expected output to be a big.Int")
				s.Require().NoError(err)
				s.Require().Equal(supply.Int64(), tc.expSupply.Int64())
			}
		})
	}
}
