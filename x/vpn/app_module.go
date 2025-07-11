package vpn

import (
	"context"
	"encoding/json"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"
	sdksimulation "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

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
	"github.com/qubetics/qubetics-blockchain/v2/x/vpn/client/cli"
	"github.com/qubetics/qubetics-blockchain/v2/x/vpn/keeper"
	"github.com/qubetics/qubetics-blockchain/v2/x/vpn/migrations"
	"github.com/qubetics/qubetics-blockchain/v2/x/vpn/services"
	"github.com/qubetics/qubetics-blockchain/v2/x/vpn/types"
	v1 "github.com/qubetics/qubetics-blockchain/v2/x/vpn/types/v1"
)

var (
	_ sdkmodule.AppModuleBasic   = AppModuleBasic{}
	_ sdkmodule.HasGenesisBasics = AppModuleBasic{}

	_ sdkmodule.AppModuleGenesis    = AppModule{}
	_ sdkmodule.AppModuleSimulation = AppModule{}
	_ sdkmodule.BeginBlockAppModule = AppModule{}
	_ sdkmodule.EndBlockAppModule   = AppModule{}
	_ sdkmodule.HasConsensusVersion = AppModule{}
	_ sdkmodule.HasInvariants       = AppModule{}
	_ sdkmodule.HasServices         = AppModule{}
)

type AppModuleBasic struct{}

func (amb AppModuleBasic) Name() string { return types.ModuleName }

func (amb AppModuleBasic) RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

func (amb AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	v1.RegisterInterfaces(registry)
}

func (amb AppModuleBasic) RegisterGRPCGatewayRoutes(ctx client.Context, mux *runtime.ServeMux) {
	_ = v1deposittypes.RegisterQueryServiceHandlerClient(context.Background(), mux, v1deposittypes.NewQueryServiceClient(ctx))
	_ = v1leasetypes.RegisterQueryServiceHandlerClient(context.Background(), mux, v1leasetypes.NewQueryServiceClient(ctx))
	_ = v1nodetypes.RegisterQueryServiceHandlerClient(context.Background(), mux, v1nodetypes.NewQueryServiceClient(ctx))
	_ = v1plantypes.RegisterQueryServiceHandlerClient(context.Background(), mux, v1plantypes.NewQueryServiceClient(ctx))
	_ = v1providertypes.RegisterQueryServiceHandlerClient(context.Background(), mux, v1providertypes.NewQueryServiceClient(ctx))
	_ = v1sessiontypes.RegisterQueryServiceHandlerClient(context.Background(), mux, v1sessiontypes.NewQueryServiceClient(ctx))
	_ = v1subscriptiontypes.RegisterQueryServiceHandlerClient(context.Background(), mux, v1subscriptiontypes.NewQueryServiceClient(ctx))

	_ = v2nodetypes.RegisterQueryServiceHandlerClient(context.Background(), mux, v2nodetypes.NewQueryServiceClient(ctx))
	_ = v2plantypes.RegisterQueryServiceHandlerClient(context.Background(), mux, v2plantypes.NewQueryServiceClient(ctx))
	_ = v2providertypes.RegisterQueryServiceHandlerClient(context.Background(), mux, v2providertypes.NewQueryServiceClient(ctx))
	_ = v2sessiontypes.RegisterQueryServiceHandlerClient(context.Background(), mux, v2sessiontypes.NewQueryServiceClient(ctx))
	_ = v2subscriptiontypes.RegisterQueryServiceHandlerClient(context.Background(), mux, v2subscriptiontypes.NewQueryServiceClient(ctx))

	_ = v3nodetypes.RegisterQueryServiceHandlerClient(context.Background(), mux, v3nodetypes.NewQueryServiceClient(ctx))
	_ = v3sessiontypes.RegisterQueryServiceHandlerClient(context.Background(), mux, v3sessiontypes.NewQueryServiceClient(ctx))
	_ = v3subscriptiontypes.RegisterQueryServiceHandlerClient(context.Background(), mux, v3subscriptiontypes.NewQueryServiceClient(ctx))
	_ = v3plantypes.RegisterQueryServiceHandlerClient(context.Background(), mux, v3plantypes.NewQueryServiceClient(ctx))
	_ = v3providertypes.RegisterQueryServiceHandlerClient(context.Background(), mux, v3providertypes.NewQueryServiceClient(ctx))
}

func (amb AppModuleBasic) GetTxCmd() *cobra.Command { return cli.GetTxCmd() }

func (amb AppModuleBasic) GetQueryCmd() *cobra.Command { return cli.GetQueryCmd() }

func (amb AppModuleBasic) DefaultGenesis(jsonCodec codec.JSONCodec) json.RawMessage {
	state := v1.DefaultGenesisState()
	return jsonCodec.MustMarshalJSON(state)
}

func (amb AppModuleBasic) ValidateGenesis(jsonCodec codec.JSONCodec, _ client.TxEncodingConfig, message json.RawMessage) error {
	var state v1.GenesisState
	if err := jsonCodec.UnmarshalJSON(message, &state); err != nil {
		return err
	}

	return state.Validate()
}

type AppModule struct {
	AppModuleBasic
	cdc     codec.Codec
	account keeper.AccountKeeper
	bank    keeper.BankKeeper
	keeper  keeper.Keeper
}

func NewAppModule(cdc codec.Codec, account keeper.AccountKeeper, bank keeper.BankKeeper, k keeper.Keeper) AppModule {
	return AppModule{
		cdc:     cdc,
		account: account,
		bank:    bank,
		keeper:  k,
	}
}

func (am AppModule) InitGenesis(ctx sdk.Context, jsonCodec codec.JSONCodec, message json.RawMessage) []abcitypes.ValidatorUpdate {
	var state v1.GenesisState
	jsonCodec.MustUnmarshalJSON(message, &state)
	am.keeper.InitGenesis(ctx, &state)

	return nil
}

func (am AppModule) ExportGenesis(ctx sdk.Context, jsonCodec codec.JSONCodec) json.RawMessage {
	state := am.keeper.ExportGenesis(ctx)
	return jsonCodec.MustMarshalJSON(state)
}

func (am AppModule) GenerateGenesisState(_ *sdkmodule.SimulationState) {}

func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

func (am AppModule) WeightedOperations(_ sdkmodule.SimulationState) []sdksimulation.WeightedOperation {
	return nil
}

func (am AppModule) BeginBlock(ctx sdk.Context, _ abcitypes.RequestBeginBlock) {
	am.keeper.BeginBlock(ctx)
}

func (am AppModule) EndBlock(_ sdk.Context, _ abcitypes.RequestEndBlock) []abcitypes.ValidatorUpdate {
	return nil
}

func (am AppModule) ConsensusVersion() uint64 { return 4 }

func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (am AppModule) RegisterServices(configurator sdkmodule.Configurator) {
	services.RegisterServices(configurator, am.keeper)

	m := migrations.NewMigrator(am.cdc, am.keeper)
	if err := configurator.RegisterMigration(types.ModuleName, 3, m.Migrate); err != nil {
		panic(err)
	}
}
