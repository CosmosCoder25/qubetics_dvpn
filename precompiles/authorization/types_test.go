package authorization_test

import (
	"fmt"
	"testing"

	"cosmossdk.io/math"
	"github.com/qubetics/qubetics-blockchain/v2/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/qubetics/qubetics-blockchain/v2/precompiles/authorization"
	cmn "github.com/qubetics/qubetics-blockchain/v2/precompiles/common"
	testutiltx "github.com/qubetics/qubetics-blockchain/v2/testutil/tx"
	"github.com/stretchr/testify/require"
)

const validTypeURL = "/cosmos.bank.v1beta1.MsgSend"

func TestCheckApprovalArgs(t *testing.T) {
	addr := testutiltx.GenerateAddress()
	testcases := []struct {
		name        string
		args        []interface{}
		expErr      bool
		ErrContains string
	}{
		{
			name:        "invalid number of arguments",
			args:        []interface{}{addr, common.Address{}, math.NewInt(100), "abc"},
			expErr:      true,
			ErrContains: fmt.Sprintf(cmn.ErrInvalidNumberOfArgs, 3, 4),
		},
		{
			name:        "invalid methods defined",
			args:        []interface{}{addr, 100, nil},
			expErr:      true,
			ErrContains: fmt.Sprintf(authorization.ErrInvalidMethods, nil),
		},
		{
			name:        "no methods defined",
			args:        []interface{}{addr, 100, []string{}},
			expErr:      true,
			ErrContains: authorization.ErrEmptyMethods,
		},
		{
			name:        "empty string found in methods array",
			args:        []interface{}{addr, 100, []string{"", validTypeURL}},
			expErr:      true,
			ErrContains: fmt.Sprintf(authorization.ErrEmptyStringInMethods, []string{"", validTypeURL}),
		},
		{
			name:   "valid arguments",
			args:   []interface{}{addr, 100, []string{validTypeURL}},
			expErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			_, _, _, err := authorization.CheckApprovalArgs(tc.args, utils.BaseDenom)
			if tc.expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.ErrContains)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCheckAllowanceArgs(t *testing.T) {
	addr := testutiltx.GenerateAddress()

	testcases := []struct {
		name        string
		args        []interface{}
		expErr      bool
		ErrContains string
	}{
		{
			name:        "invalid number of arguments",
			args:        []interface{}{"", "", "", ""},
			expErr:      true,
			ErrContains: fmt.Sprintf(cmn.ErrInvalidNumberOfArgs, 3, 4),
		},
		{
			name:        "invalid grantee address",
			args:        []interface{}{common.Address{}, addr, validTypeURL},
			expErr:      true,
			ErrContains: fmt.Sprintf(authorization.ErrInvalidGrantee, common.Address{}),
		},
		{
			name:        "invalid granter address",
			args:        []interface{}{addr, common.Address{}, validTypeURL},
			expErr:      true,
			ErrContains: fmt.Sprintf(authorization.ErrInvalidGranter, common.Address{}),
		},
		{
			name:   "valid arguments",
			args:   []interface{}{addr, addr, validTypeURL},
			expErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			_, _, _, err := authorization.CheckAllowanceArgs(tc.args)
			if tc.expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.ErrContains)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
