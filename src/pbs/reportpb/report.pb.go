// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: src/pbs/reportpb/report.proto

package reportpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Report_Status int32

const (
	Report_Default Report_Status = 0
	Report_LOW     Report_Status = 1
	Report_MEDIUM  Report_Status = 2
	Report_HIGH    Report_Status = 3
)

// Enum value maps for Report_Status.
var (
	Report_Status_name = map[int32]string{
		0: "Default",
		1: "LOW",
		2: "MEDIUM",
		3: "HIGH",
	}
	Report_Status_value = map[string]int32{
		"Default": 0,
		"LOW":     1,
		"MEDIUM":  2,
		"HIGH":    3,
	}
)

func (x Report_Status) Enum() *Report_Status {
	p := new(Report_Status)
	*p = x
	return p
}

func (x Report_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Report_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_src_pbs_reportpb_report_proto_enumTypes[0].Descriptor()
}

func (Report_Status) Type() protoreflect.EnumType {
	return &file_src_pbs_reportpb_report_proto_enumTypes[0]
}

func (x Report_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Report_Status.Descriptor instead.
func (Report_Status) EnumDescriptor() ([]byte, []int) {
	return file_src_pbs_reportpb_report_proto_rawDescGZIP(), []int{0, 0}
}

type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FaultId  int32  `protobuf:"varint,2,opt,name=fault_id,json=faultId,proto3" json:"fault_id,omitempty"`
	SensorId int32  `protobuf:"varint,3,opt,name=sensor_id,json=sensorId,proto3" json:"sensor_id,omitempty"`
	// if not specified => field will be LOW (1)
	Status Report_Status `protobuf:"varint,4,opt,name=status,proto3,enum=report.Report_Status" json:"status,omitempty"`
	Tags   []string      `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	// if sensor is creator of the report then user_id will be 0
	// otherwise user_id has value
	UserId int32 `protobuf:"varint,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// these feilds are not needed for requests
	CreatedAt string     `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string     `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Commands  []*Command `protobuf:"bytes,9,rep,name=commands,proto3" json:"commands,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pbs_reportpb_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_src_pbs_reportpb_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_src_pbs_reportpb_report_proto_rawDescGZIP(), []int{0}
}

func (x *Report) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Report) GetFaultId() int32 {
	if x != nil {
		return x.FaultId
	}
	return 0
}

func (x *Report) GetSensorId() int32 {
	if x != nil {
		return x.SensorId
	}
	return 0
}

func (x *Report) GetStatus() Report_Status {
	if x != nil {
		return x.Status
	}
	return Report_Default
}

func (x *Report) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Report) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Report) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Report) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Report) GetCommands() []*Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

type ReportLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CREATE, UPDATE, DELETE
	Operation   string `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	UserId      int32  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ReportId    string `protobuf:"bytes,3,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
	SendsorId   int32  `protobuf:"varint,4,opt,name=sendsor_id,json=sendsorId,proto3" json:"sendsor_id,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// not needed for requests
	CreatedAt string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *ReportLog) Reset() {
	*x = ReportLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pbs_reportpb_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportLog) ProtoMessage() {}

func (x *ReportLog) ProtoReflect() protoreflect.Message {
	mi := &file_src_pbs_reportpb_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportLog.ProtoReflect.Descriptor instead.
func (*ReportLog) Descriptor() ([]byte, []int) {
	return file_src_pbs_reportpb_report_proto_rawDescGZIP(), []int{1}
}

func (x *ReportLog) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *ReportLog) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ReportLog) GetReportId() string {
	if x != nil {
		return x.ReportId
	}
	return ""
}

func (x *ReportLog) GetSendsorId() int32 {
	if x != nil {
		return x.SendsorId
	}
	return 0
}

func (x *ReportLog) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ReportLog) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandId int32 `protobuf:"varint,1,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	// FIRE, ELECTICY, ...
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// is manual or auto
	Auto bool `protobuf:"varint,3,opt,name=Auto,proto3" json:"Auto,omitempty"`
	Done bool `protobuf:"varint,4,opt,name=Done,proto3" json:"Done,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pbs_reportpb_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_src_pbs_reportpb_report_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_src_pbs_reportpb_report_proto_rawDescGZIP(), []int{2}
}

func (x *Command) GetCommandId() int32 {
	if x != nil {
		return x.CommandId
	}
	return 0
}

func (x *Command) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Command) GetAuto() bool {
	if x != nil {
		return x.Auto
	}
	return false
}

