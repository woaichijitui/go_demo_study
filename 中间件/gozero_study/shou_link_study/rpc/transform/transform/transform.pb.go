// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0
// source: transform.proto

package transform

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

type ExpandReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Shorten       string                 `protobuf:"bytes,1,opt,name=shorten,proto3" json:"shorten,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExpandReq) Reset() {
	*x = ExpandReq{}
	mi := &file_transform_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExpandReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandReq) ProtoMessage() {}

func (x *ExpandReq) ProtoReflect() protoreflect.Message {
	mi := &file_transform_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandReq.ProtoReflect.Descriptor instead.
func (*ExpandReq) Descriptor() ([]byte, []int) {
	return file_transform_proto_rawDescGZIP(), []int{0}
}

func (x *ExpandReq) GetShorten() string {
	if x != nil {
		return x.Shorten
	}
	return ""
}

type ExpandResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExpandResp) Reset() {
	*x = ExpandResp{}
	mi := &file_transform_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExpandResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandResp) ProtoMessage() {}

func (x *ExpandResp) ProtoReflect() protoreflect.Message {
	mi := &file_transform_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandResp.ProtoReflect.Descriptor instead.
func (*ExpandResp) Descriptor() ([]byte, []int) {
	return file_transform_proto_rawDescGZIP(), []int{1}
}

func (x *ExpandResp) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ShortenReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortenReq) Reset() {
	*x = ShortenReq{}
	mi := &file_transform_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenReq) ProtoMessage() {}

func (x *ShortenReq) ProtoReflect() protoreflect.Message {
	mi := &file_transform_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenReq.ProtoReflect.Descriptor instead.
func (*ShortenReq) Descriptor() ([]byte, []int) {
	return file_transform_proto_rawDescGZIP(), []int{2}
}

func (x *ShortenReq) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ShortenResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Shorten       string                 `protobuf:"bytes,1,opt,name=shorten,proto3" json:"shorten,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortenResp) Reset() {
	*x = ShortenResp{}
	mi := &file_transform_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenResp) ProtoMessage() {}

func (x *ShortenResp) ProtoReflect() protoreflect.Message {
	mi := &file_transform_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenResp.ProtoReflect.Descriptor instead.
func (*ShortenResp) Descriptor() ([]byte, []int) {
	return file_transform_proto_rawDescGZIP(), []int{3}
}

func (x *ShortenResp) GetShorten() string {
	if x != nil {
		return x.Shorten
	}
	return ""
}

var File_transform_proto protoreflect.FileDescriptor

const file_transform_proto_rawDesc = "" +
	"\n" +
	"\x0ftransform.proto\x12\ttransform\"%\n" +
	"\texpandReq\x12\x18\n" +
	"\ashorten\x18\x01 \x01(\tR\ashorten\"\x1e\n" +
	"\n" +
	"expandResp\x12\x10\n" +
	"\x03url\x18\x01 \x01(\tR\x03url\"\x1e\n" +
	"\n" +
	"shortenReq\x12\x10\n" +
	"\x03url\x18\x01 \x01(\tR\x03url\"'\n" +
	"\vshortenResp\x12\x18\n" +
	"\ashorten\x18\x01 \x01(\tR\ashorten2~\n" +
	"\vtransformer\x125\n" +
	"\x06expand\x12\x14.transform.expandReq\x1a\x15.transform.expandResp\x128\n" +
	"\ashorten\x12\x15.transform.shortenReq\x1a\x16.transform.shortenRespB\rZ\v./transformb\x06proto3"

var (
	file_transform_proto_rawDescOnce sync.Once
	file_transform_proto_rawDescData []byte
)

func file_transform_proto_rawDescGZIP() []byte {
	file_transform_proto_rawDescOnce.Do(func() {
		file_transform_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_transform_proto_rawDesc), len(file_transform_proto_rawDesc)))
	})
	return file_transform_proto_rawDescData
}

var file_transform_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_transform_proto_goTypes = []any{
	(*ExpandReq)(nil),   // 0: transform.expandReq
	(*ExpandResp)(nil),  // 1: transform.expandResp
	(*ShortenReq)(nil),  // 2: transform.shortenReq
	(*ShortenResp)(nil), // 3: transform.shortenResp
}
var file_transform_proto_depIdxs = []int32{
	0, // 0: transform.transformer.expand:input_type -> transform.expandReq
	2, // 1: transform.transformer.shorten:input_type -> transform.shortenReq
	1, // 2: transform.transformer.expand:output_type -> transform.expandResp
	3, // 3: transform.transformer.shorten:output_type -> transform.shortenResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transform_proto_init() }
func file_transform_proto_init() {
	if File_transform_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_transform_proto_rawDesc), len(file_transform_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transform_proto_goTypes,
		DependencyIndexes: file_transform_proto_depIdxs,
		MessageInfos:      file_transform_proto_msgTypes,
	}.Build()
	File_transform_proto = out.File
	file_transform_proto_goTypes = nil
	file_transform_proto_depIdxs = nil
}
