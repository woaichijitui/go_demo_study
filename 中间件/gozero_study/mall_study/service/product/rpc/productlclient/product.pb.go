// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0
// source: rpc/product.proto

package productlclient

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 产品创建
type CreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Desc          string                 `protobuf:"bytes,2,opt,name=Desc,proto3" json:"Desc,omitempty"`
	Stock         int64                  `protobuf:"varint,3,opt,name=Stock,proto3" json:"Stock,omitempty"`
	Amount        int64                  `protobuf:"varint,4,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Status        int64                  `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_rpc_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_rpc_product_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *CreateRequest) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *CreateRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type CreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	mi := &file_rpc_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_rpc_product_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 产品修改
type UpdateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Desc          string                 `protobuf:"bytes,3,opt,name=Desc,proto3" json:"Desc,omitempty"`
	Stock         int64                  `protobuf:"varint,4,opt,name=Stock,proto3" json:"Stock,omitempty"`
	Amount        int64                  `protobuf:"varint,5,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Status        int64                  `protobuf:"varint,6,opt,name=Status,proto3" json:"Status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	mi := &file_rpc_product_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_product_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_rpc_product_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *UpdateRequest) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *UpdateRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *UpdateRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type UpdateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	mi := &file_rpc_product_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_product_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_rpc_product_proto_rawDescGZIP(), []int{3}
}

// 产品删除
type RemoveRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveRequest) Reset() {
	*x = RemoveRequest{}
	mi := &file_rpc_product_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRequest) ProtoMessage() {}

func (x *RemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_product_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRequest.ProtoReflect.Descriptor instead.
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return file_rpc_product_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RemoveResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveResponse) Reset() {
	*x = RemoveResponse{}
	mi := &file_rpc_product_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveResponse) ProtoMessage() {}

func (x *RemoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_product_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveResponse.ProtoReflect.Descriptor instead.
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return file_rpc_product_proto_rawDescGZIP(), []int{5}
}

// 产品详情
type DetailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DetailRequest) Reset() {
	*x = DetailRequest{}
	mi := &file_rpc_product_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailRequest) ProtoMessage() {}

func (x *DetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_product_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailRequest.ProtoReflect.Descriptor instead.
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return file_rpc_product_proto_rawDescGZIP(), []int{6}
}

func (x *DetailRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DetailResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Desc          string                 `protobuf:"bytes,3,opt,name=Desc,proto3" json:"Desc,omitempty"`
	Stock         int64                  `protobuf:"varint,4,opt,name=Stock,proto3" json:"Stock,omitempty"`
	Amount        int64                  `protobuf:"varint,5,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Status        int64                  `protobuf:"varint,6,opt,name=Status,proto3" json:"Status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DetailResponse) Reset() {
	*x = DetailResponse{}
	mi := &file_rpc_product_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailResponse) ProtoMessage() {}

func (x *DetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_product_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailResponse.ProtoReflect.Descriptor instead.
func (*DetailResponse) Descriptor() ([]byte, []int) {
	return file_rpc_product_proto_rawDescGZIP(), []int{7}
}

func (x *DetailResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DetailResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DetailResponse) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *DetailResponse) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *DetailResponse) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *DetailResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_rpc_product_proto protoreflect.FileDescriptor

