// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: compute.proto

package compute

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

// ComputeClient is the client API for Compute service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComputeClient interface {
	// LaunchSatellite starts a new Earthly Satellite instance on the latest version.
	// The instance can be used by users to build their local, Cloud, or 3rd-party CI builds.
	LaunchSatellite(ctx context.Context, in *LaunchSatelliteRequest, opts ...grpc.CallOption) (*LaunchSatelliteResponse, error)
	// ListSatellites returns a list of Earthly Satellite instances available in the organization.
	ListSatellites(ctx context.Context, in *ListSatellitesRequest, opts ...grpc.CallOption) (*ListSatellitesResponse, error)
	// UpdateSatellite updates a Satellite instance to the latest version.
	// (I.e. the latest AMI we have in AWS).
	// Calling this may result in some down-time on the instance while it updates.
	UpdateSatellite(ctx context.Context, in *UpdateSatelliteRequest, opts ...grpc.CallOption) (*UpdateSatelliteResponse, error)
	// DeleteSatellite decommissions a Satellite instance.
	DeleteSatellite(ctx context.Context, in *DeleteSatelliteRequest, opts ...grpc.CallOption) (*DeleteSatelliteResponse, error)
	// GetSatellite retrieves the details of a particular Satellite instance.
	// Mainly intended for use by Buildkit Proxy when establishing a new connection to an instance.
	GetSatellite(ctx context.Context, in *GetSatelliteRequest, opts ...grpc.CallOption) (*GetSatelliteResponse, error)
	// GetSatelliteMetrics retrieves the recent requested metrics for the given satellites. This is not meant to be a
	// definitive historical record of build metrics, but instead a way to catch a glimpse into the current or recent
	// state of the satellite.
	ListSatelliteMetrics(ctx context.Context, in *ListSatellitesMetricsRequest, opts ...grpc.CallOption) (*ListSatellitesMetricsResponse, error)
	// WakeSatellite wakes a satellite that is in a sleep state.
	// The response returns a stream that sends updates as the satellite wakes up.
	// For example, the stream may send the following statuses:
	//    SLEEP -> STARTING -> ... -> STARTING -> OPERATIONAL -> EOF
	WakeSatellite(ctx context.Context, in *WakeSatelliteRequest, opts ...grpc.CallOption) (Compute_WakeSatelliteClient, error)
	// SleepSatellite puts a satellite to sleep when it is awake.
	// The response is a stream which looks like the inverse of a WakeSatellite response.
	// Example when satellite is awake:
	//   OPERATIONAL -> STOPPING -> ... -> STOPPING -> SLEEP -> EOF
	// Example when satellite is already asleep:
	//   SLEEP -> EOF
	SleepSatellite(ctx context.Context, in *SleepSatelliteRequest, opts ...grpc.CallOption) (Compute_SleepSatelliteClient, error)
	// ReserveSatellite both wakes the instance (if necessary) and calls reserve on buildkit.
	// It is a streaming call which returns status updates during the wake up process,
	// which can take a few moments. Some examples:
	// When the satellite is already awake, a single event is emitted before closing the stream:
	//   OPERATIONAL -> EOF
	// When the satellite is asleep, several events are emitted in the following sequence:
	//   SLEEP -> STARTING -> ... -> STARTING -> OPERATIONAL -> EOF
	// When the satellite is already being woken up, or launching for the first time:
	//   STARTING -> ... -> STARTING -> OPERATIONAL -> EOF
	// When the satellite is actively being put to sleep and needs to finish that process first:
	//   STOPPING -> ... -> STARTING -> ... -> OPERATIONAL -> EOF
	ReserveSatellite(ctx context.Context, in *ReserveSatelliteRequest, opts ...grpc.CallOption) (Compute_ReserveSatelliteClient, error)
}

type computeClient struct {
	cc grpc.ClientConnInterface
}

func NewComputeClient(cc grpc.ClientConnInterface) ComputeClient {
	return &computeClient{cc}
}

