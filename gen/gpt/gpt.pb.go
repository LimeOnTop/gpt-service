// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: gpt.proto

package gpt

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

type UserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []string               `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	Preference    string                 `protobuf:"bytes,2,opt,name=preference,proto3" json:"preference,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	mi := &file_gpt_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gpt_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_gpt_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetProducts() []string {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *UserRequest) GetPreference() string {
	if x != nil {
		return x.Preference
	}
	return ""
}

type GPTResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ImageData     []byte                 `protobuf:"bytes,2,opt,name=image_data,json=imageData,proto3" json:"image_data,omitempty"`
	ImageFormat   string                 `protobuf:"bytes,3,opt,name=image_format,json=imageFormat,proto3" json:"image_format,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GPTResponse) Reset() {
	*x = GPTResponse{}
	mi := &file_gpt_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GPTResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GPTResponse) ProtoMessage() {}

func (x *GPTResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gpt_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GPTResponse.ProtoReflect.Descriptor instead.
func (*GPTResponse) Descriptor() ([]byte, []int) {
	return file_gpt_proto_rawDescGZIP(), []int{1}
}

func (x *GPTResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GPTResponse) GetImageData() []byte {
	if x != nil {
		return x.ImageData
	}
	return nil
}

func (x *GPTResponse) GetImageFormat() string {
	if x != nil {
		return x.ImageFormat
	}
	return ""
}

var File_gpt_proto protoreflect.FileDescriptor

const file_gpt_proto_rawDesc = "" +
	"\n" +
	"\tgpt.proto\x12\x03gpt\"I\n" +
	"\vUserRequest\x12\x1a\n" +
	"\bproducts\x18\x01 \x03(\tR\bproducts\x12\x1e\n" +
	"\n" +
	"preference\x18\x02 \x01(\tR\n" +
	"preference\"i\n" +
	"\vGPTResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x12\x1d\n" +
	"\n" +
	"image_data\x18\x02 \x01(\fR\timageData\x12!\n" +
	"\fimage_format\x18\x03 \x01(\tR\vimageFormat2L\n" +
	"\x0eRecommendation\x12:\n" +
	"\x14GetGPTRecommendation\x12\x10.gpt.UserRequest\x1a\x10.gpt.GPTResponseB\x15Z\x13gpt-service/gen/gptb\x06proto3"

var (
	file_gpt_proto_rawDescOnce sync.Once
	file_gpt_proto_rawDescData []byte
)

func file_gpt_proto_rawDescGZIP() []byte {
	file_gpt_proto_rawDescOnce.Do(func() {
		file_gpt_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_gpt_proto_rawDesc), len(file_gpt_proto_rawDesc)))
	})
	return file_gpt_proto_rawDescData
}

var file_gpt_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gpt_proto_goTypes = []any{
	(*UserRequest)(nil), // 0: gpt.UserRequest
	(*GPTResponse)(nil), // 1: gpt.GPTResponse
}
var file_gpt_proto_depIdxs = []int32{
	0, // 0: gpt.Recommendation.GetGPTRecommendation:input_type -> gpt.UserRequest
	1, // 1: gpt.Recommendation.GetGPTRecommendation:output_type -> gpt.GPTResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gpt_proto_init() }
func file_gpt_proto_init() {
	if File_gpt_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_gpt_proto_rawDesc), len(file_gpt_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gpt_proto_goTypes,
		DependencyIndexes: file_gpt_proto_depIdxs,
		MessageInfos:      file_gpt_proto_msgTypes,
	}.Build()
	File_gpt_proto = out.File
	file_gpt_proto_goTypes = nil
	file_gpt_proto_depIdxs = nil
}
