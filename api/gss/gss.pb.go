// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gss.proto

package gss

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetQueryBySpreadSheetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetQueryBySpreadSheetRequest) Reset()         { *m = GetQueryBySpreadSheetRequest{} }
func (m *GetQueryBySpreadSheetRequest) String() string { return proto.CompactTextString(m) }
func (*GetQueryBySpreadSheetRequest) ProtoMessage()    {}
func (*GetQueryBySpreadSheetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30f83ece174e3468, []int{0}
}

func (m *GetQueryBySpreadSheetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetQueryBySpreadSheetRequest.Unmarshal(m, b)
}
func (m *GetQueryBySpreadSheetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetQueryBySpreadSheetRequest.Marshal(b, m, deterministic)
}
func (m *GetQueryBySpreadSheetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQueryBySpreadSheetRequest.Merge(m, src)
}
func (m *GetQueryBySpreadSheetRequest) XXX_Size() int {
	return xxx_messageInfo_GetQueryBySpreadSheetRequest.Size(m)
}
func (m *GetQueryBySpreadSheetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQueryBySpreadSheetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetQueryBySpreadSheetRequest proto.InternalMessageInfo

type GetQueryBySpreadSheetResponce struct {
	Res                  string   `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	Query                string   `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetQueryBySpreadSheetResponce) Reset()         { *m = GetQueryBySpreadSheetResponce{} }
func (m *GetQueryBySpreadSheetResponce) String() string { return proto.CompactTextString(m) }
func (*GetQueryBySpreadSheetResponce) ProtoMessage()    {}
func (*GetQueryBySpreadSheetResponce) Descriptor() ([]byte, []int) {
	return fileDescriptor_30f83ece174e3468, []int{1}
}

func (m *GetQueryBySpreadSheetResponce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetQueryBySpreadSheetResponce.Unmarshal(m, b)
}
func (m *GetQueryBySpreadSheetResponce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetQueryBySpreadSheetResponce.Marshal(b, m, deterministic)
}
func (m *GetQueryBySpreadSheetResponce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQueryBySpreadSheetResponce.Merge(m, src)
}
func (m *GetQueryBySpreadSheetResponce) XXX_Size() int {
	return xxx_messageInfo_GetQueryBySpreadSheetResponce.Size(m)
}
func (m *GetQueryBySpreadSheetResponce) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQueryBySpreadSheetResponce.DiscardUnknown(m)
}

var xxx_messageInfo_GetQueryBySpreadSheetResponce proto.InternalMessageInfo

func (m *GetQueryBySpreadSheetResponce) GetRes() string {
	if m != nil {
		return m.Res
	}
	return ""
}

func (m *GetQueryBySpreadSheetResponce) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func init() {
	proto.RegisterType((*GetQueryBySpreadSheetRequest)(nil), "gss.GetQueryBySpreadSheetRequest")
	proto.RegisterType((*GetQueryBySpreadSheetResponce)(nil), "gss.GetQueryBySpreadSheetResponce")
}

func init() { proto.RegisterFile("gss.proto", fileDescriptor_30f83ece174e3468) }

var fileDescriptor_30f83ece174e3468 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2f, 0x2e, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2f, 0x2e, 0x56, 0x92, 0xe3, 0x92, 0x71, 0x4f,
	0x2d, 0x09, 0x2c, 0x4d, 0x2d, 0xaa, 0x74, 0xaa, 0x0c, 0x2e, 0x28, 0x4a, 0x4d, 0x4c, 0x09, 0xce,
	0x48, 0x4d, 0x2d, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0x72, 0xe7, 0x92, 0xc5, 0x21,
	0x5f, 0x5c, 0x90, 0x9f, 0x97, 0x9c, 0x2a, 0x24, 0xc0, 0xc5, 0x5c, 0x94, 0x5a, 0x2c, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x62, 0x0a, 0x89, 0x70, 0xb1, 0x16, 0x82, 0xd4, 0x4b, 0x30, 0x81,
	0xc5, 0x20, 0x1c, 0xa3, 0x54, 0x2e, 0x66, 0xf7, 0xe2, 0x62, 0xa1, 0x38, 0x2e, 0x51, 0xac, 0xe6,
	0x09, 0x29, 0xea, 0x81, 0x5c, 0x86, 0xcf, 0x2d, 0x52, 0x4a, 0xf8, 0x94, 0x40, 0x9c, 0x93, 0xc4,
	0x06, 0xf6, 0x9b, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x69, 0xaa, 0x9e, 0x75, 0xe8, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GssClient is the client API for Gss service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GssClient interface {
	GetQueryBySpreadSheet(ctx context.Context, in *GetQueryBySpreadSheetRequest, opts ...grpc.CallOption) (*GetQueryBySpreadSheetResponce, error)
}

type gssClient struct {
	cc *grpc.ClientConn
}

func NewGssClient(cc *grpc.ClientConn) GssClient {
	return &gssClient{cc}
}

func (c *gssClient) GetQueryBySpreadSheet(ctx context.Context, in *GetQueryBySpreadSheetRequest, opts ...grpc.CallOption) (*GetQueryBySpreadSheetResponce, error) {
	out := new(GetQueryBySpreadSheetResponce)
	err := c.cc.Invoke(ctx, "/gss.Gss/GetQueryBySpreadSheet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GssServer is the server API for Gss service.
type GssServer interface {
	GetQueryBySpreadSheet(context.Context, *GetQueryBySpreadSheetRequest) (*GetQueryBySpreadSheetResponce, error)
}

// UnimplementedGssServer can be embedded to have forward compatible implementations.
type UnimplementedGssServer struct {
}

func (*UnimplementedGssServer) GetQueryBySpreadSheet(ctx context.Context, req *GetQueryBySpreadSheetRequest) (*GetQueryBySpreadSheetResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueryBySpreadSheet not implemented")
}

func RegisterGssServer(s *grpc.Server, srv GssServer) {
	s.RegisterService(&_Gss_serviceDesc, srv)
}

func _Gss_GetQueryBySpreadSheet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQueryBySpreadSheetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GssServer).GetQueryBySpreadSheet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gss.Gss/GetQueryBySpreadSheet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GssServer).GetQueryBySpreadSheet(ctx, req.(*GetQueryBySpreadSheetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gss_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gss.Gss",
	HandlerType: (*GssServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQueryBySpreadSheet",
			Handler:    _Gss_GetQueryBySpreadSheet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gss.proto",
}
