package v1

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	v1deposittypes "github.com/qubetics/qubetics-blockchain/v2/x/deposit/types/v1"
	v1leasetypes "github.com/qubetics/qubetics-blockchain/v2/x/lease/types/v1"
	v1nodetypes "github.com/qubetics/qubetics-blockchain/v2/x/node/types/v1"
	v2nodetypes "github.com/qubetics/qubetics-blockchain/v2/x/node/types/v2"
	v3nodetypes "github.com/qubetics/qubetics-blockchain/v2/x/node/types/v3"
	v1plantypes "github.com/qubetics/qubetics-blockchain/v2/x/plan/types/v1"
	v2plantypes "github.com/qubetics/qubetics-blockchain/v2/x/plan/types/v2"
	v3plantypes "github.com/qubetics/qubetics-blockchain/v2/x/plan/types/v3"
	v1providertypes "github.com/qubetics/qubetics-blockchain/v2/x/provider/types/v1"
	v2providertypes "github.com/qubetics/qubetics-blockchain/v2/x/provider/types/v2"
	v3providertypes "github.com/qubetics/qubetics-blockchain/v2/x/provider/types/v3"
	v1sessiontypes "github.com/qubetics/qubetics-blockchain/v2/x/session/types/v1"
	v2sessiontypes "github.com/qubetics/qubetics-blockchain/v2/x/session/types/v2"
	v3sessiontypes "github.com/qubetics/qubetics-blockchain/v2/x/session/types/v3"
	v1subscriptiontypes "github.com/qubetics/qubetics-blockchain/v2/x/subscription/types/v1"
	v2subscriptiontypes "github.com/qubetics/qubetics-blockchain/v2/x/subscription/types/v2"
	v3subscriptiontypes "github.com/qubetics/qubetics-blockchain/v2/x/subscription/types/v3"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	v1deposittypes.RegisterInterfaces(registry)
	v1leasetypes.RegisterInterfaces(registry)
	v1nodetypes.RegisterInterfaces(registry)
	v1plantypes.RegisterInterfaces(registry)
	v1providertypes.RegisterInterfaces(registry)
	v1sessiontypes.RegisterInterfaces(registry)
	v1subscriptiontypes.RegisterInterfaces(registry)

	v2nodetypes.RegisterInterfaces(registry)
	v2plantypes.RegisterInterfaces(registry)
	v2providertypes.RegisterInterfaces(registry)
	v2sessiontypes.RegisterInterfaces(registry)
	v2subscriptiontypes.RegisterInterfaces(registry)

	v3nodetypes.RegisterInterfaces(registry)
	v3plantypes.RegisterInterfaces(registry)
	v3providertypes.RegisterInterfaces(registry)
	v3sessiontypes.RegisterInterfaces(registry)
	v3subscriptiontypes.RegisterInterfaces(registry)
}
