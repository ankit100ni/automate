// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interservice/cfgmgmt/request/inventory_nodes.proto

package request

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type InventoryNodes struct {
	Start                *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty" toml:"start,omitempty" mapstructure:"start,omitempty"`
	End                  *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty" toml:"end,omitempty" mapstructure:"end,omitempty"`
	Filter               []string             `protobuf:"bytes,3,rep,name=filter,proto3" json:"filter,omitempty" toml:"filter,omitempty" mapstructure:"filter,omitempty"`
	CursorDate           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=cursor_date,json=cursorDate,proto3" json:"cursor_date,omitempty" toml:"cursor_date,omitempty" mapstructure:"cursor_date,omitempty"`
	CursorId             string               `protobuf:"bytes,5,opt,name=cursor_id,json=cursorId,proto3" json:"cursor_id,omitempty" toml:"cursor_id,omitempty" mapstructure:"cursor_id,omitempty"`
	PageSize             int32                `protobuf:"varint,6,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty" toml:"page_size,omitempty" mapstructure:"page_size,omitempty"`
	Sorting              *Sorting             `protobuf:"bytes,7,opt,name=sorting,proto3" json:"sorting,omitempty" toml:"sorting,omitempty" mapstructure:"sorting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *InventoryNodes) Reset()         { *m = InventoryNodes{} }
func (m *InventoryNodes) String() string { return proto.CompactTextString(m) }
func (*InventoryNodes) ProtoMessage()    {}
func (*InventoryNodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_599302790193da7f, []int{0}
}

func (m *InventoryNodes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InventoryNodes.Unmarshal(m, b)
}
func (m *InventoryNodes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InventoryNodes.Marshal(b, m, deterministic)
}
func (m *InventoryNodes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InventoryNodes.Merge(m, src)
}
func (m *InventoryNodes) XXX_Size() int {
	return xxx_messageInfo_InventoryNodes.Size(m)
}
func (m *InventoryNodes) XXX_DiscardUnknown() {
	xxx_messageInfo_InventoryNodes.DiscardUnknown(m)
}

var xxx_messageInfo_InventoryNodes proto.InternalMessageInfo

func (m *InventoryNodes) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *InventoryNodes) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *InventoryNodes) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *InventoryNodes) GetCursorDate() *timestamp.Timestamp {
	if m != nil {
		return m.CursorDate
	}
	return nil
}

func (m *InventoryNodes) GetCursorId() string {
	if m != nil {
		return m.CursorId
	}
	return ""
}

func (m *InventoryNodes) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *InventoryNodes) GetSorting() *Sorting {
	if m != nil {
		return m.Sorting
	}
	return nil
}

func init() {
	proto.RegisterType((*InventoryNodes)(nil), "chef.automate.domain.cfgmgmt.request.InventoryNodes")
}

func init() {
	proto.RegisterFile("interservice/cfgmgmt/request/inventory_nodes.proto", fileDescriptor_599302790193da7f)
}

var fileDescriptor_599302790193da7f = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0x95, 0x96, 0x7e, 0xb9, 0x12, 0x43, 0x06, 0x64, 0x95, 0x81, 0x08, 0x31, 0x64, 0xa0,
	0x36, 0x2a, 0x13, 0xea, 0x86, 0x90, 0x50, 0x17, 0x86, 0x94, 0x89, 0xa5, 0x72, 0x93, 0x8b, 0x6b,
	0xa9, 0xb6, 0x83, 0x7d, 0xa9, 0x44, 0x7f, 0x21, 0x3f, 0x0b, 0xa5, 0x4e, 0xd6, 0x96, 0x31, 0x97,
	0xe7, 0x79, 0xdf, 0xb3, 0x8e, 0x2c, 0x94, 0x41, 0x70, 0x1e, 0xdc, 0x41, 0xe5, 0xc0, 0xf3, 0x52,
	0x6a, 0xa9, 0x91, 0x3b, 0xf8, 0xae, 0xc1, 0x23, 0x57, 0xe6, 0x00, 0x06, 0xad, 0xfb, 0xd9, 0x18,
	0x5b, 0x80, 0x67, 0x95, 0xb3, 0x68, 0xe3, 0x87, 0x7c, 0x07, 0x25, 0x13, 0x35, 0x5a, 0x2d, 0x10,
	0x58, 0x61, 0xb5, 0x50, 0x86, 0xb5, 0x2e, 0x6b, 0xdd, 0xd9, 0x9d, 0xb4, 0x56, 0xee, 0x81, 0x9f,
	0x9c, 0x6d, 0x5d, 0x72, 0x54, 0x1a, 0x3c, 0x0a, 0x5d, 0x85, 0x98, 0xd9, 0xfc, 0x6c, 0x75, 0x25,
	0x9c, 0xd0, 0xd0, 0x00, 0x01, 0xbf, 0xff, 0xed, 0x91, 0xeb, 0x55, 0xb7, 0xcf, 0x47, 0xb3, 0x4e,
	0xfc, 0x44, 0x06, 0x1e, 0x85, 0x43, 0x1a, 0x25, 0x51, 0x3a, 0x5d, 0xcc, 0x58, 0xa8, 0x64, 0x5d,
	0x25, 0xfb, 0xec, 0x2a, 0xb3, 0x00, 0xc6, 0x8f, 0xa4, 0x0f, 0xa6, 0xa0, 0xbd, 0x8b, 0x7c, 0x83,
	0xc5, 0x37, 0x64, 0x58, 0xaa, 0x3d, 0x82, 0xa3, 0xfd, 0xa4, 0x9f, 0x4e, 0xb2, 0xf6, 0x2b, 0x5e,
	0x92, 0x69, 0x5e, 0x3b, 0x6f, 0xdd, 0xa6, 0x10, 0x08, 0xf4, 0xea, 0x62, 0x1a, 0x09, 0xf8, 0x9b,
	0x40, 0x88, 0x6f, 0xc9, 0xa4, 0x95, 0x55, 0x41, 0x07, 0x49, 0x94, 0x4e, 0xb2, 0x71, 0x18, 0xac,
	0x8a, 0xe6, 0x67, 0x25, 0x24, 0x6c, 0xbc, 0x3a, 0x02, 0x1d, 0x26, 0x51, 0x3a, 0xc8, 0xc6, 0xcd,
	0x60, 0xad, 0x8e, 0x10, 0xbf, 0x93, 0x91, 0xb7, 0x0e, 0x95, 0x91, 0x74, 0x74, 0xaa, 0x9c, 0xb3,
	0xff, 0x5c, 0x82, 0xad, 0x83, 0x94, 0x75, 0xf6, 0xeb, 0xf2, 0xeb, 0x45, 0x2a, 0xdc, 0xd5, 0x5b,
	0x96, 0x5b, 0xcd, 0x9b, 0x0c, 0xde, 0x65, 0x70, 0x51, 0x29, 0x7e, 0xee, 0x30, 0xdb, 0xe1, 0xe9,
	0x7d, 0xcf, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x1c, 0x9d, 0x7c, 0x3a, 0x02, 0x00, 0x00,
}
