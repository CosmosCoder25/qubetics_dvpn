package v2

import (
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"

	base "github.com/qubetics/qubetics-blockchain/v2/types"
	v1base "github.com/qubetics/qubetics-blockchain/v2/types/v1"
)

func NewQueryProviderRequest(addr base.ProvAddress) *QueryProviderRequest {
	return &QueryProviderRequest{
		Address: addr.String(),
	}
}

func NewQueryProvidersRequest(status v1base.Status, pagination *sdkquery.PageRequest) *QueryProvidersRequest {
	return &QueryProvidersRequest{
		Status:     status,
		Pagination: pagination,
	}
}
