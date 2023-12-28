// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.1
// source: proto/security.proto

package security

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
		mi := &file_proto_security_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_security_proto_msgTypes[0]
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
	return file_proto_security_proto_rawDescGZIP(), []int{0}
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
		mi := &file_proto_security_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_security_proto_msgTypes[1]
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
	return file_proto_security_proto_rawDescGZIP(), []int{1}
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
		mi := &file_proto_security_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamRequest) ProtoMessage() {}

func (x *ClientStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_security_proto_msgTypes[2]
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
	return file_proto_security_proto_rawDescGZIP(), []int{2}
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
		mi := &file_proto_security_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamResponse) ProtoMessage() {}

func (x *ClientStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_security_proto_msgTypes[3]
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
	return file_proto_security_proto_rawDescGZIP(), []int{3}
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
		mi := &file_proto_security_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStreamRequest) ProtoMessage() {}

func (x *ServerStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_security_proto_msgTypes[4]
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
	return file_proto_security_proto_rawDescGZIP(), []int{4}
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
		mi := &file_proto_security_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStreamResponse) ProtoMessage() {}

func (x *ServerStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_security_proto_msgTypes[5]
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
	return file_proto_security_proto_rawDescGZIP(), []int{5}
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
		mi := &file_proto_security_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidiStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidiStreamRequest) ProtoMessage() {}

func (x *BidiStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_security_proto_msgTypes[6]
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
	return file_proto_security_proto_rawDescGZIP(), []int{6}
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
		mi := &file_proto_security_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidiStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidiStreamResponse) ProtoMessage() {}

func (x *BidiStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_security_proto_msgTypes[7]
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
	return file_proto_security_proto_rawDescGZIP(), []int{7}
}

func (x *BidiStreamResponse) GetStroke() int64 {
	if x != nil {
		return x.Stroke
	}
	return 0
}

type CheckTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CheckTokenRequest) Reset() {
	*x = CheckTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_security_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTokenRequest) ProtoMessage() {}

