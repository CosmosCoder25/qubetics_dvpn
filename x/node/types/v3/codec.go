package v3

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmsgservice "github.com/cosmos/cosmos-sdk/types/msgservice"

	sessiontypes "github.com/qubetics/qubetics-blockchain/v2/x/session/types/v3"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sessiontypes.Session)(nil),
		&Session{},
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgRegisterNodeRequest{},
		&MsgUpdateNodeDetailsRequest{},
		&MsgUpdateNodeStatusRequest{},
		&MsgStartSessionRequest{},
		&MsgUpdateParamsRequest{},
	)

	sdkmsgservice.RegisterMsgServiceDesc(registry, &_MsgService_serviceDesc)
}
