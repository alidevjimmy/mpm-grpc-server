// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package reportpb

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

// ReportServiceClient is the client API for ReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportServiceClient interface {
	// create new record perhaps by sensor or a person
	CreateReport(ctx context.Context, in *SensorReportRequest, opts ...grpc.CallOption) (*SensorReportResponse, error)
	// get report information by id
	GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*GetReportResponse, error)
	// get all uncompleted reports by server streaming technique
	GetUnCompletedReports(ctx context.Context, in *GetUnCompletedReportsRequest, opts ...grpc.CallOption) (ReportService_GetUnCompletedReportsClient, error)
	// get all logs which are relevant to specified report with "id"
	GetReportLogs(ctx context.Context, in *GetReportLogRequest, opts ...grpc.CallOption) (ReportService_GetReportLogsClient, error)
}

type reportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportServiceClient(cc grpc.ClientConnInterface) ReportServiceClient {
	return &reportServiceClient{cc}
}

func (c *reportServiceClient) CreateReport(ctx context.Context, in *SensorReportRequest, opts ...grpc.CallOption) (*SensorReportResponse, error) {
	out := new(SensorReportResponse)
	err := c.cc.Invoke(ctx, "/report.ReportService/CreateReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*GetReportResponse, error) {
	out := new(GetReportResponse)
	err := c.cc.Invoke(ctx, "/report.ReportService/GetReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) GetUnCompletedReports(ctx context.Context, in *GetUnCompletedReportsRequest, opts ...grpc.CallOption) (ReportService_GetUnCompletedReportsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReportService_ServiceDesc.Streams[0], "/report.ReportService/GetUnCompletedReports", opts...)
	if err != nil {
		return nil, err
	}
	x := &reportServiceGetUnCompletedReportsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReportService_GetUnCompletedReportsClient interface {
	Recv() (*GetUnCompletedReportsResponse, error)
	grpc.ClientStream
}

type reportServiceGetUnCompletedReportsClient struct {
	grpc.ClientStream
}

func (x *reportServiceGetUnCompletedReportsClient) Recv() (*GetUnCompletedReportsResponse, error) {
	m := new(GetUnCompletedReportsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *reportServiceClient) GetReportLogs(ctx context.Context, in *GetReportLogRequest, opts ...grpc.CallOption) (ReportService_GetReportLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReportService_ServiceDesc.Streams[1], "/report.ReportService/GetReportLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &reportServiceGetReportLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReportService_GetReportLogsClient interface {
	Recv() (*GetReportLogResponse, error)
	grpc.ClientStream
}

type reportServiceGetReportLogsClient struct {
	grpc.ClientStream
}

func (x *reportServiceGetReportLogsClient) Recv() (*GetReportLogResponse, error) {
	m := new(GetReportLogResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReportServiceServer is the server API for ReportService service.
// All implementations must embed UnimplementedReportServiceServer
// for forward compatibility
type ReportServiceServer interface {
	// create new record perhaps by sensor or a person
	CreateReport(context.Context, *SensorReportRequest) (*SensorReportResponse, error)
	// get report information by id
	GetReport(context.Context, *GetReportRequest) (*GetReportResponse, error)
	// get all uncompleted reports by server streaming technique
	GetUnCompletedReports(*GetUnCompletedReportsRequest, ReportService_GetUnCompletedReportsServer) error
	// get all logs which are relevant to specified report with "id"
	GetReportLogs(*GetReportLogRequest, ReportService_GetReportLogsServer) error
	mustEmbedUnimplementedReportServiceServer()
}

// UnimplementedReportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReportServiceServer struct {
}

func (UnimplementedReportServiceServer) CreateReport(context.Context, *SensorReportRequest) (*SensorReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReport not implemented")
}
func (UnimplementedReportServiceServer) GetReport(context.Context, *GetReportRequest) (*GetReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReport not implemented")
}
func (UnimplementedReportServiceServer) GetUnCompletedReports(*GetUnCompletedReportsRequest, ReportService_GetUnCompletedReportsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUnCompletedReports not implemented")
}
func (UnimplementedReportServiceServer) GetReportLogs(*GetReportLogRequest, ReportService_GetReportLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetReportLogs not implemented")
}
func (UnimplementedReportServiceServer) mustEmbedUnimplementedReportServiceServer() {}

// UnsafeReportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportServiceServer will
// result in compilation errors.
type UnsafeReportServiceServer interface {
	mustEmbedUnimplementedReportServiceServer()
}

func RegisterReportServiceServer(s grpc.ServiceRegistrar, srv ReportServiceServer) {
	s.RegisterService(&ReportService_ServiceDesc, srv)
}

func _ReportService_CreateReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SensorReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).CreateReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/report.ReportService/CreateReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).CreateReport(ctx, req.(*SensorReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportService_GetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).GetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/report.ReportService/GetReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).GetReport(ctx, req.(*GetReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportService_GetUnCompletedReports_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetUnCompletedReportsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReportServiceServer).GetUnCompletedReports(m, &reportServiceGetUnCompletedReportsServer{stream})
}

type ReportService_GetUnCompletedReportsServer interface {
	Send(*GetUnCompletedReportsResponse) error
	grpc.ServerStream
}

type reportServiceGetUnCompletedReportsServer struct {
	grpc.ServerStream
}

func (x *reportServiceGetUnCompletedReportsServer) Send(m *GetUnCompletedReportsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ReportService_GetReportLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetReportLogRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReportServiceServer).GetReportLogs(m, &reportServiceGetReportLogsServer{stream})
}

type ReportService_GetReportLogsServer interface {
	Send(*GetReportLogResponse) error
	grpc.ServerStream
}

type reportServiceGetReportLogsServer struct {
	grpc.ServerStream
}

func (x *reportServiceGetReportLogsServer) Send(m *GetReportLogResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ReportService_ServiceDesc is the grpc.ServiceDesc for ReportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "report.ReportService",
	HandlerType: (*ReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReport",
			Handler:    _ReportService_CreateReport_Handler,
		},
		{
			MethodName: "GetReport",
			Handler:    _ReportService_GetReport_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUnCompletedReports",
			Handler:       _ReportService_GetUnCompletedReports_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetReportLogs",
			Handler:       _ReportService_GetReportLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "src/pbs/reportpb/report.proto",
}
