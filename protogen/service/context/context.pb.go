// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/context.proto

package context // import "github.com/becosuke/tasks-api/protogen/service/context"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/becosuke/tasks-api/protogen/message/common"
import context1 "github.com/becosuke/tasks-api/protogen/message/context"
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

// ContextClient is the client API for Context service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContextClient interface {
	GetDocument(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*context1.Document, error)
	GetDocuments(ctx context.Context, in *common.Ids, opts ...grpc.CallOption) (Context_GetDocumentsClient, error)
	GetDocumentsAll(ctx context.Context, in *common.Pagination, opts ...grpc.CallOption) (Context_GetDocumentsAllClient, error)
	GetCountAll(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*common.Count, error)
}

type contextClient struct {
	cc *grpc.ClientConn
}

func NewContextClient(cc *grpc.ClientConn) ContextClient {
	return &contextClient{cc}
}

func (c *contextClient) GetDocument(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*context1.Document, error) {
	out := new(context1.Document)
	err := c.cc.Invoke(ctx, "/service.context.Context/GetDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextClient) GetDocuments(ctx context.Context, in *common.Ids, opts ...grpc.CallOption) (Context_GetDocumentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Context_serviceDesc.Streams[0], "/service.context.Context/GetDocuments", opts...)
	if err != nil {
		return nil, err
	}
	x := &contextGetDocumentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Context_GetDocumentsClient interface {
	Recv() (*context1.Document, error)
	grpc.ClientStream
}

type contextGetDocumentsClient struct {
	grpc.ClientStream
}

func (x *contextGetDocumentsClient) Recv() (*context1.Document, error) {
	m := new(context1.Document)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *contextClient) GetDocumentsAll(ctx context.Context, in *common.Pagination, opts ...grpc.CallOption) (Context_GetDocumentsAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Context_serviceDesc.Streams[1], "/service.context.Context/GetDocumentsAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &contextGetDocumentsAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Context_GetDocumentsAllClient interface {
	Recv() (*context1.Document, error)
	grpc.ClientStream
}

type contextGetDocumentsAllClient struct {
	grpc.ClientStream
}

func (x *contextGetDocumentsAllClient) Recv() (*context1.Document, error) {
	m := new(context1.Document)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *contextClient) GetCountAll(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*common.Count, error) {
	out := new(common.Count)
	err := c.cc.Invoke(ctx, "/service.context.Context/GetCountAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContextServer is the server API for Context service.
type ContextServer interface {
	GetDocument(context.Context, *common.Id) (*context1.Document, error)
	GetDocuments(*common.Ids, Context_GetDocumentsServer) error
	GetDocumentsAll(*common.Pagination, Context_GetDocumentsAllServer) error
	GetCountAll(context.Context, *common.Empty) (*common.Count, error)
}

func RegisterContextServer(s *grpc.Server, srv ContextServer) {
	s.RegisterService(&_Context_serviceDesc, srv)
}

func _Context_GetDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextServer).GetDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.context.Context/GetDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextServer).GetDocument(ctx, req.(*common.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Context_GetDocuments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(common.Ids)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContextServer).GetDocuments(m, &contextGetDocumentsServer{stream})
}

type Context_GetDocumentsServer interface {
	Send(*context1.Document) error
	grpc.ServerStream
}

type contextGetDocumentsServer struct {
	grpc.ServerStream
}

func (x *contextGetDocumentsServer) Send(m *context1.Document) error {
	return x.ServerStream.SendMsg(m)
}

func _Context_GetDocumentsAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(common.Pagination)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContextServer).GetDocumentsAll(m, &contextGetDocumentsAllServer{stream})
}

type Context_GetDocumentsAllServer interface {
	Send(*context1.Document) error
	grpc.ServerStream
}

type contextGetDocumentsAllServer struct {
	grpc.ServerStream
}

func (x *contextGetDocumentsAllServer) Send(m *context1.Document) error {
	return x.ServerStream.SendMsg(m)
}

