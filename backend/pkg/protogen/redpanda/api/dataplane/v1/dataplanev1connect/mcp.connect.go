// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: redpanda/api/dataplane/v1/mcp.proto

package dataplanev1connect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect "connectrpc.com/connect"

	v1 "github.com/redpanda-data/console/backend/pkg/protogen/redpanda/api/dataplane/v1"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// MCPServerServiceName is the fully-qualified name of the MCPServerService service.
	MCPServerServiceName = "redpanda.api.dataplane.v1.MCPServerService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MCPServerServiceCreateMCPServerProcedure is the fully-qualified name of the MCPServerService's
	// CreateMCPServer RPC.
	MCPServerServiceCreateMCPServerProcedure = "/redpanda.api.dataplane.v1.MCPServerService/CreateMCPServer"
	// MCPServerServiceGetMCPServerProcedure is the fully-qualified name of the MCPServerService's
	// GetMCPServer RPC.
	MCPServerServiceGetMCPServerProcedure = "/redpanda.api.dataplane.v1.MCPServerService/GetMCPServer"
	// MCPServerServiceListMCPServersProcedure is the fully-qualified name of the MCPServerService's
	// ListMCPServers RPC.
	MCPServerServiceListMCPServersProcedure = "/redpanda.api.dataplane.v1.MCPServerService/ListMCPServers"
	// MCPServerServiceUpdateMCPServerProcedure is the fully-qualified name of the MCPServerService's
	// UpdateMCPServer RPC.
	MCPServerServiceUpdateMCPServerProcedure = "/redpanda.api.dataplane.v1.MCPServerService/UpdateMCPServer"
	// MCPServerServiceDeleteMCPServerProcedure is the fully-qualified name of the MCPServerService's
	// DeleteMCPServer RPC.
	MCPServerServiceDeleteMCPServerProcedure = "/redpanda.api.dataplane.v1.MCPServerService/DeleteMCPServer"
	// MCPServerServiceStopMCPServerProcedure is the fully-qualified name of the MCPServerService's
	// StopMCPServer RPC.
	MCPServerServiceStopMCPServerProcedure = "/redpanda.api.dataplane.v1.MCPServerService/StopMCPServer"
	// MCPServerServiceStartMCPServerProcedure is the fully-qualified name of the MCPServerService's
	// StartMCPServer RPC.
	MCPServerServiceStartMCPServerProcedure = "/redpanda.api.dataplane.v1.MCPServerService/StartMCPServer"
	// MCPServerServiceGetMCPServerServiceConfigSchemaProcedure is the fully-qualified name of the
	// MCPServerService's GetMCPServerServiceConfigSchema RPC.
	MCPServerServiceGetMCPServerServiceConfigSchemaProcedure = "/redpanda.api.dataplane.v1.MCPServerService/GetMCPServerServiceConfigSchema"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	mCPServerServiceServiceDescriptor                               = v1.File_redpanda_api_dataplane_v1_mcp_proto.Services().ByName("MCPServerService")
	mCPServerServiceCreateMCPServerMethodDescriptor                 = mCPServerServiceServiceDescriptor.Methods().ByName("CreateMCPServer")
	mCPServerServiceGetMCPServerMethodDescriptor                    = mCPServerServiceServiceDescriptor.Methods().ByName("GetMCPServer")
	mCPServerServiceListMCPServersMethodDescriptor                  = mCPServerServiceServiceDescriptor.Methods().ByName("ListMCPServers")
	mCPServerServiceUpdateMCPServerMethodDescriptor                 = mCPServerServiceServiceDescriptor.Methods().ByName("UpdateMCPServer")
	mCPServerServiceDeleteMCPServerMethodDescriptor                 = mCPServerServiceServiceDescriptor.Methods().ByName("DeleteMCPServer")
	mCPServerServiceStopMCPServerMethodDescriptor                   = mCPServerServiceServiceDescriptor.Methods().ByName("StopMCPServer")
	mCPServerServiceStartMCPServerMethodDescriptor                  = mCPServerServiceServiceDescriptor.Methods().ByName("StartMCPServer")
	mCPServerServiceGetMCPServerServiceConfigSchemaMethodDescriptor = mCPServerServiceServiceDescriptor.Methods().ByName("GetMCPServerServiceConfigSchema")
)

