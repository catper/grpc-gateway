// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/internal/proto/examplepb/echo_service.proto

// Echo Service
//
// Echo Service API consists of a single service which returns
// a message.

package examplepb

import (
	context "context"
	fmt "fmt"
	proto "github.com/catper/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Embedded represents a message embedded in SimpleMessage.
type Embedded struct {
	// Types that are valid to be assigned to Mark:
	//	*Embedded_Progress
	//	*Embedded_Note
	Mark                 isEmbedded_Mark `protobuf_oneof:"mark"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Embedded) Reset()         { *m = Embedded{} }
func (m *Embedded) String() string { return proto.CompactTextString(m) }
func (*Embedded) ProtoMessage()    {}
func (*Embedded) Descriptor() ([]byte, []int) {
	return fileDescriptor_2919b738e91bb561, []int{0}
}

func (m *Embedded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Embedded.Unmarshal(m, b)
}
func (m *Embedded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Embedded.Marshal(b, m, deterministic)
}
func (m *Embedded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Embedded.Merge(m, src)
}
func (m *Embedded) XXX_Size() int {
	return xxx_messageInfo_Embedded.Size(m)
}
func (m *Embedded) XXX_DiscardUnknown() {
	xxx_messageInfo_Embedded.DiscardUnknown(m)
}

var xxx_messageInfo_Embedded proto.InternalMessageInfo

type isEmbedded_Mark interface {
	isEmbedded_Mark()
}

type Embedded_Progress struct {
	Progress int64 `protobuf:"varint,1,opt,name=progress,proto3,oneof"`
}

type Embedded_Note struct {
	Note string `protobuf:"bytes,2,opt,name=note,proto3,oneof"`
}

func (*Embedded_Progress) isEmbedded_Mark() {}

func (*Embedded_Note) isEmbedded_Mark() {}

func (m *Embedded) GetMark() isEmbedded_Mark {
	if m != nil {
		return m.Mark
	}
	return nil
}

func (m *Embedded) GetProgress() int64 {
	if x, ok := m.GetMark().(*Embedded_Progress); ok {
		return x.Progress
	}
	return 0
}

func (m *Embedded) GetNote() string {
	if x, ok := m.GetMark().(*Embedded_Note); ok {
		return x.Note
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Embedded) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Embedded_Progress)(nil),
		(*Embedded_Note)(nil),
	}
}

// SimpleMessage represents a simple message sent to the Echo service.
type SimpleMessage struct {
	// Id represents the message identifier.
	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Num int64  `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	// Types that are valid to be assigned to Code:
	//	*SimpleMessage_LineNum
	//	*SimpleMessage_Lang
	Code   isSimpleMessage_Code `protobuf_oneof:"code"`
	Status *Embedded            `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	// Types that are valid to be assigned to Ext:
	//	*SimpleMessage_En
	//	*SimpleMessage_No
	Ext                  isSimpleMessage_Ext `protobuf_oneof:"ext"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SimpleMessage) Reset()         { *m = SimpleMessage{} }
func (m *SimpleMessage) String() string { return proto.CompactTextString(m) }
func (*SimpleMessage) ProtoMessage()    {}
func (*SimpleMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_2919b738e91bb561, []int{1}
}

func (m *SimpleMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleMessage.Unmarshal(m, b)
}
func (m *SimpleMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleMessage.Marshal(b, m, deterministic)
}
func (m *SimpleMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleMessage.Merge(m, src)
}
func (m *SimpleMessage) XXX_Size() int {
	return xxx_messageInfo_SimpleMessage.Size(m)
}
func (m *SimpleMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleMessage proto.InternalMessageInfo

func (m *SimpleMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SimpleMessage) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type isSimpleMessage_Code interface {
	isSimpleMessage_Code()
}

type SimpleMessage_LineNum struct {
	LineNum int64 `protobuf:"varint,3,opt,name=line_num,json=lineNum,proto3,oneof"`
}

type SimpleMessage_Lang struct {
	Lang string `protobuf:"bytes,4,opt,name=lang,proto3,oneof"`
}

func (*SimpleMessage_LineNum) isSimpleMessage_Code() {}

func (*SimpleMessage_Lang) isSimpleMessage_Code() {}

func (m *SimpleMessage) GetCode() isSimpleMessage_Code {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *SimpleMessage) GetLineNum() int64 {
	if x, ok := m.GetCode().(*SimpleMessage_LineNum); ok {
		return x.LineNum
	}
	return 0
}

func (m *SimpleMessage) GetLang() string {
	if x, ok := m.GetCode().(*SimpleMessage_Lang); ok {
		return x.Lang
	}
	return ""
}

func (m *SimpleMessage) GetStatus() *Embedded {
	if m != nil {
		return m.Status
	}
	return nil
}

type isSimpleMessage_Ext interface {
	isSimpleMessage_Ext()
}

type SimpleMessage_En struct {
	En int64 `protobuf:"varint,6,opt,name=en,proto3,oneof"`
}

type SimpleMessage_No struct {
	No *Embedded `protobuf:"bytes,7,opt,name=no,proto3,oneof"`
}

func (*SimpleMessage_En) isSimpleMessage_Ext() {}

func (*SimpleMessage_No) isSimpleMessage_Ext() {}

func (m *SimpleMessage) GetExt() isSimpleMessage_Ext {
	if m != nil {
		return m.Ext
	}
	return nil
}

func (m *SimpleMessage) GetEn() int64 {
	if x, ok := m.GetExt().(*SimpleMessage_En); ok {
		return x.En
	}
	return 0
}

func (m *SimpleMessage) GetNo() *Embedded {
	if x, ok := m.GetExt().(*SimpleMessage_No); ok {
		return x.No
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SimpleMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SimpleMessage_LineNum)(nil),
		(*SimpleMessage_Lang)(nil),
		(*SimpleMessage_En)(nil),
		(*SimpleMessage_No)(nil),
	}
}

func init() {
	proto.RegisterType((*Embedded)(nil), "grpc.gateway.examples.internal.examplepb.Embedded")
	proto.RegisterType((*SimpleMessage)(nil), "grpc.gateway.examples.internal.examplepb.SimpleMessage")
}

func init() {
	proto.RegisterFile("examples/internal/proto/examplepb/echo_service.proto", fileDescriptor_2919b738e91bb561)
}

var fileDescriptor_2919b738e91bb561 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0xcf, 0x6a, 0x13, 0x41,
	0x18, 0xef, 0xec, 0xa6, 0x69, 0xf2, 0x05, 0xa5, 0x0c, 0x8a, 0x6b, 0x5a, 0x31, 0x2c, 0x1e, 0x42,
	0x0f, 0x3b, 0x24, 0x0a, 0x82, 0xc7, 0xd0, 0x4a, 0x11, 0xf4, 0xb0, 0xbd, 0xe5, 0x12, 0x26, 0x3b,
	0x1f, 0xdb, 0xc5, 0xdd, 0x99, 0x65, 0x67, 0x52, 0x1b, 0x96, 0x5c, 0xf4, 0x11, 0x3c, 0x78, 0xd3,
	0x37, 0x10, 0x2f, 0x3e, 0x89, 0xaf, 0xe0, 0x83, 0xc8, 0x4c, 0xb2, 0x01, 0x6d, 0x91, 0xd2, 0x4b,
	0x6f, 0xfb, 0xfd, 0xfb, 0xfd, 0x7e, 0xfb, 0xfb, 0xbe, 0x81, 0x17, 0x78, 0xc9, 0x8b, 0x32, 0x47,
	0xcd, 0x32, 0x69, 0xb0, 0x92, 0x3c, 0x67, 0x65, 0xa5, 0x8c, 0x62, 0x9b, 0x7c, 0x39, 0x67, 0x98,
	0x9c, 0xab, 0x99, 0xc6, 0xea, 0x22, 0x4b, 0x30, 0x72, 0x45, 0x3a, 0x4c, 0xab, 0x32, 0x89, 0x52,
	0x6e, 0xf0, 0x03, 0x5f, 0x46, 0x0d, 0x44, 0xd4, 0x40, 0x44, 0xdb, 0xe1, 0xfe, 0x61, 0xaa, 0x54,
	0x9a, 0x23, 0xe3, 0x65, 0xc6, 0xb8, 0x94, 0xca, 0x70, 0x93, 0x29, 0xa9, 0xd7, 0x38, 0xe1, 0x6b,
	0xe8, 0x9c, 0x14, 0x73, 0x14, 0x02, 0x05, 0x3d, 0x84, 0x4e, 0x59, 0xa9, 0xb4, 0x42, 0xad, 0x03,
	0x32, 0x20, 0x43, 0xff, 0x74, 0x27, 0xde, 0x66, 0xe8, 0x03, 0x68, 0x49, 0x65, 0x30, 0xf0, 0x06,
	0x64, 0xd8, 0x3d, 0xdd, 0x89, 0x5d, 0x34, 0x69, 0x43, 0xab, 0xe0, 0xd5, 0xfb, 0xf0, 0x8b, 0x07,
	0xf7, 0xce, 0x32, 0x4b, 0xf9, 0x16, 0xb5, 0xe6, 0x29, 0xd2, 0xfb, 0xe0, 0x65, 0xc2, 0xe1, 0x74,
	0x63, 0x2f, 0x13, 0x74, 0x1f, 0x7c, 0xb9, 0x28, 0xdc, 0xb8, 0x1f, 0xdb, 0x4f, 0x7a, 0x00, 0x9d,
	0x3c, 0x93, 0x38, 0xb3, 0x69, 0x7f, 0xc3, 0xb7, 0x67, 0x33, 0xef, 0x16, 0x85, 0xa5, 0xcb, 0xb9,
	0x4c, 0x83, 0x56, 0x43, 0x67, 0x23, 0xfa, 0x06, 0xda, 0xda, 0x70, 0xb3, 0xd0, 0xc1, 0xee, 0x80,
	0x0c, 0x7b, 0xe3, 0x71, 0x74, 0x53, 0x1f, 0xa2, 0xe6, 0x37, 0xe3, 0x0d, 0x02, 0xdd, 0x07, 0x0f,
	0x65, 0xd0, 0x76, 0xc4, 0x24, 0xf6, 0x50, 0xd2, 0x63, 0xf0, 0xa4, 0x0a, 0xf6, 0x6e, 0x8b, 0x6c,
	0x51, 0xa4, 0xb2, 0x96, 0x24, 0x4a, 0xe0, 0x64, 0x17, 0x7c, 0xbc, 0x34, 0xe3, 0x4f, 0xbb, 0xd0,
	0x3b, 0x49, 0xce, 0xd5, 0xd9, 0x7a, 0x7f, 0xf4, 0x87, 0x07, 0x2d, 0x1b, 0xd3, 0x97, 0x37, 0x67,
	0xf8, 0xcb, 0xd9, 0xfe, 0x6d, 0x07, 0xc3, 0x9f, 0xe4, 0xe3, 0xaf, 0xdf, 0x9f, 0xbd, 0xef, 0x24,
	0x7c, 0xc8, 0x2e, 0x46, 0xcd, 0x81, 0xb9, 0xf3, 0x62, 0x75, 0x26, 0x56, 0xd3, 0x27, 0xf4, 0xe0,
	0xda, 0x02, 0xab, 0xe5, 0xa2, 0x58, 0x4d, 0x9f, 0xd1, 0xf0, 0x3f, 0x65, 0x56, 0xdb, 0x15, 0xad,
	0xa6, 0x23, 0xca, 0xfe, 0xed, 0x1a, 0x6d, 0xda, 0x9a, 0x75, 0xaf, 0x58, 0xbd, 0x5e, 0x42, 0x64,
	0x8f, 0xe8, 0x5a, 0xde, 0x31, 0xab, 0xa5, 0x5a, 0x97, 0xe9, 0x57, 0x02, 0x1d, 0x6b, 0xd9, 0x44,
	0x89, 0xe5, 0x1d, 0xd8, 0x36, 0x70, 0xae, 0xf5, 0xaf, 0x9a, 0x36, 0x9b, 0x2b, 0xb1, 0x7c, 0x45,
	0x8e, 0xe8, 0x37, 0x02, 0x60, 0x05, 0x1e, 0x63, 0x8e, 0x06, 0xef, 0x40, 0xe2, 0x53, 0x27, 0xf1,
	0xf1, 0xd1, 0xa3, 0x2b, 0x12, 0x85, 0x93, 0x34, 0xe9, 0x4d, 0xbb, 0xdb, 0xd9, 0x79, 0xdb, 0xbd,
	0xfd, 0xe7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x28, 0x11, 0xad, 0x90, 0x7b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoServiceClient interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
	// EchoDelete method receives a simple message and returns it.
	EchoDelete(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.internal.examplepb.EchoService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoBody(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.internal.examplepb.EchoService/EchoBody", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoDelete(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.internal.examplepb.EchoService/EchoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServiceServer is the server API for EchoService service.
type EchoServiceServer interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(context.Context, *SimpleMessage) (*SimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(context.Context, *SimpleMessage) (*SimpleMessage, error)
	// EchoDelete method receives a simple message and returns it.
	EchoDelete(context.Context, *SimpleMessage) (*SimpleMessage, error)
}

// UnimplementedEchoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServiceServer struct {
}

func (*UnimplementedEchoServiceServer) Echo(ctx context.Context, req *SimpleMessage) (*SimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (*UnimplementedEchoServiceServer) EchoBody(ctx context.Context, req *SimpleMessage) (*SimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoBody not implemented")
}
func (*UnimplementedEchoServiceServer) EchoDelete(ctx context.Context, req *SimpleMessage) (*SimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoDelete not implemented")
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.internal.examplepb.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.internal.examplepb.EchoService/EchoBody",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoBody(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.internal.examplepb.EchoService/EchoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoDelete(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.internal.examplepb.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
		{
			MethodName: "EchoBody",
			Handler:    _EchoService_EchoBody_Handler,
		},
		{
			MethodName: "EchoDelete",
			Handler:    _EchoService_EchoDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/internal/proto/examplepb/echo_service.proto",
}
