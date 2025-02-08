// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: sso/sso.proto

package link

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
	Link_GetOriginalLink_FullMethodName   = "/link.Link/GetOriginalLink"
	Link_CreateShorterLink_FullMethodName = "/link.Link/CreateShorterLink"
)

// LinkClient is the client API for Link service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkClient interface {
	GetOriginalLink(ctx context.Context, in *GetOriginalLinkRequest, opts ...grpc.CallOption) (*GetOriginalLinkResponse, error)
	CreateShorterLink(ctx context.Context, in *CreateShorterLinkRequest, opts ...grpc.CallOption) (*CreateShorterLinkResponse, error)
}

type linkClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkClient(cc grpc.ClientConnInterface) LinkClient {
	return &linkClient{cc}
}

func (c *linkClient) GetOriginalLink(ctx context.Context, in *GetOriginalLinkRequest, opts ...grpc.CallOption) (*GetOriginalLinkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOriginalLinkResponse)
	err := c.cc.Invoke(ctx, Link_GetOriginalLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkClient) CreateShorterLink(ctx context.Context, in *CreateShorterLinkRequest, opts ...grpc.CallOption) (*CreateShorterLinkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateShorterLinkResponse)
	err := c.cc.Invoke(ctx, Link_CreateShorterLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkServer is the server API for Link service.
// All implementations must embed UnimplementedLinkServer
// for forward compatibility.
type LinkServer interface {
	GetOriginalLink(context.Context, *GetOriginalLinkRequest) (*GetOriginalLinkResponse, error)
	CreateShorterLink(context.Context, *CreateShorterLinkRequest) (*CreateShorterLinkResponse, error)
	mustEmbedUnimplementedLinkServer()
}

// UnimplementedLinkServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLinkServer struct{}

func (UnimplementedLinkServer) GetOriginalLink(context.Context, *GetOriginalLinkRequest) (*GetOriginalLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOriginalLink not implemented")
}
func (UnimplementedLinkServer) CreateShorterLink(context.Context, *CreateShorterLinkRequest) (*CreateShorterLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShorterLink not implemented")
}
func (UnimplementedLinkServer) mustEmbedUnimplementedLinkServer() {}
func (UnimplementedLinkServer) testEmbeddedByValue()              {}

// UnsafeLinkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkServer will
// result in compilation errors.
type UnsafeLinkServer interface {
	mustEmbedUnimplementedLinkServer()
}

func RegisterLinkServer(s grpc.ServiceRegistrar, srv LinkServer) {
	// If the following call pancis, it indicates UnimplementedLinkServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Link_ServiceDesc, srv)
}

func _Link_GetOriginalLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOriginalLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServer).GetOriginalLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Link_GetOriginalLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServer).GetOriginalLink(ctx, req.(*GetOriginalLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Link_CreateShorterLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShorterLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServer).CreateShorterLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Link_CreateShorterLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServer).CreateShorterLink(ctx, req.(*CreateShorterLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Link_ServiceDesc is the grpc.ServiceDesc for Link service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Link_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "link.Link",
	HandlerType: (*LinkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOriginalLink",
			Handler:    _Link_GetOriginalLink_Handler,
		},
		{
			MethodName: "CreateShorterLink",
			Handler:    _Link_CreateShorterLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/sso.proto",
}
