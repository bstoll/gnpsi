// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gnpsi

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

// GNPSIClient is the client API for GNPSI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GNPSIClient interface {
	// gNPSI subscription allows client to subscribe to SFlow/NetFlow/IPFIX
	// updates from the device.  Past updates, i.e., updates before the
	// subscription is received, will not be presented to the subscribing client.
	Subscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (GNPSI_SubscribeClient, error)
}

type gNPSIClient struct {
	cc grpc.ClientConnInterface
}

func NewGNPSIClient(cc grpc.ClientConnInterface) GNPSIClient {
	return &gNPSIClient{cc}
}

func (c *gNPSIClient) Subscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (GNPSI_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &GNPSI_ServiceDesc.Streams[0], "/gnpsi.gNPSI/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &gNPSISubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GNPSI_SubscribeClient interface {
	Recv() (*Sample, error)
	grpc.ClientStream
}

type gNPSISubscribeClient struct {
	grpc.ClientStream
}

func (x *gNPSISubscribeClient) Recv() (*Sample, error) {
	m := new(Sample)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GNPSIServer is the server API for GNPSI service.
// All implementations must embed UnimplementedGNPSIServer
// for forward compatibility
type GNPSIServer interface {
	// gNPSI subscription allows client to subscribe to SFlow/NetFlow/IPFIX
	// updates from the device.  Past updates, i.e., updates before the
	// subscription is received, will not be presented to the subscribing client.
	Subscribe(*Request, GNPSI_SubscribeServer) error
	mustEmbedUnimplementedGNPSIServer()
}

// UnimplementedGNPSIServer must be embedded to have forward compatible implementations.
type UnimplementedGNPSIServer struct {
}

func (UnimplementedGNPSIServer) Subscribe(*Request, GNPSI_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedGNPSIServer) mustEmbedUnimplementedGNPSIServer() {}

// UnsafeGNPSIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GNPSIServer will
// result in compilation errors.
type UnsafeGNPSIServer interface {
	mustEmbedUnimplementedGNPSIServer()
}

func RegisterGNPSIServer(s grpc.ServiceRegistrar, srv GNPSIServer) {
	s.RegisterService(&GNPSI_ServiceDesc, srv)
}

func _GNPSI_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GNPSIServer).Subscribe(m, &gNPSISubscribeServer{stream})
}

type GNPSI_SubscribeServer interface {
	Send(*Sample) error
	grpc.ServerStream
}

type gNPSISubscribeServer struct {
	grpc.ServerStream
}

func (x *gNPSISubscribeServer) Send(m *Sample) error {
	return x.ServerStream.SendMsg(m)
}

// GNPSI_ServiceDesc is the grpc.ServiceDesc for GNPSI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GNPSI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnpsi.gNPSI",
	HandlerType: (*GNPSIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _GNPSI_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/gnpsi/gnpsi.proto",
}