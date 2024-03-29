// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: baaa.proto

package v1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ListTitleServiceClient is the client API for ListTitleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListTitleServiceClient interface {
	GetTitle(ctx context.Context, in *GetTitleReq, opts ...grpc.CallOption) (*GetTitleResp, error)
}

type listTitleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewListTitleServiceClient(cc grpc.ClientConnInterface) ListTitleServiceClient {
	return &listTitleServiceClient{cc}
}

func (c *listTitleServiceClient) GetTitle(ctx context.Context, in *GetTitleReq, opts ...grpc.CallOption) (*GetTitleResp, error) {
	out := new(GetTitleResp)
	err := c.cc.Invoke(ctx, "/qh.samples.ListTitleService/GetTitle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListTitleServiceServer is the server API for ListTitleService service.
// All implementations must embed UnimplementedListTitleServiceServer
// for forward compatibility
type ListTitleServiceServer interface {
	GetTitle(context.Context, *GetTitleReq) (*GetTitleResp, error)
	mustEmbedUnimplementedListTitleServiceServer()
}

// UnimplementedListTitleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedListTitleServiceServer struct {
}

func (UnimplementedListTitleServiceServer) GetTitle(context.Context, *GetTitleReq) (*GetTitleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTitle not implemented")
}
func (UnimplementedListTitleServiceServer) mustEmbedUnimplementedListTitleServiceServer() {}

// UnsafeListTitleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListTitleServiceServer will
// result in compilation errors.
type UnsafeListTitleServiceServer interface {
	mustEmbedUnimplementedListTitleServiceServer()
}

func RegisterListTitleServiceServer(s grpc.ServiceRegistrar, srv ListTitleServiceServer) {
	s.RegisterService(&ListTitleService_ServiceDesc, srv)
}

func _ListTitleService_GetTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTitleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListTitleServiceServer).GetTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qh.samples.ListTitleService/GetTitle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListTitleServiceServer).GetTitle(ctx, req.(*GetTitleReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ListTitleService_ServiceDesc is the grpc.ServiceDesc for ListTitleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListTitleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qh.samples.ListTitleService",
	HandlerType: (*ListTitleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTitle",
			Handler:    _ListTitleService_GetTitle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "baaa.proto",
}
