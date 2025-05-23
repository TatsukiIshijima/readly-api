// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: readly/v1/rpc_refresh_token.proto

package pb

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

type RefreshTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RefreshToken  string                 `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshTokenRequest) Reset() {
	*x = RefreshTokenRequest{}
	mi := &file_readly_v1_rpc_refresh_token_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRequest) ProtoMessage() {}

func (x *RefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_readly_v1_rpc_refresh_token_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_readly_v1_rpc_refresh_token_proto_rawDescGZIP(), []int{0}
}

func (x *RefreshTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RefreshTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshTokenResponse) Reset() {
	*x = RefreshTokenResponse{}
	mi := &file_readly_v1_rpc_refresh_token_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenResponse) ProtoMessage() {}

func (x *RefreshTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_readly_v1_rpc_refresh_token_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenResponse.ProtoReflect.Descriptor instead.
func (*RefreshTokenResponse) Descriptor() ([]byte, []int) {
	return file_readly_v1_rpc_refresh_token_proto_rawDescGZIP(), []int{1}
}

func (x *RefreshTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

var File_readly_v1_rpc_refresh_token_proto protoreflect.FileDescriptor

const file_readly_v1_rpc_refresh_token_proto_rawDesc = "" +
	"\n" +
	"!readly/v1/rpc_refresh_token.proto\x12\treadly.v1\":\n" +
	"\x13RefreshTokenRequest\x12#\n" +
	"\rrefresh_token\x18\x01 \x01(\tR\frefreshToken\"9\n" +
	"\x14RefreshTokenResponse\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessTokenB\x82\x01\n" +
	"\rcom.readly.v1B\x14RpcRefreshTokenProtoP\x01Z\x16readly/pb/readly/v1;pb\xa2\x02\x03RXX\xaa\x02\tReadly.V1\xca\x02\tReadly\\V1\xe2\x02\x15Readly\\V1\\GPBMetadata\xea\x02\n" +
	"Readly::V1b\x06proto3"

var (
	file_readly_v1_rpc_refresh_token_proto_rawDescOnce sync.Once
	file_readly_v1_rpc_refresh_token_proto_rawDescData []byte
)

func file_readly_v1_rpc_refresh_token_proto_rawDescGZIP() []byte {
	file_readly_v1_rpc_refresh_token_proto_rawDescOnce.Do(func() {
		file_readly_v1_rpc_refresh_token_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_readly_v1_rpc_refresh_token_proto_rawDesc), len(file_readly_v1_rpc_refresh_token_proto_rawDesc)))
	})
	return file_readly_v1_rpc_refresh_token_proto_rawDescData
}

var file_readly_v1_rpc_refresh_token_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_readly_v1_rpc_refresh_token_proto_goTypes = []any{
	(*RefreshTokenRequest)(nil),  // 0: readly.v1.RefreshTokenRequest
	(*RefreshTokenResponse)(nil), // 1: readly.v1.RefreshTokenResponse
}
var file_readly_v1_rpc_refresh_token_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_readly_v1_rpc_refresh_token_proto_init() }
func file_readly_v1_rpc_refresh_token_proto_init() {
	if File_readly_v1_rpc_refresh_token_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_readly_v1_rpc_refresh_token_proto_rawDesc), len(file_readly_v1_rpc_refresh_token_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_readly_v1_rpc_refresh_token_proto_goTypes,
		DependencyIndexes: file_readly_v1_rpc_refresh_token_proto_depIdxs,
		MessageInfos:      file_readly_v1_rpc_refresh_token_proto_msgTypes,
	}.Build()
	File_readly_v1_rpc_refresh_token_proto = out.File
	file_readly_v1_rpc_refresh_token_proto_goTypes = nil
	file_readly_v1_rpc_refresh_token_proto_depIdxs = nil
}
