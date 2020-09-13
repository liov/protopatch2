// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: tests/enum/enum_stringer.proto

// clang-format off

package enum

import (
	_ "github.com/alta/protopatch/patch/gopb"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CustomStringerEnum int32

const (
	CustomStringerEnum_CUSTOM_STRINGER_INVALID CustomStringerEnum = 0
	CustomStringerEnum_CUSTOM_STRINGER_A       CustomStringerEnum = 1
	CustomStringerEnum_CUSTOM_STRINGER_B       CustomStringerEnum = 2
	CustomStringerEnum_CUSTOM_STRINGER_C       CustomStringerEnum = 3
)

// Enum value maps for CustomStringerEnum.
var (
	CustomStringerEnum_name = map[int32]string{
		0: "CUSTOM_STRINGER_INVALID",
		1: "CUSTOM_STRINGER_A",
		2: "CUSTOM_STRINGER_B",
		3: "CUSTOM_STRINGER_C",
	}
	CustomStringerEnum_value = map[string]int32{
		"CUSTOM_STRINGER_INVALID": 0,
		"CUSTOM_STRINGER_A":       1,
		"CUSTOM_STRINGER_B":       2,
		"CUSTOM_STRINGER_C":       3,
	}
)

func (x CustomStringerEnum) Enum() *CustomStringerEnum {
	p := new(CustomStringerEnum)
	*p = x
	return p
}

func (x CustomStringerEnum) OrigString() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomStringerEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_enum_enum_stringer_proto_enumTypes[0].Descriptor()
}

func (CustomStringerEnum) Type() protoreflect.EnumType {
	return &file_tests_enum_enum_stringer_proto_enumTypes[0]
}

func (x CustomStringerEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomStringerEnum.Descriptor instead.
func (CustomStringerEnum) EnumDescriptor() ([]byte, []int) {
	return file_tests_enum_enum_stringer_proto_rawDescGZIP(), []int{0}
}

type DefaultStringerEnum int32

const (
	DefaultStringerEnum_DEFAULT_STRINGER_UNSET   DefaultStringerEnum = 0
	DefaultStringerEnum_DEFAULT_STRINGER_EXAMPLE DefaultStringerEnum = 1
)

// Enum value maps for DefaultStringerEnum.
var (
	DefaultStringerEnum_name = map[int32]string{
		0: "DEFAULT_STRINGER_UNSET",
		1: "DEFAULT_STRINGER_EXAMPLE",
	}
	DefaultStringerEnum_value = map[string]int32{
		"DEFAULT_STRINGER_UNSET":   0,
		"DEFAULT_STRINGER_EXAMPLE": 1,
	}
)

func (x DefaultStringerEnum) Enum() *DefaultStringerEnum {
	p := new(DefaultStringerEnum)
	*p = x
	return p
}

func (x DefaultStringerEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DefaultStringerEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_enum_enum_stringer_proto_enumTypes[1].Descriptor()
}

func (DefaultStringerEnum) Type() protoreflect.EnumType {
	return &file_tests_enum_enum_stringer_proto_enumTypes[1]
}

func (x DefaultStringerEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DefaultStringerEnum.Descriptor instead.
func (DefaultStringerEnum) EnumDescriptor() ([]byte, []int) {
	return file_tests_enum_enum_stringer_proto_rawDescGZIP(), []int{1}
}

var File_tests_enum_enum_stringer_proto protoreflect.FileDescriptor

var file_tests_enum_enum_stringer_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x1a, 0x0e, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x2f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x89, 0x01, 0x0a,
	0x12, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x45,
	0x6e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x53, 0x54,
	0x52, 0x49, 0x4e, 0x47, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x15, 0x0a, 0x11, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e,
	0x47, 0x45, 0x52, 0x5f, 0x41, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x55, 0x53, 0x54, 0x4f,
	0x4d, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x45, 0x52, 0x5f, 0x42, 0x10, 0x02, 0x12, 0x15,
	0x0a, 0x11, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x45,
	0x52, 0x5f, 0x43, 0x10, 0x03, 0x1a, 0x11, 0xca, 0xb5, 0x03, 0x0d, 0xf2, 0x01, 0x0a, 0x4f, 0x72,
	0x69, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2a, 0x4f, 0x0a, 0x13, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x12,
	0x1a, 0x0a, 0x16, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e,
	0x47, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x44,
	0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x45, 0x52, 0x5f,
	0x45, 0x58, 0x41, 0x4d, 0x50, 0x4c, 0x45, 0x10, 0x01, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_enum_enum_stringer_proto_rawDescOnce sync.Once
	file_tests_enum_enum_stringer_proto_rawDescData = file_tests_enum_enum_stringer_proto_rawDesc
)

func file_tests_enum_enum_stringer_proto_rawDescGZIP() []byte {
	file_tests_enum_enum_stringer_proto_rawDescOnce.Do(func() {
		file_tests_enum_enum_stringer_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_enum_enum_stringer_proto_rawDescData)
	})
	return file_tests_enum_enum_stringer_proto_rawDescData
}

var file_tests_enum_enum_stringer_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_tests_enum_enum_stringer_proto_goTypes = []interface{}{
	(CustomStringerEnum)(0),  // 0: tests.enum.CustomStringerEnum
	(DefaultStringerEnum)(0), // 1: tests.enum.DefaultStringerEnum
}
var file_tests_enum_enum_stringer_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tests_enum_enum_stringer_proto_init() }
func file_tests_enum_enum_stringer_proto_init() {
	if File_tests_enum_enum_stringer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_enum_enum_stringer_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_enum_enum_stringer_proto_goTypes,
		DependencyIndexes: file_tests_enum_enum_stringer_proto_depIdxs,
		EnumInfos:         file_tests_enum_enum_stringer_proto_enumTypes,
	}.Build()
	File_tests_enum_enum_stringer_proto = out.File
	file_tests_enum_enum_stringer_proto_rawDesc = nil
	file_tests_enum_enum_stringer_proto_goTypes = nil
	file_tests_enum_enum_stringer_proto_depIdxs = nil
}