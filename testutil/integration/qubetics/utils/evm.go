package utils

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/qubetics/qubetics-blockchain/v2/contracts"
	"github.com/qubetics/qubetics-blockchain/v2/testutil/integration/qubetics/network"
	evmtypes "github.com/qubetics/qubetics-blockchain/v2/x/evm/types"
)

// GetERC20Balance returns the token balance of a given account address for
// an ERC-20 token at the given contract address.
func GetERC20Balance(nw network.Network, tokenAddress, accountAddress common.Address) (*big.Int, error) {
	input, err := contracts.ERC20MinterBurnerDecimalsContract.ABI.Pack(
		"balanceOf",
		accountAddress,
	)
	if err != nil {
		return nil, err
	}

	callData, err := json.Marshal(evmtypes.TransactionArgs{
		To:    &tokenAddress,
		Input: (*hexutil.Bytes)(&input),
	})
	if err != nil {
		return nil, err
	}

	ethRes, err := nw.GetEvmClient().EthCall(
		nw.GetContext(),
		&evmtypes.EthCallRequest{
			Args: callData,
		},
	)
	if err != nil {
		return nil, err
	}

	var balance *big.Int
	err = contracts.ERC20MinterBurnerDecimalsContract.ABI.UnpackIntoInterface(&balance, "balanceOf", ethRes.Ret)
	if err != nil {
		return nil, err
	}

	return balance, nil
}
