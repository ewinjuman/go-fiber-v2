// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/grpcHandler/user.proto

package user

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

// The request message
type RequestSignUp struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	UserRole             string   `protobuf:"bytes,3,opt,name=userRole,proto3" json:"userRole,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestSignUp) Reset()         { *m = RequestSignUp{} }
func (m *RequestSignUp) String() string { return proto.CompactTextString(m) }
func (*RequestSignUp) ProtoMessage()    {}
func (*RequestSignUp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c459adec6645bd, []int{0}
}

func (m *RequestSignUp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestSignUp.Unmarshal(m, b)
}
func (m *RequestSignUp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestSignUp.Marshal(b, m, deterministic)
}
func (m *RequestSignUp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestSignUp.Merge(m, src)
}
func (m *RequestSignUp) XXX_Size() int {
	return xxx_messageInfo_RequestSignUp.Size(m)
}
func (m *RequestSignUp) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestSignUp.DiscardUnknown(m)
}

var xxx_messageInfo_RequestSignUp proto.InternalMessageInfo

func (m *RequestSignUp) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RequestSignUp) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RequestSignUp) GetUserRole() string {
	if m != nil {
		return m.UserRole
	}
	return ""
}

// The response message
type ResponseSignUp struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Status               int32    `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	UserRole             string   `protobuf:"bytes,4,opt,name=userRole,proto3" json:"userRole,omitempty"`
	OldId                int32    `protobuf:"varint,5,opt,name=oldId,proto3" json:"oldId,omitempty"`
	MobilePhoneNumber    string   `protobuf:"bytes,6,opt,name=mobilePhoneNumber,proto3" json:"mobilePhoneNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseSignUp) Reset()         { *m = ResponseSignUp{} }
func (m *ResponseSignUp) String() string { return proto.CompactTextString(m) }
func (*ResponseSignUp) ProtoMessage()    {}
func (*ResponseSignUp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c459adec6645bd, []int{1}
}

func (m *ResponseSignUp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseSignUp.Unmarshal(m, b)
}
func (m *ResponseSignUp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseSignUp.Marshal(b, m, deterministic)
}
func (m *ResponseSignUp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseSignUp.Merge(m, src)
}
func (m *ResponseSignUp) XXX_Size() int {
	return xxx_messageInfo_ResponseSignUp.Size(m)
}
func (m *ResponseSignUp) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseSignUp.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseSignUp proto.InternalMessageInfo

func (m *ResponseSignUp) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ResponseSignUp) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ResponseSignUp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ResponseSignUp) GetUserRole() string {
	if m != nil {
		return m.UserRole
	}
	return ""
}

func (m *ResponseSignUp) GetOldId() int32 {
	if m != nil {
		return m.OldId
	}
	return 0
}

func (m *ResponseSignUp) GetMobilePhoneNumber() string {
	if m != nil {
		return m.MobilePhoneNumber
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestSignUp)(nil), "user.RequestSignUp")
	proto.RegisterType((*ResponseSignUp)(nil), "user.ResponseSignUp")
}

func init() { proto.RegisterFile("app/grpcHandler/user.proto", fileDescriptor_b1c459adec6645bd) }

var fileDescriptor_b1c459adec6645bd = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x6d, 0x6d, 0x8b, 0x0e, 0xb8, 0x60, 0x5c, 0x24, 0xf4, 0x24, 0x3d, 0x79, 0x90, 0x5d,
	0x50, 0x3c, 0x7a, 0xd7, 0x8b, 0x48, 0x64, 0x8f, 0x1e, 0x52, 0x33, 0xac, 0x81, 0xb4, 0x13, 0x33,
	0x2d, 0xbe, 0x95, 0xcf, 0x28, 0x4d, 0xbb, 0xda, 0xe2, 0x2d, 0x5f, 0x7e, 0xe6, 0xff, 0x92, 0x81,
	0x52, 0x7b, 0xbf, 0xdd, 0x07, 0xff, 0xfe, 0xa8, 0x5b, 0xe3, 0x30, 0x6c, 0x7b, 0xc6, 0xb0, 0xf1,
	0x81, 0x3a, 0x12, 0xd9, 0x70, 0xae, 0xde, 0xe0, 0x4c, 0xe1, 0x67, 0x8f, 0xdc, 0xbd, 0xda, 0x7d,
	0xbb, 0xf3, 0x62, 0x0d, 0x39, 0x36, 0xda, 0x3a, 0x99, 0x5c, 0x25, 0xd7, 0xa7, 0x6a, 0x04, 0x51,
	0xc2, 0x89, 0xd7, 0xcc, 0x5f, 0x14, 0x8c, 0x4c, 0x63, 0xf0, 0xcb, 0x43, 0x36, 0x54, 0x29, 0x72,
	0x28, 0x8f, 0xc7, 0xec, 0xc0, 0xd5, 0x77, 0x02, 0x2b, 0x85, 0xec, 0xa9, 0x65, 0x9c, 0x04, 0x2b,
	0x48, 0xad, 0x89, 0xed, 0xb9, 0x4a, 0xad, 0xf9, 0x13, 0xa6, 0x73, 0xe1, 0x25, 0x14, 0xdc, 0xe9,
	0xae, 0xe7, 0x58, 0x99, 0xab, 0x89, 0x16, 0xb2, 0x6c, 0x29, 0x1b, 0x9a, 0xc8, 0x99, 0x27, 0x23,
	0xf3, 0x38, 0x32, 0x82, 0xb8, 0x81, 0xf3, 0x86, 0x6a, 0xeb, 0xf0, 0xe5, 0x83, 0x5a, 0x7c, 0xee,
	0x9b, 0x1a, 0x83, 0x2c, 0xe2, 0xe8, 0xff, 0xe0, 0xf6, 0x01, 0xb2, 0x1d, 0x63, 0x10, 0xf7, 0x50,
	0x4c, 0xef, 0xbd, 0xd8, 0xc4, 0xa5, 0x2d, 0xb6, 0x54, 0xae, 0x0f, 0x97, 0xf3, 0xaf, 0x55, 0x47,
	0x75, 0x11, 0x77, 0x7b, 0xf7, 0x13, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x27, 0x39, 0x79, 0x79, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	SignUp(ctx context.Context, in *RequestSignUp, opts ...grpc.CallOption) (*ResponseSignUp, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) SignUp(ctx context.Context, in *RequestSignUp, opts ...grpc.CallOption) (*ResponseSignUp, error) {
	out := new(ResponseSignUp)
	err := c.cc.Invoke(ctx, "/user.User/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	SignUp(context.Context, *RequestSignUp) (*ResponseSignUp, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) SignUp(ctx context.Context, req *RequestSignUp) (*ResponseSignUp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSignUp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SignUp(ctx, req.(*RequestSignUp))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _User_SignUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/grpcHandler/user.proto",
}
