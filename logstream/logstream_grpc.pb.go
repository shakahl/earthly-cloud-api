// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: api/public/logstream/logstream.proto

package logstream

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

// IngesterClient is the client API for Ingester service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IngesterClient interface {
	PassDelta(ctx context.Context, opts ...grpc.CallOption) (Ingester_PassDeltaClient, error)
	Snapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*SnapshotReply, error)
}

type ingesterClient struct {
	cc grpc.ClientConnInterface
}

func NewIngesterClient(cc grpc.ClientConnInterface) IngesterClient {
	return &ingesterClient{cc}
}

func (c *ingesterClient) PassDelta(ctx context.Context, opts ...grpc.CallOption) (Ingester_PassDeltaClient, error) {
	stream, err := c.cc.NewStream(ctx, &Ingester_ServiceDesc.Streams[0], "/api.public.logstream.Ingester/PassDelta", opts...)
	if err != nil {
		return nil, err
	}
	x := &ingesterPassDeltaClient{stream}
	return x, nil
}

type Ingester_PassDeltaClient interface {
	Send(*PassDeltaRequest) error
	Recv() (*PassDeltaReply, error)
	grpc.ClientStream
}

type ingesterPassDeltaClient struct {
	grpc.ClientStream
}

func (x *ingesterPassDeltaClient) Send(m *PassDeltaRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ingesterPassDeltaClient) Recv() (*PassDeltaReply, error) {
	m := new(PassDeltaReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ingesterClient) Snapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*SnapshotReply, error) {
	out := new(SnapshotReply)
	err := c.cc.Invoke(ctx, "/api.public.logstream.Ingester/Snapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IngesterServer is the server API for Ingester service.
// All implementations must embed UnimplementedIngesterServer
// for forward compatibility
type IngesterServer interface {
	PassDelta(Ingester_PassDeltaServer) error
	Snapshot(context.Context, *SnapshotRequest) (*SnapshotReply, error)
	mustEmbedUnimplementedIngesterServer()
}

// UnimplementedIngesterServer must be embedded to have forward compatible implementations.
type UnimplementedIngesterServer struct {
}

func (UnimplementedIngesterServer) PassDelta(Ingester_PassDeltaServer) error {
	return status.Errorf(codes.Unimplemented, "method PassDelta not implemented")
}
func (UnimplementedIngesterServer) Snapshot(context.Context, *SnapshotRequest) (*SnapshotReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Snapshot not implemented")
}
func (UnimplementedIngesterServer) mustEmbedUnimplementedIngesterServer() {}

// UnsafeIngesterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IngesterServer will
// result in compilation errors.
type UnsafeIngesterServer interface {
	mustEmbedUnimplementedIngesterServer()
}

func RegisterIngesterServer(s grpc.ServiceRegistrar, srv IngesterServer) {
	s.RegisterService(&Ingester_ServiceDesc, srv)
}

func _Ingester_PassDelta_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IngesterServer).PassDelta(&ingesterPassDeltaServer{stream})
}

type Ingester_PassDeltaServer interface {
	Send(*PassDeltaReply) error
	Recv() (*PassDeltaRequest, error)
	grpc.ServerStream
}

type ingesterPassDeltaServer struct {
	grpc.ServerStream
}

