// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: rpc_delete_book.proto

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

type DeleteBookRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookId        int64                  `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteBookRequest) Reset() {
	*x = DeleteBookRequest{}
	mi := &file_rpc_delete_book_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookRequest) ProtoMessage() {}

func (x *DeleteBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_delete_book_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookRequest.ProtoReflect.Descriptor instead.
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return file_rpc_delete_book_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteBookRequest) GetBookId() int64 {
	if x != nil {
		return x.BookId
	}
	return 0
}

var File_rpc_delete_book_proto protoreflect.FileDescriptor

const file_rpc_delete_book_proto_rawDesc = "" +
	"\n" +
	"\x15rpc_delete_book.proto\x12\x02pb\",\n" +
	"\x11DeleteBookRequest\x12\x17\n" +
	"\abook_id\x18\x01 \x01(\x03R\x06bookIdBO\n" +
	"\x06com.pbB\x12RpcDeleteBookProtoP\x01Z\treadly/pb\xa2\x02\x03PXX\xaa\x02\x02Pb\xca\x02\x02Pb\xe2\x02\x0ePb\\GPBMetadata\xea\x02\x02Pbb\x06proto3"

var (
	file_rpc_delete_book_proto_rawDescOnce sync.Once
	file_rpc_delete_book_proto_rawDescData []byte
)

func file_rpc_delete_book_proto_rawDescGZIP() []byte {
	file_rpc_delete_book_proto_rawDescOnce.Do(func() {
		file_rpc_delete_book_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_rpc_delete_book_proto_rawDesc), len(file_rpc_delete_book_proto_rawDesc)))
	})
	return file_rpc_delete_book_proto_rawDescData
}

var file_rpc_delete_book_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rpc_delete_book_proto_goTypes = []any{
	(*DeleteBookRequest)(nil), // 0: pb.DeleteBookRequest
}
var file_rpc_delete_book_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_delete_book_proto_init() }
func file_rpc_delete_book_proto_init() {
	if File_rpc_delete_book_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_rpc_delete_book_proto_rawDesc), len(file_rpc_delete_book_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_delete_book_proto_goTypes,
		DependencyIndexes: file_rpc_delete_book_proto_depIdxs,
		MessageInfos:      file_rpc_delete_book_proto_msgTypes,
	}.Build()
	File_rpc_delete_book_proto = out.File
	file_rpc_delete_book_proto_goTypes = nil
	file_rpc_delete_book_proto_depIdxs = nil
}
