package v3

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/qubetics/qubetics-blockchain/v2/x/node/keeper"
	v3 "github.com/qubetics/qubetics-blockchain/v2/x/node/types/v3"
)

var (
	_ v3.QueryServiceServer = (*queryServer)(nil)
)

type queryServer struct {
	keeper.Keeper
}

func NewQueryServiceServer(k keeper.Keeper) v3.QueryServiceServer {
	return &queryServer{k}
}

func (q *queryServer) QueryNode(c context.Context, req *v3.QueryNodeRequest) (*v3.QueryNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryNode(ctx, req)
}

func (q *queryServer) QueryNodes(c context.Context, req *v3.QueryNodesRequest) (res *v3.QueryNodesResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryNodes(ctx, req)
}

func (q *queryServer) QueryNodesForPlan(c context.Context, req *v3.QueryNodesForPlanRequest) (*v3.QueryNodesForPlanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryNodesForPlan(ctx, req)
}

func (q *queryServer) QueryParams(c context.Context, req *v3.QueryParamsRequest) (*v3.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryParams(ctx, req)
}
