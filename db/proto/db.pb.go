// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.1
// source: proto/db.proto

package db

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
		mi := &file_proto_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_proto_msgTypes[0]
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
	return file_proto_db_proto_rawDescGZIP(), []int{0}
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
		mi := &file_proto_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_proto_msgTypes[1]
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
	return file_proto_db_proto_rawDescGZIP(), []int{1}
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
		mi := &file_proto_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamRequest) ProtoMessage() {}

func (x *ClientStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_proto_msgTypes[2]
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
	return file_proto_db_proto_rawDescGZIP(), []int{2}
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
		mi := &file_proto_db_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamResponse) ProtoMessage() {}

func (x *ClientStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_proto_msgTypes[3]
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
	return file_proto_db_proto_rawDescGZIP(), []int{3}
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
		mi := &file_proto_db_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStreamRequest) ProtoMessage() {}

func (x *ServerStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_proto_msgTypes[4]
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
	return file_proto_db_proto_rawDescGZIP(), []int{4}
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
		mi := &file_proto_db_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStreamResponse) ProtoMessage() {}

func (x *ServerStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_proto_msgTypes[5]
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
	return file_proto_db_proto_rawDescGZIP(), []int{5}
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
		mi := &file_proto_db_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidiStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidiStreamRequest) ProtoMessage() {}

func (x *BidiStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_proto_msgTypes[6]
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
	return file_proto_db_proto_rawDescGZIP(), []int{6}
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
		mi := &file_proto_db_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidiStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidiStreamResponse) ProtoMessage() {}

func (x *BidiStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_proto_msgTypes[7]
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
	return file_proto_db_proto_rawDescGZIP(), []int{7}
}

func (x *BidiStreamResponse) GetStroke() int64 {
	if x != nil {
		return x.Stroke
	}
	return 0
}

type GetPushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int64             `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page  int64             `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Param map[string]string `protobuf:"bytes,3,rep,name=param,proto3" json:"param,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetPushRequest) Reset() {
	*x = GetPushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_db_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPushRequest) ProtoMessage() {}

func (x *GetPushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPushRequest.ProtoReflect.Descriptor instead.
func (*GetPushRequest) Descriptor() ([]byte, []int) {
	return file_proto_db_proto_rawDescGZIP(), []int{8}
}

func (x *GetPushRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetPushRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetPushRequest) GetParam() map[string]string {
	if x != nil {
		return x.Param
	}
	return nil
}

type PushData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol    string  `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Token     string  `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Vol       float32 `protobuf:"fixed32,3,opt,name=vol,proto3" json:"vol,omitempty"`
	Trader    int64   `protobuf:"varint,4,opt,name=trader,proto3" json:"trader,omitempty"`
	Tx        int64   `protobuf:"varint,5,opt,name=tx,proto3" json:"tx,omitempty"`
	Tweet     int64   `protobuf:"varint,6,opt,name=tweet,proto3" json:"tweet,omitempty"`
	Liquidity float32 `protobuf:"fixed32,7,opt,name=liquidity,proto3" json:"liquidity,omitempty"`
	Age       string  `protobuf:"bytes,8,opt,name=age,proto3" json:"age,omitempty"`
	Times     int64   `protobuf:"varint,9,opt,name=times,proto3" json:"times,omitempty"`
}

func (x *PushData) Reset() {
	*x = PushData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_db_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushData) ProtoMessage() {}

func (x *PushData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushData.ProtoReflect.Descriptor instead.
func (*PushData) Descriptor() ([]byte, []int) {
	return file_proto_db_proto_rawDescGZIP(), []int{9}
}

func (x *PushData) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *PushData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *PushData) GetVol() float32 {
	if x != nil {
		return x.Vol
	}
	return 0
}

func (x *PushData) GetTrader() int64 {
	if x != nil {
		return x.Trader
	}
	return 0
}

func (x *PushData) GetTx() int64 {
	if x != nil {
		return x.Tx
	}
	return 0
}

func (x *PushData) GetTweet() int64 {
	if x != nil {
		return x.Tweet
	}
	return 0
}

func (x *PushData) GetLiquidity() float32 {
	if x != nil {
		return x.Liquidity
	}
	return 0
}

func (x *PushData) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

func (x *PushData) GetTimes() int64 {
	if x != nil {
		return x.Times
	}
	return 0
}

type GetPushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PushData []*PushData `protobuf:"bytes,1,rep,name=pushData,proto3" json:"pushData,omitempty"`
}

func (x *GetPushResponse) Reset() {
	*x = GetPushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_db_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPushResponse) ProtoMessage() {}

func (x *GetPushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPushResponse.ProtoReflect.Descriptor instead.
func (*GetPushResponse) Descriptor() ([]byte, []int) {
	return file_proto_db_proto_rawDescGZIP(), []int{10}
}

func (x *GetPushResponse) GetPushData() []*PushData {
	if x != nil {
		return x.PushData
	}
	return nil
}

type GetPushDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Param map[string]string `protobuf:"bytes,1,rep,name=param,proto3" json:"param,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetPushDetailRequest) Reset() {
	*x = GetPushDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_db_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPushDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPushDetailRequest) ProtoMessage() {}

