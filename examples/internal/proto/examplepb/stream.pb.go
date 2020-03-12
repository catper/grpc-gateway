// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/internal/proto/examplepb/stream.proto

package examplepb

import (
	context "context"
	fmt "fmt"
	proto "github.com/catper/protobuf/proto"
	empty "github.com/catper/protobuf/ptypes/empty"
	sub "github.com/catper/grpc-gateway/examples/internal/proto/sub"
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

func init() {
	proto.RegisterFile("examples/internal/proto/examplepb/stream.proto", fileDescriptor_cc5dba844cf4f624)
}

var fileDescriptor_cc5dba844cf4f624 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0xe5, 0x7b, 0x11, 0x82, 0x20, 0x16, 0x0f, 0x0c, 0x01, 0x09, 0x51, 0x21, 0xd1, 0x32,
	0xd8, 0x2d, 0x6c, 0x65, 0xa2, 0xa8, 0x1b, 0x88, 0xa1, 0x1b, 0x4b, 0x65, 0x47, 0xa7, 0xae, 0xd5,
	0xc4, 0x8e, 0xec, 0x93, 0x42, 0x57, 0x24, 0x9e, 0xa0, 0xcf, 0xc0, 0x13, 0x21, 0xde, 0x80, 0x07,
	0x41, 0x75, 0xd2, 0x4c, 0x54, 0x29, 0x62, 0xcc, 0x39, 0xf9, 0x8e, 0xff, 0xef, 0x8f, 0x18, 0xbc,
	0x88, 0x2c, 0x4f, 0xc1, 0x73, 0x6d, 0x10, 0x9c, 0x11, 0x29, 0xcf, 0x9d, 0x45, 0xcb, 0xab, 0x79,
	0x2e, 0xb9, 0x47, 0x07, 0x22, 0x63, 0x61, 0x4c, 0xdb, 0xca, 0xe5, 0x09, 0x53, 0x02, 0xe1, 0x59,
	0x2c, 0x6a, 0x98, 0xad, 0x61, 0x56, 0x63, 0xf1, 0x89, 0xb2, 0x56, 0xa5, 0xc0, 0x45, 0xae, 0xb9,
	0x30, 0xc6, 0xa2, 0x40, 0x6d, 0x8d, 0x2f, 0xef, 0xc4, 0xc7, 0xd5, 0x36, 0x7c, 0xc9, 0x62, 0xc2,
	0x21, 0xcb, 0x71, 0x51, 0x2d, 0x6f, 0x9a, 0x43, 0x89, 0xb1, 0xd4, 0x38, 0xb6, 0x93, 0x31, 0xcc,
	0xc1, 0x2d, 0x70, 0xaa, 0x8d, 0xaa, 0xe0, 0xce, 0x26, 0xd8, 0x17, 0x92, 0x67, 0xe0, 0xbd, 0x50,
	0x50, 0xfe, 0x7a, 0xf5, 0xf9, 0x3f, 0x3a, 0x1c, 0x05, 0xbb, 0x11, 0xb8, 0xb9, 0x4e, 0x80, 0x2e,
	0x49, 0x14, 0x0d, 0x8a, 0x74, 0x76, 0xe7, 0x40, 0x20, 0xd0, 0x3e, 0xdb, 0x56, 0x97, 0xdd, 0x0e,
	0x34, 0x3e, 0x4e, 0x86, 0x75, 0x9a, 0xf8, 0x88, 0x95, 0x8a, 0x6c, 0xad, 0xc8, 0x86, 0x2b, 0xc5,
	0x16, 0x7f, 0xfd, 0xf8, 0x5a, 0xfe, 0xeb, 0xb4, 0xce, 0xf9, 0xbc, 0xb7, 0x16, 0xfa, 0x49, 0x87,
	0xcb, 0x22, 0x9d, 0xf5, 0xc9, 0x65, 0x9b, 0xd0, 0x37, 0x12, 0xed, 0xdc, 0x6b, 0x8f, 0x74, 0xc3,
	0xcd, 0xf8, 0x0f, 0x39, 0x5b, 0x17, 0x21, 0xcf, 0x19, 0x3d, 0x6d, 0xc8, 0xd3, 0x25, 0xf4, 0x9d,
	0x44, 0x7b, 0xab, 0x76, 0x86, 0xc9, 0xd4, 0xd2, 0x5e, 0xd3, 0x9b, 0xbe, 0x90, 0x6c, 0x84, 0x4e,
	0x1b, 0xf5, 0x50, 0xb6, 0x1e, 0xff, 0x1e, 0xd9, 0xbe, 0x2d, 0x48, 0xa6, 0x36, 0xb4, 0xd5, 0x25,
	0x83, 0x83, 0xa7, 0xfd, 0x5a, 0x58, 0xee, 0x86, 0xae, 0xae, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xb8, 0x9c, 0x62, 0x91, 0xea, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamServiceClient interface {
	BulkCreate(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkCreateClient, error)
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (StreamService_ListClient, error)
	BulkEcho(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkEchoClient, error)
}

type streamServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreamServiceClient(cc *grpc.ClientConn) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) BulkCreate(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkCreateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[0], "/grpc.gateway.examples.internal.examplepb.StreamService/BulkCreate", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceBulkCreateClient{stream}
	return x, nil
}