const file_rpc_product_proto_rawDesc = "" +
	"\n" +
	"\x11rpc/product.proto\x12\x0eproductlclient\"}\n" +
	"\rCreateRequest\x12\x12\n" +
	"\x04Name\x18\x01 \x01(\tR\x04Name\x12\x12\n" +
	"\x04Desc\x18\x02 \x01(\tR\x04Desc\x12\x14\n" +
	"\x05Stock\x18\x03 \x01(\x03R\x05Stock\x12\x16\n" +
	"\x06Amount\x18\x04 \x01(\x03R\x06Amount\x12\x16\n" +
	"\x06Status\x18\x05 \x01(\x03R\x06Status\" \n" +
	"\x0eCreateResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"\x8d\x01\n" +
	"\rUpdateRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x12\n" +
	"\x04Name\x18\x02 \x01(\tR\x04Name\x12\x12\n" +
	"\x04Desc\x18\x03 \x01(\tR\x04Desc\x12\x14\n" +
	"\x05Stock\x18\x04 \x01(\x03R\x05Stock\x12\x16\n" +
	"\x06Amount\x18\x05 \x01(\x03R\x06Amount\x12\x16\n" +
	"\x06Status\x18\x06 \x01(\x03R\x06Status\"\x10\n" +
	"\x0eUpdateResponse\"\x1f\n" +
	"\rRemoveRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"\x10\n" +
	"\x0eRemoveResponse\"\x1f\n" +
	"\rDetailRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"\x8e\x01\n" +
	"\x0eDetailResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x12\n" +
	"\x04Name\x18\x02 \x01(\tR\x04Name\x12\x12\n" +
	"\x04Desc\x18\x03 \x01(\tR\x04Desc\x12\x14\n" +
	"\x05Stock\x18\x04 \x01(\x03R\x05Stock\x12\x16\n" +
	"\x06Amount\x18\x05 \x01(\x03R\x06Amount\x12\x16\n" +
	"\x06Status\x18\x06 \x01(\x03R\x06Status2\xad\x02\n" +
	"\aProduct\x12G\n" +
	"\x06Create\x12\x1d.productlclient.CreateRequest\x1a\x1e.productlclient.CreateResponse\x12G\n" +
	"\x06Update\x12\x1d.productlclient.UpdateRequest\x1a\x1e.productlclient.UpdateResponse\x12G\n" +
	"\x06Remove\x12\x1d.productlclient.RemoveRequest\x1a\x1e.productlclient.RemoveResponse\x12G\n" +
	"\x06Detail\x12\x1d.productlclient.DetailRequest\x1a\x1e.productlclient.DetailResponseB\x12Z\x10./productlclientb\x06proto3"

var (
	file_rpc_product_proto_rawDescOnce sync.Once
	file_rpc_product_proto_rawDescData []byte
)

func file_rpc_product_proto_rawDescGZIP() []byte {
	file_rpc_product_proto_rawDescOnce.Do(func() {
		file_rpc_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_rpc_product_proto_rawDesc), len(file_rpc_product_proto_rawDesc)))
	})
	return file_rpc_product_proto_rawDescData
}

var file_rpc_product_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_rpc_product_proto_goTypes = []any{
	(*CreateRequest)(nil),  // 0: productlclient.CreateRequest
	(*CreateResponse)(nil), // 1: productlclient.CreateResponse
	(*UpdateRequest)(nil),  // 2: productlclient.UpdateRequest
	(*UpdateResponse)(nil), // 3: productlclient.UpdateResponse
	(*RemoveRequest)(nil),  // 4: productlclient.RemoveRequest
	(*RemoveResponse)(nil), // 5: productlclient.RemoveResponse
	(*DetailRequest)(nil),  // 6: productlclient.DetailRequest
	(*DetailResponse)(nil), // 7: productlclient.DetailResponse
}
var file_rpc_product_proto_depIdxs = []int32{
	0, // 0: productlclient.Product.Create:input_type -> productlclient.CreateRequest
	2, // 1: productlclient.Product.Update:input_type -> productlclient.UpdateRequest
	4, // 2: productlclient.Product.Remove:input_type -> productlclient.RemoveRequest
	6, // 3: productlclient.Product.Detail:input_type -> productlclient.DetailRequest
	1, // 4: productlclient.Product.Create:output_type -> productlclient.CreateResponse
	3, // 5: productlclient.Product.Update:output_type -> productlclient.UpdateResponse
	5, // 6: productlclient.Product.Remove:output_type -> productlclient.RemoveResponse
	7, // 7: productlclient.Product.Detail:output_type -> productlclient.DetailResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_product_proto_init() }
func file_rpc_product_proto_init() {
	if File_rpc_product_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_rpc_product_proto_rawDesc), len(file_rpc_product_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_product_proto_goTypes,
		DependencyIndexes: file_rpc_product_proto_depIdxs,
		MessageInfos:      file_rpc_product_proto_msgTypes,
	}.Build()
	File_rpc_product_proto = out.File
	file_rpc_product_proto_goTypes = nil
	file_rpc_product_proto_depIdxs = nil
}