func (c *computeClient) LaunchSatellite(ctx context.Context, in *LaunchSatelliteRequest, opts ...grpc.CallOption) (*LaunchSatelliteResponse, error) {
	out := new(LaunchSatelliteResponse)
	err := c.cc.Invoke(ctx, "/api.public.compute.Compute/LaunchSatellite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeClient) ListSatellites(ctx context.Context, in *ListSatellitesRequest, opts ...grpc.CallOption) (*ListSatellitesResponse, error) {
	out := new(ListSatellitesResponse)
	err := c.cc.Invoke(ctx, "/api.public.compute.Compute/ListSatellites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeClient) UpdateSatellite(ctx context.Context, in *UpdateSatelliteRequest, opts ...grpc.CallOption) (*UpdateSatelliteResponse, error) {
	out := new(UpdateSatelliteResponse)
	err := c.cc.Invoke(ctx, "/api.public.compute.Compute/UpdateSatellite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeClient) DeleteSatellite(ctx context.Context, in *DeleteSatelliteRequest, opts ...grpc.CallOption) (*DeleteSatelliteResponse, error) {
	out := new(DeleteSatelliteResponse)
	err := c.cc.Invoke(ctx, "/api.public.compute.Compute/DeleteSatellite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeClient) GetSatellite(ctx context.Context, in *GetSatelliteRequest, opts ...grpc.CallOption) (*GetSatelliteResponse, error) {
	out := new(GetSatelliteResponse)
	err := c.cc.Invoke(ctx, "/api.public.compute.Compute/GetSatellite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeClient) ListSatelliteMetrics(ctx context.Context, in *ListSatellitesMetricsRequest, opts ...grpc.CallOption) (*ListSatellitesMetricsResponse, error) {
	out := new(ListSatellitesMetricsResponse)
	err := c.cc.Invoke(ctx, "/api.public.compute.Compute/ListSatelliteMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeClient) WakeSatellite(ctx context.Context, in *WakeSatelliteRequest, opts ...grpc.CallOption) (Compute_WakeSatelliteClient, error) {
	stream, err := c.cc.NewStream(ctx, &Compute_ServiceDesc.Streams[0], "/api.public.compute.Compute/WakeSatellite", opts...)
	if err != nil {
		return nil, err
	}
	x := &computeWakeSatelliteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Compute_WakeSatelliteClient interface {
	Recv() (*WakeSatelliteResponse, error)
	grpc.ClientStream
}

type computeWakeSatelliteClient struct {
	grpc.ClientStream
}

func (x *computeWakeSatelliteClient) Recv() (*WakeSatelliteResponse, error) {
	m := new(WakeSatelliteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *computeClient) SleepSatellite(ctx context.Context, in *SleepSatelliteRequest, opts ...grpc.CallOption) (Compute_SleepSatelliteClient, error) {
	stream, err := c.cc.NewStream(ctx, &Compute_ServiceDesc.Streams[1], "/api.public.compute.Compute/SleepSatellite", opts...)
	if err != nil {
		return nil, err
	}
	x := &computeSleepSatelliteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Compute_SleepSatelliteClient interface {
	Recv() (*SleepSatelliteResponse, error)
	grpc.ClientStream
}

type computeSleepSatelliteClient struct {
	grpc.ClientStream
}

func (x *computeSleepSatelliteClient) Recv() (*SleepSatelliteResponse, error) {
	m := new(SleepSatelliteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *computeClient) ReserveSatellite(ctx context.Context, in *ReserveSatelliteRequest, opts ...grpc.CallOption) (Compute_ReserveSatelliteClient, error) {
	stream, err := c.cc.NewStream(ctx, &Compute_ServiceDesc.Streams[2], "/api.public.compute.Compute/ReserveSatellite", opts...)
	if err != nil {
		return nil, err
	}
	x := &computeReserveSatelliteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Compute_ReserveSatelliteClient interface {
	Recv() (*ReserveSatelliteResponse, error)
	grpc.ClientStream
}

type computeReserveSatelliteClient struct {
	grpc.ClientStream
}

func (x *computeReserveSatelliteClient) Recv() (*ReserveSatelliteResponse, error) {
	m := new(ReserveSatelliteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ComputeServer is the server API for Compute service.
// All implementations must embed UnimplementedComputeServer
// for forward compatibility
type ComputeServer interface {
	// LaunchSatellite starts a new Earthly Satellite instance on the latest version.
	// The instance can be used by users to build their local, Cloud, or 3rd-party CI builds.
	LaunchSatellite(context.Context, *LaunchSatelliteRequest) (*LaunchSatelliteResponse, error)
	// ListSatellites returns a list of Earthly Satellite instances available in the organization.
	ListSatellites(context.Context, *ListSatellitesRequest) (*ListSatellitesResponse, error)
	// UpdateSatellite updates a Satellite instance to the latest version.
	// (I.e. the latest AMI we have in AWS).
	// Calling this may result in some down-time on the instance while it updates.
	UpdateSatellite(context.Context, *UpdateSatelliteRequest) (*UpdateSatelliteResponse, error)
	// DeleteSatellite decommissions a Satellite instance.
	DeleteSatellite(context.Context, *DeleteSatelliteRequest) (*DeleteSatelliteResponse, error)
	// GetSatellite retrieves the details of a particular Satellite instance.
	// Mainly intended for use by Buildkit Proxy when establishing a new connection to an instance.
	GetSatellite(context.Context, *GetSatelliteRequest) (*GetSatelliteResponse, error)
	// GetSatelliteMetrics retrieves the recent requested metrics for the given satellites. This is not meant to be a
	// definitive historical record of build metrics, but instead a way to catch a glimpse into the current or recent
	// state of the satellite.
	ListSatelliteMetrics(context.Context, *ListSatellitesMetricsRequest) (*ListSatellitesMetricsResponse, error)
	// WakeSatellite wakes a satellite that is in a sleep state.
	// The response returns a stream that sends updates as the satellite wakes up.
	// For example, the stream may send the following statuses:
	//    SLEEP -> STARTING -> ... -> STARTING -> OPERATIONAL -> EOF
	WakeSatellite(*WakeSatelliteRequest, Compute_WakeSatelliteServer) error
	// SleepSatellite puts a satellite to sleep when it is awake.
	// The response is a stream which looks like the inverse of a WakeSatellite response.
	// Example when satellite is awake:
	//   OPERATIONAL -> STOPPING -> ... -> STOPPING -> SLEEP -> EOF
	// Example when satellite is already asleep:
	//   SLEEP -> EOF
	SleepSatellite(*SleepSatelliteRequest, Compute_SleepSatelliteServer) error
	// ReserveSatellite both wakes the instance (if necessary) and calls reserve on buildkit.
	// It is a streaming call which returns status updates during the wake up process,
	// which can take a few moments. Some examples:
	// When the satellite is already awake, a single event is emitted before closing the stream:
	//   OPERATIONAL -> EOF
	// When the satellite is asleep, several events are emitted in the following sequence:
	//   SLEEP -> STARTING -> ... -> STARTING -> OPERATIONAL -> EOF
	// When the satellite is already being woken up, or launching for the first time:
	//   STARTING -> ... -> STARTING -> OPERATIONAL -> EOF
	// When the satellite is actively being put to sleep and needs to finish that process first:
	//   STOPPING -> ... -> STARTING -> ... -> OPERATIONAL -> EOF
	ReserveSatellite(*ReserveSatelliteRequest, Compute_ReserveSatelliteServer) error
	mustEmbedUnimplementedComputeServer()
}

// UnimplementedComputeServer must be embedded to have forward compatible implementations.
type UnimplementedComputeServer struct {
}

func (UnimplementedComputeServer) LaunchSatellite(context.Context, *LaunchSatelliteRequest) (*LaunchSatelliteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaunchSatellite not implemented")
}
func (UnimplementedComputeServer) ListSatellites(context.Context, *ListSatellitesRequest) (*ListSatellitesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSatellites not implemented")
}
func (UnimplementedComputeServer) UpdateSatellite(context.Context, *UpdateSatelliteRequest) (*UpdateSatelliteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSatellite not implemented")
}
func (UnimplementedComputeServer) DeleteSatellite(context.Context, *DeleteSatelliteRequest) (*DeleteSatelliteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSatellite not implemented")
}
func (UnimplementedComputeServer) GetSatellite(context.Context, *GetSatelliteRequest) (*GetSatelliteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSatellite not implemented")
}
func (UnimplementedComputeServer) ListSatelliteMetrics(context.Context, *ListSatellitesMetricsRequest) (*ListSatellitesMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSatelliteMetrics not implemented")
}
func (UnimplementedComputeServer) WakeSatellite(*WakeSatelliteRequest, Compute_WakeSatelliteServer) error {
	return status.Errorf(codes.Unimplemented, "method WakeSatellite not implemented")
}
func (UnimplementedComputeServer) SleepSatellite(*SleepSatelliteRequest, Compute_SleepSatelliteServer) error {
	return status.Errorf(codes.Unimplemented, "method SleepSatellite not implemented")
}
func (UnimplementedComputeServer) ReserveSatellite(*ReserveSatelliteRequest, Compute_ReserveSatelliteServer) error {
	return status.Errorf(codes.Unimplemented, "method ReserveSatellite not implemented")
}
func (UnimplementedComputeServer) mustEmbedUnimplementedComputeServer() {}

// UnsafeComputeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComputeServer will
// result in compilation errors.
type UnsafeComputeServer interface {
	mustEmbedUnimplementedComputeServer()
}

func RegisterComputeServer(s grpc.ServiceRegistrar, srv ComputeServer) {
	s.RegisterService(&Compute_ServiceDesc, srv)
}

func _Compute_LaunchSatellite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LaunchSatelliteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeServer).LaunchSatellite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.compute.Compute/LaunchSatellite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeServer).LaunchSatellite(ctx, req.(*LaunchSatelliteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Compute_ListSatellites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSatellitesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeServer).ListSatellites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.compute.Compute/ListSatellites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeServer).ListSatellites(ctx, req.(*ListSatellitesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Compute_UpdateSatellite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSatelliteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeServer).UpdateSatellite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.compute.Compute/UpdateSatellite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeServer).UpdateSatellite(ctx, req.(*UpdateSatelliteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Compute_DeleteSatellite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSatelliteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeServer).DeleteSatellite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.compute.Compute/DeleteSatellite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeServer).DeleteSatellite(ctx, req.(*DeleteSatelliteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Compute_GetSatellite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSatelliteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeServer).GetSatellite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.compute.Compute/GetSatellite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeServer).GetSatellite(ctx, req.(*GetSatelliteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Compute_ListSatelliteMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSatellitesMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeServer).ListSatelliteMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.compute.Compute/ListSatelliteMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeServer).ListSatelliteMetrics(ctx, req.(*ListSatellitesMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Compute_WakeSatellite_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WakeSatelliteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ComputeServer).WakeSatellite(m, &computeWakeSatelliteServer{stream})
}

type Compute_WakeSatelliteServer interface {
	Send(*WakeSatelliteResponse) error
	grpc.ServerStream
}

type computeWakeSatelliteServer struct {
	grpc.ServerStream
}

func (x *computeWakeSatelliteServer) Send(m *WakeSatelliteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Compute_SleepSatellite_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SleepSatelliteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ComputeServer).SleepSatellite(m, &computeSleepSatelliteServer{stream})
}

type Compute_SleepSatelliteServer interface {
	Send(*SleepSatelliteResponse) error
	grpc.ServerStream
}

type computeSleepSatelliteServer struct {
	grpc.ServerStream
}

func (x *computeSleepSatelliteServer) Send(m *SleepSatelliteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Compute_ReserveSatellite_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReserveSatelliteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ComputeServer).ReserveSatellite(m, &computeReserveSatelliteServer{stream})
}

type Compute_ReserveSatelliteServer interface {
	Send(*ReserveSatelliteResponse) error
	grpc.ServerStream
}

type computeReserveSatelliteServer struct {
	grpc.ServerStream
}

func (x *computeReserveSatelliteServer) Send(m *ReserveSatelliteResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Compute_ServiceDesc is the grpc.ServiceDesc for Compute service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Compute_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.public.compute.Compute",
	HandlerType: (*ComputeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LaunchSatellite",
			Handler:    _Compute_LaunchSatellite_Handler,
		},
		{
			MethodName: "ListSatellites",
			Handler:    _Compute_ListSatellites_Handler,
		},
		{
			MethodName: "UpdateSatellite",
			Handler:    _Compute_UpdateSatellite_Handler,
		},
		{
			MethodName: "DeleteSatellite",
			Handler:    _Compute_DeleteSatellite_Handler,
		},
		{
			MethodName: "GetSatellite",
			Handler:    _Compute_GetSatellite_Handler,
		},
		{
			MethodName: "ListSatelliteMetrics",
			Handler:    _Compute_ListSatelliteMetrics_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WakeSatellite",
			Handler:       _Compute_WakeSatellite_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SleepSatellite",
			Handler:       _Compute_SleepSatellite_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReserveSatellite",
			Handler:       _Compute_ReserveSatellite_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "compute.proto",
}