func (x *Command) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type SensorReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report *Report `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *SensorReportRequest) Reset() {
	*x = SensorReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pbs_reportpb_report_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SensorReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SensorReportRequest) ProtoMessage() {}

func (x *SensorReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_pbs_reportpb_report_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SensorReportRequest.ProtoReflect.Descriptor instead.
func (*SensorReportRequest) Descriptor() ([]byte, []int) {
	return file_src_pbs_reportpb_report_proto_rawDescGZIP(), []int{3}
}

func (x *SensorReportRequest) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

type SensorReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report *Report `protobuf:"bytes,1,opt,name=Report,proto3" json:"Report,omitempty"`
}

func (x *SensorReportResponse) Reset() {
	*x = SensorReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pbs_reportpb_report_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SensorReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SensorReportResponse) ProtoMessage() {}

func (x *SensorReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_pbs_reportpb_report_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SensorReportResponse.ProtoReflect.Descriptor instead.
func (*SensorReportResponse) Descriptor() ([]byte, []int) {
	return file_src_pbs_reportpb_report_proto_rawDescGZIP(), []int{4}
}

func (x *SensorReportResponse) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

type GetReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportId string `protobuf:"bytes,1,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
}

func (x *GetReportRequest) Reset() {
	*x = GetReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pbs_reportpb_report_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportRequest) ProtoMessage() {}

func (x *GetReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_pbs_reportpb_report_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportRequest.ProtoReflect.Descriptor instead.
func (*GetReportRequest) Descriptor() ([]byte, []int) {
	return file_src_pbs_reportpb_report_proto_rawDescGZIP(), []int{5}
}

func (x *GetReportRequest) GetReportId() string {
	if x != nil {
		return x.ReportId
	}
	return ""
}

type GetReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report *Report `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *GetReportResponse) Reset() {
	*x = GetReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pbs_reportpb_report_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportResponse) ProtoMessage() {}

func (x *GetReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_pbs_reportpb_report_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportResponse.ProtoReflect.Descriptor instead.
func (*GetReportResponse) Descriptor() ([]byte, []int) {
	return file_src_pbs_reportpb_report_proto_rawDescGZIP(), []int{6}
}

func (x *GetReportResponse) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

type GetUnCompletedReportsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUnCompletedReportsRequest) Reset() {
	*x = GetUnCompletedReportsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pbs_reportpb_report_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUnCompletedReportsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnCompletedReportsRequest) ProtoMessage() {}

func (x *GetUnCompletedReportsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_pbs_reportpb_report_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnCompletedReportsRequest.ProtoReflect.Descriptor instead.
func (*GetUnCompletedReportsRequest) Descriptor() ([]byte, []int) {
	return file_src_pbs_reportpb_report_proto_rawDescGZIP(), []int{7}
}

type GetUnCompletedReportsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report []*Report `protobuf:"bytes,1,rep,name=report,proto3" json:"report,omitempty"`
}

func (x *GetUnCompletedReportsResponse) Reset() {
	*x = GetUnCompletedReportsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pbs_reportpb_report_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUnCompletedReportsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnCompletedReportsResponse) ProtoMessage() {}

func (x *GetUnCompletedReportsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_pbs_reportpb_report_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnCompletedReportsResponse.ProtoReflect.Descriptor instead.
func (*GetUnCompletedReportsResponse) Descriptor() ([]byte, []int) {
	return file_src_pbs_reportpb_report_proto_rawDescGZIP(), []int{8}
}

func (x *GetUnCompletedReportsResponse) GetReport() []*Report {
	if x != nil {
		return x.Report
	}
	return nil
}

type GetReportLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportId string `protobuf:"bytes,1,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
}

func (x *GetReportLogRequest) Reset() {
	*x = GetReportLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pbs_reportpb_report_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReportLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportLogRequest) ProtoMessage() {}

func (x *GetReportLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_pbs_reportpb_report_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportLogRequest.ProtoReflect.Descriptor instead.
func (*GetReportLogRequest) Descriptor() ([]byte, []int) {
	return file_src_pbs_reportpb_report_proto_rawDescGZIP(), []int{9}
}

func (x *GetReportLogRequest) GetReportId() string {
	if x != nil {
		return x.ReportId
	}
	return ""
}

type GetReportLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportLog []*ReportLog `protobuf:"bytes,1,rep,name=report_log,json=reportLog,proto3" json:"report_log,omitempty"`
}

func (x *GetReportLogResponse) Reset() {
	*x = GetReportLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_pbs_reportpb_report_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReportLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportLogResponse) ProtoMessage() {}

func (x *GetReportLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_pbs_reportpb_report_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportLogResponse.ProtoReflect.Descriptor instead.
func (*GetReportLogResponse) Descriptor() ([]byte, []int) {
	return file_src_pbs_reportpb_report_proto_rawDescGZIP(), []int{10}
}

func (x *GetReportLogResponse) GetReportLog() []*ReportLog {
	if x != nil {
		return x.ReportLog
	}
	return nil
}

var File_src_pbs_reportpb_report_proto protoreflect.FileDescriptor

var file_src_pbs_reportpb_report_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x73, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x70, 0x62, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0xcd, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x22, 0x34, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x4f, 0x57, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x02, 0x12, 0x08, 0x0a,
	0x04, 0x48, 0x49, 0x47, 0x48, 0x10, 0x03, 0x22, 0xbf, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x6e,
	0x64, 0x73, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73,
	0x65, 0x6e, 0x64, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x66, 0x0a, 0x07, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x75, 0x74,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x41, 0x75, 0x74, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x44, 0x6f, 0x6e,
	0x65, 0x22, 0x3d, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x22, 0x3e, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49,
	0x64, 0x22, 0x3b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x1e,
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x47,
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x32, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6c, 0x6f,
	0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x4c, 0x6f, 0x67, 0x32, 0xd6, 0x02, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x53, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55,
	0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x12, 0x24, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x6e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x67,
	0x73, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x18,
	0x5a, 0x09, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x62, 0xaa, 0x02, 0x0a, 0x47, 0x72,
	0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_pbs_reportpb_report_proto_rawDescOnce sync.Once
	file_src_pbs_reportpb_report_proto_rawDescData = file_src_pbs_reportpb_report_proto_rawDesc
)

