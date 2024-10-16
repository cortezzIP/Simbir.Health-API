// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: timetable-service/timetable.proto

package timetable

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
	Timetable_CreateTimetableRecord_FullMethodName               = "/timetable.Timetable/CreateTimetableRecord"
	Timetable_CreateAppointmentRecord_FullMethodName             = "/timetable.Timetable/CreateAppointmentRecord"
	Timetable_GetTimetableByHospitalId_FullMethodName            = "/timetable.Timetable/GetTimetableByHospitalId"
	Timetable_GetTimetableByDoctorId_FullMethodName              = "/timetable.Timetable/GetTimetableByDoctorId"
	Timetable_GetTimetableByHospitalRoom_FullMethodName          = "/timetable.Timetable/GetTimetableByHospitalRoom"
	Timetable_GetFreeCouponsForDoctorsAppointment_FullMethodName = "/timetable.Timetable/GetFreeCouponsForDoctorsAppointment"
	Timetable_UpdateTimetableRecord_FullMethodName               = "/timetable.Timetable/UpdateTimetableRecord"
	Timetable_DeleteTimetableRecord_FullMethodName               = "/timetable.Timetable/DeleteTimetableRecord"
	Timetable_DeleteTimetableRecordByDoctorId_FullMethodName     = "/timetable.Timetable/DeleteTimetableRecordByDoctorId"
	Timetable_DeleteTimetableRecordByHospitalId_FullMethodName   = "/timetable.Timetable/DeleteTimetableRecordByHospitalId"
	Timetable_DeleteAppointmentRecord_FullMethodName             = "/timetable.Timetable/DeleteAppointmentRecord"
)

