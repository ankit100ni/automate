// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interservice/infra_proxy/response/orgs.proto

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

type CreateOrg struct {
	// Chef organization.
	Org                  *Org     `protobuf:"bytes,1,opt,name=org,proto3" json:"org,omitempty" toml:"org,omitempty" mapstructure:"org,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *CreateOrg) Reset()         { *m = CreateOrg{} }
func (m *CreateOrg) String() string { return proto.CompactTextString(m) }
func (*CreateOrg) ProtoMessage()    {}
func (*CreateOrg) Descriptor() ([]byte, []int) {
	return fileDescriptor_a35bb51085006cf6, []int{0}
}

func (m *CreateOrg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrg.Unmarshal(m, b)
}
func (m *CreateOrg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrg.Marshal(b, m, deterministic)
}
func (m *CreateOrg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrg.Merge(m, src)
}
func (m *CreateOrg) XXX_Size() int {
	return xxx_messageInfo_CreateOrg.Size(m)
}
func (m *CreateOrg) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrg.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrg proto.InternalMessageInfo

func (m *CreateOrg) GetOrg() *Org {
	if m != nil {
		return m.Org
	}
	return nil
}

type DeleteOrg struct {
	// Chef organization.
	Org                  *Org     `protobuf:"bytes,1,opt,name=org,proto3" json:"org,omitempty" toml:"org,omitempty" mapstructure:"org,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *DeleteOrg) Reset()         { *m = DeleteOrg{} }
func (m *DeleteOrg) String() string { return proto.CompactTextString(m) }
func (*DeleteOrg) ProtoMessage()    {}
func (*DeleteOrg) Descriptor() ([]byte, []int) {
	return fileDescriptor_a35bb51085006cf6, []int{1}
}

func (m *DeleteOrg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteOrg.Unmarshal(m, b)
}
func (m *DeleteOrg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteOrg.Marshal(b, m, deterministic)
}
func (m *DeleteOrg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteOrg.Merge(m, src)
}
func (m *DeleteOrg) XXX_Size() int {
	return xxx_messageInfo_DeleteOrg.Size(m)
}
func (m *DeleteOrg) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteOrg.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteOrg proto.InternalMessageInfo

func (m *DeleteOrg) GetOrg() *Org {
	if m != nil {
		return m.Org
	}
	return nil
}

type UpdateOrg struct {
	// Chef organization.
	Org                  *Org     `protobuf:"bytes,1,opt,name=org,proto3" json:"org,omitempty" toml:"org,omitempty" mapstructure:"org,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *UpdateOrg) Reset()         { *m = UpdateOrg{} }
func (m *UpdateOrg) String() string { return proto.CompactTextString(m) }
func (*UpdateOrg) ProtoMessage()    {}
func (*UpdateOrg) Descriptor() ([]byte, []int) {
	return fileDescriptor_a35bb51085006cf6, []int{2}
}

func (m *UpdateOrg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateOrg.Unmarshal(m, b)
}
func (m *UpdateOrg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateOrg.Marshal(b, m, deterministic)
}
func (m *UpdateOrg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateOrg.Merge(m, src)
}
func (m *UpdateOrg) XXX_Size() int {
	return xxx_messageInfo_UpdateOrg.Size(m)
}
func (m *UpdateOrg) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateOrg.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateOrg proto.InternalMessageInfo

func (m *UpdateOrg) GetOrg() *Org {
	if m != nil {
		return m.Org
	}
	return nil
}

type GetOrgs struct {
	// Chef organization list.
	Orgs                 []*Org   `protobuf:"bytes,1,rep,name=orgs,proto3" json:"orgs,omitempty" toml:"orgs,omitempty" mapstructure:"orgs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *GetOrgs) Reset()         { *m = GetOrgs{} }
func (m *GetOrgs) String() string { return proto.CompactTextString(m) }
func (*GetOrgs) ProtoMessage()    {}
func (*GetOrgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_a35bb51085006cf6, []int{3}
}

func (m *GetOrgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrgs.Unmarshal(m, b)
}
func (m *GetOrgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrgs.Marshal(b, m, deterministic)
}
func (m *GetOrgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrgs.Merge(m, src)
}
func (m *GetOrgs) XXX_Size() int {
	return xxx_messageInfo_GetOrgs.Size(m)
}
func (m *GetOrgs) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrgs.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrgs proto.InternalMessageInfo

func (m *GetOrgs) GetOrgs() []*Org {
	if m != nil {
		return m.Orgs
	}
	return nil
}

type GetOrg struct {
	// Chef organization.
	Org                  *Org     `protobuf:"bytes,1,opt,name=org,proto3" json:"org,omitempty" toml:"org,omitempty" mapstructure:"org,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *GetOrg) Reset()         { *m = GetOrg{} }
func (m *GetOrg) String() string { return proto.CompactTextString(m) }
func (*GetOrg) ProtoMessage()    {}
func (*GetOrg) Descriptor() ([]byte, []int) {
	return fileDescriptor_a35bb51085006cf6, []int{4}
}

func (m *GetOrg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrg.Unmarshal(m, b)
}
func (m *GetOrg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrg.Marshal(b, m, deterministic)
}
func (m *GetOrg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrg.Merge(m, src)
}
func (m *GetOrg) XXX_Size() int {
	return xxx_messageInfo_GetOrg.Size(m)
}
func (m *GetOrg) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrg.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrg proto.InternalMessageInfo

func (m *GetOrg) GetOrg() *Org {
	if m != nil {
		return m.Org
	}
	return nil
}

type Org struct {
	// Chef organization ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	// Chef organization name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Chef organization admin user.
	AdminUser string `protobuf:"bytes,3,opt,name=admin_user,json=adminUser,proto3" json:"admin_user,omitempty" toml:"admin_user,omitempty" mapstructure:"admin_user,omitempty"`
	// Chef organization credential ID.
	CredentialId string `protobuf:"bytes,4,opt,name=credential_id,json=credentialId,proto3" json:"credential_id,omitempty" toml:"credential_id,omitempty" mapstructure:"credential_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,5,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// List of projects this chef organization belongs to. May be empty.
	Projects             []string `protobuf:"bytes,6,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Org) Reset()         { *m = Org{} }
func (m *Org) String() string { return proto.CompactTextString(m) }
func (*Org) ProtoMessage()    {}
func (*Org) Descriptor() ([]byte, []int) {
	return fileDescriptor_a35bb51085006cf6, []int{5}
}

func (m *Org) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Org.Unmarshal(m, b)
}
func (m *Org) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Org.Marshal(b, m, deterministic)
}
func (m *Org) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Org.Merge(m, src)
}
func (m *Org) XXX_Size() int {
	return xxx_messageInfo_Org.Size(m)
}
func (m *Org) XXX_DiscardUnknown() {
	xxx_messageInfo_Org.DiscardUnknown(m)
}

var xxx_messageInfo_Org proto.InternalMessageInfo

func (m *Org) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Org) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Org) GetAdminUser() string {
	if m != nil {
		return m.AdminUser
	}
	return ""
}

func (m *Org) GetCredentialId() string {
	if m != nil {
		return m.CredentialId
	}
	return ""
}

func (m *Org) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *Org) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type ResetOrgAdminKey struct {
	// Chef organization.
	Org                  *Org     `protobuf:"bytes,1,opt,name=org,proto3" json:"org,omitempty" toml:"org,omitempty" mapstructure:"org,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ResetOrgAdminKey) Reset()         { *m = ResetOrgAdminKey{} }
func (m *ResetOrgAdminKey) String() string { return proto.CompactTextString(m) }
func (*ResetOrgAdminKey) ProtoMessage()    {}
func (*ResetOrgAdminKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_a35bb51085006cf6, []int{6}
}

func (m *ResetOrgAdminKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetOrgAdminKey.Unmarshal(m, b)
}
func (m *ResetOrgAdminKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetOrgAdminKey.Marshal(b, m, deterministic)
}
func (m *ResetOrgAdminKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetOrgAdminKey.Merge(m, src)
}
func (m *ResetOrgAdminKey) XXX_Size() int {
	return xxx_messageInfo_ResetOrgAdminKey.Size(m)
}
func (m *ResetOrgAdminKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetOrgAdminKey.DiscardUnknown(m)
}

var xxx_messageInfo_ResetOrgAdminKey proto.InternalMessageInfo

func (m *ResetOrgAdminKey) GetOrg() *Org {
	if m != nil {
		return m.Org
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateOrg)(nil), "chef.automate.domain.infra_proxy.response.CreateOrg")
	proto.RegisterType((*DeleteOrg)(nil), "chef.automate.domain.infra_proxy.response.DeleteOrg")
	proto.RegisterType((*UpdateOrg)(nil), "chef.automate.domain.infra_proxy.response.UpdateOrg")
	proto.RegisterType((*GetOrgs)(nil), "chef.automate.domain.infra_proxy.response.GetOrgs")
	proto.RegisterType((*GetOrg)(nil), "chef.automate.domain.infra_proxy.response.GetOrg")
	proto.RegisterType((*Org)(nil), "chef.automate.domain.infra_proxy.response.Org")
	proto.RegisterType((*ResetOrgAdminKey)(nil), "chef.automate.domain.infra_proxy.response.ResetOrgAdminKey")
}

func init() {
	proto.RegisterFile("interservice/infra_proxy/response/orgs.proto", fileDescriptor_a35bb51085006cf6)
}

var fileDescriptor_a35bb51085006cf6 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4b, 0x33, 0x31,
	0x10, 0xc6, 0xd9, 0x6e, 0xdf, 0xbe, 0xdd, 0xf1, 0x0f, 0x92, 0xd3, 0xa2, 0x08, 0xa5, 0x5e, 0x14,
	0x24, 0x01, 0xbd, 0x4b, 0xad, 0x82, 0x54, 0x91, 0x42, 0xb1, 0x17, 0x2f, 0x25, 0xdd, 0x4c, 0xb7,
	0x91, 0x6e, 0xb2, 0x4c, 0x52, 0xb1, 0x5f, 0xc8, 0xcf, 0x29, 0x49, 0x29, 0x7a, 0x53, 0xb4, 0xb7,
	0xe4, 0x99, 0x79, 0x7e, 0x99, 0x09, 0x0f, 0x9c, 0x6b, 0xe3, 0x91, 0x1c, 0xd2, 0xab, 0x2e, 0x50,
	0x68, 0x33, 0x23, 0x39, 0xa9, 0xc9, 0xbe, 0xad, 0x04, 0xa1, 0xab, 0xad, 0x71, 0x28, 0x2c, 0x95,
	0x8e, 0xd7, 0x64, 0xbd, 0x65, 0x67, 0xc5, 0x1c, 0x67, 0x5c, 0x2e, 0xbd, 0xad, 0xa4, 0x47, 0xae,
	0x6c, 0x25, 0xb5, 0xe1, 0x5f, 0x5c, 0x7c, 0xe3, 0xea, 0x3e, 0x42, 0x76, 0x43, 0x28, 0x3d, 0x0e,
	0xa9, 0x64, 0x3d, 0x48, 0x2d, 0x95, 0x79, 0xd2, 0x49, 0x4e, 0x77, 0x2e, 0x38, 0xff, 0x31, 0x85,
	0x0f, 0xa9, 0x1c, 0x05, 0x6b, 0xc0, 0xdd, 0xe2, 0x02, 0xb7, 0x88, 0x1b, 0xd7, 0x6a, 0x8b, 0xd3,
	0xfd, 0xbf, 0x43, 0x3f, 0xa4, 0xd2, 0xb1, 0x3e, 0x34, 0xc3, 0x87, 0xe5, 0x49, 0x27, 0xfd, 0x05,
	0x2d, 0x7a, 0xbb, 0xf7, 0xd0, 0x5a, 0xe3, 0xb6, 0x30, 0xda, 0x7b, 0x02, 0x69, 0x20, 0xed, 0x43,
	0x43, 0xab, 0x08, 0xca, 0x46, 0x0d, 0xad, 0x18, 0x83, 0xa6, 0x91, 0x15, 0xe6, 0x8d, 0xa8, 0xc4,
	0x33, 0x3b, 0x06, 0x90, 0xaa, 0xd2, 0x66, 0xb2, 0x74, 0x48, 0x79, 0x1a, 0x2b, 0x59, 0x54, 0xc6,
	0x0e, 0x89, 0x9d, 0xc0, 0x5e, 0x41, 0xa8, 0xd0, 0x78, 0x2d, 0x17, 0x13, 0xad, 0xf2, 0x66, 0xec,
	0xd8, 0xfd, 0x14, 0x07, 0x8a, 0x1d, 0x41, 0x16, 0xd2, 0x84, 0x14, 0x1a, 0xfe, 0xc5, 0x86, 0xf6,
	0x5a, 0x18, 0x28, 0x76, 0x08, 0xed, 0x9a, 0xec, 0x0b, 0x16, 0xde, 0xe5, 0xad, 0x4e, 0x1a, 0x6a,
	0x9b, 0x7b, 0xf7, 0x09, 0x0e, 0x46, 0xe8, 0xe2, 0xda, 0xd7, 0xe1, 0xc9, 0x07, 0x5c, 0xfd, 0x7d,
	0xfd, 0x7e, 0xef, 0xf9, 0xaa, 0xd4, 0x7e, 0xbe, 0x9c, 0xf2, 0xc2, 0x56, 0x22, 0x00, 0xc4, 0x06,
	0x20, 0x64, 0xad, 0xc5, 0xb7, 0xf1, 0x9f, 0xb6, 0x62, 0xf4, 0x2f, 0x3f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xf0, 0x8b, 0x30, 0xae, 0x2a, 0x03, 0x00, 0x00,
}