func file_src_pbs_reportpb_report_proto_rawDescGZIP() []byte {
	file_src_pbs_reportpb_report_proto_rawDescOnce.Do(func() {
		file_src_pbs_reportpb_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_pbs_reportpb_report_proto_rawDescData)
	})
	return file_src_pbs_reportpb_report_proto_rawDescData
}

var file_src_pbs_reportpb_report_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_src_pbs_reportpb_report_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_src_pbs_reportpb_report_proto_goTypes = []interface{}{
	(Report_Status)(0),                    // 0: report.Report.Status
	(*Report)(nil),                        // 1: report.Report
	(*ReportLog)(nil),                     // 2: report.ReportLog
	(*Command)(nil),                       // 3: report.Command
	(*SensorReportRequest)(nil),           // 4: report.SensorReportRequest
	(*SensorReportResponse)(nil),          // 5: report.SensorReportResponse
	(*GetReportRequest)(nil),              // 6: report.GetReportRequest
	(*GetReportResponse)(nil),             // 7: report.GetReportResponse
	(*GetUnCompletedReportsRequest)(nil),  // 8: report.GetUnCompletedReportsRequest
	(*GetUnCompletedReportsResponse)(nil), // 9: report.GetUnCompletedReportsResponse
	(*GetReportLogRequest)(nil),           // 10: report.GetReportLogRequest
	(*GetReportLogResponse)(nil),          // 11: report.GetReportLogResponse
}
var file_src_pbs_reportpb_report_proto_depIdxs = []int32{
	0,  // 0: report.Report.status:type_name -> report.Report.Status
	3,  // 1: report.Report.commands:type_name -> report.Command
	1,  // 2: report.SensorReportRequest.report:type_name -> report.Report
	1,  // 3: report.SensorReportResponse.Report:type_name -> report.Report
	1,  // 4: report.GetReportResponse.report:type_name -> report.Report
	1,  // 5: report.GetUnCompletedReportsResponse.report:type_name -> report.Report
	2,  // 6: report.GetReportLogResponse.report_log:type_name -> report.ReportLog
	4,  // 7: report.ReportService.CreateReport:input_type -> report.SensorReportRequest
	6,  // 8: report.ReportService.GetReport:input_type -> report.GetReportRequest
	8,  // 9: report.ReportService.GetUnCompletedReports:input_type -> report.GetUnCompletedReportsRequest
	10, // 10: report.ReportService.GetReportLogs:input_type -> report.GetReportLogRequest
	5,  // 11: report.ReportService.CreateReport:output_type -> report.SensorReportResponse
	7,  // 12: report.ReportService.GetReport:output_type -> report.GetReportResponse
	9,  // 13: report.ReportService.GetUnCompletedReports:output_type -> report.GetUnCompletedReportsResponse
	11, // 14: report.ReportService.GetReportLogs:output_type -> report.GetReportLogResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_src_pbs_reportpb_report_proto_init() }
func file_src_pbs_reportpb_report_proto_init() {
	if File_src_pbs_reportpb_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_pbs_reportpb_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_pbs_reportpb_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportLog); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_pbs_reportpb_report_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_pbs_reportpb_report_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SensorReportRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_pbs_reportpb_report_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SensorReportResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_pbs_reportpb_report_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReportRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_pbs_reportpb_report_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReportResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_pbs_reportpb_report_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUnCompletedReportsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_pbs_reportpb_report_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUnCompletedReportsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_pbs_reportpb_report_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReportLogRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_pbs_reportpb_report_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReportLogResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_pbs_reportpb_report_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_pbs_reportpb_report_proto_goTypes,
		DependencyIndexes: file_src_pbs_reportpb_report_proto_depIdxs,
		EnumInfos:         file_src_pbs_reportpb_report_proto_enumTypes,
		MessageInfos:      file_src_pbs_reportpb_report_proto_msgTypes,
	}.Build()
	File_src_pbs_reportpb_report_proto = out.File
	file_src_pbs_reportpb_report_proto_rawDesc = nil
	file_src_pbs_reportpb_report_proto_goTypes = nil
	file_src_pbs_reportpb_report_proto_depIdxs = nil
}
