// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protoc-gen-swagger/options/annotations.proto

package options

import (
	fmt "fmt"
	proto "github.com/catper/protobuf/proto"
	descriptor "github.com/catper/protobuf/protoc-gen-go/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

var E_Openapiv2Swagger = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*Swagger)(nil),
	Field:         1042,
	Name:          "grpc.gateway.protoc_gen_swagger.options.openapiv2_swagger",
	Tag:           "bytes,1042,opt,name=openapiv2_swagger",
	Filename:      "protoc-gen-swagger/options/annotations.proto",
}

var E_Openapiv2Operation = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*Operation)(nil),
	Field:         1042,
	Name:          "grpc.gateway.protoc_gen_swagger.options.openapiv2_operation",
	Tag:           "bytes,1042,opt,name=openapiv2_operation",
	Filename:      "protoc-gen-swagger/options/annotations.proto",
}

var E_Openapiv2Schema = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*Schema)(nil),
	Field:         1042,
	Name:          "grpc.gateway.protoc_gen_swagger.options.openapiv2_schema",
	Tag:           "bytes,1042,opt,name=openapiv2_schema",
	Filename:      "protoc-gen-swagger/options/annotations.proto",
}

var E_Openapiv2Tag = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*Tag)(nil),
	Field:         1042,
	Name:          "grpc.gateway.protoc_gen_swagger.options.openapiv2_tag",
	Tag:           "bytes,1042,opt,name=openapiv2_tag",
	Filename:      "protoc-gen-swagger/options/annotations.proto",
}

var E_Openapiv2Field = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*JSONSchema)(nil),
	Field:         1042,
	Name:          "grpc.gateway.protoc_gen_swagger.options.openapiv2_field",
	Tag:           "bytes,1042,opt,name=openapiv2_field",
	Filename:      "protoc-gen-swagger/options/annotations.proto",
}

func init() {
	proto.RegisterExtension(E_Openapiv2Swagger)
	proto.RegisterExtension(E_Openapiv2Operation)
	proto.RegisterExtension(E_Openapiv2Schema)
	proto.RegisterExtension(E_Openapiv2Tag)
	proto.RegisterExtension(E_Openapiv2Field)
}

func init() {
	proto.RegisterFile("protoc-gen-swagger/options/annotations.proto", fileDescriptor_a6a34ca6badab664)
}

var fileDescriptor_a6a34ca6badab664 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4f, 0xea, 0x40,
	0x14, 0xc5, 0xc3, 0xe6, 0xe5, 0xa5, 0xef, 0xa9, 0x58, 0x37, 0x86, 0xf8, 0x87, 0x9d, 0xc6, 0xc0,
	0x8c, 0x81, 0x5d, 0x77, 0x6a, 0xe2, 0xc2, 0x44, 0x49, 0x0a, 0x2b, 0x37, 0x64, 0x18, 0x2e, 0x97,
	0x49, 0x4a, 0xef, 0x64, 0x66, 0x80, 0x90, 0xb0, 0xf4, 0x13, 0xf8, 0x89, 0x8d, 0xd3, 0xd2, 0x9a,
	0x8a, 0xa6, 0xbb, 0xce, 0xe9, 0xbd, 0xe7, 0x77, 0x7a, 0x3a, 0x41, 0x47, 0x1b, 0x72, 0x24, 0xbb,
	0x08, 0x69, 0xd7, 0xae, 0x05, 0x22, 0x18, 0x4e, 0xda, 0x29, 0x4a, 0x2d, 0x17, 0x69, 0x4a, 0x4e,
	0xf8, 0x67, 0xe6, 0xc7, 0xc2, 0x2b, 0x34, 0x5a, 0x32, 0x14, 0x0e, 0xd6, 0x62, 0x93, 0x69, 0x72,
	0x8c, 0x90, 0x8e, 0xf3, 0x55, 0x96, 0xaf, 0xb6, 0xda, 0x48, 0x84, 0x09, 0x70, 0x3f, 0x32, 0x59,
	0xce, 0xf8, 0x14, 0xac, 0x34, 0x4a, 0x3b, 0x32, 0xd9, 0x5a, 0xeb, 0xe6, 0x17, 0x30, 0x69, 0x48,
	0x85, 0x56, 0xab, 0x5e, 0x36, 0x1b, 0x6d, 0x83, 0xe3, 0x42, 0xda, 0xa1, 0xc2, 0x33, 0x96, 0x31,
	0xd8, 0x8e, 0xc1, 0x1e, 0x55, 0x02, 0x83, 0xcc, 0xe2, 0xf4, 0xfd, 0x6f, 0xbb, 0x71, 0xfd, 0xaf,
	0x77, 0xcb, 0x6a, 0x26, 0x66, 0xc3, 0xec, 0x1c, 0x37, 0x0b, 0x52, 0xae, 0x44, 0x6f, 0x8d, 0xe0,
	0xa4, 0xc4, 0x93, 0x06, 0xe3, 0x3b, 0x09, 0x2f, 0xbe, 0x05, 0x78, 0x06, 0x37, 0xa7, 0x69, 0x25,
	0x42, 0xaf, 0x76, 0x84, 0xc1, 0xce, 0x3a, 0x0e, 0x0b, 0x5e, 0xa1, 0x45, 0xdb, 0xa0, 0xf9, 0xa5,
	0x04, 0x39, 0x87, 0x85, 0x08, 0x2f, 0xf7, 0x44, 0xb0, 0x56, 0x60, 0xb5, 0x06, 0x5e, 0xbf, 0x06,
	0x6f, 0x1c, 0x1f, 0x95, 0x2d, 0x78, 0x21, 0xb2, 0xc1, 0x41, 0x49, 0x77, 0x02, 0xf7, 0xa0, 0x87,
	0x60, 0x56, 0x4a, 0x56, 0xd1, 0x9d, 0xda, 0xe8, 0x91, 0xc0, 0xf8, 0x7f, 0x01, 0x19, 0x09, 0x8c,
	0xb6, 0x41, 0x99, 0x63, 0x3c, 0x53, 0x90, 0x4c, 0xc3, 0xf3, 0x3d, 0x7f, 0x1d, 0x92, 0x6a, 0xe7,
	0xfd, 0xda, 0xd0, 0xa7, 0xe1, 0xe0, 0x25, 0xff, 0xe6, 0xc3, 0x82, 0xe5, 0x2d, 0xef, 0x1f, 0x5e,
	0xef, 0x50, 0xb9, 0xf9, 0x72, 0xc2, 0x24, 0x2d, 0xf8, 0xa7, 0x61, 0x17, 0x24, 0xd9, 0x8d, 0x75,
	0x90, 0x1f, 0x73, 0x7f, 0xfe, 0xf3, 0x55, 0x9e, 0xfc, 0xf1, 0xef, 0xfa, 0x1f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x59, 0x78, 0xb0, 0x03, 0x68, 0x03, 0x00, 0x00,
}
