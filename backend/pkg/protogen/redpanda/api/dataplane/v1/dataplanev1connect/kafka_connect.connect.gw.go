// Code generated by protoc-gen-connect-gateway. DO NOT EDIT.
//
// Source: redpanda/api/dataplane/v1/kafka_connect.proto

package dataplanev1connect

import (
	context "context"
	fmt "fmt"

	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	connect_gateway "go.vallahaye.net/connect-gateway"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/redpanda-data/console/backend/pkg/protogen/redpanda/api/dataplane/v1"
)

// KafkaConnectServiceGatewayServer implements the gRPC server API for the KafkaConnectService
// service.
type KafkaConnectServiceGatewayServer struct {
	v1.UnimplementedKafkaConnectServiceServer
	listConnectClusters  connect_gateway.UnaryHandler[v1.ListConnectClustersRequest, v1.ListConnectClustersResponse]
	getConnectCluster    connect_gateway.UnaryHandler[v1.GetConnectClusterRequest, v1.GetConnectClusterResponse]
	listConnectors       connect_gateway.UnaryHandler[v1.ListConnectorsRequest, v1.ListConnectorsResponse]
	createConnector      connect_gateway.UnaryHandler[v1.CreateConnectorRequest, v1.CreateConnectorResponse]
	restartConnector     connect_gateway.UnaryHandler[v1.RestartConnectorRequest, emptypb.Empty]
	getConnector         connect_gateway.UnaryHandler[v1.GetConnectorRequest, v1.GetConnectorResponse]
	getConnectorStatus   connect_gateway.UnaryHandler[v1.GetConnectorStatusRequest, v1.GetConnectorStatusResponse]
	pauseConnector       connect_gateway.UnaryHandler[v1.PauseConnectorRequest, emptypb.Empty]
	resumeConnector      connect_gateway.UnaryHandler[v1.ResumeConnectorRequest, emptypb.Empty]
	stopConnector        connect_gateway.UnaryHandler[v1.StopConnectorRequest, emptypb.Empty]
	deleteConnector      connect_gateway.UnaryHandler[v1.DeleteConnectorRequest, emptypb.Empty]
	upsertConnector      connect_gateway.UnaryHandler[v1.UpsertConnectorRequest, v1.UpsertConnectorResponse]
	getConnectorConfig   connect_gateway.UnaryHandler[v1.GetConnectorConfigRequest, v1.GetConnectorConfigResponse]
	listConnectorTopics  connect_gateway.UnaryHandler[v1.ListConnectorTopicsRequest, v1.ListConnectorTopicsResponse]
	resetConnectorTopics connect_gateway.UnaryHandler[v1.ResetConnectorTopicsRequest, emptypb.Empty]
}

// NewKafkaConnectServiceGatewayServer constructs a Connect-Gateway gRPC server for the
// KafkaConnectService service.
func NewKafkaConnectServiceGatewayServer(svc KafkaConnectServiceHandler, opts ...connect_gateway.HandlerOption) *KafkaConnectServiceGatewayServer {
	return &KafkaConnectServiceGatewayServer{
		listConnectClusters:  connect_gateway.NewUnaryHandler(KafkaConnectServiceListConnectClustersProcedure, svc.ListConnectClusters, opts...),
		getConnectCluster:    connect_gateway.NewUnaryHandler(KafkaConnectServiceGetConnectClusterProcedure, svc.GetConnectCluster, opts...),
		listConnectors:       connect_gateway.NewUnaryHandler(KafkaConnectServiceListConnectorsProcedure, svc.ListConnectors, opts...),
		createConnector:      connect_gateway.NewUnaryHandler(KafkaConnectServiceCreateConnectorProcedure, svc.CreateConnector, opts...),
		restartConnector:     connect_gateway.NewUnaryHandler(KafkaConnectServiceRestartConnectorProcedure, svc.RestartConnector, opts...),
		getConnector:         connect_gateway.NewUnaryHandler(KafkaConnectServiceGetConnectorProcedure, svc.GetConnector, opts...),
		getConnectorStatus:   connect_gateway.NewUnaryHandler(KafkaConnectServiceGetConnectorStatusProcedure, svc.GetConnectorStatus, opts...),
		pauseConnector:       connect_gateway.NewUnaryHandler(KafkaConnectServicePauseConnectorProcedure, svc.PauseConnector, opts...),
		resumeConnector:      connect_gateway.NewUnaryHandler(KafkaConnectServiceResumeConnectorProcedure, svc.ResumeConnector, opts...),
		stopConnector:        connect_gateway.NewUnaryHandler(KafkaConnectServiceStopConnectorProcedure, svc.StopConnector, opts...),
		deleteConnector:      connect_gateway.NewUnaryHandler(KafkaConnectServiceDeleteConnectorProcedure, svc.DeleteConnector, opts...),
		upsertConnector:      connect_gateway.NewUnaryHandler(KafkaConnectServiceUpsertConnectorProcedure, svc.UpsertConnector, opts...),
		getConnectorConfig:   connect_gateway.NewUnaryHandler(KafkaConnectServiceGetConnectorConfigProcedure, svc.GetConnectorConfig, opts...),
		listConnectorTopics:  connect_gateway.NewUnaryHandler(KafkaConnectServiceListConnectorTopicsProcedure, svc.ListConnectorTopics, opts...),
		resetConnectorTopics: connect_gateway.NewUnaryHandler(KafkaConnectServiceResetConnectorTopicsProcedure, svc.ResetConnectorTopics, opts...),
	}
}

