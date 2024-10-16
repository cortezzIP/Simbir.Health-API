// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: hospital-service/hospital.proto

package hospital

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Hospital_CreateHospital_FullMethodName   = "/hospital.Hospital/CreateHospital"
	Hospital_GetHospitals_FullMethodName     = "/hospital.Hospital/GetHospitals"
	Hospital_GetHospitalById_FullMethodName  = "/hospital.Hospital/GetHospitalById"
	Hospital_GetHospitalRooms_FullMethodName = "/hospital.Hospital/GetHospitalRooms"
	Hospital_UpdateHospital_FullMethodName   = "/hospital.Hospital/UpdateHospital"
	Hospital_DeleteHospital_FullMethodName   = "/hospital.Hospital/DeleteHospital"
)

// HospitalClient is the client API for Hospital service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HospitalClient interface {
	CreateHospital(ctx context.Context, in *CreateHospitalRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetHospitals(ctx context.Context, in *GetAllHospitalsRequest, opts ...grpc.CallOption) (*GetAllHospitalsResponse, error)
	GetHospitalById(ctx context.Context, in *GetHospitalByIdRequest, opts ...grpc.CallOption) (*GetHospitalByIdResponse, error)
	GetHospitalRooms(ctx context.Context, in *GetHospitalsRoomsRequest, opts ...grpc.CallOption) (*GetHospitalsRoomsResponse, error)
	UpdateHospital(ctx context.Context, in *UpdateHospitalRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteHospital(ctx context.Context, in *DeleteHospitalRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type hospitalClient struct {
	cc grpc.ClientConnInterface
}

func NewHospitalClient(cc grpc.ClientConnInterface) HospitalClient {
	return &hospitalClient{cc}
}

func (c *hospitalClient) CreateHospital(ctx context.Context, in *CreateHospitalRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Hospital_CreateHospital_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalClient) GetHospitals(ctx context.Context, in *GetAllHospitalsRequest, opts ...grpc.CallOption) (*GetAllHospitalsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllHospitalsResponse)
	err := c.cc.Invoke(ctx, Hospital_GetHospitals_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalClient) GetHospitalById(ctx context.Context, in *GetHospitalByIdRequest, opts ...grpc.CallOption) (*GetHospitalByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHospitalByIdResponse)
	err := c.cc.Invoke(ctx, Hospital_GetHospitalById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalClient) GetHospitalRooms(ctx context.Context, in *GetHospitalsRoomsRequest, opts ...grpc.CallOption) (*GetHospitalsRoomsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHospitalsRoomsResponse)
	err := c.cc.Invoke(ctx, Hospital_GetHospitalRooms_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalClient) UpdateHospital(ctx context.Context, in *UpdateHospitalRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Hospital_UpdateHospital_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalClient) DeleteHospital(ctx context.Context, in *DeleteHospitalRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Hospital_DeleteHospital_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HospitalServer is the server API for Hospital service.
// All implementations must embed UnimplementedHospitalServer
// for forward compatibility.
type HospitalServer interface {
	CreateHospital(context.Context, *CreateHospitalRequest) (*emptypb.Empty, error)
	GetHospitals(context.Context, *GetAllHospitalsRequest) (*GetAllHospitalsResponse, error)
	GetHospitalById(context.Context, *GetHospitalByIdRequest) (*GetHospitalByIdResponse, error)
	GetHospitalRooms(context.Context, *GetHospitalsRoomsRequest) (*GetHospitalsRoomsResponse, error)
	UpdateHospital(context.Context, *UpdateHospitalRequest) (*emptypb.Empty, error)
	DeleteHospital(context.Context, *DeleteHospitalRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedHospitalServer()
}

// UnimplementedHospitalServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHospitalServer struct{}

func (UnimplementedHospitalServer) CreateHospital(context.Context, *CreateHospitalRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHospital not implemented")
}
func (UnimplementedHospitalServer) GetHospitals(context.Context, *GetAllHospitalsRequest) (*GetAllHospitalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHospitals not implemented")
}
func (UnimplementedHospitalServer) GetHospitalById(context.Context, *GetHospitalByIdRequest) (*GetHospitalByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHospitalById not implemented")
}
func (UnimplementedHospitalServer) GetHospitalRooms(context.Context, *GetHospitalsRoomsRequest) (*GetHospitalsRoomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHospitalRooms not implemented")
}
func (UnimplementedHospitalServer) UpdateHospital(context.Context, *UpdateHospitalRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHospital not implemented")
}
func (UnimplementedHospitalServer) DeleteHospital(context.Context, *DeleteHospitalRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHospital not implemented")
}
func (UnimplementedHospitalServer) mustEmbedUnimplementedHospitalServer() {}
func (UnimplementedHospitalServer) testEmbeddedByValue()                  {}

// UnsafeHospitalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HospitalServer will
// result in compilation errors.
type UnsafeHospitalServer interface {
	mustEmbedUnimplementedHospitalServer()
}

func RegisterHospitalServer(s grpc.ServiceRegistrar, srv HospitalServer) {
	// If the following call pancis, it indicates UnimplementedHospitalServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Hospital_ServiceDesc, srv)
}

func _Hospital_CreateHospital_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHospitalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServer).CreateHospital(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hospital_CreateHospital_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServer).CreateHospital(ctx, req.(*CreateHospitalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hospital_GetHospitals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllHospitalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServer).GetHospitals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hospital_GetHospitals_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServer).GetHospitals(ctx, req.(*GetAllHospitalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hospital_GetHospitalById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHospitalByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServer).GetHospitalById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hospital_GetHospitalById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServer).GetHospitalById(ctx, req.(*GetHospitalByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hospital_GetHospitalRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHospitalsRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServer).GetHospitalRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hospital_GetHospitalRooms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServer).GetHospitalRooms(ctx, req.(*GetHospitalsRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hospital_UpdateHospital_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHospitalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServer).UpdateHospital(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hospital_UpdateHospital_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServer).UpdateHospital(ctx, req.(*UpdateHospitalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hospital_DeleteHospital_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHospitalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServer).DeleteHospital(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hospital_DeleteHospital_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServer).DeleteHospital(ctx, req.(*DeleteHospitalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hospital_ServiceDesc is the grpc.ServiceDesc for Hospital service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hospital_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hospital.Hospital",
	HandlerType: (*HospitalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHospital",
			Handler:    _Hospital_CreateHospital_Handler,
		},
		{
			MethodName: "GetHospitals",
			Handler:    _Hospital_GetHospitals_Handler,
		},
		{
			MethodName: "GetHospitalById",
			Handler:    _Hospital_GetHospitalById_Handler,
		},
		{
			MethodName: "GetHospitalRooms",
			Handler:    _Hospital_GetHospitalRooms_Handler,
		},
		{
			MethodName: "UpdateHospital",
			Handler:    _Hospital_UpdateHospital_Handler,
		},
		{
			MethodName: "DeleteHospital",
			Handler:    _Hospital_DeleteHospital_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hospital-service/hospital.proto",
}
