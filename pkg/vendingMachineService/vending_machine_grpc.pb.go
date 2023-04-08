// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: vending_machine.proto

package vendingMachineService

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

// VendingMachineServiceClient is the client API for VendingMachineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VendingMachineServiceClient interface {
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error)
	ExecuteCommand(ctx context.Context, opts ...grpc.CallOption) (VendingMachineService_ExecuteCommandClient, error)
	NotifyEvent(ctx context.Context, in *NotifyEventRequest, opts ...grpc.CallOption) (VendingMachineService_NotifyEventClient, error)
}

type vendingMachineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVendingMachineServiceClient(cc grpc.ClientConnInterface) VendingMachineServiceClient {
	return &vendingMachineServiceClient{cc}
}

func (c *vendingMachineServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error) {
	out := new(GetProductResponse)
	err := c.cc.Invoke(ctx, "/vending_machine.v1.VendingMachineService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendingMachineServiceClient) ExecuteCommand(ctx context.Context, opts ...grpc.CallOption) (VendingMachineService_ExecuteCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &VendingMachineService_ServiceDesc.Streams[0], "/vending_machine.v1.VendingMachineService/ExecuteCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &vendingMachineServiceExecuteCommandClient{stream}
	return x, nil
}

type VendingMachineService_ExecuteCommandClient interface {
	Send(*ExecuteCommandRequest) error
	Recv() (*ExecuteCommandResponse, error)
	grpc.ClientStream
}

type vendingMachineServiceExecuteCommandClient struct {
	grpc.ClientStream
}

func (x *vendingMachineServiceExecuteCommandClient) Send(m *ExecuteCommandRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *vendingMachineServiceExecuteCommandClient) Recv() (*ExecuteCommandResponse, error) {
	m := new(ExecuteCommandResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vendingMachineServiceClient) NotifyEvent(ctx context.Context, in *NotifyEventRequest, opts ...grpc.CallOption) (VendingMachineService_NotifyEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &VendingMachineService_ServiceDesc.Streams[1], "/vending_machine.v1.VendingMachineService/NotifyEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &vendingMachineServiceNotifyEventClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VendingMachineService_NotifyEventClient interface {
	Recv() (*NotifyEventResponse, error)
	grpc.ClientStream
}

type vendingMachineServiceNotifyEventClient struct {
	grpc.ClientStream
}

func (x *vendingMachineServiceNotifyEventClient) Recv() (*NotifyEventResponse, error) {
	m := new(NotifyEventResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VendingMachineServiceServer is the server API for VendingMachineService service.
// All implementations should embed UnimplementedVendingMachineServiceServer
// for forward compatibility
type VendingMachineServiceServer interface {
	GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error)
	ExecuteCommand(VendingMachineService_ExecuteCommandServer) error
	NotifyEvent(*NotifyEventRequest, VendingMachineService_NotifyEventServer) error
}

// UnimplementedVendingMachineServiceServer should be embedded to have forward compatible implementations.
type UnimplementedVendingMachineServiceServer struct{}

func (UnimplementedVendingMachineServiceServer) GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}

func (UnimplementedVendingMachineServiceServer) ExecuteCommand(VendingMachineService_ExecuteCommandServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecuteCommand not implemented")
}

func (UnimplementedVendingMachineServiceServer) NotifyEvent(*NotifyEventRequest, VendingMachineService_NotifyEventServer) error {
	return status.Errorf(codes.Unimplemented, "method NotifyEvent not implemented")
}

// UnsafeVendingMachineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VendingMachineServiceServer will
// result in compilation errors.
type UnsafeVendingMachineServiceServer interface {
	mustEmbedUnimplementedVendingMachineServiceServer()
}

func RegisterVendingMachineServiceServer(s grpc.ServiceRegistrar, srv VendingMachineServiceServer) {
	s.RegisterService(&VendingMachineService_ServiceDesc, srv)
}

func _VendingMachineService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendingMachineServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vending_machine.v1.VendingMachineService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendingMachineServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendingMachineService_ExecuteCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VendingMachineServiceServer).ExecuteCommand(&vendingMachineServiceExecuteCommandServer{stream})
}

type VendingMachineService_ExecuteCommandServer interface {
	Send(*ExecuteCommandResponse) error
	Recv() (*ExecuteCommandRequest, error)
	grpc.ServerStream
}

type vendingMachineServiceExecuteCommandServer struct {
	grpc.ServerStream
}

func (x *vendingMachineServiceExecuteCommandServer) Send(m *ExecuteCommandResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *vendingMachineServiceExecuteCommandServer) Recv() (*ExecuteCommandRequest, error) {
	m := new(ExecuteCommandRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _VendingMachineService_NotifyEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NotifyEventRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VendingMachineServiceServer).NotifyEvent(m, &vendingMachineServiceNotifyEventServer{stream})
}

type VendingMachineService_NotifyEventServer interface {
	Send(*NotifyEventResponse) error
	grpc.ServerStream
}

type vendingMachineServiceNotifyEventServer struct {
	grpc.ServerStream
}

func (x *vendingMachineServiceNotifyEventServer) Send(m *NotifyEventResponse) error {
	return x.ServerStream.SendMsg(m)
}

// VendingMachineService_ServiceDesc is the grpc.ServiceDesc for VendingMachineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VendingMachineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vending_machine.v1.VendingMachineService",
	HandlerType: (*VendingMachineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _VendingMachineService_GetProduct_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecuteCommand",
			Handler:       _VendingMachineService_ExecuteCommand_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "NotifyEvent",
			Handler:       _VendingMachineService_NotifyEvent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "vending_machine.proto",
}
