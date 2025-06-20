// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: v1/user_error.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type UserServiceErrorReason int32

const (
	UserServiceErrorReason_UNKNOWN_ERROR   UserServiceErrorReason = 0
	UserServiceErrorReason_LOGIN_FAILED    UserServiceErrorReason = 1
	UserServiceErrorReason_REGISTER_FAILED UserServiceErrorReason = 2
)

// Enum value maps for UserServiceErrorReason.
var (
	UserServiceErrorReason_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "LOGIN_FAILED",
		2: "REGISTER_FAILED",
	}
	UserServiceErrorReason_value = map[string]int32{
		"UNKNOWN_ERROR":   0,
		"LOGIN_FAILED":    1,
		"REGISTER_FAILED": 2,
	}
)

func (x UserServiceErrorReason) Enum() *UserServiceErrorReason {
	p := new(UserServiceErrorReason)
	*p = x
	return p
}

func (x UserServiceErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserServiceErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_user_error_proto_enumTypes[0].Descriptor()
}

func (UserServiceErrorReason) Type() protoreflect.EnumType {
	return &file_v1_user_error_proto_enumTypes[0]
}

func (x UserServiceErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserServiceErrorReason.Descriptor instead.
func (UserServiceErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_v1_user_error_proto_rawDescGZIP(), []int{0}
}

var File_v1_user_error_proto protoreflect.FileDescriptor

const file_v1_user_error_proto_rawDesc = "" +
	"\n" +
	"\x13v1/user_error.proto\x12\x0fuser.service.v1\x1a\x13errors/errors.proto*X\n" +
	"\x16UserServiceErrorReason\x12\x11\n" +
	"\rUNKNOWN_ERROR\x10\x00\x12\x10\n" +
	"\fLOGIN_FAILED\x10\x01\x12\x13\n" +
	"\x0fREGISTER_FAILED\x10\x02\x1a\x04\xa0E\xf4\x03B\x16P\x01Z\x12user/service/v1;v1b\x06proto3"

var (
	file_v1_user_error_proto_rawDescOnce sync.Once
	file_v1_user_error_proto_rawDescData []byte
)

func file_v1_user_error_proto_rawDescGZIP() []byte {
	file_v1_user_error_proto_rawDescOnce.Do(func() {
		file_v1_user_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_user_error_proto_rawDesc), len(file_v1_user_error_proto_rawDesc)))
	})
	return file_v1_user_error_proto_rawDescData
}

var file_v1_user_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_user_error_proto_goTypes = []any{
	(UserServiceErrorReason)(0), // 0: user.service.v1.UserServiceErrorReason
}
var file_v1_user_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_user_error_proto_init() }
func file_v1_user_error_proto_init() {
	if File_v1_user_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_user_error_proto_rawDesc), len(file_v1_user_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_user_error_proto_goTypes,
		DependencyIndexes: file_v1_user_error_proto_depIdxs,
		EnumInfos:         file_v1_user_error_proto_enumTypes,
	}.Build()
	File_v1_user_error_proto = out.File
	file_v1_user_error_proto_goTypes = nil
	file_v1_user_error_proto_depIdxs = nil
}