func (x *GetPushDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPushDetailRequest.ProtoReflect.Descriptor instead.
func (*GetPushDetailRequest) Descriptor() ([]byte, []int) {
	return file_proto_db_proto_rawDescGZIP(), []int{11}
}

func (x *GetPushDetailRequest) GetParam() map[string]string {
	if x != nil {
		return x.Param
	}
	return nil
}

type PushDetailData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time      string  `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Level     string  `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	Vol       float32 `protobuf:"fixed32,3,opt,name=vol,proto3" json:"vol,omitempty"`
	Trader    int64   `protobuf:"varint,4,opt,name=trader,proto3" json:"trader,omitempty"`
	Tx        int64   `protobuf:"varint,5,opt,name=tx,proto3" json:"tx,omitempty"`
	Tweet     int64   `protobuf:"varint,6,opt,name=tweet,proto3" json:"tweet,omitempty"`
	Liquidity float32 `protobuf:"fixed32,7,opt,name=liquidity,proto3" json:"liquidity,omitempty"`
	Price     float32 `protobuf:"fixed32,8,opt,name=price,proto3" json:"price,omitempty"`
	Age       string  `protobuf:"bytes,9,opt,name=age,proto3" json:"age,omitempty"`
	Times     int64   `protobuf:"varint,10,opt,name=times,proto3" json:"times,omitempty"`
	Type      int64   `protobuf:"varint,11,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *PushDetailData) Reset() {
	*x = PushDetailData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_db_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushDetailData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushDetailData) ProtoMessage() {}

func (x *PushDetailData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushDetailData.ProtoReflect.Descriptor instead.
func (*PushDetailData) Descriptor() ([]byte, []int) {
	return file_proto_db_proto_rawDescGZIP(), []int{12}
}

func (x *PushDetailData) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *PushDetailData) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *PushDetailData) GetVol() float32 {
	if x != nil {
		return x.Vol
	}
	return 0
}

func (x *PushDetailData) GetTrader() int64 {
	if x != nil {
		return x.Trader
	}
	return 0
}

func (x *PushDetailData) GetTx() int64 {
	if x != nil {
		return x.Tx
	}
	return 0
}

func (x *PushDetailData) GetTweet() int64 {
	if x != nil {
		return x.Tweet
	}
	return 0
}

func (x *PushDetailData) GetLiquidity() float32 {
	if x != nil {
		return x.Liquidity
	}
	return 0
}

func (x *PushDetailData) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *PushDetailData) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

func (x *PushDetailData) GetTimes() int64 {
	if x != nil {
		return x.Times
	}
	return 0
}

func (x *PushDetailData) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

type GetPushDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PushDetailData []*PushDetailData `protobuf:"bytes,1,rep,name=pushDetailData,proto3" json:"pushDetailData,omitempty"`
}

func (x *GetPushDetailResponse) Reset() {
	*x = GetPushDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_db_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPushDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPushDetailResponse) ProtoMessage() {}

func (x *GetPushDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPushDetailResponse.ProtoReflect.Descriptor instead.
func (*GetPushDetailResponse) Descriptor() ([]byte, []int) {
	return file_proto_db_proto_rawDescGZIP(), []int{13}
}

func (x *GetPushDetailResponse) GetPushDetailData() []*PushDetailData {
	if x != nil {
		return x.PushDetailData
	}
	return nil
}

var File_proto_db_proto protoreflect.FileDescriptor

var file_proto_db_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x64, 0x62, 0x22, 0x21, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x2d, 0x0a, 0x13, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x22, 0x2c, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x14, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x2b, 0x0a, 0x11, 0x42, 0x69, 0x64, 0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x22, 0x2c,
	0x0a, 0x12, 0x42, 0x69, 0x64, 0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x22, 0xa9, 0x01, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x38,
	0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xce, 0x01, 0x0a, 0x08, 0x50, 0x75, 0x73,
	0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x03, 0x76, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x78, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x77, 0x65, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x77,
	0x65, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x08,
	0x70, 0x75, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x64, 0x62, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x70, 0x75,
	0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x22, 0x8b, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x75,
	0x73, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x39, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x73, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x38, 0x0a, 0x0a, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xfa, 0x01, 0x0a, 0x0e, 0x50, 0x75, 0x73, 0x68, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03,
	0x76, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x77, 0x65, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x77, 0x65, 0x65,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x53, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x75, 0x73, 0x68, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x70, 0x75,
	0x73, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x62, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0e, 0x70, 0x75, 0x73, 0x68, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x32, 0x80, 0x03, 0x0a, 0x02, 0x44, 0x62, 0x12, 0x2b, 0x0a,
	0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x0f, 0x2e, 0x64, 0x62, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x64, 0x62, 0x2e, 0x43, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x17, 0x2e, 0x64, 0x62, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64, 0x62, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x12, 0x45, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x17, 0x2e, 0x64, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64, 0x62, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x41, 0x0a, 0x0a, 0x42, 0x69, 0x64, 0x69,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x64, 0x62, 0x2e, 0x42, 0x69, 0x64, 0x69,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x64, 0x62, 0x2e, 0x42, 0x69, 0x64, 0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x34, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x50, 0x75, 0x73, 0x68, 0x12, 0x12, 0x2e, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x64, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x46, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x75, 0x73, 0x68, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x18, 0x2e, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x73, 0x68, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x73, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_db_proto_rawDescOnce sync.Once
	file_proto_db_proto_rawDescData = file_proto_db_proto_rawDesc
)

func file_proto_db_proto_rawDescGZIP() []byte {
	file_proto_db_proto_rawDescOnce.Do(func() {
		file_proto_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_db_proto_rawDescData)
	})
	return file_proto_db_proto_rawDescData
}

var file_proto_db_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_proto_db_proto_goTypes = []interface{}{
	(*CallRequest)(nil),           // 0: db.CallRequest
	(*CallResponse)(nil),          // 1: db.CallResponse
	(*ClientStreamRequest)(nil),   // 2: db.ClientStreamRequest
	(*ClientStreamResponse)(nil),  // 3: db.ClientStreamResponse
	(*ServerStreamRequest)(nil),   // 4: db.ServerStreamRequest
	(*ServerStreamResponse)(nil),  // 5: db.ServerStreamResponse
	(*BidiStreamRequest)(nil),     // 6: db.BidiStreamRequest
	(*BidiStreamResponse)(nil),    // 7: db.BidiStreamResponse
	(*GetPushRequest)(nil),        // 8: db.GetPushRequest
	(*PushData)(nil),              // 9: db.PushData
	(*GetPushResponse)(nil),       // 10: db.GetPushResponse
	(*GetPushDetailRequest)(nil),  // 11: db.GetPushDetailRequest
	(*PushDetailData)(nil),        // 12: db.PushDetailData
	(*GetPushDetailResponse)(nil), // 13: db.GetPushDetailResponse
	nil,                           // 14: db.GetPushRequest.ParamEntry
	nil,                           // 15: db.GetPushDetailRequest.ParamEntry
}
var file_proto_db_proto_depIdxs = []int32{
	14, // 0: db.GetPushRequest.param:type_name -> db.GetPushRequest.ParamEntry
	9,  // 1: db.GetPushResponse.pushData:type_name -> db.PushData
	15, // 2: db.GetPushDetailRequest.param:type_name -> db.GetPushDetailRequest.ParamEntry
	12, // 3: db.GetPushDetailResponse.pushDetailData:type_name -> db.PushDetailData
	0,  // 4: db.Db.Call:input_type -> db.CallRequest
	2,  // 5: db.Db.ClientStream:input_type -> db.ClientStreamRequest
	4,  // 6: db.Db.ServerStream:input_type -> db.ServerStreamRequest
	6,  // 7: db.Db.BidiStream:input_type -> db.BidiStreamRequest
	8,  // 8: db.Db.GetPush:input_type -> db.GetPushRequest
	11, // 9: db.Db.GetPushDetail:input_type -> db.GetPushDetailRequest
	1,  // 10: db.Db.Call:output_type -> db.CallResponse
	3,  // 11: db.Db.ClientStream:output_type -> db.ClientStreamResponse
	5,  // 12: db.Db.ServerStream:output_type -> db.ServerStreamResponse
	7,  // 13: db.Db.BidiStream:output_type -> db.BidiStreamResponse
	10, // 14: db.Db.GetPush:output_type -> db.GetPushResponse
	13, // 15: db.Db.GetPushDetail:output_type -> db.GetPushDetailResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_db_proto_init() }
func file_proto_db_proto_init() {
	if File_proto_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_db_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_db_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_db_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_db_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_db_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_db_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPushRequest); i {
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
		file_proto_db_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushData); i {
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
		file_proto_db_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPushResponse); i {
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
		file_proto_db_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPushDetailRequest); i {
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
		file_proto_db_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushDetailData); i {
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
		file_proto_db_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPushDetailResponse); i {
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
			RawDescriptor: file_proto_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_db_proto_goTypes,
		DependencyIndexes: file_proto_db_proto_depIdxs,
		MessageInfos:      file_proto_db_proto_msgTypes,
	}.Build()
	File_proto_db_proto = out.File
	file_proto_db_proto_rawDesc = nil
	file_proto_db_proto_goTypes = nil
	file_proto_db_proto_depIdxs = nil
}
