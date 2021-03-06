// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream.proto

package friday

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// 請求使用者資訊
type UserInfoRequest struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoRequest) Reset()         { *m = UserInfoRequest{} }
func (m *UserInfoRequest) String() string { return proto.CompactTextString(m) }
func (*UserInfoRequest) ProtoMessage()    {}
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_15bd7721444db361, []int{0}
}
func (m *UserInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoRequest.Unmarshal(m, b)
}
func (m *UserInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoRequest.Marshal(b, m, deterministic)
}
func (dst *UserInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoRequest.Merge(dst, src)
}
func (m *UserInfoRequest) XXX_Size() int {
	return xxx_messageInfo_UserInfoRequest.Size(m)
}
func (m *UserInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoRequest proto.InternalMessageInfo

func (m *UserInfoRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

// 請求使用者資訊的結果
type UserInfoResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  uint32   `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Sex                  uint32   `protobuf:"varint,3,opt,name=sex,proto3" json:"sex,omitempty"`
	Count                uint32   `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoResponse) Reset()         { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()    {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_15bd7721444db361, []int{1}
}
func (m *UserInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResponse.Unmarshal(m, b)
}
func (m *UserInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResponse.Marshal(b, m, deterministic)
}
func (dst *UserInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResponse.Merge(dst, src)
}
func (m *UserInfoResponse) XXX_Size() int {
	return xxx_messageInfo_UserInfoResponse.Size(m)
}
func (m *UserInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResponse proto.InternalMessageInfo

func (m *UserInfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfoResponse) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserInfoResponse) GetSex() uint32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *UserInfoResponse) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*UserInfoRequest)(nil), "friday.UserInfoRequest")
	proto.RegisterType((*UserInfoResponse)(nil), "friday.UserInfoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataClient is the client API for Data service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataClient interface {
	// 簡單Rpc
	// 獲取使用者資料
	GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	//  修改使用者 雙向流模式
	ChangeUserInfo(ctx context.Context, opts ...grpc.CallOption) (Data_ChangeUserInfoClient, error)
}

type dataClient struct {
	cc *grpc.ClientConn
}

func NewDataClient(cc *grpc.ClientConn) DataClient {
	return &dataClient{cc}
}

func (c *dataClient) GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/friday.Data/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) ChangeUserInfo(ctx context.Context, opts ...grpc.CallOption) (Data_ChangeUserInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Data_serviceDesc.Streams[0], "/friday.Data/ChangeUserInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataChangeUserInfoClient{stream}
	return x, nil
}

type Data_ChangeUserInfoClient interface {
	Send(*UserInfoResponse) error
	Recv() (*UserInfoResponse, error)
	grpc.ClientStream
}

type dataChangeUserInfoClient struct {
	grpc.ClientStream
}

func (x *dataChangeUserInfoClient) Send(m *UserInfoResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataChangeUserInfoClient) Recv() (*UserInfoResponse, error) {
	m := new(UserInfoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataServer is the server API for Data service.
type DataServer interface {
	// 簡單Rpc
	// 獲取使用者資料
	GetUserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
	//  修改使用者 雙向流模式
	ChangeUserInfo(Data_ChangeUserInfoServer) error
}

func RegisterDataServer(s *grpc.Server, srv DataServer) {
	s.RegisterService(&_Data_serviceDesc, srv)
}

func _Data_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/friday.Data/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).GetUserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_ChangeUserInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataServer).ChangeUserInfo(&dataChangeUserInfoServer{stream})
}

type Data_ChangeUserInfoServer interface {
	Send(*UserInfoResponse) error
	Recv() (*UserInfoResponse, error)
	grpc.ServerStream
}

type dataChangeUserInfoServer struct {
	grpc.ServerStream
}

func (x *dataChangeUserInfoServer) Send(m *UserInfoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataChangeUserInfoServer) Recv() (*UserInfoResponse, error) {
	m := new(UserInfoResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Data_serviceDesc = grpc.ServiceDesc{
	ServiceName: "friday.Data",
	HandlerType: (*DataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _Data_GetUserInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ChangeUserInfo",
			Handler:       _Data_ChangeUserInfo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}

func init() { proto.RegisterFile("stream.proto", fileDescriptor_stream_15bd7721444db361) }

var fileDescriptor_stream_15bd7721444db361 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x2b, 0xca, 0x4c, 0x49, 0xac,
	0x54, 0x52, 0xe6, 0xe2, 0x0f, 0x2d, 0x4e, 0x2d, 0xf2, 0xcc, 0x4b, 0xcb, 0x0f, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe0, 0x62, 0x2e, 0xcd, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60,
	0x0e, 0x02, 0x31, 0x95, 0x12, 0xb8, 0x04, 0x10, 0x8a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85,
	0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0xc1, 0xca, 0x38, 0x83, 0xc0, 0x6c, 0x90, 0xce, 0xc4,
	0xf4, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xde, 0x20, 0x10, 0x13, 0x24, 0x52, 0x9c, 0x5a, 0x21,
	0xc1, 0x0c, 0x11, 0x29, 0x4e, 0xad, 0x10, 0x12, 0xe1, 0x62, 0x4d, 0xce, 0x2f, 0xcd, 0x2b, 0x91,
	0x60, 0x01, 0x8b, 0x41, 0x38, 0x46, 0xd3, 0x18, 0xb9, 0x58, 0x5c, 0x12, 0x4b, 0x12, 0x85, 0x9c,
	0xb8, 0xb8, 0xdd, 0x53, 0x4b, 0x60, 0xb6, 0x09, 0x89, 0xeb, 0x41, 0xdc, 0xa9, 0x87, 0xe6, 0x48,
	0x29, 0x09, 0x4c, 0x09, 0x88, 0xc3, 0x94, 0x18, 0x84, 0xbc, 0xb8, 0xf8, 0x9c, 0x33, 0x12, 0xf3,
	0xd2, 0x53, 0xe1, 0xc6, 0xe0, 0x54, 0x8d, 0xcf, 0x1c, 0x0d, 0x46, 0x03, 0xc6, 0x24, 0x36, 0x70,
	0x70, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x9a, 0x17, 0xaa, 0x3e, 0x01, 0x00, 0x00,
}
