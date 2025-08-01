// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: redpanda/api/dataplane/v1/mcp.proto

package dataplanev1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MCPServerService_CreateMCPServer_FullMethodName                 = "/redpanda.api.dataplane.v1.MCPServerService/CreateMCPServer"
	MCPServerService_GetMCPServer_FullMethodName                    = "/redpanda.api.dataplane.v1.MCPServerService/GetMCPServer"
	MCPServerService_ListMCPServers_FullMethodName                  = "/redpanda.api.dataplane.v1.MCPServerService/ListMCPServers"
	MCPServerService_UpdateMCPServer_FullMethodName                 = "/redpanda.api.dataplane.v1.MCPServerService/UpdateMCPServer"
	MCPServerService_DeleteMCPServer_FullMethodName                 = "/redpanda.api.dataplane.v1.MCPServerService/DeleteMCPServer"
	MCPServerService_StopMCPServer_FullMethodName                   = "/redpanda.api.dataplane.v1.MCPServerService/StopMCPServer"
	MCPServerService_StartMCPServer_FullMethodName                  = "/redpanda.api.dataplane.v1.MCPServerService/StartMCPServer"
	MCPServerService_GetMCPServerServiceConfigSchema_FullMethodName = "/redpanda.api.dataplane.v1.MCPServerService/GetMCPServerServiceConfigSchema"
)

// MCPServerServiceClient is the client API for MCPServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// MCPServer is the service for Redpanda Connect MCP Servers.
// It exposes the API for creating and managing Redpanda Connect MCP servers and their configurations.
type MCPServerServiceClient interface {
	// CreateMCPServer creates a Redpanda Connect MCP Server in the Redpanda cluster.
	CreateMCPServer(ctx context.Context, in *CreateMCPServerRequest, opts ...grpc.CallOption) (*CreateMCPServerResponse, error)
	// GetMCPServer gets a specific Redpanda Connect MCP Server.
	GetMCPServer(ctx context.Context, in *GetMCPServerRequest, opts ...grpc.CallOption) (*GetMCPServerResponse, error)
	// ListMCPServers implements the list mcp_servers method which lists the MCP servers
	// in the Redpanda cluster.
	ListMCPServers(ctx context.Context, in *ListMCPServersRequest, opts ...grpc.CallOption) (*ListMCPServersResponse, error)
	// Update MCPServer updates a specific Redpanda Connect MCP server configuration.
	UpdateMCPServer(ctx context.Context, in *UpdateMCPServerRequest, opts ...grpc.CallOption) (*UpdateMCPServerResponse, error)
	// DeleteMCPServer deletes a specific Redpanda Connect MCP server.
	DeleteMCPServer(ctx context.Context, in *DeleteMCPServerRequest, opts ...grpc.CallOption) (*DeleteMCPServerResponse, error)
	// StopMCPServer stops a specific Redpanda Connect MCP server.
	StopMCPServer(ctx context.Context, in *StopMCPServerRequest, opts ...grpc.CallOption) (*StopMCPServerResponse, error)
	// StartMCPServer starts a specific Redpanda Connect MCP server that has been previously stopped.
	StartMCPServer(ctx context.Context, in *StartMCPServerRequest, opts ...grpc.CallOption) (*StartMCPServerResponse, error)
	// The configuration schema includes available components and processors in this Redpanda Connect MCP Server instance.
	GetMCPServerServiceConfigSchema(ctx context.Context, in *GetMCPServerServiceConfigSchemaRequest, opts ...grpc.CallOption) (*GetMCPServerServiceConfigSchemaResponse, error)
}

type mCPServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMCPServerServiceClient(cc grpc.ClientConnInterface) MCPServerServiceClient {
	return &mCPServerServiceClient{cc}
}

