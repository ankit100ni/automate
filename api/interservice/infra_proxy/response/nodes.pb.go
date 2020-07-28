// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interservice/infra_proxy/response/nodes.proto

package response

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

type AffectedNodes struct {
	// List of the nodes which are affected by the chef object.
	Nodes                []*NodeAttribute `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty" toml:"nodes,omitempty" mapstructure:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte           `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32            `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *AffectedNodes) Reset()         { *m = AffectedNodes{} }
func (m *AffectedNodes) String() string { return proto.CompactTextString(m) }
func (*AffectedNodes) ProtoMessage()    {}
func (*AffectedNodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_3abe3cd8dc5dfbb0, []int{0}
}

func (m *AffectedNodes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AffectedNodes.Unmarshal(m, b)
}
func (m *AffectedNodes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AffectedNodes.Marshal(b, m, deterministic)
}
func (m *AffectedNodes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AffectedNodes.Merge(m, src)
}
func (m *AffectedNodes) XXX_Size() int {
	return xxx_messageInfo_AffectedNodes.Size(m)
}
func (m *AffectedNodes) XXX_DiscardUnknown() {
	xxx_messageInfo_AffectedNodes.DiscardUnknown(m)
}

var xxx_messageInfo_AffectedNodes proto.InternalMessageInfo

func (m *AffectedNodes) GetNodes() []*NodeAttribute {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type NodeAttribute struct {
	// Node ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	// Node name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Node last checkin.
	CheckIn string `protobuf:"bytes,3,opt,name=check_in,json=checkIn,proto3" json:"check_in,omitempty" toml:"check_in,omitempty" mapstructure:"check_in,omitempty"`
	// Node uptime.
	Uptime string `protobuf:"bytes,4,opt,name=uptime,proto3" json:"uptime,omitempty" toml:"uptime,omitempty" mapstructure:"uptime,omitempty"`
	// Node platform.
	Platform string `protobuf:"bytes,5,opt,name=platform,proto3" json:"platform,omitempty" toml:"platform,omitempty" mapstructure:"platform,omitempty"`
	// Node environment name.
	Environment string `protobuf:"bytes,6,opt,name=environment,proto3" json:"environment,omitempty" toml:"environment,omitempty" mapstructure:"environment,omitempty"`
	// Node policy group.
	PolicyGroup          string   `protobuf:"bytes,7,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty" toml:"policy_group,omitempty" mapstructure:"policy_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *NodeAttribute) Reset()         { *m = NodeAttribute{} }
func (m *NodeAttribute) String() string { return proto.CompactTextString(m) }
func (*NodeAttribute) ProtoMessage()    {}
func (*NodeAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_3abe3cd8dc5dfbb0, []int{1}
}

func (m *NodeAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeAttribute.Unmarshal(m, b)
}
func (m *NodeAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeAttribute.Marshal(b, m, deterministic)
}
func (m *NodeAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeAttribute.Merge(m, src)
}
func (m *NodeAttribute) XXX_Size() int {
	return xxx_messageInfo_NodeAttribute.Size(m)
}
func (m *NodeAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_NodeAttribute proto.InternalMessageInfo

func (m *NodeAttribute) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NodeAttribute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeAttribute) GetCheckIn() string {
	if m != nil {
		return m.CheckIn
	}
	return ""
}

func (m *NodeAttribute) GetUptime() string {
	if m != nil {
		return m.Uptime
	}
	return ""
}

func (m *NodeAttribute) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *NodeAttribute) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *NodeAttribute) GetPolicyGroup() string {
	if m != nil {
		return m.PolicyGroup
	}
	return ""
}

type DeleteNode struct {
	// Node name.
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *DeleteNode) Reset()         { *m = DeleteNode{} }
func (m *DeleteNode) String() string { return proto.CompactTextString(m) }
func (*DeleteNode) ProtoMessage()    {}
func (*DeleteNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_3abe3cd8dc5dfbb0, []int{2}
}

func (m *DeleteNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNode.Unmarshal(m, b)
}
func (m *DeleteNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNode.Marshal(b, m, deterministic)
}
func (m *DeleteNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNode.Merge(m, src)
}
func (m *DeleteNode) XXX_Size() int {
	return xxx_messageInfo_DeleteNode.Size(m)
}
func (m *DeleteNode) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNode.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNode proto.InternalMessageInfo

