// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: cgentron/amazon/annotations.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
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

var file_cgentron_amazon_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Methods)(nil),
		Field:         12345679,
		Name:          "cgentron.amazon.methods",
		Tag:           "bytes,12345679,opt,name=methods",
		Filename:      "cgentron/amazon/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*Fields)(nil),
		Field:         12345679,
		Name:          "cgentron.amazon.fields",
		Tag:           "bytes,12345679,opt,name=fields",
		Filename:      "cgentron/amazon/annotations.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional cgentron.amazon.Methods methods = 12345679;
	E_Methods = &file_cgentron_amazon_annotations_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional cgentron.amazon.Fields fields = 12345679;
	E_Fields = &file_cgentron_amazon_annotations_proto_extTypes[1]
)

var File_cgentron_amazon_annotations_proto protoreflect.FileDescriptor

var file_cgentron_amazon_annotations_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x67, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x6e, 0x2f, 0x61, 0x6d, 0x61, 0x7a, 0x6f,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x67, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x61, 0x6d,
	0x61, 0x7a, 0x6f, 0x6e, 0x1a, 0x1d, 0x63, 0x67, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x6e, 0x2f, 0x61,
	0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x63, 0x67, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x6e, 0x2f, 0x61, 0x6d,
	0x61, 0x7a, 0x6f, 0x6e, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3a, 0x55, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x1e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xcf,
	0xc2, 0xf1, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x67, 0x65, 0x6e, 0x74, 0x72,
	0x6f, 0x6e, 0x2e, 0x61, 0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x3a, 0x51, 0x0a, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xcf, 0xc2, 0xf1, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63,
	0x67, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x61, 0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x42, 0x2c, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x67, 0x65, 0x6e,
	0x74, 0x72, 0x6f, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x61, 0x6d, 0x7a, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_cgentron_amazon_annotations_proto_goTypes = []interface{}{
	(*descriptorpb.MethodOptions)(nil), // 0: google.protobuf.MethodOptions
	(*descriptorpb.FieldOptions)(nil),  // 1: google.protobuf.FieldOptions
	(*Methods)(nil),                    // 2: cgentron.amazon.Methods
	(*Fields)(nil),                     // 3: cgentron.amazon.Fields
}
var file_cgentron_amazon_annotations_proto_depIdxs = []int32{
	0, // 0: cgentron.amazon.methods:extendee -> google.protobuf.MethodOptions
	1, // 1: cgentron.amazon.fields:extendee -> google.protobuf.FieldOptions
	2, // 2: cgentron.amazon.methods:type_name -> cgentron.amazon.Methods
	3, // 3: cgentron.amazon.fields:type_name -> cgentron.amazon.Fields
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cgentron_amazon_annotations_proto_init() }
func file_cgentron_amazon_annotations_proto_init() {
	if File_cgentron_amazon_annotations_proto != nil {
		return
	}
	file_cgentron_amazon_methods_proto_init()
	file_cgentron_amazon_fields_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cgentron_amazon_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_cgentron_amazon_annotations_proto_goTypes,
		DependencyIndexes: file_cgentron_amazon_annotations_proto_depIdxs,
		ExtensionInfos:    file_cgentron_amazon_annotations_proto_extTypes,
	}.Build()
	File_cgentron_amazon_annotations_proto = out.File
	file_cgentron_amazon_annotations_proto_rawDesc = nil
	file_cgentron_amazon_annotations_proto_goTypes = nil
	file_cgentron_amazon_annotations_proto_depIdxs = nil
}