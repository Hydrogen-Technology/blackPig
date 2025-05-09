// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.2
// source: blackpig/v2/reserve.proto

package v2

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TimeTable struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StartTime     string                 `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime       string                 `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Day           string                 `protobuf:"bytes,4,opt,name=day,proto3" json:"day,omitempty"`
	UserId        int32                  `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // 用户id
	OrderedId     int32                  `protobuf:"varint,6,opt,name=ordered_id,json=orderedId,proto3" json:"ordered_id,omitempty"` // 被预约的导员的id
	CreatedAt     string                 `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimeTable) Reset() {
	*x = TimeTable{}
	mi := &file_blackpig_v2_reserve_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeTable) ProtoMessage() {}

func (x *TimeTable) ProtoReflect() protoreflect.Message {
	mi := &file_blackpig_v2_reserve_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeTable.ProtoReflect.Descriptor instead.
func (*TimeTable) Descriptor() ([]byte, []int) {
	return file_blackpig_v2_reserve_proto_rawDescGZIP(), []int{0}
}

func (x *TimeTable) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TimeTable) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *TimeTable) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *TimeTable) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *TimeTable) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *TimeTable) GetOrderedId() int32 {
	if x != nil {
		return x.OrderedId
	}
	return 0
}

func (x *TimeTable) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *TimeTable) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type Ordered struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TimeTables    []*TimeTable           `protobuf:"bytes,2,rep,name=time_tables,json=timeTables,proto3" json:"time_tables,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ordered) Reset() {
	*x = Ordered{}
	mi := &file_blackpig_v2_reserve_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ordered) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ordered) ProtoMessage() {}

func (x *Ordered) ProtoReflect() protoreflect.Message {
	mi := &file_blackpig_v2_reserve_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ordered.ProtoReflect.Descriptor instead.
func (*Ordered) Descriptor() ([]byte, []int) {
	return file_blackpig_v2_reserve_proto_rawDescGZIP(), []int{1}
}

func (x *Ordered) GetTimeTables() []*TimeTable {
	if x != nil {
		return x.TimeTables
	}
	return nil
}

type DeleteTimeTableRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTimeTableRequest) Reset() {
	*x = DeleteTimeTableRequest{}
	mi := &file_blackpig_v2_reserve_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTimeTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTimeTableRequest) ProtoMessage() {}

func (x *DeleteTimeTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blackpig_v2_reserve_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTimeTableRequest.ProtoReflect.Descriptor instead.
func (*DeleteTimeTableRequest) Descriptor() ([]byte, []int) {
	return file_blackpig_v2_reserve_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteTimeTableRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteTimeTableResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTimeTableResponse) Reset() {
	*x = DeleteTimeTableResponse{}
	mi := &file_blackpig_v2_reserve_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTimeTableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTimeTableResponse) ProtoMessage() {}

func (x *DeleteTimeTableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blackpig_v2_reserve_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTimeTableResponse.ProtoReflect.Descriptor instead.
func (*DeleteTimeTableResponse) Descriptor() ([]byte, []int) {
	return file_blackpig_v2_reserve_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteTimeTableResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetOrderedRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderedId     int32                  `protobuf:"varint,1,opt,name=ordered_id,json=orderedId,proto3" json:"ordered_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderedRequest) Reset() {
	*x = GetOrderedRequest{}
	mi := &file_blackpig_v2_reserve_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderedRequest) ProtoMessage() {}

