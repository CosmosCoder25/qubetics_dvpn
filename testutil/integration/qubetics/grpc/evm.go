package grpc

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/qubetics/qubetics-blockchain/v2/x/evm/types"
)

// GetEvmAccount returns the EVM account for the given address.
func (gqh *IntegrationHandler) GetEvmAccount(address common.Address) (*evmtypes.QueryAccountResponse, error) {
	evmClient := gqh.network.GetEvmClient()
	return evmClient.Account(context.Background(), &evmtypes.QueryAccountRequest{
		Address: address.String(),
	})
}

// EstimateGas returns the estimated gas for the given call args.
func (gqh *IntegrationHandler) EstimateGas(args []byte, gasCap uint64) (*evmtypes.EstimateGasResponse, error) {
	evmClient := gqh.network.GetEvmClient()
	return evmClient.EstimateGas(context.Background(), &evmtypes.EthCallRequest{
		Args:   args,
		GasCap: gasCap,
	})
}

// GetEvmParams returns the EVM module params.
func (gqh *IntegrationHandler) GetEvmParams() (*evmtypes.QueryParamsResponse, error) {
	evmClient := gqh.network.GetEvmClient()
	return evmClient.Params(context.Background(), &evmtypes.QueryParamsRequest{})
}