type StreamService_BulkCreateClient interface {
	Send(*ABitOfEverything) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type streamServiceBulkCreateClient struct {
	grpc.ClientStream
}

func (x *streamServiceBulkCreateClient) Send(m *ABitOfEverything) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceBulkCreateClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (StreamService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[1], "/grpc.gateway.examples.internal.examplepb.StreamService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_ListClient interface {
	Recv() (*ABitOfEverything, error)
	grpc.ClientStream
}

type streamServiceListClient struct {
	grpc.ClientStream
}

func (x *streamServiceListClient) Recv() (*ABitOfEverything, error) {
	m := new(ABitOfEverything)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) BulkEcho(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkEchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[2], "/grpc.gateway.examples.internal.examplepb.StreamService/BulkEcho", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceBulkEchoClient{stream}
	return x, nil
}

type StreamService_BulkEchoClient interface {
	Send(*sub.StringMessage) error
	Recv() (*sub.StringMessage, error)
	grpc.ClientStream
}

type streamServiceBulkEchoClient struct {
	grpc.ClientStream
}

func (x *streamServiceBulkEchoClient) Send(m *sub.StringMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceBulkEchoClient) Recv() (*sub.StringMessage, error) {
	m := new(sub.StringMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
type StreamServiceServer interface {
	BulkCreate(StreamService_BulkCreateServer) error
	List(*empty.Empty, StreamService_ListServer) error
	BulkEcho(StreamService_BulkEchoServer) error
}

// UnimplementedStreamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (*UnimplementedStreamServiceServer) BulkCreate(srv StreamService_BulkCreateServer) error {
	return status.Errorf(codes.Unimplemented, "method BulkCreate not implemented")
}
func (*UnimplementedStreamServiceServer) List(req *empty.Empty, srv StreamService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedStreamServiceServer) BulkEcho(srv StreamService_BulkEchoServer) error {
	return status.Errorf(codes.Unimplemented, "method BulkEcho not implemented")
}

func RegisterStreamServiceServer(s *grpc.Server, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

func _StreamService_BulkCreate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).BulkCreate(&streamServiceBulkCreateServer{stream})
}

type StreamService_BulkCreateServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*ABitOfEverything, error)
	grpc.ServerStream
}

type streamServiceBulkCreateServer struct {
	grpc.ServerStream
}

func (x *streamServiceBulkCreateServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceBulkCreateServer) Recv() (*ABitOfEverything, error) {
	m := new(ABitOfEverything)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).List(m, &streamServiceListServer{stream})
}

type StreamService_ListServer interface {
	Send(*ABitOfEverything) error
	grpc.ServerStream
}

type streamServiceListServer struct {
	grpc.ServerStream
}

func (x *streamServiceListServer) Send(m *ABitOfEverything) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamService_BulkEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).BulkEcho(&streamServiceBulkEchoServer{stream})
}

type StreamService_BulkEchoServer interface {
	Send(*sub.StringMessage) error
	Recv() (*sub.StringMessage, error)
	grpc.ServerStream
}

type streamServiceBulkEchoServer struct {
	grpc.ServerStream
}

func (x *streamServiceBulkEchoServer) Send(m *sub.StringMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceBulkEchoServer) Recv() (*sub.StringMessage, error) {
	m := new(sub.StringMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.internal.examplepb.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BulkCreate",
			Handler:       _StreamService_BulkCreate_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "List",
			Handler:       _StreamService_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BulkEcho",
			Handler:       _StreamService_BulkEcho_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "examples/internal/proto/examplepb/stream.proto",
}
