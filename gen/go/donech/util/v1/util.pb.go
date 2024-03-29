// Code generated by protoc-gen-go. DO NOT EDIT.
// source: donech/util/v1/util.proto

package utilv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

//Pager 通用分页信息.
type Pager struct {
	PageSize             int32    `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageNum              int32    `protobuf:"varint,2,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	Cursor               uint64   `protobuf:"varint,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
	HasMore              bool     `protobuf:"varint,4,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	Total                int32    `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pager) Reset()         { *m = Pager{} }
func (m *Pager) String() string { return proto.CompactTextString(m) }
func (*Pager) ProtoMessage()    {}
func (*Pager) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d2cbc1f2d396714, []int{0}
}

func (m *Pager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pager.Unmarshal(m, b)
}
func (m *Pager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pager.Marshal(b, m, deterministic)
}
func (m *Pager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pager.Merge(m, src)
}
func (m *Pager) XXX_Size() int {
	return xxx_messageInfo_Pager.Size(m)
}
func (m *Pager) XXX_DiscardUnknown() {
	xxx_messageInfo_Pager.DiscardUnknown(m)
}

var xxx_messageInfo_Pager proto.InternalMessageInfo

func (m *Pager) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Pager) GetPageNum() int32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *Pager) GetCursor() uint64 {
	if m != nil {
		return m.Cursor
	}
	return 0
}

func (m *Pager) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

func (m *Pager) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

//CursorPager 游标分页信息.
type CursorPager struct {
	PageSize             int32    `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Cursor               uint64   `protobuf:"varint,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	HasMore              bool     `protobuf:"varint,3,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CursorPager) Reset()         { *m = CursorPager{} }
func (m *CursorPager) String() string { return proto.CompactTextString(m) }
func (*CursorPager) ProtoMessage()    {}
func (*CursorPager) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d2cbc1f2d396714, []int{1}
}

func (m *CursorPager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CursorPager.Unmarshal(m, b)
}
func (m *CursorPager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CursorPager.Marshal(b, m, deterministic)
}
func (m *CursorPager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CursorPager.Merge(m, src)
}
func (m *CursorPager) XXX_Size() int {
	return xxx_messageInfo_CursorPager.Size(m)
}
func (m *CursorPager) XXX_DiscardUnknown() {
	xxx_messageInfo_CursorPager.DiscardUnknown(m)
}

var xxx_messageInfo_CursorPager proto.InternalMessageInfo

func (m *CursorPager) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *CursorPager) GetCursor() uint64 {
	if m != nil {
		return m.Cursor
	}
	return 0
}

func (m *CursorPager) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

func init() {
	proto.RegisterType((*Pager)(nil), "donech.util.v1.Pager")
	proto.RegisterType((*CursorPager)(nil), "donech.util.v1.CursorPager")
}

func init() { proto.RegisterFile("donech/util/v1/util.proto", fileDescriptor_5d2cbc1f2d396714) }

var fileDescriptor_5d2cbc1f2d396714 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0xc9, 0xcf, 0x4b,
	0x4d, 0xce, 0xd0, 0x2f, 0x2d, 0xc9, 0xcc, 0xd1, 0x2f, 0x33, 0x04, 0xd3, 0x7a, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0x42, 0x7c, 0x10, 0x29, 0x3d, 0xb0, 0x50, 0x99, 0xa1, 0x52, 0x07, 0x23, 0x17, 0x6b,
	0x40, 0x62, 0x7a, 0x6a, 0x91, 0x90, 0x34, 0x17, 0x67, 0x41, 0x62, 0x7a, 0x6a, 0x7c, 0x71, 0x66,
	0x55, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x07, 0x48, 0x20, 0x38, 0xb3, 0x2a, 0x55,
	0x48, 0x92, 0x0b, 0xcc, 0x8e, 0xcf, 0x2b, 0xcd, 0x95, 0x60, 0x02, 0xcb, 0xb1, 0x83, 0xf8, 0x7e,
	0xa5, 0xb9, 0x42, 0x62, 0x5c, 0x6c, 0xc9, 0xa5, 0x45, 0xc5, 0xf9, 0x45, 0x12, 0xcc, 0x0a, 0x8c,
	0x1a, 0x2c, 0x41, 0x50, 0x1e, 0x48, 0x4b, 0x46, 0x62, 0x71, 0x7c, 0x6e, 0x7e, 0x51, 0xaa, 0x04,
	0x8b, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x7b, 0x46, 0x62, 0xb1, 0x6f, 0x7e, 0x51, 0xaa, 0x90, 0x08,
	0x17, 0x6b, 0x49, 0x7e, 0x49, 0x62, 0x8e, 0x04, 0x2b, 0xd8, 0x28, 0x08, 0x47, 0x29, 0x96, 0x8b,
	0xdb, 0x19, 0xac, 0x95, 0x08, 0xf7, 0x20, 0x2c, 0x65, 0xc2, 0x69, 0x29, 0x33, 0x8a, 0xa5, 0x4e,
	0x56, 0x5c, 0x42, 0xc9, 0xf9, 0xb9, 0x7a, 0xa8, 0xfe, 0x77, 0xe2, 0x0c, 0x2d, 0xc9, 0xcc, 0x09,
	0x00, 0x05, 0x4d, 0x00, 0x63, 0x14, 0x1b, 0x48, 0xb4, 0xcc, 0xf0, 0x14, 0x13, 0x9f, 0x0b, 0x58,
	0x55, 0x0c, 0x48, 0x32, 0x26, 0xcc, 0x30, 0x89, 0x0d, 0x1c, 0x78, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xfe, 0x3a, 0x41, 0xfb, 0x59, 0x01, 0x00, 0x00,
}
