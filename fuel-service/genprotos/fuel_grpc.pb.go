// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: protos/fuel.proto

package genprotos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	FuelService_CreateFuel_FullMethodName                = "/FuelService/CreateFuel"
	FuelService_GetFuel_FullMethodName                   = "/FuelService/GetFuel"
	FuelService_UpdateFuel_FullMethodName                = "/FuelService/UpdateFuel"
	FuelService_DeleteFuel_FullMethodName                = "/FuelService/DeleteFuel"
	FuelService_ListFuels_FullMethodName                 = "/FuelService/ListFuels"
	FuelService_CreateFuelHistory_FullMethodName         = "/FuelService/CreateFuelHistory"
	FuelService_GetFuelHistory_FullMethodName            = "/FuelService/GetFuelHistory"
	FuelService_ListFuelHistoriesByFuelID_FullMethodName = "/FuelService/ListFuelHistoriesByFuelID"
	FuelService_ListFuelHistories_FullMethodName         = "/FuelService/ListFuelHistories"
)

// FuelServiceClient is the client API for FuelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FuelServiceClient interface {
	// fuel_management
	CreateFuel(ctx context.Context, in *CreateFuelRequest, opts ...grpc.CallOption) (*FuelResponse, error)
	GetFuel(ctx context.Context, in *GetFuelRequest, opts ...grpc.CallOption) (*FuelResponse, error)
	UpdateFuel(ctx context.Context, in *UpdateFuelRequest, opts ...grpc.CallOption) (*FuelResponse, error)
	DeleteFuel(ctx context.Context, in *DeleteFuelRequest, opts ...grpc.CallOption) (*Empty, error)
	ListFuels(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListFuelsResponse, error)
	// fuel_history
	CreateFuelHistory(ctx context.Context, in *CreateFuelHistoryRequest, opts ...grpc.CallOption) (*FuelHistoryResponse, error)
	GetFuelHistory(ctx context.Context, in *GetFuelHistoryRequest, opts ...grpc.CallOption) (*FuelHistoryResponse, error)
	ListFuelHistoriesByFuelID(ctx context.Context, in *ListFuelHistoriesByFuelIDRequest, opts ...grpc.CallOption) (*ListFuelHistoriesByFuelIDResponse, error)
	ListFuelHistories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListFuelHistoriesResponse, error)
}

type fuelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFuelServiceClient(cc grpc.ClientConnInterface) FuelServiceClient {
	return &fuelServiceClient{cc}
}

func (c *fuelServiceClient) CreateFuel(ctx context.Context, in *CreateFuelRequest, opts ...grpc.CallOption) (*FuelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FuelResponse)
	err := c.cc.Invoke(ctx, FuelService_CreateFuel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fuelServiceClient) GetFuel(ctx context.Context, in *GetFuelRequest, opts ...grpc.CallOption) (*FuelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FuelResponse)
	err := c.cc.Invoke(ctx, FuelService_GetFuel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fuelServiceClient) UpdateFuel(ctx context.Context, in *UpdateFuelRequest, opts ...grpc.CallOption) (*FuelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FuelResponse)
	err := c.cc.Invoke(ctx, FuelService_UpdateFuel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fuelServiceClient) DeleteFuel(ctx context.Context, in *DeleteFuelRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, FuelService_DeleteFuel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fuelServiceClient) ListFuels(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListFuelsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListFuelsResponse)
	err := c.cc.Invoke(ctx, FuelService_ListFuels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fuelServiceClient) CreateFuelHistory(ctx context.Context, in *CreateFuelHistoryRequest, opts ...grpc.CallOption) (*FuelHistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FuelHistoryResponse)
	err := c.cc.Invoke(ctx, FuelService_CreateFuelHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fuelServiceClient) GetFuelHistory(ctx context.Context, in *GetFuelHistoryRequest, opts ...grpc.CallOption) (*FuelHistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FuelHistoryResponse)
	err := c.cc.Invoke(ctx, FuelService_GetFuelHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fuelServiceClient) ListFuelHistoriesByFuelID(ctx context.Context, in *ListFuelHistoriesByFuelIDRequest, opts ...grpc.CallOption) (*ListFuelHistoriesByFuelIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListFuelHistoriesByFuelIDResponse)
	err := c.cc.Invoke(ctx, FuelService_ListFuelHistoriesByFuelID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fuelServiceClient) ListFuelHistories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListFuelHistoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListFuelHistoriesResponse)
	err := c.cc.Invoke(ctx, FuelService_ListFuelHistories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FuelServiceServer is the server API for FuelService service.
// All implementations must embed UnimplementedFuelServiceServer
// for forward compatibility
type FuelServiceServer interface {
	// fuel_management
	CreateFuel(context.Context, *CreateFuelRequest) (*FuelResponse, error)
	GetFuel(context.Context, *GetFuelRequest) (*FuelResponse, error)
	UpdateFuel(context.Context, *UpdateFuelRequest) (*FuelResponse, error)
	DeleteFuel(context.Context, *DeleteFuelRequest) (*Empty, error)
	ListFuels(context.Context, *Empty) (*ListFuelsResponse, error)
	// fuel_history
	CreateFuelHistory(context.Context, *CreateFuelHistoryRequest) (*FuelHistoryResponse, error)
	GetFuelHistory(context.Context, *GetFuelHistoryRequest) (*FuelHistoryResponse, error)
	ListFuelHistoriesByFuelID(context.Context, *ListFuelHistoriesByFuelIDRequest) (*ListFuelHistoriesByFuelIDResponse, error)
	ListFuelHistories(context.Context, *Empty) (*ListFuelHistoriesResponse, error)
	mustEmbedUnimplementedFuelServiceServer()
}

// UnimplementedFuelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFuelServiceServer struct {
}

