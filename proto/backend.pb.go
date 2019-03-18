// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backend.proto

package backend

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AssetValue struct {
	Key                  string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value                float32  `protobuf:"fixed32,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetValue) Reset()         { *m = AssetValue{} }
func (m *AssetValue) String() string { return proto.CompactTextString(m) }
func (*AssetValue) ProtoMessage()    {}
func (*AssetValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ab9ba5b8d8b2ba5, []int{0}
}

func (m *AssetValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetValue.Unmarshal(m, b)
}
func (m *AssetValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetValue.Marshal(b, m, deterministic)
}
func (m *AssetValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetValue.Merge(m, src)
}
func (m *AssetValue) XXX_Size() int {
	return xxx_messageInfo_AssetValue.Size(m)
}
func (m *AssetValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetValue.DiscardUnknown(m)
}

var xxx_messageInfo_AssetValue proto.InternalMessageInfo

func (m *AssetValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AssetValue) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Data struct {
	Data                 []*AssetValue `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ab9ba5b8d8b2ba5, []int{1}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetData() []*AssetValue {
	if m != nil {
		return m.Data
	}
	return nil
}

type Reply struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ab9ba5b8d8b2ba5, []int{2}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*AssetValue)(nil), "backend.AssetValue")
	proto.RegisterType((*Data)(nil), "backend.Data")
	proto.RegisterType((*Reply)(nil), "backend.Reply")
}

func init() { proto.RegisterFile("backend.proto", fileDescriptor_5ab9ba5b8d8b2ba5) }

var fileDescriptor_5ab9ba5b8d8b2ba5 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4a, 0x4c, 0xce,
	0x4e, 0xcd, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x4c, 0xb8,
	0xb8, 0x1c, 0x8b, 0x8b, 0x53, 0x4b, 0xc2, 0x12, 0x73, 0x4a, 0x53, 0x85, 0x04, 0xb8, 0x98, 0xbd,
	0x53, 0x2b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x11, 0x2e, 0x56, 0xb0,
	0x94, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x53, 0x10, 0x84, 0xa3, 0xa4, 0xcf, 0xc5, 0xe2, 0x92, 0x58,
	0x92, 0x28, 0xa4, 0x0e, 0xa1, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x84, 0xf5, 0x60, 0x96,
	0x20, 0x8c, 0x0c, 0x02, 0x2b, 0x50, 0x92, 0xe7, 0x62, 0x0d, 0x4a, 0x2d, 0xc8, 0xa9, 0x14, 0x12,
	0xe3, 0x62, 0x2b, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x86, 0x5a, 0x02, 0xe5, 0x19, 0x59, 0x72, 0xb1,
	0xbb, 0x17, 0xa5, 0xa6, 0x96, 0xa4, 0x16, 0x09, 0xe9, 0x71, 0x71, 0x87, 0x16, 0xa4, 0x24, 0x96,
	0xa4, 0x82, 0x4d, 0x11, 0xe2, 0x85, 0x9b, 0x0a, 0x32, 0x49, 0x8a, 0x0f, 0xce, 0x05, 0x1b, 0xa8,
	0xc4, 0xe0, 0xa4, 0xce, 0x25, 0x9a, 0x99, 0xaf, 0x97, 0x5e, 0x54, 0x90, 0xac, 0x57, 0x9e, 0x58,
	0x52, 0x52, 0x01, 0x53, 0xe0, 0xc4, 0xe3, 0x04, 0x61, 0x04, 0x80, 0xbc, 0x1c, 0xc0, 0x98, 0xc4,
	0x06, 0xf6, 0xbb, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x13, 0xf0, 0x7f, 0xb8, 0x0c, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

func errUnimplemented(methodName string) error {
	return status.Errorf(codes.Unimplemented, "method %s not implemented", methodName)
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	UpdateAsset(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Reply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) UpdateAsset(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/backend.Greeter/UpdateAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	UpdateAsset(context.Context, *Data) (*Reply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) UpdateAsset(ctx context.Context, req *Data) (*Reply, error) {
	return nil, errUnimplemented("UpdateAsset")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_UpdateAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).UpdateAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.Greeter/UpdateAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).UpdateAsset(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backend.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateAsset",
			Handler:    _Greeter_UpdateAsset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend.proto",
}