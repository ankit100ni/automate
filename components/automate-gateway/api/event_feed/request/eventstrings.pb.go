// Code generated by protoc-gen-go. DO NOT EDIT.
// source: automate-gateway/api/event_feed/request/eventstrings.proto

package request

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

type EventStrings struct {
	// Earliest events to return.
	Start string `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	// Latest events to return.
	End string `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	// User timezone to apply to request.
	Timezone string `protobuf:"bytes,3,opt,name=timezone,proto3" json:"timezone,omitempty"`
	// Interval for returned results, for example: 1 hour buckets.
	HoursBetween int32 `protobuf:"varint,4,opt,name=hours_between,json=hoursBetween,proto3" json:"hours_between,omitempty"`
	// Filters to be applied to the request.
	Filter               []string `protobuf:"bytes,5,rep,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventStrings) Reset()         { *m = EventStrings{} }
func (m *EventStrings) String() string { return proto.CompactTextString(m) }
func (*EventStrings) ProtoMessage()    {}
func (*EventStrings) Descriptor() ([]byte, []int) {
	return fileDescriptor_c713500844a16f14, []int{0}
}

func (m *EventStrings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventStrings.Unmarshal(m, b)
}
func (m *EventStrings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventStrings.Marshal(b, m, deterministic)
}
func (m *EventStrings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventStrings.Merge(m, src)
}
func (m *EventStrings) XXX_Size() int {
	return xxx_messageInfo_EventStrings.Size(m)
}
func (m *EventStrings) XXX_DiscardUnknown() {
	xxx_messageInfo_EventStrings.DiscardUnknown(m)
}

var xxx_messageInfo_EventStrings proto.InternalMessageInfo

func (m *EventStrings) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *EventStrings) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func (m *EventStrings) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *EventStrings) GetHoursBetween() int32 {
	if m != nil {
		return m.HoursBetween
	}
	return 0
}

func (m *EventStrings) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func init() {
	proto.RegisterType((*EventStrings)(nil), "chef.automate.api.event_feed.request.EventStrings")
}

func init() {
	proto.RegisterFile("automate-gateway/api/event_feed/request/eventstrings.proto", fileDescriptor_c713500844a16f14)
}

var fileDescriptor_c713500844a16f14 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4b, 0xc4, 0x30,
	0x18, 0xc5, 0xa9, 0xb5, 0x87, 0x17, 0x4e, 0x90, 0x20, 0x12, 0x9c, 0x8a, 0x3a, 0x74, 0x31, 0x19,
	0xdc, 0x1c, 0x0f, 0x9c, 0xc4, 0xa5, 0x6e, 0x2e, 0x47, 0xda, 0x7b, 0x6d, 0x03, 0x36, 0xa9, 0xc9,
	0x57, 0x0f, 0xfd, 0x27, 0xfc, 0x97, 0xa5, 0xcd, 0x9d, 0xae, 0xb7, 0x7d, 0xef, 0x07, 0xef, 0x83,
	0xdf, 0x63, 0x8f, 0x7a, 0x24, 0xd7, 0x6b, 0xc2, 0x7d, 0xab, 0x09, 0x3b, 0xfd, 0xa5, 0xf4, 0x60,
	0x14, 0x3e, 0x61, 0x69, 0xd3, 0x00, 0x5b, 0xe5, 0xf1, 0x31, 0x22, 0x50, 0x44, 0x81, 0xbc, 0xb1,
	0x6d, 0x90, 0x83, 0x77, 0xe4, 0xf8, 0x5d, 0xdd, 0xa1, 0x91, 0x87, 0x07, 0x52, 0x0f, 0x46, 0xfe,
	0x17, 0xe5, 0xbe, 0x78, 0xf3, 0x93, 0xb0, 0xd5, 0xd3, 0x84, 0x5f, 0x63, 0x99, 0x5f, 0xb2, 0x2c,
	0x90, 0xf6, 0x24, 0x92, 0x3c, 0x29, 0x96, 0x65, 0x0c, 0xfc, 0x82, 0xa5, 0xb0, 0x5b, 0x71, 0x32,
	0xb3, 0xe9, 0xe4, 0xd7, 0xec, 0x8c, 0x4c, 0x8f, 0x6f, 0x67, 0x21, 0xd2, 0x19, 0xff, 0x65, 0x7e,
	0xcb, 0xce, 0x3b, 0x37, 0xfa, 0xb0, 0xa9, 0x40, 0x3b, 0xc0, 0x8a, 0xd3, 0x3c, 0x29, 0xb2, 0x72,
	0x35, 0xc3, 0x75, 0x64, 0xfc, 0x8a, 0x2d, 0x1a, 0xf3, 0x4e, 0xf0, 0x22, 0xcb, 0xd3, 0x62, 0x59,
	0xee, 0xd3, 0xfa, 0xe5, 0xed, 0xb9, 0x35, 0xd4, 0x8d, 0x95, 0xac, 0x5d, 0xaf, 0x26, 0x09, 0x75,
	0x90, 0x50, 0xb5, 0xeb, 0x07, 0x67, 0x27, 0x55, 0x75, 0xe4, 0x32, 0xd5, 0x62, 0x5e, 0xe3, 0xe1,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x34, 0x4c, 0xda, 0x4e, 0x4b, 0x01, 0x00, 0x00,
}
