// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: enums/unit_enum.proto

package enums

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

type Unit int32

const (
	Unit_UNKNOWN   Unit = 0
	Unit_BIT       Unit = 1
	Unit_BYTE      Unit = 2
	Unit_KILO_BYTE Unit = 3
	Unit_MEGA_BYTE Unit = 4
	Unit_GIGA_BYTE Unit = 5
	Unit_TERA_BYTE Unit = 6
)

// Enum value maps for Unit.
var (
	Unit_name = map[int32]string{
		0: "UNKNOWN",
		1: "BIT",
		2: "BYTE",
		3: "KILO_BYTE",
		4: "MEGA_BYTE",
		5: "GIGA_BYTE",
		6: "TERA_BYTE",
	}
	Unit_value = map[string]int32{
		"UNKNOWN":   0,
		"BIT":       1,
		"BYTE":      2,
		"KILO_BYTE": 3,
		"MEGA_BYTE": 4,
		"GIGA_BYTE": 5,
		"TERA_BYTE": 6,
	}
)

func (x Unit) Enum() *Unit {
	p := new(Unit)
	*p = x
	return p
}

func (x Unit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Unit) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_unit_enum_proto_enumTypes[0].Descriptor()
}

func (Unit) Type() protoreflect.EnumType {
	return &file_enums_unit_enum_proto_enumTypes[0]
}

func (x Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Unit.Descriptor instead.
func (Unit) EnumDescriptor() ([]byte, []int) {
	return file_enums_unit_enum_proto_rawDescGZIP(), []int{0}
}

var File_enums_unit_enum_proto protoreflect.FileDescriptor

var file_enums_unit_enum_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x65, 0x6e, 0x75,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2a, 0x62, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x49,
	0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x59, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0d, 0x0a,
	0x09, 0x4b, 0x49, 0x4c, 0x4f, 0x5f, 0x42, 0x59, 0x54, 0x45, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09,
	0x4d, 0x45, 0x47, 0x41, 0x5f, 0x42, 0x59, 0x54, 0x45, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x47,
	0x49, 0x47, 0x41, 0x5f, 0x42, 0x59, 0x54, 0x45, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x45,
	0x52, 0x41, 0x5f, 0x42, 0x59, 0x54, 0x45, 0x10, 0x06, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_enums_unit_enum_proto_rawDescOnce sync.Once
	file_enums_unit_enum_proto_rawDescData []byte
)

func file_enums_unit_enum_proto_rawDescGZIP() []byte {
	file_enums_unit_enum_proto_rawDescOnce.Do(func() {
		file_enums_unit_enum_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_enums_unit_enum_proto_rawDesc), len(file_enums_unit_enum_proto_rawDesc)))
	})
	return file_enums_unit_enum_proto_rawDescData
}

var file_enums_unit_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_unit_enum_proto_goTypes = []any{
	(Unit)(0), // 0: protos.enums.Unit
}
var file_enums_unit_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_unit_enum_proto_init() }
func file_enums_unit_enum_proto_init() {
	if File_enums_unit_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_enums_unit_enum_proto_rawDesc), len(file_enums_unit_enum_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_unit_enum_proto_goTypes,
		DependencyIndexes: file_enums_unit_enum_proto_depIdxs,
		EnumInfos:         file_enums_unit_enum_proto_enumTypes,
	}.Build()
	File_enums_unit_enum_proto = out.File
	file_enums_unit_enum_proto_goTypes = nil
	file_enums_unit_enum_proto_depIdxs = nil
}