func (c *mCPServerServiceClient) CreateMCPServer(ctx context.Context, in *CreateMCPServerRequest, opts ...grpc.CallOption) (*CreateMCPServerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateMCPServerResponse)
	err := c.cc.Invoke(ctx, MCPServerService_CreateMCPServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mCPServerServiceClient) GetMCPServer(ctx context.Context, in *GetMCPServerRequest, opts ...grpc.CallOption) (*GetMCPServerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMCPServerResponse)
	err := c.cc.Invoke(ctx, MCPServerService_GetMCPServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mCPServerServiceClient) ListMCPServers(ctx context.Context, in *ListMCPServersRequest, opts ...grpc.CallOption) (*ListMCPServersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListMCPServersResponse)
	err := c.cc.Invoke(ctx, MCPServerService_ListMCPServers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mCPServerServiceClient) UpdateMCPServer(ctx context.Context, in *UpdateMCPServerRequest, opts ...grpc.CallOption) (*UpdateMCPServerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateMCPServerResponse)
	err := c.cc.Invoke(ctx, MCPServerService_UpdateMCPServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mCPServerServiceClient) DeleteMCPServer(ctx context.Context, in *DeleteMCPServerRequest, opts ...grpc.CallOption) (*DeleteMCPServerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMCPServerResponse)
	err := c.cc.Invoke(ctx, MCPServerService_DeleteMCPServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mCPServerServiceClient) StopMCPServer(ctx context.Context, in *StopMCPServerRequest, opts ...grpc.CallOption) (*StopMCPServerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StopMCPServerResponse)
	err := c.cc.Invoke(ctx, MCPServerService_StopMCPServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mCPServerServiceClient) StartMCPServer(ctx context.Context, in *StartMCPServerRequest, opts ...grpc.CallOption) (*StartMCPServerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartMCPServerResponse)
	err := c.cc.Invoke(ctx, MCPServerService_StartMCPServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mCPServerServiceClient) GetMCPServerServiceConfigSchema(ctx context.Context, in *GetMCPServerServiceConfigSchemaRequest, opts ...grpc.CallOption) (*GetMCPServerServiceConfigSchemaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMCPServerServiceConfigSchemaResponse)
	err := c.cc.Invoke(ctx, MCPServerService_GetMCPServerServiceConfigSchema_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MCPServerServiceServer is the server API for MCPServerService service.
// All implementations must embed UnimplementedMCPServerServiceServer
// for forward compatibility.
//
// MCPServer is the service for Redpanda Connect MCP Servers.
// It exposes the API for creating and managing Redpanda Connect MCP servers and their configurations.
type MCPServerServiceServer interface {
	// CreateMCPServer creates a Redpanda Connect MCP Server in the Redpanda cluster.
	CreateMCPServer(context.Context, *CreateMCPServerRequest) (*CreateMCPServerResponse, error)
	// GetMCPServer gets a specific Redpanda Connect MCP Server.
	GetMCPServer(context.Context, *GetMCPServerRequest) (*GetMCPServerResponse, error)
	// ListMCPServers implements the list mcp_servers method which lists the MCP servers
	// in the Redpanda cluster.
	ListMCPServers(context.Context, *ListMCPServersRequest) (*ListMCPServersResponse, error)
	// Update MCPServer updates a specific Redpanda Connect MCP server configuration.
	UpdateMCPServer(context.Context, *UpdateMCPServerRequest) (*UpdateMCPServerResponse, error)
	// DeleteMCPServer deletes a specific Redpanda Connect MCP server.
	DeleteMCPServer(context.Context, *DeleteMCPServerRequest) (*DeleteMCPServerResponse, error)
	// StopMCPServer stops a specific Redpanda Connect MCP server.
	StopMCPServer(context.Context, *StopMCPServerRequest) (*StopMCPServerResponse, error)
	// StartMCPServer starts a specific Redpanda Connect MCP server that has been previously stopped.
	StartMCPServer(context.Context, *StartMCPServerRequest) (*StartMCPServerResponse, error)
	// The configuration schema includes available components and processors in this Redpanda Connect MCP Server instance.
	GetMCPServerServiceConfigSchema(context.Context, *GetMCPServerServiceConfigSchemaRequest) (*GetMCPServerServiceConfigSchemaResponse, error)
	mustEmbedUnimplementedMCPServerServiceServer()
}

// UnimplementedMCPServerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMCPServerServiceServer struct{}

func (UnimplementedMCPServerServiceServer) CreateMCPServer(context.Context, *CreateMCPServerRequest) (*CreateMCPServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMCPServer not implemented")
}
func (UnimplementedMCPServerServiceServer) GetMCPServer(context.Context, *GetMCPServerRequest) (*GetMCPServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMCPServer not implemented")
}
func (UnimplementedMCPServerServiceServer) ListMCPServers(context.Context, *ListMCPServersRequest) (*ListMCPServersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMCPServers not implemented")
}
func (UnimplementedMCPServerServiceServer) UpdateMCPServer(context.Context, *UpdateMCPServerRequest) (*UpdateMCPServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMCPServer not implemented")
}
func (UnimplementedMCPServerServiceServer) DeleteMCPServer(context.Context, *DeleteMCPServerRequest) (*DeleteMCPServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMCPServer not implemented")
}
func (UnimplementedMCPServerServiceServer) StopMCPServer(context.Context, *StopMCPServerRequest) (*StopMCPServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopMCPServer not implemented")
}
func (UnimplementedMCPServerServiceServer) StartMCPServer(context.Context, *StartMCPServerRequest) (*StartMCPServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMCPServer not implemented")
}
func (UnimplementedMCPServerServiceServer) GetMCPServerServiceConfigSchema(context.Context, *GetMCPServerServiceConfigSchemaRequest) (*GetMCPServerServiceConfigSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMCPServerServiceConfigSchema not implemented")
}
func (UnimplementedMCPServerServiceServer) mustEmbedUnimplementedMCPServerServiceServer() {}
func (UnimplementedMCPServerServiceServer) testEmbeddedByValue()                          {}

// UnsafeMCPServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MCPServerServiceServer will
// result in compilation errors.
type UnsafeMCPServerServiceServer interface {
	mustEmbedUnimplementedMCPServerServiceServer()
}

func RegisterMCPServerServiceServer(s grpc.ServiceRegistrar, srv MCPServerServiceServer) {
	// If the following call pancis, it indicates UnimplementedMCPServerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MCPServerService_ServiceDesc, srv)
}

func _MCPServerService_CreateMCPServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMCPServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MCPServerServiceServer).CreateMCPServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MCPServerService_CreateMCPServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MCPServerServiceServer).CreateMCPServer(ctx, req.(*CreateMCPServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MCPServerService_GetMCPServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMCPServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MCPServerServiceServer).GetMCPServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MCPServerService_GetMCPServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MCPServerServiceServer).GetMCPServer(ctx, req.(*GetMCPServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MCPServerService_ListMCPServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMCPServersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MCPServerServiceServer).ListMCPServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MCPServerService_ListMCPServers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MCPServerServiceServer).ListMCPServers(ctx, req.(*ListMCPServersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MCPServerService_UpdateMCPServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMCPServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MCPServerServiceServer).UpdateMCPServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MCPServerService_UpdateMCPServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MCPServerServiceServer).UpdateMCPServer(ctx, req.(*UpdateMCPServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MCPServerService_DeleteMCPServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMCPServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MCPServerServiceServer).DeleteMCPServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MCPServerService_DeleteMCPServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MCPServerServiceServer).DeleteMCPServer(ctx, req.(*DeleteMCPServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MCPServerService_StopMCPServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopMCPServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MCPServerServiceServer).StopMCPServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MCPServerService_StopMCPServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MCPServerServiceServer).StopMCPServer(ctx, req.(*StopMCPServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MCPServerService_StartMCPServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMCPServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MCPServerServiceServer).StartMCPServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MCPServerService_StartMCPServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MCPServerServiceServer).StartMCPServer(ctx, req.(*StartMCPServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MCPServerService_GetMCPServerServiceConfigSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMCPServerServiceConfigSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MCPServerServiceServer).GetMCPServerServiceConfigSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MCPServerService_GetMCPServerServiceConfigSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MCPServerServiceServer).GetMCPServerServiceConfigSchema(ctx, req.(*GetMCPServerServiceConfigSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MCPServerService_ServiceDesc is the grpc.ServiceDesc for MCPServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MCPServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "redpanda.api.dataplane.v1.MCPServerService",
	HandlerType: (*MCPServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMCPServer",
			Handler:    _MCPServerService_CreateMCPServer_Handler,
		},
		{
			MethodName: "GetMCPServer",
			Handler:    _MCPServerService_GetMCPServer_Handler,
		},
		{
			MethodName: "ListMCPServers",
			Handler:    _MCPServerService_ListMCPServers_Handler,
		},
		{
			MethodName: "UpdateMCPServer",
			Handler:    _MCPServerService_UpdateMCPServer_Handler,
		},
		{
			MethodName: "DeleteMCPServer",
			Handler:    _MCPServerService_DeleteMCPServer_Handler,
		},
		{
			MethodName: "StopMCPServer",
			Handler:    _MCPServerService_StopMCPServer_Handler,
		},
		{
			MethodName: "StartMCPServer",
			Handler:    _MCPServerService_StartMCPServer_Handler,
		},
		{
			MethodName: "GetMCPServerServiceConfigSchema",
			Handler:    _MCPServerService_GetMCPServerServiceConfigSchema_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "redpanda/api/dataplane/v1/mcp.proto",
}