// MCPServerServiceClient is a client for the redpanda.api.dataplane.v1.MCPServerService service.
type MCPServerServiceClient interface {
	// CreateMCPServer creates a Redpanda Connect MCP Server in the Redpanda cluster.
	CreateMCPServer(context.Context, *connect.Request[v1.CreateMCPServerRequest]) (*connect.Response[v1.CreateMCPServerResponse], error)
	// GetMCPServer gets a specific Redpanda Connect MCP Server.
	GetMCPServer(context.Context, *connect.Request[v1.GetMCPServerRequest]) (*connect.Response[v1.GetMCPServerResponse], error)
	// ListMCPServers implements the list mcp_servers method which lists the MCP servers
	// in the Redpanda cluster.
	ListMCPServers(context.Context, *connect.Request[v1.ListMCPServersRequest]) (*connect.Response[v1.ListMCPServersResponse], error)
	// Update MCPServer updates a specific Redpanda Connect MCP server configuration.
	UpdateMCPServer(context.Context, *connect.Request[v1.UpdateMCPServerRequest]) (*connect.Response[v1.UpdateMCPServerResponse], error)
	// DeleteMCPServer deletes a specific Redpanda Connect MCP server.
	DeleteMCPServer(context.Context, *connect.Request[v1.DeleteMCPServerRequest]) (*connect.Response[v1.DeleteMCPServerResponse], error)
	// StopMCPServer stops a specific Redpanda Connect MCP server.
	StopMCPServer(context.Context, *connect.Request[v1.StopMCPServerRequest]) (*connect.Response[v1.StopMCPServerResponse], error)
	// StartMCPServer starts a specific Redpanda Connect MCP server that has been previously stopped.
	StartMCPServer(context.Context, *connect.Request[v1.StartMCPServerRequest]) (*connect.Response[v1.StartMCPServerResponse], error)
	// The configuration schema includes available components and processors in this Redpanda Connect MCP Server instance.
	GetMCPServerServiceConfigSchema(context.Context, *connect.Request[v1.GetMCPServerServiceConfigSchemaRequest]) (*connect.Response[v1.GetMCPServerServiceConfigSchemaResponse], error)
}