func (x *GetOrderedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blackpig_v2_reserve_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderedRequest.ProtoReflect.Descriptor instead.
func (*GetOrderedRequest) Descriptor() ([]byte, []int) {
	return file_blackpig_v2_reserve_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrderedRequest) GetOrderedId() int32 {
	if x != nil {
		return x.OrderedId
	}
	return 0
}

type GetTimeTablesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Day           string                 `protobuf:"bytes,2,opt,name=day,proto3" json:"day,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTimeTablesRequest) Reset() {
	*x = GetTimeTablesRequest{}
	mi := &file_blackpig_v2_reserve_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTimeTablesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTimeTablesRequest) ProtoMessage() {}

func (x *GetTimeTablesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blackpig_v2_reserve_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTimeTablesRequest.ProtoReflect.Descriptor instead.
func (*GetTimeTablesRequest) Descriptor() ([]byte, []int) {
	return file_blackpig_v2_reserve_proto_rawDescGZIP(), []int{5}
}

func (x *GetTimeTablesRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetTimeTablesRequest) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

var File_blackpig_v2_reserve_proto protoreflect.FileDescriptor

var file_blackpig_v2_reserve_proto_rawDesc = []byte{
	0x0a, 0x19, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x70, 0x69, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x6c, 0x61,
	0x63, 0x6b, 0x70, 0x69, 0x67, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x01, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x42, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65,
	0x64, 0x12, 0x37, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x70, 0x69,
	0x67, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x0a,
	0x74, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x49, 0x64, 0x22, 0x41, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79,
	0x32, 0xd0, 0x04, 0x0a, 0x10, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x70, 0x69,
	0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x1a, 0x14, 0x2e, 0x62,
	0x6c, 0x61, 0x63, 0x6b, 0x70, 0x69, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x65, 0x64, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f,
	0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2f, 0x61, 0x64, 0x64, 0x12, 0x60,
	0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x16, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x70, 0x69, 0x67, 0x2e, 0x76, 0x32, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x1a, 0x16, 0x2e, 0x62, 0x6c, 0x61, 0x63,
	0x6b, 0x70, 0x69, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x32, 0x12, 0x2f, 0x76,
	0x32, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x7d, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x23, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x70, 0x69, 0x67, 0x2e, 0x76,
	0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b,
	0x70, 0x69, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x2a, 0x17, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x7b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x42, 0x79, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x49, 0x64, 0x12, 0x1e, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b,
	0x70, 0x69, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b,
	0x70, 0x69, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x22, 0x2c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x2f,
	0x7b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x84, 0x01, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x65, 0x64, 0x12, 0x21, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x70, 0x69, 0x67,
	0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b,
	0x70, 0x69, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x22, 0x33,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x12, 0x2b, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x64,
	0x61, 0x79, 0x7d, 0x42, 0x14, 0x5a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x61, 0x63, 0x6b,
	0x50, 0x69, 0x67, 0x2f, 0x76, 0x32, 0x3b, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_blackpig_v2_reserve_proto_rawDescOnce sync.Once
	file_blackpig_v2_reserve_proto_rawDescData = file_blackpig_v2_reserve_proto_rawDesc
)

func file_blackpig_v2_reserve_proto_rawDescGZIP() []byte {
	file_blackpig_v2_reserve_proto_rawDescOnce.Do(func() {
		file_blackpig_v2_reserve_proto_rawDescData = protoimpl.X.CompressGZIP(file_blackpig_v2_reserve_proto_rawDescData)
	})
	return file_blackpig_v2_reserve_proto_rawDescData
}

var file_blackpig_v2_reserve_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_blackpig_v2_reserve_proto_goTypes = []any{
	(*TimeTable)(nil),               // 0: blackpig.v2.TimeTable
	(*Ordered)(nil),                 // 1: blackpig.v2.Ordered
	(*DeleteTimeTableRequest)(nil),  // 2: blackpig.v2.DeleteTimeTableRequest
	(*DeleteTimeTableResponse)(nil), // 3: blackpig.v2.DeleteTimeTableResponse
	(*GetOrderedRequest)(nil),       // 4: blackpig.v2.GetOrderedRequest
	(*GetTimeTablesRequest)(nil),    // 5: blackpig.v2.GetTimeTablesRequest
}
var file_blackpig_v2_reserve_proto_depIdxs = []int32{
	0, // 0: blackpig.v2.Ordered.time_tables:type_name -> blackpig.v2.TimeTable
	1, // 1: blackpig.v2.TimeTableService.AddTimeTables:input_type -> blackpig.v2.Ordered
	0, // 2: blackpig.v2.TimeTableService.UpdateTimeTable:input_type -> blackpig.v2.TimeTable
	2, // 3: blackpig.v2.TimeTableService.DeleteTimeTable:input_type -> blackpig.v2.DeleteTimeTableRequest
	4, // 4: blackpig.v2.TimeTableService.GetOrderedByOrderedId:input_type -> blackpig.v2.GetOrderedRequest
	5, // 5: blackpig.v2.TimeTableService.GetTimeTablesOrdered:input_type -> blackpig.v2.GetTimeTablesRequest
	1, // 6: blackpig.v2.TimeTableService.AddTimeTables:output_type -> blackpig.v2.Ordered
	0, // 7: blackpig.v2.TimeTableService.UpdateTimeTable:output_type -> blackpig.v2.TimeTable
	3, // 8: blackpig.v2.TimeTableService.DeleteTimeTable:output_type -> blackpig.v2.DeleteTimeTableResponse
	1, // 9: blackpig.v2.TimeTableService.GetOrderedByOrderedId:output_type -> blackpig.v2.Ordered
	1, // 10: blackpig.v2.TimeTableService.GetTimeTablesOrdered:output_type -> blackpig.v2.Ordered
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_blackpig_v2_reserve_proto_init() }
func file_blackpig_v2_reserve_proto_init() {
	if File_blackpig_v2_reserve_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blackpig_v2_reserve_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blackpig_v2_reserve_proto_goTypes,
		DependencyIndexes: file_blackpig_v2_reserve_proto_depIdxs,
		MessageInfos:      file_blackpig_v2_reserve_proto_msgTypes,
	}.Build()
	File_blackpig_v2_reserve_proto = out.File
	file_blackpig_v2_reserve_proto_rawDesc = nil
	file_blackpig_v2_reserve_proto_goTypes = nil
	file_blackpig_v2_reserve_proto_depIdxs = nil
}
