// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: skeleton.proto

package skeleton

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

// SkeletonClient is the client API for Skeleton service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SkeletonClient interface {
	DoIt(ctx context.Context, in *DoItRequest, opts ...grpc.CallOption) (*DoItResponse, error)
}

type skeletonClient struct {
	cc grpc.ClientConnInterface
}

func NewSkeletonClient(cc grpc.ClientConnInterface) SkeletonClient {
	return &skeletonClient{cc}
}

func (c *skeletonClient) DoIt(ctx context.Context, in *DoItRequest, opts ...grpc.CallOption) (*DoItResponse, error) {
	out := new(DoItResponse)
	err := c.cc.Invoke(ctx, "/api.public.skeleton.Skeleton/DoIt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SkeletonServer is the server API for Skeleton service.
// All implementations must embed UnimplementedSkeletonServer
// for forward compatibility
type SkeletonServer interface {
	DoIt(context.Context, *DoItRequest) (*DoItResponse, error)
	mustEmbedUnimplementedSkeletonServer()
}

// UnimplementedSkeletonServer must be embedded to have forward compatible implementations.
type UnimplementedSkeletonServer struct {
}

func (UnimplementedSkeletonServer) DoIt(context.Context, *DoItRequest) (*DoItResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoIt not implemented")
}
func (UnimplementedSkeletonServer) mustEmbedUnimplementedSkeletonServer() {}

// UnsafeSkeletonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SkeletonServer will
// result in compilation errors.
type UnsafeSkeletonServer interface {
	mustEmbedUnimplementedSkeletonServer()
}

func RegisterSkeletonServer(s grpc.ServiceRegistrar, srv SkeletonServer) {
	s.RegisterService(&Skeleton_ServiceDesc, srv)
}

func _Skeleton_DoIt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoItRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkeletonServer).DoIt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.skeleton.Skeleton/DoIt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkeletonServer).DoIt(ctx, req.(*DoItRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Skeleton_ServiceDesc is the grpc.ServiceDesc for Skeleton service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Skeleton_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.public.skeleton.Skeleton",
	HandlerType: (*SkeletonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoIt",
			Handler:    _Skeleton_DoIt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "skeleton.proto",
}
