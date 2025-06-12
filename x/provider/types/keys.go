package types

import (
	sdkaddress "github.com/cosmos/cosmos-sdk/types/address"

	base "github.com/qubetics/qubetics-blockchain/v2/types"
)

const (
	ModuleName = "provider"
)

var (
	ParamsKey = []byte{0x00}

	ProviderKeyPrefix         = []byte{0x10}
	ActiveProviderKeyPrefix   = append(ProviderKeyPrefix, 0x01)
	InactiveProviderKeyPrefix = append(ProviderKeyPrefix, 0x02)
)

func ActiveProviderKey(addr base.ProvAddress) []byte {
	return append(ActiveProviderKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func InactiveProviderKey(addr base.ProvAddress) (v []byte) {
	return append(InactiveProviderKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}