// TimetableClient is the client API for Timetable service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimetableClient interface {
	CreateTimetableRecord(ctx context.Context, in *CreateTimetableRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateAppointmentRecord(ctx context.Context, in *CreateAppointmentRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTimetableByHospitalId(ctx context.Context, in *GetTimetableByHospitalIdRequest, opts ...grpc.CallOption) (*GetTimetableByHospitalIdResponse, error)
	GetTimetableByDoctorId(ctx context.Context, in *GetTimetableByDoctorIdRequest, opts ...grpc.CallOption) (*GetTimetableByDoctorIdResponse, error)
	GetTimetableByHospitalRoom(ctx context.Context, in *GetTimetableByHospitalRoomRequest, opts ...grpc.CallOption) (*GetTimetableByHospitalRoomResponse, error)
	GetFreeCouponsForDoctorsAppointment(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetFreeCouponsForDoctorsAppointmentResponse, error)
	UpdateTimetableRecord(ctx context.Context, in *UpdateTimetableRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteTimetableRecord(ctx context.Context, in *DeleteTimetableRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteTimetableRecordByDoctorId(ctx context.Context, in *DeleteTimetableRecordByDoctorIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteTimetableRecordByHospitalId(ctx context.Context, in *DeleteTimetableRecordByHospitalIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteAppointmentRecord(ctx context.Context, in *DeleteAppointmentRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type timetableClient struct {
	cc grpc.ClientConnInterface
}

func NewTimetableClient(cc grpc.ClientConnInterface) TimetableClient {
	return &timetableClient{cc}
}

func (c *timetableClient) CreateTimetableRecord(ctx context.Context, in *CreateTimetableRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Timetable_CreateTimetableRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timetableClient) CreateAppointmentRecord(ctx context.Context, in *CreateAppointmentRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Timetable_CreateAppointmentRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timetableClient) GetTimetableByHospitalId(ctx context.Context, in *GetTimetableByHospitalIdRequest, opts ...grpc.CallOption) (*GetTimetableByHospitalIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTimetableByHospitalIdResponse)
	err := c.cc.Invoke(ctx, Timetable_GetTimetableByHospitalId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timetableClient) GetTimetableByDoctorId(ctx context.Context, in *GetTimetableByDoctorIdRequest, opts ...grpc.CallOption) (*GetTimetableByDoctorIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTimetableByDoctorIdResponse)
	err := c.cc.Invoke(ctx, Timetable_GetTimetableByDoctorId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timetableClient) GetTimetableByHospitalRoom(ctx context.Context, in *GetTimetableByHospitalRoomRequest, opts ...grpc.CallOption) (*GetTimetableByHospitalRoomResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTimetableByHospitalRoomResponse)
	err := c.cc.Invoke(ctx, Timetable_GetTimetableByHospitalRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timetableClient) GetFreeCouponsForDoctorsAppointment(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetFreeCouponsForDoctorsAppointmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFreeCouponsForDoctorsAppointmentResponse)
	err := c.cc.Invoke(ctx, Timetable_GetFreeCouponsForDoctorsAppointment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timetableClient) UpdateTimetableRecord(ctx context.Context, in *UpdateTimetableRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Timetable_UpdateTimetableRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timetableClient) DeleteTimetableRecord(ctx context.Context, in *DeleteTimetableRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Timetable_DeleteTimetableRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timetableClient) DeleteTimetableRecordByDoctorId(ctx context.Context, in *DeleteTimetableRecordByDoctorIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Timetable_DeleteTimetableRecordByDoctorId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timetableClient) DeleteTimetableRecordByHospitalId(ctx context.Context, in *DeleteTimetableRecordByHospitalIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Timetable_DeleteTimetableRecordByHospitalId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timetableClient) DeleteAppointmentRecord(ctx context.Context, in *DeleteAppointmentRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Timetable_DeleteAppointmentRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimetableServer is the server API for Timetable service.
// All implementations must embed UnimplementedTimetableServer
// for forward compatibility.
type TimetableServer interface {
	CreateTimetableRecord(context.Context, *CreateTimetableRecordRequest) (*emptypb.Empty, error)
	CreateAppointmentRecord(context.Context, *CreateAppointmentRecordRequest) (*emptypb.Empty, error)
	GetTimetableByHospitalId(context.Context, *GetTimetableByHospitalIdRequest) (*GetTimetableByHospitalIdResponse, error)
	GetTimetableByDoctorId(context.Context, *GetTimetableByDoctorIdRequest) (*GetTimetableByDoctorIdResponse, error)
	GetTimetableByHospitalRoom(context.Context, *GetTimetableByHospitalRoomRequest) (*GetTimetableByHospitalRoomResponse, error)
	GetFreeCouponsForDoctorsAppointment(context.Context, *emptypb.Empty) (*GetFreeCouponsForDoctorsAppointmentResponse, error)
	UpdateTimetableRecord(context.Context, *UpdateTimetableRecordRequest) (*emptypb.Empty, error)
	DeleteTimetableRecord(context.Context, *DeleteTimetableRecordRequest) (*emptypb.Empty, error)
	DeleteTimetableRecordByDoctorId(context.Context, *DeleteTimetableRecordByDoctorIdRequest) (*emptypb.Empty, error)
	DeleteTimetableRecordByHospitalId(context.Context, *DeleteTimetableRecordByHospitalIdRequest) (*emptypb.Empty, error)
	DeleteAppointmentRecord(context.Context, *DeleteAppointmentRecordRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTimetableServer()
}

// UnimplementedTimetableServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTimetableServer struct{}

func (UnimplementedTimetableServer) CreateTimetableRecord(context.Context, *CreateTimetableRecordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTimetableRecord not implemented")
}
func (UnimplementedTimetableServer) CreateAppointmentRecord(context.Context, *CreateAppointmentRecordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppointmentRecord not implemented")
}
func (UnimplementedTimetableServer) GetTimetableByHospitalId(context.Context, *GetTimetableByHospitalIdRequest) (*GetTimetableByHospitalIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimetableByHospitalId not implemented")
}
func (UnimplementedTimetableServer) GetTimetableByDoctorId(context.Context, *GetTimetableByDoctorIdRequest) (*GetTimetableByDoctorIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimetableByDoctorId not implemented")
}
func (UnimplementedTimetableServer) GetTimetableByHospitalRoom(context.Context, *GetTimetableByHospitalRoomRequest) (*GetTimetableByHospitalRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimetableByHospitalRoom not implemented")
}
func (UnimplementedTimetableServer) GetFreeCouponsForDoctorsAppointment(context.Context, *emptypb.Empty) (*GetFreeCouponsForDoctorsAppointmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFreeCouponsForDoctorsAppointment not implemented")
}
func (UnimplementedTimetableServer) UpdateTimetableRecord(context.Context, *UpdateTimetableRecordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTimetableRecord not implemented")
}
func (UnimplementedTimetableServer) DeleteTimetableRecord(context.Context, *DeleteTimetableRecordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTimetableRecord not implemented")
}
func (UnimplementedTimetableServer) DeleteTimetableRecordByDoctorId(context.Context, *DeleteTimetableRecordByDoctorIdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTimetableRecordByDoctorId not implemented")
}
func (UnimplementedTimetableServer) DeleteTimetableRecordByHospitalId(context.Context, *DeleteTimetableRecordByHospitalIdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTimetableRecordByHospitalId not implemented")
}
func (UnimplementedTimetableServer) DeleteAppointmentRecord(context.Context, *DeleteAppointmentRecordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppointmentRecord not implemented")
}
func (UnimplementedTimetableServer) mustEmbedUnimplementedTimetableServer() {}
func (UnimplementedTimetableServer) testEmbeddedByValue()                   {}

// UnsafeTimetableServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimetableServer will
// result in compilation errors.
type UnsafeTimetableServer interface {
	mustEmbedUnimplementedTimetableServer()
}

func RegisterTimetableServer(s grpc.ServiceRegistrar, srv TimetableServer) {
	// If the following call pancis, it indicates UnimplementedTimetableServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Timetable_ServiceDesc, srv)
}

func _Timetable_CreateTimetableRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTimetableRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimetableServer).CreateTimetableRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Timetable_CreateTimetableRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimetableServer).CreateTimetableRecord(ctx, req.(*CreateTimetableRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Timetable_CreateAppointmentRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppointmentRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimetableServer).CreateAppointmentRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Timetable_CreateAppointmentRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimetableServer).CreateAppointmentRecord(ctx, req.(*CreateAppointmentRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Timetable_GetTimetableByHospitalId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimetableByHospitalIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimetableServer).GetTimetableByHospitalId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Timetable_GetTimetableByHospitalId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimetableServer).GetTimetableByHospitalId(ctx, req.(*GetTimetableByHospitalIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Timetable_GetTimetableByDoctorId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimetableByDoctorIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimetableServer).GetTimetableByDoctorId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Timetable_GetTimetableByDoctorId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimetableServer).GetTimetableByDoctorId(ctx, req.(*GetTimetableByDoctorIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Timetable_GetTimetableByHospitalRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimetableByHospitalRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimetableServer).GetTimetableByHospitalRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Timetable_GetTimetableByHospitalRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimetableServer).GetTimetableByHospitalRoom(ctx, req.(*GetTimetableByHospitalRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Timetable_GetFreeCouponsForDoctorsAppointment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimetableServer).GetFreeCouponsForDoctorsAppointment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Timetable_GetFreeCouponsForDoctorsAppointment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimetableServer).GetFreeCouponsForDoctorsAppointment(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Timetable_UpdateTimetableRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTimetableRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimetableServer).UpdateTimetableRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Timetable_UpdateTimetableRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimetableServer).UpdateTimetableRecord(ctx, req.(*UpdateTimetableRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Timetable_DeleteTimetableRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTimetableRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimetableServer).DeleteTimetableRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Timetable_DeleteTimetableRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimetableServer).DeleteTimetableRecord(ctx, req.(*DeleteTimetableRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Timetable_DeleteTimetableRecordByDoctorId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTimetableRecordByDoctorIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimetableServer).DeleteTimetableRecordByDoctorId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Timetable_DeleteTimetableRecordByDoctorId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimetableServer).DeleteTimetableRecordByDoctorId(ctx, req.(*DeleteTimetableRecordByDoctorIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Timetable_DeleteTimetableRecordByHospitalId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTimetableRecordByHospitalIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimetableServer).DeleteTimetableRecordByHospitalId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Timetable_DeleteTimetableRecordByHospitalId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimetableServer).DeleteTimetableRecordByHospitalId(ctx, req.(*DeleteTimetableRecordByHospitalIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Timetable_DeleteAppointmentRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppointmentRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimetableServer).DeleteAppointmentRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Timetable_DeleteAppointmentRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimetableServer).DeleteAppointmentRecord(ctx, req.(*DeleteAppointmentRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Timetable_ServiceDesc is the grpc.ServiceDesc for Timetable service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Timetable_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "timetable.Timetable",
	HandlerType: (*TimetableServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTimetableRecord",
			Handler:    _Timetable_CreateTimetableRecord_Handler,
		},
		{
			MethodName: "CreateAppointmentRecord",
			Handler:    _Timetable_CreateAppointmentRecord_Handler,
		},
		{
			MethodName: "GetTimetableByHospitalId",
			Handler:    _Timetable_GetTimetableByHospitalId_Handler,
		},
		{
			MethodName: "GetTimetableByDoctorId",
			Handler:    _Timetable_GetTimetableByDoctorId_Handler,
		},
		{
			MethodName: "GetTimetableByHospitalRoom",
			Handler:    _Timetable_GetTimetableByHospitalRoom_Handler,
		},
		{
			MethodName: "GetFreeCouponsForDoctorsAppointment",
			Handler:    _Timetable_GetFreeCouponsForDoctorsAppointment_Handler,
		},
		{
			MethodName: "UpdateTimetableRecord",
			Handler:    _Timetable_UpdateTimetableRecord_Handler,
		},
		{
			MethodName: "DeleteTimetableRecord",
			Handler:    _Timetable_DeleteTimetableRecord_Handler,
		},
		{
			MethodName: "DeleteTimetableRecordByDoctorId",
			Handler:    _Timetable_DeleteTimetableRecordByDoctorId_Handler,
		},
		{
			MethodName: "DeleteTimetableRecordByHospitalId",
			Handler:    _Timetable_DeleteTimetableRecordByHospitalId_Handler,
		},
		{
			MethodName: "DeleteAppointmentRecord",
			Handler:    _Timetable_DeleteAppointmentRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "timetable-service/timetable.proto",
}
