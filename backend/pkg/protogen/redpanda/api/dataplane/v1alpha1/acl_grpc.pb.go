// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: redpanda/api/dataplane/v1alpha1/acl.proto

package dataplanev1alpha1

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
	ACLService_ListACLs_FullMethodName   = "/redpanda.api.dataplane.v1alpha1.ACLService/ListACLs"
	ACLService_CreateACL_FullMethodName  = "/redpanda.api.dataplane.v1alpha1.ACLService/CreateACL"
	ACLService_DeleteACLs_FullMethodName = "/redpanda.api.dataplane.v1alpha1.ACLService/DeleteACLs"
)

// ACLServiceClient is the client API for ACLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Deprecated: Do not use.
type ACLServiceClient interface {
	// Deprecated: Do not use.
	ListACLs(ctx context.Context, in *ListACLsRequest, opts ...grpc.CallOption) (*ListACLsResponse, error)
	// Deprecated: Do not use.
	CreateACL(ctx context.Context, in *CreateACLRequest, opts ...grpc.CallOption) (*CreateACLResponse, error)
	// Deprecated: Do not use.
	DeleteACLs(ctx context.Context, in *DeleteACLsRequest, opts ...grpc.CallOption) (*DeleteACLsResponse, error)
}

type aCLServiceClient struct {
	cc grpc.ClientConnInterface
}

// Deprecated: Do not use.
func NewACLServiceClient(cc grpc.ClientConnInterface) ACLServiceClient {
	return &aCLServiceClient{cc}
}

// Deprecated: Do not use.
func (c *aCLServiceClient) ListACLs(ctx context.Context, in *ListACLsRequest, opts ...grpc.CallOption) (*ListACLsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListACLsResponse)
	err := c.cc.Invoke(ctx, ACLService_ListACLs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *aCLServiceClient) CreateACL(ctx context.Context, in *CreateACLRequest, opts ...grpc.CallOption) (*CreateACLResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateACLResponse)
	err := c.cc.Invoke(ctx, ACLService_CreateACL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *aCLServiceClient) DeleteACLs(ctx context.Context, in *DeleteACLsRequest, opts ...grpc.CallOption) (*DeleteACLsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteACLsResponse)
	err := c.cc.Invoke(ctx, ACLService_DeleteACLs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ACLServiceServer is the server API for ACLService service.
// All implementations must embed UnimplementedACLServiceServer
// for forward compatibility.
//
// Deprecated: Do not use.
type ACLServiceServer interface {
	// Deprecated: Do not use.
	ListACLs(context.Context, *ListACLsRequest) (*ListACLsResponse, error)
	// Deprecated: Do not use.
	CreateACL(context.Context, *CreateACLRequest) (*CreateACLResponse, error)
	// Deprecated: Do not use.
	DeleteACLs(context.Context, *DeleteACLsRequest) (*DeleteACLsResponse, error)
	mustEmbedUnimplementedACLServiceServer()
}

// UnimplementedACLServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedACLServiceServer struct{}

func (UnimplementedACLServiceServer) ListACLs(context.Context, *ListACLsRequest) (*ListACLsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListACLs not implemented")
}
func (UnimplementedACLServiceServer) CreateACL(context.Context, *CreateACLRequest) (*CreateACLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateACL not implemented")
}
func (UnimplementedACLServiceServer) DeleteACLs(context.Context, *DeleteACLsRequest) (*DeleteACLsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteACLs not implemented")
}
func (UnimplementedACLServiceServer) mustEmbedUnimplementedACLServiceServer() {}
func (UnimplementedACLServiceServer) testEmbeddedByValue()                    {}

// UnsafeACLServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ACLServiceServer will
// result in compilation errors.
type UnsafeACLServiceServer interface {
	mustEmbedUnimplementedACLServiceServer()
}

// Deprecated: Do not use.
func RegisterACLServiceServer(s grpc.ServiceRegistrar, srv ACLServiceServer) {
	// If the following call pancis, it indicates UnimplementedACLServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ACLService_ServiceDesc, srv)
}

func _ACLService_ListACLs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListACLsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACLServiceServer).ListACLs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ACLService_ListACLs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACLServiceServer).ListACLs(ctx, req.(*ListACLsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACLService_CreateACL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateACLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACLServiceServer).CreateACL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ACLService_CreateACL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACLServiceServer).CreateACL(ctx, req.(*CreateACLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACLService_DeleteACLs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteACLsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACLServiceServer).DeleteACLs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ACLService_DeleteACLs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACLServiceServer).DeleteACLs(ctx, req.(*DeleteACLsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ACLService_ServiceDesc is the grpc.ServiceDesc for ACLService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ACLService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "redpanda.api.dataplane.v1alpha1.ACLService",
	HandlerType: (*ACLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListACLs",
			Handler:    _ACLService_ListACLs_Handler,
		},
		{
			MethodName: "CreateACL",
			Handler:    _ACLService_CreateACL_Handler,
		},
		{
			MethodName: "DeleteACLs",
			Handler:    _ACLService_DeleteACLs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "redpanda/api/dataplane/v1alpha1/acl.proto",
}
