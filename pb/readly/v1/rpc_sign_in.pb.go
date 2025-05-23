// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: readly/v1/rpc_sign_in.proto

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

type SignInRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	mi := &file_readly_v1_rpc_sign_in_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_readly_v1_rpc_sign_in_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_readly_v1_rpc_sign_in_proto_rawDescGZIP(), []int{0}
}

func (x *SignInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignInResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	UserId        int64                  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name          string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignInResponse) Reset() {
	*x = SignInResponse{}
	mi := &file_readly_v1_rpc_sign_in_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResponse) ProtoMessage() {}

func (x *SignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_readly_v1_rpc_sign_in_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResponse.ProtoReflect.Descriptor instead.
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return file_readly_v1_rpc_sign_in_proto_rawDescGZIP(), []int{1}
}

func (x *SignInResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *SignInResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *SignInResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SignInResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SignInResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_readly_v1_rpc_sign_in_proto protoreflect.FileDescriptor

const file_readly_v1_rpc_sign_in_proto_rawDesc = "" +
	"\n" +
	"\x1breadly/v1/rpc_sign_in.proto\x12\treadly.v1\"A\n" +
	"\rSignInRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"\x9b\x01\n" +
	"\x0eSignInResponse\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x12#\n" +
	"\rrefresh_token\x18\x02 \x01(\tR\frefreshToken\x12\x17\n" +
	"\auser_id\x18\x03 \x01(\x03R\x06userId\x12\x12\n" +
	"\x04name\x18\x04 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x05 \x01(\tR\x05emailB|\n" +
	"\rcom.readly.v1B\x0eRpcSignInProtoP\x01Z\x16readly/pb/readly/v1;pb\xa2\x02\x03RXX\xaa\x02\tReadly.V1\xca\x02\tReadly\\V1\xe2\x02\x15Readly\\V1\\GPBMetadata\xea\x02\n" +
	"Readly::V1b\x06proto3"

var (
	file_readly_v1_rpc_sign_in_proto_rawDescOnce sync.Once
	file_readly_v1_rpc_sign_in_proto_rawDescData []byte
)

func file_readly_v1_rpc_sign_in_proto_rawDescGZIP() []byte {
	file_readly_v1_rpc_sign_in_proto_rawDescOnce.Do(func() {
		file_readly_v1_rpc_sign_in_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_readly_v1_rpc_sign_in_proto_rawDesc), len(file_readly_v1_rpc_sign_in_proto_rawDesc)))
	})
	return file_readly_v1_rpc_sign_in_proto_rawDescData
}

var file_readly_v1_rpc_sign_in_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_readly_v1_rpc_sign_in_proto_goTypes = []any{
	(*SignInRequest)(nil),  // 0: readly.v1.SignInRequest
	(*SignInResponse)(nil), // 1: readly.v1.SignInResponse
}
var file_readly_v1_rpc_sign_in_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_readly_v1_rpc_sign_in_proto_init() }
func file_readly_v1_rpc_sign_in_proto_init() {
	if File_readly_v1_rpc_sign_in_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_readly_v1_rpc_sign_in_proto_rawDesc), len(file_readly_v1_rpc_sign_in_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_readly_v1_rpc_sign_in_proto_goTypes,
		DependencyIndexes: file_readly_v1_rpc_sign_in_proto_depIdxs,
		MessageInfos:      file_readly_v1_rpc_sign_in_proto_msgTypes,
	}.Build()
	File_readly_v1_rpc_sign_in_proto = out.File
	file_readly_v1_rpc_sign_in_proto_goTypes = nil
	file_readly_v1_rpc_sign_in_proto_depIdxs = nil
}
