// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.1
// source: proto/consul.proto

package consul

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

type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consul_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consul_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_proto_consul_proto_rawDescGZIP(), []int{0}
}

func (x *CallRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CallResponse) Reset() {
	*x = CallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consul_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consul_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResponse.ProtoReflect.Descriptor instead.
func (*CallResponse) Descriptor() ([]byte, []int) {
	return file_proto_consul_proto_rawDescGZIP(), []int{1}
}

func (x *CallResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ClientStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stroke int64 `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
}

func (x *ClientStreamRequest) Reset() {
	*x = ClientStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consul_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamRequest) ProtoMessage() {}

func (x *ClientStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consul_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStreamRequest.ProtoReflect.Descriptor instead.
func (*ClientStreamRequest) Descriptor() ([]byte, []int) {
	return file_proto_consul_proto_rawDescGZIP(), []int{2}
}

func (x *ClientStreamRequest) GetStroke() int64 {
	if x != nil {
		return x.Stroke
	}
	return 0
}

type ClientStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ClientStreamResponse) Reset() {
	*x = ClientStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consul_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamResponse) ProtoMessage() {}

func (x *ClientStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consul_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStreamResponse.ProtoReflect.Descriptor instead.
func (*ClientStreamResponse) Descriptor() ([]byte, []int) {
	return file_proto_consul_proto_rawDescGZIP(), []int{3}
}

func (x *ClientStreamResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ServerStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ServerStreamRequest) Reset() {
	*x = ServerStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consul_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStreamRequest) ProtoMessage() {}

func (x *ServerStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consul_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStreamRequest.ProtoReflect.Descriptor instead.
func (*ServerStreamRequest) Descriptor() ([]byte, []int) {
	return file_proto_consul_proto_rawDescGZIP(), []int{4}
}

func (x *ServerStreamRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ServerStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ServerStreamResponse) Reset() {
	*x = ServerStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consul_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStreamResponse) ProtoMessage() {}

func (x *ServerStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consul_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStreamResponse.ProtoReflect.Descriptor instead.
func (*ServerStreamResponse) Descriptor() ([]byte, []int) {
	return file_proto_consul_proto_rawDescGZIP(), []int{5}
}

func (x *ServerStreamResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type BidiStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stroke int64 `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
}

func (x *BidiStreamRequest) Reset() {
	*x = BidiStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consul_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidiStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidiStreamRequest) ProtoMessage() {}

func (x *BidiStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consul_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidiStreamRequest.ProtoReflect.Descriptor instead.
func (*BidiStreamRequest) Descriptor() ([]byte, []int) {
	return file_proto_consul_proto_rawDescGZIP(), []int{6}
}

func (x *BidiStreamRequest) GetStroke() int64 {
	if x != nil {
		return x.Stroke
	}
	return 0
}

type BidiStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stroke int64 `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
}

func (x *BidiStreamResponse) Reset() {
	*x = BidiStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consul_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidiStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidiStreamResponse) ProtoMessage() {}

func (x *BidiStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consul_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidiStreamResponse.ProtoReflect.Descriptor instead.
func (*BidiStreamResponse) Descriptor() ([]byte, []int) {
	return file_proto_consul_proto_rawDescGZIP(), []int{7}
}

func (x *BidiStreamResponse) GetStroke() int64 {
	if x != nil {
		return x.Stroke
	}
	return 0
}

type RegisterServiceStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr        string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	ServiceName string `protobuf:"bytes,2,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Ip          string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Port        int64  `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Tag         string `protobuf:"bytes,5,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *RegisterServiceStreamRequest) Reset() {
	*x = RegisterServiceStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consul_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterServiceStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterServiceStreamRequest) ProtoMessage() {}

func (x *RegisterServiceStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consul_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterServiceStreamRequest.ProtoReflect.Descriptor instead.
func (*RegisterServiceStreamRequest) Descriptor() ([]byte, []int) {
	return file_proto_consul_proto_rawDescGZIP(), []int{8}
}

func (x *RegisterServiceStreamRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *RegisterServiceStreamRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *RegisterServiceStreamRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *RegisterServiceStreamRequest) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RegisterServiceStreamRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type RegisterServiceStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *RegisterServiceStreamResponse) Reset() {
	*x = RegisterServiceStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consul_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterServiceStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterServiceStreamResponse) ProtoMessage() {}

func (x *RegisterServiceStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consul_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterServiceStreamResponse.ProtoReflect.Descriptor instead.
func (*RegisterServiceStreamResponse) Descriptor() ([]byte, []int) {
	return file_proto_consul_proto_rawDescGZIP(), []int{9}
}

func (x *RegisterServiceStreamResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type DeregisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr      string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	ServiceID string `protobuf:"bytes,2,opt,name=serviceID,proto3" json:"serviceID,omitempty"`
}

func (x *DeregisterRequest) Reset() {
	*x = DeregisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consul_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeregisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeregisterRequest) ProtoMessage() {}

func (x *DeregisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consul_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeregisterRequest.ProtoReflect.Descriptor instead.
func (*DeregisterRequest) Descriptor() ([]byte, []int) {
	return file_proto_consul_proto_rawDescGZIP(), []int{10}
}

func (x *DeregisterRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *DeregisterRequest) GetServiceID() string {
	if x != nil {
		return x.ServiceID
	}
	return ""
}

type DeregisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *DeregisterResponse) Reset() {
	*x = DeregisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consul_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeregisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeregisterResponse) ProtoMessage() {}

func (x *DeregisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consul_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeregisterResponse.ProtoReflect.Descriptor instead.
func (*DeregisterResponse) Descriptor() ([]byte, []int) {
	return file_proto_consul_proto_rawDescGZIP(), []int{11}
}

func (x *DeregisterResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_proto_consul_proto protoreflect.FileDescriptor

var file_proto_consul_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x22, 0x21, 0x0a, 0x0b,
	0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x20, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x2d, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x6f,
	0x6b, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65,
	0x22, 0x2c, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2b,
	0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x14, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x11, 0x42, 0x69, 0x64,
	0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x22, 0x2c, 0x0a, 0x12, 0x42, 0x69, 0x64, 0x69, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x6f, 0x6b, 0x65, 0x22, 0x8a, 0x01, 0x0a, 0x1c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61,
	0x67, 0x22, 0x35, 0x0a, 0x1d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x45, 0x0a, 0x11, 0x44, 0x65, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x22,
	0x2a, 0x0a, 0x12, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xd5, 0x03, 0x0a, 0x06,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x33, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x13,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x43, 0x61, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1b, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x4d, 0x0a, 0x0c, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x49, 0x0a, 0x0a, 0x42, 0x69, 0x64,
	0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x2e, 0x42, 0x69, 0x64, 0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x42, 0x69, 0x64, 0x69,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x12, 0x66, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x24, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0a,
	0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x2e, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x44,
	0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_consul_proto_rawDescOnce sync.Once
	file_proto_consul_proto_rawDescData = file_proto_consul_proto_rawDesc
)

func file_proto_consul_proto_rawDescGZIP() []byte {
	file_proto_consul_proto_rawDescOnce.Do(func() {
		file_proto_consul_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_consul_proto_rawDescData)
	})
	return file_proto_consul_proto_rawDescData
}

var file_proto_consul_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_consul_proto_goTypes = []interface{}{
	(*CallRequest)(nil),                   // 0: consul.CallRequest
	(*CallResponse)(nil),                  // 1: consul.CallResponse
	(*ClientStreamRequest)(nil),           // 2: consul.ClientStreamRequest
	(*ClientStreamResponse)(nil),          // 3: consul.ClientStreamResponse
	(*ServerStreamRequest)(nil),           // 4: consul.ServerStreamRequest
	(*ServerStreamResponse)(nil),          // 5: consul.ServerStreamResponse
	(*BidiStreamRequest)(nil),             // 6: consul.BidiStreamRequest
	(*BidiStreamResponse)(nil),            // 7: consul.BidiStreamResponse
	(*RegisterServiceStreamRequest)(nil),  // 8: consul.RegisterServiceStreamRequest
	(*RegisterServiceStreamResponse)(nil), // 9: consul.RegisterServiceStreamResponse
	(*DeregisterRequest)(nil),             // 10: consul.DeregisterRequest
	(*DeregisterResponse)(nil),            // 11: consul.DeregisterResponse
}
var file_proto_consul_proto_depIdxs = []int32{
	0,  // 0: consul.Consul.Call:input_type -> consul.CallRequest
	2,  // 1: consul.Consul.ClientStream:input_type -> consul.ClientStreamRequest
	4,  // 2: consul.Consul.ServerStream:input_type -> consul.ServerStreamRequest
	6,  // 3: consul.Consul.BidiStream:input_type -> consul.BidiStreamRequest
	8,  // 4: consul.Consul.RegisterServiceStream:input_type -> consul.RegisterServiceStreamRequest
	10, // 5: consul.Consul.Deregister:input_type -> consul.DeregisterRequest
	1,  // 6: consul.Consul.Call:output_type -> consul.CallResponse
	3,  // 7: consul.Consul.ClientStream:output_type -> consul.ClientStreamResponse
	5,  // 8: consul.Consul.ServerStream:output_type -> consul.ServerStreamResponse
	7,  // 9: consul.Consul.BidiStream:output_type -> consul.BidiStreamResponse
	9,  // 10: consul.Consul.RegisterServiceStream:output_type -> consul.RegisterServiceStreamResponse
	11, // 11: consul.Consul.Deregister:output_type -> consul.DeregisterResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_consul_proto_init() }
func file_proto_consul_proto_init() {
	if File_proto_consul_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_consul_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest); i {
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
		file_proto_consul_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallResponse); i {
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
		file_proto_consul_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStreamRequest); i {
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
		file_proto_consul_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStreamResponse); i {
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
		file_proto_consul_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStreamRequest); i {
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
		file_proto_consul_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStreamResponse); i {
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
		file_proto_consul_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidiStreamRequest); i {
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
		file_proto_consul_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidiStreamResponse); i {
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
		file_proto_consul_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterServiceStreamRequest); i {
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
		file_proto_consul_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterServiceStreamResponse); i {
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
		file_proto_consul_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeregisterRequest); i {
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
		file_proto_consul_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeregisterResponse); i {
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
			RawDescriptor: file_proto_consul_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_consul_proto_goTypes,
		DependencyIndexes: file_proto_consul_proto_depIdxs,
		MessageInfos:      file_proto_consul_proto_msgTypes,
	}.Build()
	File_proto_consul_proto = out.File
	file_proto_consul_proto_rawDesc = nil
	file_proto_consul_proto_goTypes = nil
	file_proto_consul_proto_depIdxs = nil
}