// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/task.proto

package task // import "github.com/becosuke/tasks-api/protogen/service/task"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/becosuke/tasks-api/protogen/message/common"
import task "github.com/becosuke/tasks-api/protogen/message/task"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaskClient is the client API for Task service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskClient interface {
	GetDocument(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*task.Document, error)
	GetDocuments(ctx context.Context, in *common.Ids, opts ...grpc.CallOption) (Task_GetDocumentsClient, error)
	GetDocumentsAll(ctx context.Context, in *common.Pagination, opts ...grpc.CallOption) (Task_GetDocumentsAllClient, error)
	GetCountAll(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*common.Count, error)
}

type taskClient struct {
	cc *grpc.ClientConn
}

func NewTaskClient(cc *grpc.ClientConn) TaskClient {
	return &taskClient{cc}
}

func (c *taskClient) GetDocument(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*task.Document, error) {
	out := new(task.Document)
	err := c.cc.Invoke(ctx, "/service.task.Task/GetDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) GetDocuments(ctx context.Context, in *common.Ids, opts ...grpc.CallOption) (Task_GetDocumentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Task_serviceDesc.Streams[0], "/service.task.Task/GetDocuments", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskGetDocumentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Task_GetDocumentsClient interface {
	Recv() (*task.Document, error)
	grpc.ClientStream
}

type taskGetDocumentsClient struct {
	grpc.ClientStream
}

func (x *taskGetDocumentsClient) Recv() (*task.Document, error) {
	m := new(task.Document)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskClient) GetDocumentsAll(ctx context.Context, in *common.Pagination, opts ...grpc.CallOption) (Task_GetDocumentsAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Task_serviceDesc.Streams[1], "/service.task.Task/GetDocumentsAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskGetDocumentsAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Task_GetDocumentsAllClient interface {
	Recv() (*task.Document, error)
	grpc.ClientStream
}

type taskGetDocumentsAllClient struct {
	grpc.ClientStream
}

func (x *taskGetDocumentsAllClient) Recv() (*task.Document, error) {
	m := new(task.Document)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskClient) GetCountAll(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*common.Count, error) {
	out := new(common.Count)
	err := c.cc.Invoke(ctx, "/service.task.Task/GetCountAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServer is the server API for Task service.
type TaskServer interface {
	GetDocument(context.Context, *common.Id) (*task.Document, error)
	GetDocuments(*common.Ids, Task_GetDocumentsServer) error
	GetDocumentsAll(*common.Pagination, Task_GetDocumentsAllServer) error
	GetCountAll(context.Context, *common.Empty) (*common.Count, error)
}

func RegisterTaskServer(s *grpc.Server, srv TaskServer) {
	s.RegisterService(&_Task_serviceDesc, srv)
}

func _Task_GetDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).GetDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.task.Task/GetDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).GetDocument(ctx, req.(*common.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_GetDocuments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(common.Ids)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskServer).GetDocuments(m, &taskGetDocumentsServer{stream})
}

type Task_GetDocumentsServer interface {
	Send(*task.Document) error
	grpc.ServerStream
}

type taskGetDocumentsServer struct {
	grpc.ServerStream
}

func (x *taskGetDocumentsServer) Send(m *task.Document) error {
	return x.ServerStream.SendMsg(m)
}

func _Task_GetDocumentsAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(common.Pagination)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskServer).GetDocumentsAll(m, &taskGetDocumentsAllServer{stream})
}

type Task_GetDocumentsAllServer interface {
	Send(*task.Document) error
	grpc.ServerStream
}

type taskGetDocumentsAllServer struct {
	grpc.ServerStream
}

func (x *taskGetDocumentsAllServer) Send(m *task.Document) error {
	return x.ServerStream.SendMsg(m)
}