func _Context_GetCountAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextServer).GetCountAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.context.Context/GetCountAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextServer).GetCountAll(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Context_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.context.Context",
	HandlerType: (*ContextServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDocument",
			Handler:    _Context_GetDocument_Handler,
		},
		{
			MethodName: "GetCountAll",
			Handler:    _Context_GetCountAll_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDocuments",
			Handler:       _Context_GetDocuments_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetDocumentsAll",
			Handler:       _Context_GetDocumentsAll_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service/context.proto",
}

func init() { proto.RegisterFile("service/context.proto", fileDescriptor_context_f82586fc028f5c29) }

var fileDescriptor_context_f82586fc028f5c29 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xe9, 0xfb, 0x82, 0x42, 0x14, 0x2a, 0xa3, 0x95, 0xba, 0xfe, 0x39, 0xd4, 0xb3, 0xd9,
	0x5a, 0x41, 0xbc, 0x6a, 0x15, 0xf1, 0xe6, 0xc9, 0x83, 0x9e, 0x36, 0xc9, 0xb2, 0x2e, 0x4d, 0x76,
	0x42, 0x67, 0x52, 0x14, 0xe9, 0xc5, 0xab, 0x47, 0x3f, 0x9a, 0x5f, 0xc1, 0x0f, 0x22, 0xd9, 0xa4,
	0x34, 0xa4, 0x20, 0x5e, 0x9f, 0x99, 0xe7, 0xf7, 0x23, 0x3b, 0x09, 0x7a, 0xa4, 0xa7, 0x33, 0x1b,
	0x6b, 0x19, 0xa3, 0x63, 0xfd, 0xc2, 0x61, 0x3e, 0x45, 0x46, 0xe8, 0xd6, 0x71, 0x58, 0xc7, 0xe2,
	0xc0, 0x20, 0x9a, 0x54, 0x4b, 0x95, 0x5b, 0xa9, 0x9c, 0x43, 0x56, 0x6c, 0xd1, 0x51, 0xb5, 0x2e,
	0x7a, 0x99, 0x26, 0x52, 0xa6, 0x45, 0x11, 0x3b, 0xcb, 0x38, 0xcb, 0xd0, 0x55, 0xe9, 0xe8, 0xe3,
	0x7f, 0xb0, 0x3e, 0xae, 0xf6, 0xe0, 0x21, 0xd8, 0xb8, 0xd5, 0x7c, 0x8d, 0x71, 0x91, 0x69, 0xc7,
	0x00, 0x61, 0xdd, 0x08, 0xeb, 0xc6, 0x5d, 0x22, 0xf6, 0x1a, 0x59, 0x05, 0x5f, 0xac, 0x0f, 0xfa,
	0xef, 0x5f, 0xdf, 0x9f, 0xff, 0x00, 0xb6, 0xe4, 0xec, 0x74, 0xa1, 0x96, 0x6f, 0x36, 0x99, 0x83,
	0x0e, 0x36, 0x1b, 0x5c, 0x82, 0xed, 0x55, 0x30, 0xfd, 0x46, 0x3e, 0xf6, 0xe4, 0x43, 0xd8, 0x6f,
	0x90, 0x49, 0x26, 0xf5, 0xb8, 0x74, 0xd0, 0x7c, 0xd8, 0x81, 0x24, 0xe8, 0x36, 0x35, 0x97, 0x69,
	0x0a, 0xa2, 0x6d, 0xba, 0x57, 0xc6, 0x3a, 0xff, 0x5a, 0x7f, 0xfc, 0x94, 0xd1, 0x52, 0xa8, 0xd2,
	0x74, 0xd8, 0x81, 0x27, 0xff, 0x48, 0x63, 0x2c, 0x1c, 0x97, 0x86, 0x5e, 0xdb, 0x70, 0x93, 0xe5,
	0xfc, 0x2a, 0x56, 0x62, 0x5f, 0x18, 0x1c, 0x79, 0x70, 0x1f, 0x76, 0xdb, 0x60, 0x19, 0x97, 0xf3,
	0xab, 0x8b, 0xc7, 0x73, 0x63, 0xf9, 0xb9, 0x88, 0xca, 0x9a, 0x8c, 0x74, 0x8c, 0x54, 0x4c, 0xb4,
	0x64, 0x45, 0x13, 0x3a, 0x29, 0xaf, 0xed, 0x8f, 0x66, 0xb4, 0x93, 0xad, 0x3f, 0x25, 0x5a, 0xf3,
	0x93, 0xb3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0x96, 0x5f, 0x4a, 0x43, 0x02, 0x00, 0x00,
}