func (x *ingesterPassDeltaServer) Send(m *PassDeltaReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ingesterPassDeltaServer) Recv() (*PassDeltaRequest, error) {
	m := new(PassDeltaRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Ingester_Snapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngesterServer).Snapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.logstream.Ingester/Snapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngesterServer).Snapshot(ctx, req.(*SnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Ingester_ServiceDesc is the grpc.ServiceDesc for Ingester service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ingester_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.public.logstream.Ingester",
	HandlerType: (*IngesterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Snapshot",
			Handler:    _Ingester_Snapshot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PassDelta",
			Handler:       _Ingester_PassDelta_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/public/logstream/logstream.proto",
}

// LogStreamReadClient is the client API for LogStreamRead service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogStreamReadClient interface {
	// GetManifest returns the manifest for a build, from a snapshot.
	GetManifest(ctx context.Context, in *GetManifestRequest, opts ...grpc.CallOption) (*GetManifestReply, error)
	// GetLog returns the logs for a build, from a snapshot.
	GetLog(ctx context.Context, in *GetLogRequest, opts ...grpc.CallOption) (*GetLogReply, error)
	// Subscribe returns a stream of deltas for a given build.
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (LogStreamRead_SubscribeClient, error)
	// ReadLog returns a stream of log lines for a given build and target,
	// using both any available snapshot and any that might follow deltas.
	// If the follow option is used, then the stream remains open, until
	// the target has finished.
	ReadLog(ctx context.Context, in *ReadLogRequest, opts ...grpc.CallOption) (LogStreamRead_ReadLogClient, error)
}

type logStreamReadClient struct {
	cc grpc.ClientConnInterface
}

func NewLogStreamReadClient(cc grpc.ClientConnInterface) LogStreamReadClient {
	return &logStreamReadClient{cc}
}

func (c *logStreamReadClient) GetManifest(ctx context.Context, in *GetManifestRequest, opts ...grpc.CallOption) (*GetManifestReply, error) {
	out := new(GetManifestReply)
	err := c.cc.Invoke(ctx, "/api.public.logstream.LogStreamRead/GetManifest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logStreamReadClient) GetLog(ctx context.Context, in *GetLogRequest, opts ...grpc.CallOption) (*GetLogReply, error) {
	out := new(GetLogReply)
	err := c.cc.Invoke(ctx, "/api.public.logstream.LogStreamRead/GetLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logStreamReadClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (LogStreamRead_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &LogStreamRead_ServiceDesc.Streams[0], "/api.public.logstream.LogStreamRead/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &logStreamReadSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogStreamRead_SubscribeClient interface {
	Recv() (*SubscribeReply, error)
	grpc.ClientStream
}

type logStreamReadSubscribeClient struct {
	grpc.ClientStream
}

func (x *logStreamReadSubscribeClient) Recv() (*SubscribeReply, error) {
	m := new(SubscribeReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logStreamReadClient) ReadLog(ctx context.Context, in *ReadLogRequest, opts ...grpc.CallOption) (LogStreamRead_ReadLogClient, error) {
	stream, err := c.cc.NewStream(ctx, &LogStreamRead_ServiceDesc.Streams[1], "/api.public.logstream.LogStreamRead/ReadLog", opts...)
	if err != nil {
		return nil, err
	}
	x := &logStreamReadReadLogClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogStreamRead_ReadLogClient interface {
	Recv() (*ReadLogReply, error)
	grpc.ClientStream
}

type logStreamReadReadLogClient struct {
	grpc.ClientStream
}

func (x *logStreamReadReadLogClient) Recv() (*ReadLogReply, error) {
	m := new(ReadLogReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogStreamReadServer is the server API for LogStreamRead service.
// All implementations must embed UnimplementedLogStreamReadServer
// for forward compatibility
type LogStreamReadServer interface {
	// GetManifest returns the manifest for a build, from a snapshot.
	GetManifest(context.Context, *GetManifestRequest) (*GetManifestReply, error)
	// GetLog returns the logs for a build, from a snapshot.
	GetLog(context.Context, *GetLogRequest) (*GetLogReply, error)
	// Subscribe returns a stream of deltas for a given build.
	Subscribe(*SubscribeRequest, LogStreamRead_SubscribeServer) error
	// ReadLog returns a stream of log lines for a given build and target,
	// using both any available snapshot and any that might follow deltas.
	// If the follow option is used, then the stream remains open, until
	// the target has finished.
	ReadLog(*ReadLogRequest, LogStreamRead_ReadLogServer) error
	mustEmbedUnimplementedLogStreamReadServer()
}

// UnimplementedLogStreamReadServer must be embedded to have forward compatible implementations.
type UnimplementedLogStreamReadServer struct {
}

func (UnimplementedLogStreamReadServer) GetManifest(context.Context, *GetManifestRequest) (*GetManifestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManifest not implemented")
}
func (UnimplementedLogStreamReadServer) GetLog(context.Context, *GetLogRequest) (*GetLogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLog not implemented")
}
func (UnimplementedLogStreamReadServer) Subscribe(*SubscribeRequest, LogStreamRead_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedLogStreamReadServer) ReadLog(*ReadLogRequest, LogStreamRead_ReadLogServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadLog not implemented")
}
func (UnimplementedLogStreamReadServer) mustEmbedUnimplementedLogStreamReadServer() {}

// UnsafeLogStreamReadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogStreamReadServer will
// result in compilation errors.
type UnsafeLogStreamReadServer interface {
	mustEmbedUnimplementedLogStreamReadServer()
}

func RegisterLogStreamReadServer(s grpc.ServiceRegistrar, srv LogStreamReadServer) {
	s.RegisterService(&LogStreamRead_ServiceDesc, srv)
}

func _LogStreamRead_GetManifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManifestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogStreamReadServer).GetManifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.logstream.LogStreamRead/GetManifest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogStreamReadServer).GetManifest(ctx, req.(*GetManifestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogStreamRead_GetLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogStreamReadServer).GetLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.logstream.LogStreamRead/GetLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogStreamReadServer).GetLog(ctx, req.(*GetLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogStreamRead_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogStreamReadServer).Subscribe(m, &logStreamReadSubscribeServer{stream})
}

type LogStreamRead_SubscribeServer interface {
	Send(*SubscribeReply) error
	grpc.ServerStream
}

type logStreamReadSubscribeServer struct {
	grpc.ServerStream
}

func (x *logStreamReadSubscribeServer) Send(m *SubscribeReply) error {
	return x.ServerStream.SendMsg(m)
}

func _LogStreamRead_ReadLog_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadLogRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogStreamReadServer).ReadLog(m, &logStreamReadReadLogServer{stream})
}

type LogStreamRead_ReadLogServer interface {
	Send(*ReadLogReply) error
	grpc.ServerStream
}

type logStreamReadReadLogServer struct {
	grpc.ServerStream
}

func (x *logStreamReadReadLogServer) Send(m *ReadLogReply) error {
	return x.ServerStream.SendMsg(m)
}

// LogStreamRead_ServiceDesc is the grpc.ServiceDesc for LogStreamRead service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogStreamRead_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.public.logstream.LogStreamRead",
	HandlerType: (*LogStreamReadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetManifest",
			Handler:    _LogStreamRead_GetManifest_Handler,
		},
		{
			MethodName: "GetLog",
			Handler:    _LogStreamRead_GetLog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _LogStreamRead_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReadLog",
			Handler:       _LogStreamRead_ReadLog_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/public/logstream/logstream.proto",
}
