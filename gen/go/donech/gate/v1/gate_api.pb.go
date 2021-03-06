// Code generated by protoc-gen-go. DO NOT EDIT.
// source: donech/gate/v1/gate_api.proto

package gatev1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

//LoginRequest.
type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	AutoLogin            bool     `protobuf:"varint,4,opt,name=auto_login,json=autoLogin,proto3" json:"auto_login,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a44fc2c2f01e580d, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *LoginRequest) GetAutoLogin() bool {
	if m != nil {
		return m.AutoLogin
	}
	return false
}

//LoginResponse.
type LoginResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a44fc2c2f01e580d, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LoginResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

//LogoutRequest.
type LogoutRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a44fc2c2f01e580d, []int{2}
}

func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (m *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(m, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

func (m *LogoutRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

//LogoutResponse.
type LogoutResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a44fc2c2f01e580d, []int{3}
}

func (m *LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResponse.Unmarshal(m, b)
}
func (m *LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResponse.Marshal(b, m, deterministic)
}
func (m *LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResponse.Merge(m, src)
}
func (m *LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_LogoutResponse.Size(m)
}
func (m *LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResponse proto.InternalMessageInfo

func (m *LogoutResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LogoutResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

//ReadinessRequest.
type ReadinessRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadinessRequest) Reset()         { *m = ReadinessRequest{} }
func (m *ReadinessRequest) String() string { return proto.CompactTextString(m) }
func (*ReadinessRequest) ProtoMessage()    {}
func (*ReadinessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a44fc2c2f01e580d, []int{4}
}

func (m *ReadinessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadinessRequest.Unmarshal(m, b)
}
func (m *ReadinessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadinessRequest.Marshal(b, m, deterministic)
}
func (m *ReadinessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadinessRequest.Merge(m, src)
}
func (m *ReadinessRequest) XXX_Size() int {
	return xxx_messageInfo_ReadinessRequest.Size(m)
}
func (m *ReadinessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadinessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadinessRequest proto.InternalMessageInfo

//ReadinessResponse.
type ReadinessResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadinessResponse) Reset()         { *m = ReadinessResponse{} }
func (m *ReadinessResponse) String() string { return proto.CompactTextString(m) }
func (*ReadinessResponse) ProtoMessage()    {}
func (*ReadinessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a44fc2c2f01e580d, []int{5}
}

func (m *ReadinessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadinessResponse.Unmarshal(m, b)
}
func (m *ReadinessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadinessResponse.Marshal(b, m, deterministic)
}
func (m *ReadinessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadinessResponse.Merge(m, src)
}
func (m *ReadinessResponse) XXX_Size() int {
	return xxx_messageInfo_ReadinessResponse.Size(m)
}
func (m *ReadinessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadinessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadinessResponse proto.InternalMessageInfo

func (m *ReadinessResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ReadinessResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

//LivenessRequest.
type LivenessRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LivenessRequest) Reset()         { *m = LivenessRequest{} }
func (m *LivenessRequest) String() string { return proto.CompactTextString(m) }
func (*LivenessRequest) ProtoMessage()    {}
func (*LivenessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a44fc2c2f01e580d, []int{6}
}

func (m *LivenessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LivenessRequest.Unmarshal(m, b)
}
func (m *LivenessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LivenessRequest.Marshal(b, m, deterministic)
}
func (m *LivenessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LivenessRequest.Merge(m, src)
}
func (m *LivenessRequest) XXX_Size() int {
	return xxx_messageInfo_LivenessRequest.Size(m)
}
func (m *LivenessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LivenessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LivenessRequest proto.InternalMessageInfo

//LivenessResponse.
type LivenessResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LivenessResponse) Reset()         { *m = LivenessResponse{} }
func (m *LivenessResponse) String() string { return proto.CompactTextString(m) }
func (*LivenessResponse) ProtoMessage()    {}
func (*LivenessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a44fc2c2f01e580d, []int{7}
}

func (m *LivenessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LivenessResponse.Unmarshal(m, b)
}
func (m *LivenessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LivenessResponse.Marshal(b, m, deterministic)
}
func (m *LivenessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LivenessResponse.Merge(m, src)
}
func (m *LivenessResponse) XXX_Size() int {
	return xxx_messageInfo_LivenessResponse.Size(m)
}
func (m *LivenessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LivenessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LivenessResponse proto.InternalMessageInfo

func (m *LivenessResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LivenessResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "donech.gate.v1.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "donech.gate.v1.LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "donech.gate.v1.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "donech.gate.v1.LogoutResponse")
	proto.RegisterType((*ReadinessRequest)(nil), "donech.gate.v1.ReadinessRequest")
	proto.RegisterType((*ReadinessResponse)(nil), "donech.gate.v1.ReadinessResponse")
	proto.RegisterType((*LivenessRequest)(nil), "donech.gate.v1.LivenessRequest")
	proto.RegisterType((*LivenessResponse)(nil), "donech.gate.v1.LivenessResponse")
}

func init() { proto.RegisterFile("donech/gate/v1/gate_api.proto", fileDescriptor_a44fc2c2f01e580d) }

var fileDescriptor_a44fc2c2f01e580d = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x49, 0x7b, 0x5b, 0xdb, 0xc3, 0xbd, 0x35, 0x3d, 0xba, 0x88, 0xe1, 0x56, 0x63, 0x40,
	0xb8, 0xb8, 0x48, 0x88, 0x82, 0xa8, 0xe0, 0xc2, 0x8b, 0x20, 0x62, 0x17, 0x25, 0x0b, 0x17, 0x52,
	0x2c, 0x63, 0x33, 0xc4, 0xd0, 0x76, 0x26, 0x66, 0x26, 0x91, 0x6e, 0x7d, 0x05, 0xdf, 0xc0, 0xc7,
	0x71, 0xeb, 0x2b, 0xf8, 0x20, 0x32, 0x93, 0x3f, 0x4d, 0xab, 0x55, 0xba, 0xca, 0xcc, 0xf9, 0x3e,
	0xbe, 0x5f, 0xe6, 0x9c, 0x19, 0x98, 0x44, 0x9c, 0xd1, 0xe5, 0x27, 0x3f, 0x26, 0x92, 0xfa, 0x45,
	0xa0, 0xbf, 0x0b, 0x92, 0x26, 0x5e, 0x9a, 0x71, 0xc9, 0x71, 0x54, 0xca, 0x9e, 0x2a, 0x7b, 0x45,
	0x60, 0x5f, 0xc6, 0x9c, 0xc7, 0x6b, 0xea, 0x93, 0x34, 0xf1, 0x09, 0x63, 0x5c, 0x12, 0x99, 0x70,
	0x26, 0x4a, 0xb7, 0xbb, 0x85, 0xf3, 0x29, 0x8f, 0x13, 0x16, 0xd2, 0xcf, 0x39, 0x15, 0x12, 0x6d,
	0x18, 0xe4, 0x82, 0x66, 0x8c, 0x6c, 0xa8, 0x65, 0x38, 0xc6, 0xd5, 0x30, 0x6c, 0xf6, 0x4a, 0x4b,
	0x89, 0x10, 0x5f, 0x78, 0x16, 0x59, 0x9d, 0x52, 0xab, 0xf7, 0x88, 0x70, 0x26, 0xb7, 0x29, 0xb5,
	0xba, 0xba, 0xae, 0xd7, 0x38, 0x01, 0x20, 0xb9, 0xe4, 0x8b, 0xb5, 0x02, 0x58, 0x67, 0x8e, 0x71,
	0x35, 0x08, 0x87, 0xaa, 0xa2, 0x89, 0xee, 0x5b, 0xb8, 0xa8, 0xd0, 0x22, 0xe5, 0x4c, 0x50, 0x95,
	0xb1, 0xe4, 0x51, 0xc9, 0xed, 0x85, 0x7a, 0x8d, 0x26, 0x74, 0x37, 0x22, 0xae, 0x70, 0x6a, 0x89,
	0xb7, 0xa1, 0x27, 0xf9, 0x8a, 0xb2, 0x0a, 0x55, 0x6e, 0xdc, 0x07, 0x3a, 0x8c, 0xe7, 0xb2, 0x3e,
	0x48, 0x63, 0x33, 0xda, 0xb6, 0x27, 0x30, 0xaa, 0x6d, 0xa7, 0x40, 0x5d, 0x04, 0x33, 0xa4, 0x24,
	0x4a, 0x18, 0x15, 0xa2, 0x22, 0xb8, 0xcf, 0x60, 0xdc, 0xaa, 0x9d, 0x14, 0x37, 0x86, 0x9b, 0xd3,
	0xa4, 0xa0, 0xed, 0xb4, 0xa7, 0x60, 0xee, 0x4a, 0xa7, 0x84, 0x3d, 0xfa, 0xde, 0x85, 0x1b, 0xaf,
	0x89, 0xa4, 0x2f, 0x67, 0x6f, 0xf0, 0x03, 0xf4, 0x74, 0x4f, 0xf1, 0xd2, 0xdb, 0xbf, 0x06, 0x5e,
	0x7b, 0xca, 0xf6, 0xe4, 0x88, 0x5a, 0x72, 0x5d, 0xeb, 0xeb, 0xcf, 0x5f, 0xdf, 0x3a, 0xe8, 0x5e,
	0xe8, 0x4b, 0x53, 0x04, 0xbe, 0x9e, 0xe0, 0x73, 0xe3, 0x21, 0x12, 0xe8, 0x97, 0xfd, 0xc3, 0xbf,
	0x45, 0xec, 0xda, 0x6f, 0xdf, 0x3d, 0x26, 0xff, 0x17, 0xb1, 0x82, 0x61, 0xd3, 0x56, 0x74, 0x0e,
	0x63, 0x0e, 0xa7, 0x60, 0xdf, 0xff, 0x87, 0xa3, 0x62, 0xdd, 0xd1, 0xac, 0x5b, 0x38, 0xae, 0x59,
	0x59, 0x93, 0x1f, 0xc3, 0xa0, 0xee, 0x3a, 0xde, 0xfb, 0xe3, 0x97, 0xf7, 0x47, 0x64, 0x3b, 0xc7,
	0x0d, 0xfb, 0xa7, 0x42, 0xb3, 0x39, 0x55, 0xe5, 0xb8, 0x7e, 0x01, 0xb8, 0xe4, 0x9b, 0x83, 0x80,
	0xeb, 0x73, 0x3d, 0xb7, 0x34, 0x99, 0xa9, 0xb7, 0x38, 0x33, 0xde, 0xf7, 0x95, 0x50, 0x04, 0x3f,
	0x3a, 0xa3, 0x57, 0xda, 0x38, 0x57, 0xfa, 0xfc, 0x5d, 0xf0, 0xb1, 0xaf, 0x5f, 0xeb, 0xe3, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x8a, 0x17, 0xf7, 0xfc, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GateAPIClient is the client API for GateAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GateAPIClient interface {
	//Login ??????.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	//Logout.
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	//Readiness.
	Readiness(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*ReadinessResponse, error)
	//Liveness.
	Liveness(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error)
}

type gateAPIClient struct {
	cc *grpc.ClientConn
}

func NewGateAPIClient(cc *grpc.ClientConn) GateAPIClient {
	return &gateAPIClient{cc}
}

func (c *gateAPIClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/donech.gate.v1.GateAPI/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gateAPIClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/donech.gate.v1.GateAPI/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gateAPIClient) Readiness(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*ReadinessResponse, error) {
	out := new(ReadinessResponse)
	err := c.cc.Invoke(ctx, "/donech.gate.v1.GateAPI/Readiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gateAPIClient) Liveness(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error) {
	out := new(LivenessResponse)
	err := c.cc.Invoke(ctx, "/donech.gate.v1.GateAPI/Liveness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GateAPIServer is the server API for GateAPI service.
type GateAPIServer interface {
	//Login ??????.
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	//Logout.
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	//Readiness.
	Readiness(context.Context, *ReadinessRequest) (*ReadinessResponse, error)
	//Liveness.
	Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error)
}

func RegisterGateAPIServer(s *grpc.Server, srv GateAPIServer) {
	s.RegisterService(&_GateAPI_serviceDesc, srv)
}

func _GateAPI_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateAPIServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/donech.gate.v1.GateAPI/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateAPIServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GateAPI_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateAPIServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/donech.gate.v1.GateAPI/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateAPIServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GateAPI_Readiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateAPIServer).Readiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/donech.gate.v1.GateAPI/Readiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateAPIServer).Readiness(ctx, req.(*ReadinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GateAPI_Liveness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LivenessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateAPIServer).Liveness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/donech.gate.v1.GateAPI/Liveness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateAPIServer).Liveness(ctx, req.(*LivenessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GateAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "donech.gate.v1.GateAPI",
	HandlerType: (*GateAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _GateAPI_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _GateAPI_Logout_Handler,
		},
		{
			MethodName: "Readiness",
			Handler:    _GateAPI_Readiness_Handler,
		},
		{
			MethodName: "Liveness",
			Handler:    _GateAPI_Liveness_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "donech/gate/v1/gate_api.proto",
}