func (UnimplementedFuelServiceServer) CreateFuel(context.Context, *CreateFuelRequest) (*FuelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFuel not implemented")
}
func (UnimplementedFuelServiceServer) GetFuel(context.Context, *GetFuelRequest) (*FuelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFuel not implemented")
}
func (UnimplementedFuelServiceServer) UpdateFuel(context.Context, *UpdateFuelRequest) (*FuelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFuel not implemented")
}
func (UnimplementedFuelServiceServer) DeleteFuel(context.Context, *DeleteFuelRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFuel not implemented")
}
func (UnimplementedFuelServiceServer) ListFuels(context.Context, *Empty) (*ListFuelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFuels not implemented")
}
func (UnimplementedFuelServiceServer) CreateFuelHistory(context.Context, *CreateFuelHistoryRequest) (*FuelHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFuelHistory not implemented")
}
func (UnimplementedFuelServiceServer) GetFuelHistory(context.Context, *GetFuelHistoryRequest) (*FuelHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFuelHistory not implemented")
}
func (UnimplementedFuelServiceServer) ListFuelHistoriesByFuelID(context.Context, *ListFuelHistoriesByFuelIDRequest) (*ListFuelHistoriesByFuelIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFuelHistoriesByFuelID not implemented")
}
func (UnimplementedFuelServiceServer) ListFuelHistories(context.Context, *Empty) (*ListFuelHistoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFuelHistories not implemented")
}
func (UnimplementedFuelServiceServer) mustEmbedUnimplementedFuelServiceServer() {}

// UnsafeFuelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FuelServiceServer will
// result in compilation errors.
type UnsafeFuelServiceServer interface {
	mustEmbedUnimplementedFuelServiceServer()
}

func RegisterFuelServiceServer(s grpc.ServiceRegistrar, srv FuelServiceServer) {
	s.RegisterService(&FuelService_ServiceDesc, srv)
}

func _FuelService_CreateFuel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFuelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuelServiceServer).CreateFuel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FuelService_CreateFuel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuelServiceServer).CreateFuel(ctx, req.(*CreateFuelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FuelService_GetFuel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFuelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuelServiceServer).GetFuel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FuelService_GetFuel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuelServiceServer).GetFuel(ctx, req.(*GetFuelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FuelService_UpdateFuel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFuelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuelServiceServer).UpdateFuel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FuelService_UpdateFuel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuelServiceServer).UpdateFuel(ctx, req.(*UpdateFuelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FuelService_DeleteFuel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFuelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuelServiceServer).DeleteFuel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FuelService_DeleteFuel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuelServiceServer).DeleteFuel(ctx, req.(*DeleteFuelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FuelService_ListFuels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuelServiceServer).ListFuels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FuelService_ListFuels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuelServiceServer).ListFuels(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FuelService_CreateFuelHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFuelHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuelServiceServer).CreateFuelHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FuelService_CreateFuelHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuelServiceServer).CreateFuelHistory(ctx, req.(*CreateFuelHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FuelService_GetFuelHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFuelHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuelServiceServer).GetFuelHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FuelService_GetFuelHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuelServiceServer).GetFuelHistory(ctx, req.(*GetFuelHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FuelService_ListFuelHistoriesByFuelID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFuelHistoriesByFuelIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuelServiceServer).ListFuelHistoriesByFuelID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FuelService_ListFuelHistoriesByFuelID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuelServiceServer).ListFuelHistoriesByFuelID(ctx, req.(*ListFuelHistoriesByFuelIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FuelService_ListFuelHistories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuelServiceServer).ListFuelHistories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FuelService_ListFuelHistories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuelServiceServer).ListFuelHistories(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// FuelService_ServiceDesc is the grpc.ServiceDesc for FuelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FuelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FuelService",
	HandlerType: (*FuelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFuel",
			Handler:    _FuelService_CreateFuel_Handler,
		},
		{
			MethodName: "GetFuel",
			Handler:    _FuelService_GetFuel_Handler,
		},
		{
			MethodName: "UpdateFuel",
			Handler:    _FuelService_UpdateFuel_Handler,
		},
		{
			MethodName: "DeleteFuel",
			Handler:    _FuelService_DeleteFuel_Handler,
		},
		{
			MethodName: "ListFuels",
			Handler:    _FuelService_ListFuels_Handler,
		},
		{
			MethodName: "CreateFuelHistory",
			Handler:    _FuelService_CreateFuelHistory_Handler,
		},
		{
			MethodName: "GetFuelHistory",
			Handler:    _FuelService_GetFuelHistory_Handler,
		},
		{
			MethodName: "ListFuelHistoriesByFuelID",
			Handler:    _FuelService_ListFuelHistoriesByFuelID_Handler,
		},
		{
			MethodName: "ListFuelHistories",
			Handler:    _FuelService_ListFuelHistories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/fuel.proto",
}