func _Task_GetCountAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).GetCountAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.task.Task/GetCountAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).GetCountAll(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Task_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.task.Task",
	HandlerType: (*TaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDocument",
			Handler:    _Task_GetDocument_Handler,
		},
		{
			MethodName: "GetCountAll",
			Handler:    _Task_GetCountAll_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDocuments",
			Handler:       _Task_GetDocuments_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetDocumentsAll",
			Handler:       _Task_GetDocumentsAll_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service/task.proto",
}

func init() { proto.RegisterFile("service/task.proto", fileDescriptor_task_ac72541556ba9c0c) }

var fileDescriptor_task_ac72541556ba9c0c = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4a, 0xc3, 0x40,
	0x14, 0xc6, 0x69, 0x11, 0x17, 0xb1, 0x52, 0x78, 0x35, 0x22, 0xa3, 0x0b, 0xe9, 0xde, 0x4c, 0x6d,
	0xf1, 0x00, 0xfe, 0x43, 0xdc, 0x75, 0xe1, 0x42, 0x5c, 0x28, 0x93, 0x64, 0x18, 0x87, 0x24, 0xf3,
	0x42, 0xdf, 0xa4, 0x20, 0xd2, 0x8d, 0x57, 0xf0, 0x68, 0x5e, 0xc1, 0x1b, 0x78, 0x01, 0x99, 0x49,
	0x22, 0xb1, 0x45, 0x70, 0xfb, 0x9b, 0xef, 0xfd, 0xbe, 0xf0, 0x5e, 0x02, 0x20, 0xb9, 0x58, 0xea,
	0x44, 0x72, 0x2b, 0x28, 0x8b, 0xca, 0x05, 0x5a, 0x84, 0x41, 0xc3, 0x22, 0xc7, 0xd8, 0x91, 0x42,
	0x54, 0xb9, 0xe4, 0xa2, 0xd4, 0x5c, 0x18, 0x83, 0x56, 0x58, 0x8d, 0x86, 0xea, 0x2c, 0x83, 0x42,
	0x12, 0x09, 0xd5, 0x9d, 0x67, 0x7b, 0x2d, 0x4b, 0xb0, 0x28, 0xd0, 0xd4, 0x74, 0xfa, 0xd5, 0x0f,
	0xb6, 0xee, 0x04, 0x65, 0x30, 0x0f, 0x76, 0x6e, 0xa4, 0xbd, 0xc2, 0xa4, 0x2a, 0xa4, 0xb1, 0x00,
	0x51, 0x13, 0x8f, 0x9a, 0xf8, 0x6d, 0xca, 0xf6, 0x7f, 0x98, 0xd7, 0xb6, 0xd9, 0x71, 0xf8, 0xf6,
	0xf1, 0xf9, 0xde, 0x1f, 0xc2, 0x2e, 0x5f, 0x9e, 0xfa, 0x46, 0xfe, 0xaa, 0xd3, 0x15, 0x3c, 0x05,
	0x83, 0x8e, 0x91, 0x60, 0xb4, 0xa9, 0xa4, 0x3f, 0x9d, 0xc7, 0xde, 0xc9, 0xe0, 0xa0, 0x75, 0x12,
	0x4f, 0x9b, 0x37, 0x67, 0xa7, 0xd5, 0xa4, 0x07, 0x8f, 0xc1, 0xb0, 0x5b, 0x70, 0x9e, 0xe7, 0xc0,
	0xd6, 0x3b, 0xe6, 0x42, 0x69, 0xe3, 0x77, 0xf3, 0x9f, 0xcf, 0x9f, 0x36, 0x55, 0x22, 0xcf, 0x27,
	0x3d, 0xb8, 0xf7, 0x2b, 0xb9, 0xc4, 0xca, 0x58, 0xe7, 0x0e, 0xd7, 0xdd, 0xd7, 0x45, 0x69, 0x5f,
	0xd8, 0x06, 0xf6, 0x03, 0xe3, 0x43, 0x6f, 0x0d, 0x61, 0xf4, 0xcb, 0xca, 0x13, 0xf7, 0x78, 0x71,
	0xf6, 0x30, 0x53, 0xda, 0x3e, 0x57, 0xb1, 0x9b, 0xe1, 0xb1, 0x4c, 0x90, 0xaa, 0xac, 0xbe, 0x16,
	0x9d, 0xb8, 0x93, 0xfa, 0xe3, 0x28, 0x69, 0x78, 0xf7, 0x47, 0x88, 0xb7, 0x3d, 0x9e, 0x7d, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x8d, 0x57, 0x23, 0x44, 0x1f, 0x02, 0x00, 0x00,
}
