package testutil

import (
	"math/big"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/qubetics/qubetics-blockchain/v2/x/evm/core/vm"
	"github.com/stretchr/testify/require"
)

// NewPrecompileContract creates a new precompile contract and sets the gas meter.
func NewPrecompileContract(t *testing.T, ctx sdk.Context, caller common.Address, precompile vm.ContractRef, gas uint64) (*vm.Contract, sdk.Context) {
	contract := vm.NewContract(vm.AccountRef(caller), precompile, big.NewInt(0), gas)
	ctx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())
	initialGas := ctx.GasMeter().GasConsumed()
	require.Equal(t, uint64(0), initialGas)
	return contract, ctx
}
