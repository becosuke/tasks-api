// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/context.proto

package context // import "github.com/becosuke/tasks-api/protogen/service/context"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

type CreateRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_context_a8fb9df73aa13691, []int{0}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type CreateResponse struct {
	Document             *context1.Document `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_context_a8fb9df73aa13691, []int{1}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (dst *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(dst, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetDocument() *context1.Document {
	if m != nil {
		return m.Document
	}
	return nil
}

type UpdateRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_context_a8fb9df73aa13691, []int{2}
}
func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(dst, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type UpdateResponse struct {
	Document             *context1.Document `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_context_a8fb9df73aa13691, []int{3}
}
func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(dst, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetDocument() *context1.Document {
	if m != nil {
		return m.Document
	}
	return nil
}

type DeleteRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_context_a8fb9df73aa13691, []int{4}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(dst, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteResponse struct {
	Document             *context1.Document `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_context_a8fb9df73aa13691, []int{5}
}
func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(dst, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetDocument() *context1.Document {
	if m != nil {
		return m.Document
	}
	return nil
}

type GetDocumentRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDocumentRequest) Reset()         { *m = GetDocumentRequest{} }
func (m *GetDocumentRequest) String() string { return proto.CompactTextString(m) }
func (*GetDocumentRequest) ProtoMessage()    {}
func (*GetDocumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_context_a8fb9df73aa13691, []int{6}
}
func (m *GetDocumentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDocumentRequest.Unmarshal(m, b)
}
func (m *GetDocumentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDocumentRequest.Marshal(b, m, deterministic)
}
func (dst *GetDocumentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDocumentRequest.Merge(dst, src)
}
func (m *GetDocumentRequest) XXX_Size() int {
	return xxx_messageInfo_GetDocumentRequest.Size(m)
}
func (m *GetDocumentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDocumentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDocumentRequest proto.InternalMessageInfo

func (m *GetDocumentRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetDocumentResponse struct {
	Document             *context1.Document `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetDocumentResponse) Reset()         { *m = GetDocumentResponse{} }
func (m *GetDocumentResponse) String() string { return proto.CompactTextString(m) }
func (*GetDocumentResponse) ProtoMessage()    {}
func (*GetDocumentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_context_a8fb9df73aa13691, []int{7}
}
func (m *GetDocumentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDocumentResponse.Unmarshal(m, b)
}
func (m *GetDocumentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDocumentResponse.Marshal(b, m, deterministic)
}
func (dst *GetDocumentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDocumentResponse.Merge(dst, src)
}
func (m *GetDocumentResponse) XXX_Size() int {
	return xxx_messageInfo_GetDocumentResponse.Size(m)
}
func (m *GetDocumentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDocumentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDocumentResponse proto.InternalMessageInfo

func (m *GetDocumentResponse) GetDocument() *context1.Document {
	if m != nil {
		return m.Document
	}
	return nil
}

type GetDocumentsRequest struct {
	Ids                  []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDocumentsRequest) Reset()         { *m = GetDocumentsRequest{} }
func (m *GetDocumentsRequest) String() string { return proto.CompactTextString(m) }
func (*GetDocumentsRequest) ProtoMessage()    {}
func (*GetDocumentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_context_a8fb9df73aa13691, []int{8}
}
func (m *GetDocumentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDocumentsRequest.Unmarshal(m, b)
}
func (m *GetDocumentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDocumentsRequest.Marshal(b, m, deterministic)
}
func (dst *GetDocumentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDocumentsRequest.Merge(dst, src)
}
func (m *GetDocumentsRequest) XXX_Size() int {
	return xxx_messageInfo_GetDocumentsRequest.Size(m)
}
func (m *GetDocumentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDocumentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDocumentsRequest proto.InternalMessageInfo

func (m *GetDocumentsRequest) GetIds() []uint64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type GetDocumentsResponse struct {
	Documents            []*context1.Document `protobuf:"bytes,1,rep,name=documents,proto3" json:"documents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetDocumentsResponse) Reset()         { *m = GetDocumentsResponse{} }
func (m *GetDocumentsResponse) String() string { return proto.CompactTextString(m) }
func (*GetDocumentsResponse) ProtoMessage()    {}
func (*GetDocumentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_context_a8fb9df73aa13691, []int{9}
}
func (m *GetDocumentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDocumentsResponse.Unmarshal(m, b)
}
func (m *GetDocumentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDocumentsResponse.Marshal(b, m, deterministic)
}
func (dst *GetDocumentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDocumentsResponse.Merge(dst, src)
}
func (m *GetDocumentsResponse) XXX_Size() int {
	return xxx_messageInfo_GetDocumentsResponse.Size(m)
}
func (m *GetDocumentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDocumentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDocumentsResponse proto.InternalMessageInfo

func (m *GetDocumentsResponse) GetDocuments() []*context1.Document {
	if m != nil {
		return m.Documents
	}
	return nil
}

type GetDocumentsAllRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDocumentsAllRequest) Reset()         { *m = GetDocumentsAllRequest{} }
func (m *GetDocumentsAllRequest) String() string { return proto.CompactTextString(m) }
func (*GetDocumentsAllRequest) ProtoMessage()    {}
func (*GetDocumentsAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_context_a8fb9df73aa13691, []int{10}
}
func (m *GetDocumentsAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDocumentsAllRequest.Unmarshal(m, b)
}
func (m *GetDocumentsAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDocumentsAllRequest.Marshal(b, m, deterministic)
}
func (dst *GetDocumentsAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDocumentsAllRequest.Merge(dst, src)
}
func (m *GetDocumentsAllRequest) XXX_Size() int {
	return xxx_messageInfo_GetDocumentsAllRequest.Size(m)
}
func (m *GetDocumentsAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDocumentsAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDocumentsAllRequest proto.InternalMessageInfo

func (m *GetDocumentsAllRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetDocumentsAllRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type GetDocumentsAllResponse struct {
	Documents            []*context1.Document `protobuf:"bytes,1,rep,name=documents,proto3" json:"documents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetDocumentsAllResponse) Reset()         { *m = GetDocumentsAllResponse{} }
func (m *GetDocumentsAllResponse) String() string { return proto.CompactTextString(m) }
func (*GetDocumentsAllResponse) ProtoMessage()    {}
func (*GetDocumentsAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_context_a8fb9df73aa13691, []int{11}
}
func (m *GetDocumentsAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDocumentsAllResponse.Unmarshal(m, b)
}
func (m *GetDocumentsAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDocumentsAllResponse.Marshal(b, m, deterministic)
}
func (dst *GetDocumentsAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDocumentsAllResponse.Merge(dst, src)
}
func (m *GetDocumentsAllResponse) XXX_Size() int {
	return xxx_messageInfo_GetDocumentsAllResponse.Size(m)
}
func (m *GetDocumentsAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDocumentsAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDocumentsAllResponse proto.InternalMessageInfo

func (m *GetDocumentsAllResponse) GetDocuments() []*context1.Document {
	if m != nil {
		return m.Documents
	}
	return nil
}

type GetCountAllRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCountAllRequest) Reset()         { *m = GetCountAllRequest{} }
func (m *GetCountAllRequest) String() string { return proto.CompactTextString(m) }
func (*GetCountAllRequest) ProtoMessage()    {}
func (*GetCountAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_context_a8fb9df73aa13691, []int{12}
}
func (m *GetCountAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCountAllRequest.Unmarshal(m, b)
}
func (m *GetCountAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCountAllRequest.Marshal(b, m, deterministic)
}
func (dst *GetCountAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCountAllRequest.Merge(dst, src)
}
func (m *GetCountAllRequest) XXX_Size() int {
	return xxx_messageInfo_GetCountAllRequest.Size(m)
}
func (m *GetCountAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCountAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCountAllRequest proto.InternalMessageInfo

type GetCountAllResponse struct {
	Count                uint64   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCountAllResponse) Reset()         { *m = GetCountAllResponse{} }
func (m *GetCountAllResponse) String() string { return proto.CompactTextString(m) }
func (*GetCountAllResponse) ProtoMessage()    {}
func (*GetCountAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_context_a8fb9df73aa13691, []int{13}
}
func (m *GetCountAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCountAllResponse.Unmarshal(m, b)
}
func (m *GetCountAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCountAllResponse.Marshal(b, m, deterministic)
}
func (dst *GetCountAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCountAllResponse.Merge(dst, src)
}
func (m *GetCountAllResponse) XXX_Size() int {
	return xxx_messageInfo_GetCountAllResponse.Size(m)
}
func (m *GetCountAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCountAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCountAllResponse proto.InternalMessageInfo

func (m *GetCountAllResponse) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "service.context.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "service.context.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "service.context.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "service.context.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "service.context.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "service.context.DeleteResponse")
	proto.RegisterType((*GetDocumentRequest)(nil), "service.context.GetDocumentRequest")
	proto.RegisterType((*GetDocumentResponse)(nil), "service.context.GetDocumentResponse")
	proto.RegisterType((*GetDocumentsRequest)(nil), "service.context.GetDocumentsRequest")
	proto.RegisterType((*GetDocumentsResponse)(nil), "service.context.GetDocumentsResponse")
	proto.RegisterType((*GetDocumentsAllRequest)(nil), "service.context.GetDocumentsAllRequest")
	proto.RegisterType((*GetDocumentsAllResponse)(nil), "service.context.GetDocumentsAllResponse")
	proto.RegisterType((*GetCountAllRequest)(nil), "service.context.GetCountAllRequest")
	proto.RegisterType((*GetCountAllResponse)(nil), "service.context.GetCountAllResponse")
}

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
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*GetDocumentResponse, error)
	GetDocuments(ctx context.Context, in *GetDocumentsRequest, opts ...grpc.CallOption) (*GetDocumentsResponse, error)
	GetDocumentsAll(ctx context.Context, in *GetDocumentsAllRequest, opts ...grpc.CallOption) (*GetDocumentsAllResponse, error)
	GetCountAll(ctx context.Context, in *GetCountAllRequest, opts ...grpc.CallOption) (*GetCountAllResponse, error)
}

type contextClient struct {
	cc *grpc.ClientConn
}

func NewContextClient(cc *grpc.ClientConn) ContextClient {
	return &contextClient{cc}
}

func (c *contextClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/service.context.Context/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/service.context.Context/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/service.context.Context/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextClient) GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*GetDocumentResponse, error) {
	out := new(GetDocumentResponse)
	err := c.cc.Invoke(ctx, "/service.context.Context/GetDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextClient) GetDocuments(ctx context.Context, in *GetDocumentsRequest, opts ...grpc.CallOption) (*GetDocumentsResponse, error) {
	out := new(GetDocumentsResponse)
	err := c.cc.Invoke(ctx, "/service.context.Context/GetDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextClient) GetDocumentsAll(ctx context.Context, in *GetDocumentsAllRequest, opts ...grpc.CallOption) (*GetDocumentsAllResponse, error) {
	out := new(GetDocumentsAllResponse)
	err := c.cc.Invoke(ctx, "/service.context.Context/GetDocumentsAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextClient) GetCountAll(ctx context.Context, in *GetCountAllRequest, opts ...grpc.CallOption) (*GetCountAllResponse, error) {
	out := new(GetCountAllResponse)
	err := c.cc.Invoke(ctx, "/service.context.Context/GetCountAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContextServer is the server API for Context service.
type ContextServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	GetDocument(context.Context, *GetDocumentRequest) (*GetDocumentResponse, error)
	GetDocuments(context.Context, *GetDocumentsRequest) (*GetDocumentsResponse, error)
	GetDocumentsAll(context.Context, *GetDocumentsAllRequest) (*GetDocumentsAllResponse, error)
	GetCountAll(context.Context, *GetCountAllRequest) (*GetCountAllResponse, error)
}

func RegisterContextServer(s *grpc.Server, srv ContextServer) {
	s.RegisterService(&_Context_serviceDesc, srv)
}

func _Context_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.context.Context/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Context_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.context.Context/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Context_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.context.Context/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Context_GetDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDocumentRequest)
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
		return srv.(ContextServer).GetDocument(ctx, req.(*GetDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Context_GetDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextServer).GetDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.context.Context/GetDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextServer).GetDocuments(ctx, req.(*GetDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Context_GetDocumentsAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDocumentsAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextServer).GetDocumentsAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.context.Context/GetDocumentsAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextServer).GetDocumentsAll(ctx, req.(*GetDocumentsAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Context_GetCountAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCountAllRequest)
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
		return srv.(ContextServer).GetCountAll(ctx, req.(*GetCountAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Context_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.context.Context",
	HandlerType: (*ContextServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Context_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Context_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Context_Delete_Handler,
		},
		{
			MethodName: "GetDocument",
			Handler:    _Context_GetDocument_Handler,
		},
		{
			MethodName: "GetDocuments",
			Handler:    _Context_GetDocuments_Handler,
		},
		{
			MethodName: "GetDocumentsAll",
			Handler:    _Context_GetDocumentsAll_Handler,
		},
		{
			MethodName: "GetCountAll",
			Handler:    _Context_GetCountAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/context.proto",
}

func init() { proto.RegisterFile("service/context.proto", fileDescriptor_context_a8fb9df73aa13691) }

var fileDescriptor_context_a8fb9df73aa13691 = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xd5, 0x6e, 0x0d, 0xec, 0x8c, 0x6e, 0x95, 0x69, 0x4b, 0xc9, 0x60, 0x9b, 0xbc, 0x4d,
	0xab, 0x86, 0x88, 0xc5, 0xd0, 0x00, 0x71, 0x07, 0x9d, 0xd8, 0x0d, 0x12, 0x52, 0x24, 0x6e, 0xb8,
	0x4b, 0x13, 0xb7, 0x58, 0x4b, 0xe3, 0x52, 0xbb, 0xd3, 0x24, 0x34, 0x84, 0x78, 0x05, 0xde, 0x82,
	0xd7, 0xe1, 0x15, 0x78, 0x10, 0x14, 0x3b, 0x5e, 0xf3, 0xaf, 0xd9, 0x45, 0xef, 0x6a, 0xfb, 0xeb,
	0xf7, 0x73, 0xce, 0x39, 0x9f, 0x0c, 0x1d, 0x41, 0x67, 0x57, 0xcc, 0xa7, 0xc4, 0xe7, 0x91, 0xa4,
	0xd7, 0xd2, 0x99, 0xce, 0xb8, 0xe4, 0x68, 0x3b, 0xd9, 0x76, 0x92, 0x6d, 0xfb, 0xc9, 0x98, 0xf3,
	0x71, 0x48, 0x89, 0x37, 0x65, 0xc4, 0x8b, 0x22, 0x2e, 0x3d, 0xc9, 0x78, 0x24, 0xb4, 0xdc, 0xee,
	0x4c, 0xa8, 0x10, 0xde, 0x38, 0xe7, 0x82, 0x8f, 0xa0, 0x39, 0x98, 0x51, 0x4f, 0x52, 0x97, 0x7e,
	0x9b, 0x53, 0x21, 0x51, 0x1b, 0x1a, 0x92, 0xc9, 0x90, 0xf6, 0x6a, 0xfb, 0xb5, 0xfe, 0x86, 0xab,
	0x17, 0xf8, 0x02, 0xb6, 0x8c, 0x4c, 0x4c, 0x79, 0x24, 0x28, 0x3a, 0x83, 0xfb, 0x01, 0xf7, 0xe7,
	0x13, 0x1a, 0x49, 0x25, 0xdd, 0x3c, 0x7d, 0xec, 0x24, 0x08, 0x73, 0x23, 0xe7, 0x3c, 0x11, 0xb8,
	0xb7, 0x52, 0x7c, 0x06, 0xcd, 0xcf, 0xd3, 0x20, 0xc5, 0xdb, 0x82, 0x3a, 0x0b, 0x94, 0xc3, 0xba,
	0x5b, 0x67, 0xc1, 0x82, 0x5f, 0xcf, 0xf1, 0xcd, 0xdf, 0x56, 0xe3, 0xef, 0x41, 0xf3, 0x9c, 0x86,
	0x74, 0x29, 0x3f, 0x26, 0x19, 0xc1, 0x6a, 0xa4, 0x43, 0x40, 0x17, 0x54, 0xde, 0x1e, 0x2c, 0xc1,
	0x7d, 0x84, 0x87, 0x19, 0xd5, 0x6a, 0xcc, 0xe3, 0x8c, 0x9b, 0x30, 0xd0, 0x16, 0xac, 0xb1, 0x40,
	0xf4, 0x6a, 0xfb, 0x6b, 0xfd, 0x75, 0x37, 0xfe, 0x89, 0x3f, 0x41, 0x3b, 0x2b, 0x4c, 0xb8, 0xaf,
	0x61, 0xc3, 0x98, 0x69, 0x7d, 0x25, 0x78, 0xa1, 0xc5, 0x1f, 0xa0, 0x9b, 0x36, 0x7c, 0x17, 0x86,
	0xa9, 0x81, 0x0a, 0xd9, 0x84, 0xe9, 0xef, 0x68, 0xb8, 0x7a, 0x81, 0xba, 0x60, 0xf1, 0xd1, 0x48,
	0x50, 0xa9, 0xfa, 0xdc, 0x70, 0x93, 0x15, 0x76, 0xe1, 0x51, 0xc1, 0x67, 0xd5, 0xbb, 0xb5, 0x55,
	0x27, 0x06, 0x7c, 0x1e, 0xc9, 0xc5, 0xbd, 0xf0, 0x33, 0x55, 0xab, 0xc5, 0x6e, 0x42, 0x69, 0x43,
	0xc3, 0x8f, 0xf7, 0x92, 0x1e, 0xe9, 0xc5, 0xe9, 0x1f, 0x0b, 0xee, 0x0d, 0x34, 0x02, 0x79, 0x60,
	0xe9, 0x2c, 0xa0, 0x5d, 0x27, 0x97, 0x41, 0x27, 0x93, 0x25, 0x7b, 0x6f, 0xe9, 0xb9, 0x86, 0xe1,
	0xee, 0xaf, 0xbf, 0xff, 0x7e, 0xd7, 0x5b, 0x78, 0x93, 0x5c, 0xbd, 0x30, 0xc1, 0x7c, 0x5b, 0x3b,
	0x41, 0x23, 0xb0, 0xf4, 0xb8, 0x97, 0x20, 0x32, 0xf1, 0x29, 0x41, 0x64, 0x73, 0x82, 0x77, 0x14,
	0xa2, 0x63, 0xb7, 0x52, 0x08, 0xf2, 0x9d, 0x05, 0x37, 0x31, 0xc7, 0x07, 0x4b, 0x0f, 0x7b, 0x09,
	0x27, 0x13, 0x93, 0x12, 0x4e, 0x36, 0x25, 0xb8, 0xa7, 0x38, 0xe8, 0xa4, 0xc0, 0x41, 0x33, 0xd8,
	0x4c, 0xb5, 0x14, 0x1d, 0x14, 0x9c, 0x8a, 0x31, 0xb1, 0x0f, 0xab, 0x45, 0x59, 0x26, 0x2a, 0x32,
	0x7f, 0xd6, 0xe0, 0x41, 0x7a, 0x8e, 0x50, 0xa5, 0xa1, 0x09, 0x8a, 0x7d, 0x74, 0x87, 0x2a, 0xe1,
	0x1e, 0x28, 0xee, 0x53, 0xb4, 0x93, 0xe2, 0x0a, 0x62, 0x06, 0x2e, 0xbe, 0x81, 0xb8, 0x41, 0x3f,
	0x60, 0x3b, 0x37, 0xc9, 0xe8, 0xb8, 0xd2, 0x7e, 0x31, 0x9b, 0x76, 0xff, 0x6e, 0x61, 0x45, 0x09,
	0x04, 0xf1, 0xc2, 0x10, 0x5d, 0xab, 0xb2, 0x9b, 0xf9, 0x2e, 0x2f, 0x7b, 0x2e, 0x13, 0xe5, 0x65,
	0xcf, 0x47, 0x04, 0xef, 0x2a, 0x66, 0x0f, 0x75, 0xf3, 0x4c, 0xa2, 0xc2, 0xf2, 0xfe, 0xcd, 0x97,
	0x57, 0x63, 0x26, 0xbf, 0xce, 0x87, 0x8e, 0xcf, 0x27, 0x64, 0x48, 0x7d, 0x2e, 0xe6, 0x97, 0x94,
	0x48, 0x4f, 0x5c, 0x8a, 0xe7, 0xf1, 0xeb, 0xa4, 0x9e, 0x9e, 0x31, 0x8d, 0x48, 0xee, 0x65, 0x1b,
	0x5a, 0xea, 0xe4, 0xe5, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x1b, 0xa5, 0x06, 0xf3, 0x06,
	0x00, 0x00,
}