func (s *KafkaConnectServiceGatewayServer) ListConnectClusters(ctx context.Context, req *v1.ListConnectClustersRequest) (*v1.ListConnectClustersResponse, error) {
	return s.listConnectClusters(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) GetConnectCluster(ctx context.Context, req *v1.GetConnectClusterRequest) (*v1.GetConnectClusterResponse, error) {
	return s.getConnectCluster(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) ListConnectors(ctx context.Context, req *v1.ListConnectorsRequest) (*v1.ListConnectorsResponse, error) {
	return s.listConnectors(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) CreateConnector(ctx context.Context, req *v1.CreateConnectorRequest) (*v1.CreateConnectorResponse, error) {
	return s.createConnector(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) RestartConnector(ctx context.Context, req *v1.RestartConnectorRequest) (*emptypb.Empty, error) {
	return s.restartConnector(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) GetConnector(ctx context.Context, req *v1.GetConnectorRequest) (*v1.GetConnectorResponse, error) {
	return s.getConnector(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) GetConnectorStatus(ctx context.Context, req *v1.GetConnectorStatusRequest) (*v1.GetConnectorStatusResponse, error) {
	return s.getConnectorStatus(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) PauseConnector(ctx context.Context, req *v1.PauseConnectorRequest) (*emptypb.Empty, error) {
	return s.pauseConnector(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) ResumeConnector(ctx context.Context, req *v1.ResumeConnectorRequest) (*emptypb.Empty, error) {
	return s.resumeConnector(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) StopConnector(ctx context.Context, req *v1.StopConnectorRequest) (*emptypb.Empty, error) {
	return s.stopConnector(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) DeleteConnector(ctx context.Context, req *v1.DeleteConnectorRequest) (*emptypb.Empty, error) {
	return s.deleteConnector(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) UpsertConnector(ctx context.Context, req *v1.UpsertConnectorRequest) (*v1.UpsertConnectorResponse, error) {
	return s.upsertConnector(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) GetConnectorConfig(ctx context.Context, req *v1.GetConnectorConfigRequest) (*v1.GetConnectorConfigResponse, error) {
	return s.getConnectorConfig(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) ListConnectorTopics(ctx context.Context, req *v1.ListConnectorTopicsRequest) (*v1.ListConnectorTopicsResponse, error) {
	return s.listConnectorTopics(ctx, req)
}

func (s *KafkaConnectServiceGatewayServer) ResetConnectorTopics(ctx context.Context, req *v1.ResetConnectorTopicsRequest) (*emptypb.Empty, error) {
	return s.resetConnectorTopics(ctx, req)
}

// RegisterKafkaConnectServiceHandlerGatewayServer registers the Connect handlers for the
// KafkaConnectService "svc" to "mux".
func RegisterKafkaConnectServiceHandlerGatewayServer(mux *runtime.ServeMux, svc KafkaConnectServiceHandler, opts ...connect_gateway.HandlerOption) {
	if err := v1.RegisterKafkaConnectServiceHandlerServer(context.TODO(), mux, NewKafkaConnectServiceGatewayServer(svc, opts...)); err != nil {
		panic(fmt.Errorf("connect-gateway: %w", err))
	}
}