func (m *DeleteNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*AffectedNodes)(nil), "chef.automate.domain.infra_proxy.response.AffectedNodes")
	proto.RegisterType((*NodeAttribute)(nil), "chef.automate.domain.infra_proxy.response.NodeAttribute")
	proto.RegisterType((*DeleteNode)(nil), "chef.automate.domain.infra_proxy.response.DeleteNode")
}

func init() {
	proto.RegisterFile("interservice/infra_proxy/response/nodes.proto", fileDescriptor_3abe3cd8dc5dfbb0)
}

var fileDescriptor_3abe3cd8dc5dfbb0 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3b, 0x4f, 0xc3, 0x30,
	0x14, 0x85, 0x95, 0xbe, 0x71, 0x29, 0x83, 0x07, 0x64, 0x98, 0x42, 0xa7, 0x32, 0xe0, 0x48, 0xb0,
	0x30, 0x21, 0x8a, 0x90, 0x10, 0x4b, 0x87, 0x8e, 0x2c, 0x91, 0xeb, 0xdc, 0xb4, 0x57, 0xd4, 0x0f,
	0x39, 0x37, 0x15, 0xfd, 0x85, 0xfc, 0x2d, 0x14, 0x97, 0x96, 0x32, 0xc1, 0xe6, 0x73, 0x8e, 0xef,
	0xa7, 0xfb, 0x60, 0x37, 0x68, 0x09, 0x42, 0x05, 0x61, 0x83, 0x1a, 0x32, 0xb4, 0x65, 0x50, 0xb9,
	0x0f, 0xee, 0x63, 0x9b, 0x05, 0xa8, 0xbc, 0xb3, 0x15, 0x64, 0xd6, 0x15, 0x50, 0x49, 0x1f, 0x1c,
	0x39, 0x7e, 0xad, 0x57, 0x50, 0x4a, 0x55, 0x93, 0x33, 0x8a, 0x40, 0x16, 0xce, 0x28, 0xb4, 0xf2,
	0xa8, 0x4c, 0xee, 0xcb, 0xc6, 0x39, 0x1b, 0x4d, 0xcb, 0x12, 0x34, 0x41, 0x31, 0x6b, 0x08, 0x7c,
	0xc6, 0xba, 0x11, 0x25, 0x92, 0xb4, 0x3d, 0x19, 0xde, 0xde, 0xcb, 0x7f, 0xb3, 0x64, 0x03, 0x98,
	0x12, 0x05, 0x5c, 0xd4, 0x04, 0xf3, 0x1d, 0x66, 0xfc, 0x99, 0xb0, 0xd1, 0xaf, 0x80, 0x9f, 0xb1,
	0x16, 0x16, 0x22, 0x49, 0x93, 0xc9, 0xc9, 0xbc, 0x85, 0x05, 0xe7, 0xac, 0x63, 0x95, 0x01, 0xd1,
	0x8a, 0x4e, 0x7c, 0xf3, 0x0b, 0x36, 0xd0, 0x2b, 0xd0, 0xef, 0x39, 0x5a, 0xd1, 0x8e, 0x7e, 0x3f,
	0xea, 0x57, 0xcb, 0xcf, 0x59, 0xaf, 0xf6, 0x84, 0x06, 0x44, 0x27, 0x06, 0xdf, 0x8a, 0x5f, 0xb2,
	0x81, 0x5f, 0x2b, 0x2a, 0x5d, 0x30, 0xa2, 0x1b, 0x93, 0x83, 0xe6, 0x29, 0x1b, 0x82, 0xdd, 0x60,
	0x70, 0xd6, 0x80, 0x25, 0xd1, 0x8b, 0xf1, 0xb1, 0xc5, 0xaf, 0xd8, 0xa9, 0x77, 0x6b, 0xd4, 0xdb,
	0x7c, 0x19, 0x5c, 0xed, 0x45, 0x7f, 0xf7, 0x65, 0xe7, 0xbd, 0x34, 0xd6, 0x38, 0x65, 0xec, 0x19,
	0xd6, 0x40, 0xd0, 0x8c, 0x73, 0xe8, 0xba, 0xfd, 0xd3, 0xf5, 0xd3, 0xe3, 0xdb, 0xc3, 0x12, 0x69,
	0x55, 0x2f, 0xa4, 0x76, 0x26, 0x6b, 0x16, 0x97, 0xed, 0x17, 0x97, 0x29, 0x8f, 0xd9, 0x9f, 0x57,
	0x5c, 0xf4, 0xe2, 0x01, 0xef, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x24, 0x4a, 0x54, 0xf4, 0xf1,
	0x01, 0x00, 0x00,
}