func (x *CheckTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_security_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTokenRequest.ProtoReflect.Descriptor instead.
func (*CheckTokenRequest) Descriptor() ([]byte, []int) {
	return file_proto_security_proto_rawDescGZIP(), []int{8}
}

func (x *CheckTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CheckTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuyTax             float32 `protobuf:"fixed32,1,opt,name=buyTax,proto3" json:"buyTax,omitempty"`
	SellTax            float32 `protobuf:"fixed32,2,opt,name=sellTax,proto3" json:"sellTax,omitempty"`
	Owner              bool    `protobuf:"varint,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Honey              bool    `protobuf:"varint,4,opt,name=honey,proto3" json:"honey,omitempty"`
	Pause              bool    `protobuf:"varint,5,opt,name=pause,proto3" json:"pause,omitempty"`
	Lock               bool    `protobuf:"varint,6,opt,name=lock,proto3" json:"lock,omitempty"`
	CoolDown           bool    `protobuf:"varint,7,opt,name=coolDown,proto3" json:"coolDown,omitempty"`
	Mint               bool    `protobuf:"varint,8,opt,name=mint,proto3" json:"mint,omitempty"`
	Score              int64   `protobuf:"varint,9,opt,name=score,proto3" json:"score,omitempty"`
	OpenSource         bool    `protobuf:"varint,10,opt,name=openSource,proto3" json:"openSource,omitempty"`
	Proxy              bool    `protobuf:"varint,11,opt,name=proxy,proto3" json:"proxy,omitempty"`
	TransferPause      bool    `protobuf:"varint,12,opt,name=transferPause,proto3" json:"transferPause,omitempty"`
	SlippageModifiable bool    `protobuf:"varint,13,opt,name=slippageModifiable,proto3" json:"slippageModifiable,omitempty"`
	BlackList          bool    `protobuf:"varint,14,opt,name=blackList,proto3" json:"blackList,omitempty"`
	WhiteList          bool    `protobuf:"varint,15,opt,name=whiteList,proto3" json:"whiteList,omitempty"`
	IsBuy              bool    `protobuf:"varint,16,opt,name=isBuy,proto3" json:"isBuy,omitempty"`
	IsSellAll          bool    `protobuf:"varint,17,opt,name=isSellAll,proto3" json:"isSellAll,omitempty"`
	HiddenOwner        bool    `protobuf:"varint,18,opt,name=hiddenOwner,proto3" json:"hiddenOwner,omitempty"`
}

func (x *CheckTokenResponse) Reset() {
	*x = CheckTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_security_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTokenResponse) ProtoMessage() {}

func (x *CheckTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_security_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTokenResponse.ProtoReflect.Descriptor instead.
func (*CheckTokenResponse) Descriptor() ([]byte, []int) {
	return file_proto_security_proto_rawDescGZIP(), []int{9}
}

func (x *CheckTokenResponse) GetBuyTax() float32 {
	if x != nil {
		return x.BuyTax
	}
	return 0
}

func (x *CheckTokenResponse) GetSellTax() float32 {
	if x != nil {
		return x.SellTax
	}
	return 0
}

func (x *CheckTokenResponse) GetOwner() bool {
	if x != nil {
		return x.Owner
	}
	return false
}

func (x *CheckTokenResponse) GetHoney() bool {
	if x != nil {
		return x.Honey
	}
	return false
}

func (x *CheckTokenResponse) GetPause() bool {
	if x != nil {
		return x.Pause
	}
	return false
}

func (x *CheckTokenResponse) GetLock() bool {
	if x != nil {
		return x.Lock
	}
	return false
}

func (x *CheckTokenResponse) GetCoolDown() bool {
	if x != nil {
		return x.CoolDown
	}
	return false
}

func (x *CheckTokenResponse) GetMint() bool {
	if x != nil {
		return x.Mint
	}
	return false
}

func (x *CheckTokenResponse) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *CheckTokenResponse) GetOpenSource() bool {
	if x != nil {
		return x.OpenSource
	}
	return false
}

func (x *CheckTokenResponse) GetProxy() bool {
	if x != nil {
		return x.Proxy
	}
	return false
}

func (x *CheckTokenResponse) GetTransferPause() bool {
	if x != nil {
		return x.TransferPause
	}
	return false
}

func (x *CheckTokenResponse) GetSlippageModifiable() bool {
	if x != nil {
		return x.SlippageModifiable
	}
	return false
}

func (x *CheckTokenResponse) GetBlackList() bool {
	if x != nil {
		return x.BlackList
	}
	return false
}

func (x *CheckTokenResponse) GetWhiteList() bool {
	if x != nil {
		return x.WhiteList
	}
	return false
}

func (x *CheckTokenResponse) GetIsBuy() bool {
	if x != nil {
		return x.IsBuy
	}
	return false
}

func (x *CheckTokenResponse) GetIsSellAll() bool {
	if x != nil {
		return x.IsSellAll
	}
	return false
}

func (x *CheckTokenResponse) GetHiddenOwner() bool {
	if x != nil {
		return x.HiddenOwner
	}
	return false
}

var File_proto_security_proto protoreflect.FileDescriptor

var file_proto_security_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x22, 0x21, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x2d, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x6f, 0x6b, 0x65, 0x22, 0x2c, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x2c, 0x0a, 0x14, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2b, 0x0a,
	0x11, 0x42, 0x69, 0x64, 0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x22, 0x2c, 0x0a, 0x12, 0x42, 0x69,
	0x64, 0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x80, 0x04, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75,
	0x79, 0x54, 0x61, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x62, 0x75, 0x79, 0x54,
	0x61, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x6c, 0x6c, 0x54, 0x61, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x07, 0x73, 0x65, 0x6c, 0x6c, 0x54, 0x61, 0x78, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x68, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x75, 0x73,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x70, 0x61, 0x75, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6f, 0x6c, 0x44, 0x6f, 0x77, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6f, 0x6c, 0x44, 0x6f, 0x77, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x69, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6d, 0x69,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x6e,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6f, 0x70,
	0x65, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x24,
	0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x61, 0x75, 0x73, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50,
	0x61, 0x75, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x6c, 0x69, 0x70, 0x70, 0x61, 0x67, 0x65,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x12, 0x73, 0x6c, 0x69, 0x70, 0x70, 0x61, 0x67, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x77, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x42, 0x75, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x69, 0x73, 0x42, 0x75, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x53, 0x65, 0x6c, 0x6c,
	0x41, 0x6c, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x65, 0x6c,
	0x6c, 0x41, 0x6c, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x69, 0x64, 0x64, 0x65,
	0x6e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x32, 0x83, 0x03, 0x0a, 0x08, 0x53, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x61,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0c,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1d, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12,
	0x51, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x1d, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x4d, 0x0a, 0x0a, 0x42, 0x69, 0x64, 0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x1b, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x42, 0x69, 0x64, 0x69,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x42, 0x69, 0x64, 0x69, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x12, 0x49, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1b, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_security_proto_rawDescOnce sync.Once
	file_proto_security_proto_rawDescData = file_proto_security_proto_rawDesc
)

func file_proto_security_proto_rawDescGZIP() []byte {
	file_proto_security_proto_rawDescOnce.Do(func() {
		file_proto_security_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_security_proto_rawDescData)
	})
	return file_proto_security_proto_rawDescData
}

var file_proto_security_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_security_proto_goTypes = []interface{}{
	(*CallRequest)(nil),          // 0: security.CallRequest
	(*CallResponse)(nil),         // 1: security.CallResponse
	(*ClientStreamRequest)(nil),  // 2: security.ClientStreamRequest
	(*ClientStreamResponse)(nil), // 3: security.ClientStreamResponse
	(*ServerStreamRequest)(nil),  // 4: security.ServerStreamRequest
	(*ServerStreamResponse)(nil), // 5: security.ServerStreamResponse
	(*BidiStreamRequest)(nil),    // 6: security.BidiStreamRequest
	(*BidiStreamResponse)(nil),   // 7: security.BidiStreamResponse
	(*CheckTokenRequest)(nil),    // 8: security.CheckTokenRequest
	(*CheckTokenResponse)(nil),   // 9: security.CheckTokenResponse
}
var file_proto_security_proto_depIdxs = []int32{
	0, // 0: security.Security.Call:input_type -> security.CallRequest
	2, // 1: security.Security.ClientStream:input_type -> security.ClientStreamRequest
	4, // 2: security.Security.ServerStream:input_type -> security.ServerStreamRequest
	6, // 3: security.Security.BidiStream:input_type -> security.BidiStreamRequest
	8, // 4: security.Security.CheckToken:input_type -> security.CheckTokenRequest
	1, // 5: security.Security.Call:output_type -> security.CallResponse
	3, // 6: security.Security.ClientStream:output_type -> security.ClientStreamResponse
	5, // 7: security.Security.ServerStream:output_type -> security.ServerStreamResponse
	7, // 8: security.Security.BidiStream:output_type -> security.BidiStreamResponse
	9, // 9: security.Security.CheckToken:output_type -> security.CheckTokenResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_security_proto_init() }
func file_proto_security_proto_init() {
	if File_proto_security_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_security_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_security_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_security_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_security_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_security_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_security_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_security_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_security_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_security_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTokenRequest); i {
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
		file_proto_security_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTokenResponse); i {
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
			RawDescriptor: file_proto_security_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_security_proto_goTypes,
		DependencyIndexes: file_proto_security_proto_depIdxs,
		MessageInfos:      file_proto_security_proto_msgTypes,
	}.Build()
	File_proto_security_proto = out.File
	file_proto_security_proto_rawDesc = nil
	file_proto_security_proto_goTypes = nil
	file_proto_security_proto_depIdxs = nil
}