// NewMCPServerServiceClient constructs a client for the redpanda.api.dataplane.v1.MCPServerService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMCPServerServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) MCPServerServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &mCPServerServiceClient{
		createMCPServer: connect.NewClient[v1.CreateMCPServerRequest, v1.CreateMCPServerResponse](
			httpClient,
			baseURL+MCPServerServiceCreateMCPServerProcedure,
			connect.WithSchema(mCPServerServiceCreateMCPServerMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getMCPServer: connect.NewClient[v1.GetMCPServerRequest, v1.GetMCPServerResponse](
			httpClient,
			baseURL+MCPServerServiceGetMCPServerProcedure,
			connect.WithSchema(mCPServerServiceGetMCPServerMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listMCPServers: connect.NewClient[v1.ListMCPServersRequest, v1.ListMCPServersResponse](
			httpClient,
			baseURL+MCPServerServiceListMCPServersProcedure,
			connect.WithSchema(mCPServerServiceListMCPServersMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateMCPServer: connect.NewClient[v1.UpdateMCPServerRequest, v1.UpdateMCPServerResponse](
			httpClient,
			baseURL+MCPServerServiceUpdateMCPServerProcedure,
			connect.WithSchema(mCPServerServiceUpdateMCPServerMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteMCPServer: connect.NewClient[v1.DeleteMCPServerRequest, v1.DeleteMCPServerResponse](
			httpClient,
			baseURL+MCPServerServiceDeleteMCPServerProcedure,
			connect.WithSchema(mCPServerServiceDeleteMCPServerMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		stopMCPServer: connect.NewClient[v1.StopMCPServerRequest, v1.StopMCPServerResponse](
			httpClient,
			baseURL+MCPServerServiceStopMCPServerProcedure,
			connect.WithSchema(mCPServerServiceStopMCPServerMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		startMCPServer: connect.NewClient[v1.StartMCPServerRequest, v1.StartMCPServerResponse](
			httpClient,
			baseURL+MCPServerServiceStartMCPServerProcedure,
			connect.WithSchema(mCPServerServiceStartMCPServerMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getMCPServerServiceConfigSchema: connect.NewClient[v1.GetMCPServerServiceConfigSchemaRequest, v1.GetMCPServerServiceConfigSchemaResponse](
			httpClient,
			baseURL+MCPServerServiceGetMCPServerServiceConfigSchemaProcedure,
			connect.WithSchema(mCPServerServiceGetMCPServerServiceConfigSchemaMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// mCPServerServiceClient implements MCPServerServiceClient.
type mCPServerServiceClient struct {
	createMCPServer                 *connect.Client[v1.CreateMCPServerRequest, v1.CreateMCPServerResponse]
	getMCPServer                    *connect.Client[v1.GetMCPServerRequest, v1.GetMCPServerResponse]
	listMCPServers                  *connect.Client[v1.ListMCPServersRequest, v1.ListMCPServersResponse]
	updateMCPServer                 *connect.Client[v1.UpdateMCPServerRequest, v1.UpdateMCPServerResponse]
	deleteMCPServer                 *connect.Client[v1.DeleteMCPServerRequest, v1.DeleteMCPServerResponse]
	stopMCPServer                   *connect.Client[v1.StopMCPServerRequest, v1.StopMCPServerResponse]
	startMCPServer                  *connect.Client[v1.StartMCPServerRequest, v1.StartMCPServerResponse]
	getMCPServerServiceConfigSchema *connect.Client[v1.GetMCPServerServiceConfigSchemaRequest, v1.GetMCPServerServiceConfigSchemaResponse]
}

// CreateMCPServer calls redpanda.api.dataplane.v1.MCPServerService.CreateMCPServer.
func (c *mCPServerServiceClient) CreateMCPServer(ctx context.Context, req *connect.Request[v1.CreateMCPServerRequest]) (*connect.Response[v1.CreateMCPServerResponse], error) {
	return c.createMCPServer.CallUnary(ctx, req)
}

// GetMCPServer calls redpanda.api.dataplane.v1.MCPServerService.GetMCPServer.
func (c *mCPServerServiceClient) GetMCPServer(ctx context.Context, req *connect.Request[v1.GetMCPServerRequest]) (*connect.Response[v1.GetMCPServerResponse], error) {
	return c.getMCPServer.CallUnary(ctx, req)
}

// ListMCPServers calls redpanda.api.dataplane.v1.MCPServerService.ListMCPServers.
func (c *mCPServerServiceClient) ListMCPServers(ctx context.Context, req *connect.Request[v1.ListMCPServersRequest]) (*connect.Response[v1.ListMCPServersResponse], error) {
	return c.listMCPServers.CallUnary(ctx, req)
}

// UpdateMCPServer calls redpanda.api.dataplane.v1.MCPServerService.UpdateMCPServer.
func (c *mCPServerServiceClient) UpdateMCPServer(ctx context.Context, req *connect.Request[v1.UpdateMCPServerRequest]) (*connect.Response[v1.UpdateMCPServerResponse], error) {
	return c.updateMCPServer.CallUnary(ctx, req)
}

// DeleteMCPServer calls redpanda.api.dataplane.v1.MCPServerService.DeleteMCPServer.
func (c *mCPServerServiceClient) DeleteMCPServer(ctx context.Context, req *connect.Request[v1.DeleteMCPServerRequest]) (*connect.Response[v1.DeleteMCPServerResponse], error) {
	return c.deleteMCPServer.CallUnary(ctx, req)
}

// StopMCPServer calls redpanda.api.dataplane.v1.MCPServerService.StopMCPServer.
func (c *mCPServerServiceClient) StopMCPServer(ctx context.Context, req *connect.Request[v1.StopMCPServerRequest]) (*connect.Response[v1.StopMCPServerResponse], error) {
	return c.stopMCPServer.CallUnary(ctx, req)
}

// StartMCPServer calls redpanda.api.dataplane.v1.MCPServerService.StartMCPServer.
func (c *mCPServerServiceClient) StartMCPServer(ctx context.Context, req *connect.Request[v1.StartMCPServerRequest]) (*connect.Response[v1.StartMCPServerResponse], error) {
	return c.startMCPServer.CallUnary(ctx, req)
}

// GetMCPServerServiceConfigSchema calls
// redpanda.api.dataplane.v1.MCPServerService.GetMCPServerServiceConfigSchema.
func (c *mCPServerServiceClient) GetMCPServerServiceConfigSchema(ctx context.Context, req *connect.Request[v1.GetMCPServerServiceConfigSchemaRequest]) (*connect.Response[v1.GetMCPServerServiceConfigSchemaResponse], error) {
	return c.getMCPServerServiceConfigSchema.CallUnary(ctx, req)
}

// MCPServerServiceHandler is an implementation of the redpanda.api.dataplane.v1.MCPServerService
// service.
type MCPServerServiceHandler interface {
	// CreateMCPServer creates a Redpanda Connect MCP Server in the Redpanda cluster.
	CreateMCPServer(context.Context, *connect.Request[v1.CreateMCPServerRequest]) (*connect.Response[v1.CreateMCPServerResponse], error)
	// GetMCPServer gets a specific Redpanda Connect MCP Server.
	GetMCPServer(context.Context, *connect.Request[v1.GetMCPServerRequest]) (*connect.Response[v1.GetMCPServerResponse], error)
	// ListMCPServers implements the list mcp_servers method which lists the MCP servers
	// in the Redpanda cluster.
	ListMCPServers(context.Context, *connect.Request[v1.ListMCPServersRequest]) (*connect.Response[v1.ListMCPServersResponse], error)
	// Update MCPServer updates a specific Redpanda Connect MCP server configuration.
	UpdateMCPServer(context.Context, *connect.Request[v1.UpdateMCPServerRequest]) (*connect.Response[v1.UpdateMCPServerResponse], error)
	// DeleteMCPServer deletes a specific Redpanda Connect MCP server.
	DeleteMCPServer(context.Context, *connect.Request[v1.DeleteMCPServerRequest]) (*connect.Response[v1.DeleteMCPServerResponse], error)
	// StopMCPServer stops a specific Redpanda Connect MCP server.
	StopMCPServer(context.Context, *connect.Request[v1.StopMCPServerRequest]) (*connect.Response[v1.StopMCPServerResponse], error)
	// StartMCPServer starts a specific Redpanda Connect MCP server that has been previously stopped.
	StartMCPServer(context.Context, *connect.Request[v1.StartMCPServerRequest]) (*connect.Response[v1.StartMCPServerResponse], error)
	// The configuration schema includes available components and processors in this Redpanda Connect MCP Server instance.
	GetMCPServerServiceConfigSchema(context.Context, *connect.Request[v1.GetMCPServerServiceConfigSchemaRequest]) (*connect.Response[v1.GetMCPServerServiceConfigSchemaResponse], error)
}

// NewMCPServerServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMCPServerServiceHandler(svc MCPServerServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	mCPServerServiceCreateMCPServerHandler := connect.NewUnaryHandler(
		MCPServerServiceCreateMCPServerProcedure,
		svc.CreateMCPServer,
		connect.WithSchema(mCPServerServiceCreateMCPServerMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	mCPServerServiceGetMCPServerHandler := connect.NewUnaryHandler(
		MCPServerServiceGetMCPServerProcedure,
		svc.GetMCPServer,
		connect.WithSchema(mCPServerServiceGetMCPServerMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	mCPServerServiceListMCPServersHandler := connect.NewUnaryHandler(
		MCPServerServiceListMCPServersProcedure,
		svc.ListMCPServers,
		connect.WithSchema(mCPServerServiceListMCPServersMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	mCPServerServiceUpdateMCPServerHandler := connect.NewUnaryHandler(
		MCPServerServiceUpdateMCPServerProcedure,
		svc.UpdateMCPServer,
		connect.WithSchema(mCPServerServiceUpdateMCPServerMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	mCPServerServiceDeleteMCPServerHandler := connect.NewUnaryHandler(
		MCPServerServiceDeleteMCPServerProcedure,
		svc.DeleteMCPServer,
		connect.WithSchema(mCPServerServiceDeleteMCPServerMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	mCPServerServiceStopMCPServerHandler := connect.NewUnaryHandler(
		MCPServerServiceStopMCPServerProcedure,
		svc.StopMCPServer,
		connect.WithSchema(mCPServerServiceStopMCPServerMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	mCPServerServiceStartMCPServerHandler := connect.NewUnaryHandler(
		MCPServerServiceStartMCPServerProcedure,
		svc.StartMCPServer,
		connect.WithSchema(mCPServerServiceStartMCPServerMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	mCPServerServiceGetMCPServerServiceConfigSchemaHandler := connect.NewUnaryHandler(
		MCPServerServiceGetMCPServerServiceConfigSchemaProcedure,
		svc.GetMCPServerServiceConfigSchema,
		connect.WithSchema(mCPServerServiceGetMCPServerServiceConfigSchemaMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/redpanda.api.dataplane.v1.MCPServerService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MCPServerServiceCreateMCPServerProcedure:
			mCPServerServiceCreateMCPServerHandler.ServeHTTP(w, r)
		case MCPServerServiceGetMCPServerProcedure:
			mCPServerServiceGetMCPServerHandler.ServeHTTP(w, r)
		case MCPServerServiceListMCPServersProcedure:
			mCPServerServiceListMCPServersHandler.ServeHTTP(w, r)
		case MCPServerServiceUpdateMCPServerProcedure:
			mCPServerServiceUpdateMCPServerHandler.ServeHTTP(w, r)
		case MCPServerServiceDeleteMCPServerProcedure:
			mCPServerServiceDeleteMCPServerHandler.ServeHTTP(w, r)
		case MCPServerServiceStopMCPServerProcedure:
			mCPServerServiceStopMCPServerHandler.ServeHTTP(w, r)
		case MCPServerServiceStartMCPServerProcedure:
			mCPServerServiceStartMCPServerHandler.ServeHTTP(w, r)
		case MCPServerServiceGetMCPServerServiceConfigSchemaProcedure:
			mCPServerServiceGetMCPServerServiceConfigSchemaHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMCPServerServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMCPServerServiceHandler struct{}

func (UnimplementedMCPServerServiceHandler) CreateMCPServer(context.Context, *connect.Request[v1.CreateMCPServerRequest]) (*connect.Response[v1.CreateMCPServerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.MCPServerService.CreateMCPServer is not implemented"))
}

func (UnimplementedMCPServerServiceHandler) GetMCPServer(context.Context, *connect.Request[v1.GetMCPServerRequest]) (*connect.Response[v1.GetMCPServerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.MCPServerService.GetMCPServer is not implemented"))
}

func (UnimplementedMCPServerServiceHandler) ListMCPServers(context.Context, *connect.Request[v1.ListMCPServersRequest]) (*connect.Response[v1.ListMCPServersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.MCPServerService.ListMCPServers is not implemented"))
}

func (UnimplementedMCPServerServiceHandler) UpdateMCPServer(context.Context, *connect.Request[v1.UpdateMCPServerRequest]) (*connect.Response[v1.UpdateMCPServerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.MCPServerService.UpdateMCPServer is not implemented"))
}

func (UnimplementedMCPServerServiceHandler) DeleteMCPServer(context.Context, *connect.Request[v1.DeleteMCPServerRequest]) (*connect.Response[v1.DeleteMCPServerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.MCPServerService.DeleteMCPServer is not implemented"))
}

func (UnimplementedMCPServerServiceHandler) StopMCPServer(context.Context, *connect.Request[v1.StopMCPServerRequest]) (*connect.Response[v1.StopMCPServerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.MCPServerService.StopMCPServer is not implemented"))
}

func (UnimplementedMCPServerServiceHandler) StartMCPServer(context.Context, *connect.Request[v1.StartMCPServerRequest]) (*connect.Response[v1.StartMCPServerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.MCPServerService.StartMCPServer is not implemented"))
}

func (UnimplementedMCPServerServiceHandler) GetMCPServerServiceConfigSchema(context.Context, *connect.Request[v1.GetMCPServerServiceConfigSchemaRequest]) (*connect.Response[v1.GetMCPServerServiceConfigSchemaResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.MCPServerService.GetMCPServerServiceConfigSchema is not implemented"))
}
