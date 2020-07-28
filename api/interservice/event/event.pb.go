// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interservice/event/event.proto

package event

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type EventType struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty" toml:"Name,omitempty" mapstructure:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *EventType) Reset()         { *m = EventType{} }
func (m *EventType) String() string { return proto.CompactTextString(m) }
func (*EventType) ProtoMessage()    {}
func (*EventType) Descriptor() ([]byte, []int) {
	return fileDescriptor_7331b547e63846c1, []int{0}
}

func (m *EventType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventType.Unmarshal(m, b)
}
func (m *EventType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventType.Marshal(b, m, deterministic)
}
func (m *EventType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventType.Merge(m, src)
}
func (m *EventType) XXX_Size() int {
	return xxx_messageInfo_EventType.Size(m)
}
func (m *EventType) XXX_DiscardUnknown() {
	xxx_messageInfo_EventType.DiscardUnknown(m)
}

var xxx_messageInfo_EventType proto.InternalMessageInfo

func (m *EventType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Producer struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" toml:"ID,omitempty" mapstructure:"ID,omitempty"`
	ProducerName         string   `protobuf:"bytes,2,opt,name=ProducerName,proto3" json:"ProducerName,omitempty" toml:"ProducerName,omitempty" mapstructure:"ProducerName,omitempty"`
	ProducerType         string   `protobuf:"bytes,3,opt,name=ProducerType,proto3" json:"ProducerType,omitempty" toml:"ProducerType,omitempty" mapstructure:"ProducerType,omitempty"`
	Tags                 []string `protobuf:"bytes,4,rep,name=Tags,proto3" json:"Tags,omitempty" toml:"Tags,omitempty" mapstructure:"Tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Producer) Reset()         { *m = Producer{} }
func (m *Producer) String() string { return proto.CompactTextString(m) }
func (*Producer) ProtoMessage()    {}
func (*Producer) Descriptor() ([]byte, []int) {
	return fileDescriptor_7331b547e63846c1, []int{1}
}

func (m *Producer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Producer.Unmarshal(m, b)
}
func (m *Producer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Producer.Marshal(b, m, deterministic)
}
func (m *Producer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Producer.Merge(m, src)
}
func (m *Producer) XXX_Size() int {
	return xxx_messageInfo_Producer.Size(m)
}
func (m *Producer) XXX_DiscardUnknown() {
	xxx_messageInfo_Producer.DiscardUnknown(m)
}

var xxx_messageInfo_Producer proto.InternalMessageInfo

func (m *Producer) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Producer) GetProducerName() string {
	if m != nil {
		return m.ProducerName
	}
	return ""
}

func (m *Producer) GetProducerType() string {
	if m != nil {
		return m.ProducerType
	}
	return ""
}

func (m *Producer) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Actor struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" toml:"ID,omitempty" mapstructure:"ID,omitempty"`
	ObjectType           string   `protobuf:"bytes,2,opt,name=ObjectType,proto3" json:"ObjectType,omitempty" toml:"ObjectType,omitempty" mapstructure:"ObjectType,omitempty"`
	DisplayName          string   `protobuf:"bytes,3,opt,name=DisplayName,proto3" json:"DisplayName,omitempty" toml:"DisplayName,omitempty" mapstructure:"DisplayName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Actor) Reset()         { *m = Actor{} }
func (m *Actor) String() string { return proto.CompactTextString(m) }
func (*Actor) ProtoMessage()    {}
func (*Actor) Descriptor() ([]byte, []int) {
	return fileDescriptor_7331b547e63846c1, []int{2}
}

func (m *Actor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Actor.Unmarshal(m, b)
}
func (m *Actor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Actor.Marshal(b, m, deterministic)
}
func (m *Actor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Actor.Merge(m, src)
}
func (m *Actor) XXX_Size() int {
	return xxx_messageInfo_Actor.Size(m)
}
func (m *Actor) XXX_DiscardUnknown() {
	xxx_messageInfo_Actor.DiscardUnknown(m)
}

var xxx_messageInfo_Actor proto.InternalMessageInfo

func (m *Actor) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Actor) GetObjectType() string {
	if m != nil {
		return m.ObjectType
	}
	return ""
}

func (m *Actor) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

type Object struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" toml:"ID,omitempty" mapstructure:"ID,omitempty"`
	ObjectType           string   `protobuf:"bytes,2,opt,name=ObjectType,proto3" json:"ObjectType,omitempty" toml:"ObjectType,omitempty" mapstructure:"ObjectType,omitempty"`
	DisplayName          string   `protobuf:"bytes,3,opt,name=DisplayName,proto3" json:"DisplayName,omitempty" toml:"DisplayName,omitempty" mapstructure:"DisplayName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_7331b547e63846c1, []int{3}
}

func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Object) GetObjectType() string {
	if m != nil {
		return m.ObjectType
	}
	return ""
}

func (m *Object) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

type Target struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" toml:"ID,omitempty" mapstructure:"ID,omitempty"`
	ObjectType           string   `protobuf:"bytes,2,opt,name=ObjectType,proto3" json:"ObjectType,omitempty" toml:"ObjectType,omitempty" mapstructure:"ObjectType,omitempty"`
	DisplayName          string   `protobuf:"bytes,3,opt,name=DisplayName,proto3" json:"DisplayName,omitempty" toml:"DisplayName,omitempty" mapstructure:"DisplayName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Target) Reset()         { *m = Target{} }
func (m *Target) String() string { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()    {}
func (*Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_7331b547e63846c1, []int{4}
}

func (m *Target) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Target.Unmarshal(m, b)
}
func (m *Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Target.Marshal(b, m, deterministic)
}
func (m *Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Target.Merge(m, src)
}
func (m *Target) XXX_Size() int {
	return xxx_messageInfo_Target.Size(m)
}
func (m *Target) XXX_DiscardUnknown() {
	xxx_messageInfo_Target.DiscardUnknown(m)
}

var xxx_messageInfo_Target proto.InternalMessageInfo

func (m *Target) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Target) GetObjectType() string {
	if m != nil {
		return m.ObjectType
	}
	return ""
}

func (m *Target) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

type EventMsg struct {
	EventID              string               `protobuf:"bytes,1,opt,name=EventID,proto3" json:"EventID,omitempty" toml:"EventID,omitempty" mapstructure:"EventID,omitempty"`
	Type                 *EventType           `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty" toml:"Type,omitempty" mapstructure:"Type,omitempty"`
	Producer             *Producer            `protobuf:"bytes,3,opt,name=Producer,proto3" json:"Producer,omitempty" toml:"Producer,omitempty" mapstructure:"Producer,omitempty"`
	Tags                 []string             `protobuf:"bytes,4,rep,name=Tags,proto3" json:"Tags,omitempty" toml:"Tags,omitempty" mapstructure:"Tags,omitempty"`
	Published            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=Published,proto3" json:"Published,omitempty" toml:"Published,omitempty" mapstructure:"Published,omitempty"`
	Actor                *Actor               `protobuf:"bytes,6,opt,name=Actor,proto3" json:"Actor,omitempty" toml:"Actor,omitempty" mapstructure:"Actor,omitempty"`
	Verb                 string               `protobuf:"bytes,7,opt,name=Verb,proto3" json:"Verb,omitempty" toml:"Verb,omitempty" mapstructure:"Verb,omitempty"`
	Object               *Object              `protobuf:"bytes,8,opt,name=Object,proto3" json:"Object,omitempty" toml:"Object,omitempty" mapstructure:"Object,omitempty"`
	Target               *Target              `protobuf:"bytes,9,opt,name=Target,proto3" json:"Target,omitempty" toml:"Target,omitempty" mapstructure:"Target,omitempty"`
	Data                 *_struct.Struct      `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty" toml:"data,omitempty" mapstructure:"data,omitempty"`
	Projects             []string             `protobuf:"bytes,11,rep,name=Projects,proto3" json:"Projects,omitempty" toml:"Projects,omitempty" mapstructure:"Projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *EventMsg) Reset()         { *m = EventMsg{} }
func (m *EventMsg) String() string { return proto.CompactTextString(m) }
func (*EventMsg) ProtoMessage()    {}
func (*EventMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_7331b547e63846c1, []int{5}
}

func (m *EventMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventMsg.Unmarshal(m, b)
}
func (m *EventMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventMsg.Marshal(b, m, deterministic)
}
func (m *EventMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMsg.Merge(m, src)
}
func (m *EventMsg) XXX_Size() int {
	return xxx_messageInfo_EventMsg.Size(m)
}
func (m *EventMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMsg.DiscardUnknown(m)
}

var xxx_messageInfo_EventMsg proto.InternalMessageInfo

func (m *EventMsg) GetEventID() string {
	if m != nil {
		return m.EventID
	}
	return ""
}

func (m *EventMsg) GetType() *EventType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *EventMsg) GetProducer() *Producer {
	if m != nil {
		return m.Producer
	}
	return nil
}

func (m *EventMsg) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *EventMsg) GetPublished() *timestamp.Timestamp {
	if m != nil {
		return m.Published
	}
	return nil
}

func (m *EventMsg) GetActor() *Actor {
	if m != nil {
		return m.Actor
	}
	return nil
}

func (m *EventMsg) GetVerb() string {
	if m != nil {
		return m.Verb
	}
	return ""
}

func (m *EventMsg) GetObject() *Object {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *EventMsg) GetTarget() *Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *EventMsg) GetData() *_struct.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *EventMsg) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type EventResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty" toml:"Success,omitempty" mapstructure:"Success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *EventResponse) Reset()         { *m = EventResponse{} }
func (m *EventResponse) String() string { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()    {}
func (*EventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7331b547e63846c1, []int{6}
}

func (m *EventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventResponse.Unmarshal(m, b)
}
func (m *EventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventResponse.Marshal(b, m, deterministic)
}
func (m *EventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventResponse.Merge(m, src)
}
func (m *EventResponse) XXX_Size() int {
	return xxx_messageInfo_EventResponse.Size(m)
}
func (m *EventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventResponse proto.InternalMessageInfo

func (m *EventResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type PublishRequest struct {
	Msg                  *EventMsg `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty" toml:"Msg,omitempty" mapstructure:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte    `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32     `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7331b547e63846c1, []int{7}
}

func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequest.Unmarshal(m, b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
}
func (m *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(m, src)
}
func (m *PublishRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRequest.Size(m)
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

func (m *PublishRequest) GetMsg() *EventMsg {
	if m != nil {
		return m.Msg
	}
	return nil
}

type PublishResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty" toml:"Success,omitempty" mapstructure:"Success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *PublishResponse) Reset()         { *m = PublishResponse{} }
func (m *PublishResponse) String() string { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()    {}
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7331b547e63846c1, []int{8}
}

func (m *PublishResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishResponse.Unmarshal(m, b)
}
func (m *PublishResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishResponse.Marshal(b, m, deterministic)
}
func (m *PublishResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishResponse.Merge(m, src)
}
func (m *PublishResponse) XXX_Size() int {
	return xxx_messageInfo_PublishResponse.Size(m)
}
func (m *PublishResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublishResponse proto.InternalMessageInfo

func (m *PublishResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type SubscribeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7331b547e63846c1, []int{9}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

type SubscribeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *SubscribeResponse) Reset()         { *m = SubscribeResponse{} }
func (m *SubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeResponse) ProtoMessage()    {}
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7331b547e63846c1, []int{10}
}

func (m *SubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeResponse.Unmarshal(m, b)
}
func (m *SubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeResponse.Merge(m, src)
}
func (m *SubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeResponse.Size(m)
}
func (m *SubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeResponse proto.InternalMessageInfo

type StartRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *StartRequest) Reset()         { *m = StartRequest{} }
func (m *StartRequest) String() string { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()    {}
func (*StartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7331b547e63846c1, []int{11}
}

func (m *StartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartRequest.Unmarshal(m, b)
}
func (m *StartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartRequest.Marshal(b, m, deterministic)
}
func (m *StartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartRequest.Merge(m, src)
}
func (m *StartRequest) XXX_Size() int {
	return xxx_messageInfo_StartRequest.Size(m)
}
func (m *StartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartRequest proto.InternalMessageInfo

type StartResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *StartResponse) Reset()         { *m = StartResponse{} }
func (m *StartResponse) String() string { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()    {}
func (*StartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7331b547e63846c1, []int{12}
}

func (m *StartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartResponse.Unmarshal(m, b)
}
func (m *StartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartResponse.Marshal(b, m, deterministic)
}
func (m *StartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartResponse.Merge(m, src)
}
func (m *StartResponse) XXX_Size() int {
	return xxx_messageInfo_StartResponse.Size(m)
}
func (m *StartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartResponse proto.InternalMessageInfo

type StopRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *StopRequest) Reset()         { *m = StopRequest{} }
func (m *StopRequest) String() string { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()    {}
func (*StopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7331b547e63846c1, []int{13}
}

func (m *StopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopRequest.Unmarshal(m, b)
}
func (m *StopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopRequest.Marshal(b, m, deterministic)
}
func (m *StopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopRequest.Merge(m, src)
}
func (m *StopRequest) XXX_Size() int {
	return xxx_messageInfo_StopRequest.Size(m)
}
func (m *StopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopRequest proto.InternalMessageInfo

type StopResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *StopResponse) Reset()         { *m = StopResponse{} }
func (m *StopResponse) String() string { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()    {}
func (*StopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7331b547e63846c1, []int{14}
}

func (m *StopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopResponse.Unmarshal(m, b)
}
func (m *StopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopResponse.Marshal(b, m, deterministic)
}
func (m *StopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopResponse.Merge(m, src)
}
func (m *StopResponse) XXX_Size() int {
	return xxx_messageInfo_StopResponse.Size(m)
}
func (m *StopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventType)(nil), "chef.automate.domain.event.api.EventType")
	proto.RegisterType((*Producer)(nil), "chef.automate.domain.event.api.Producer")
	proto.RegisterType((*Actor)(nil), "chef.automate.domain.event.api.Actor")
	proto.RegisterType((*Object)(nil), "chef.automate.domain.event.api.Object")
	proto.RegisterType((*Target)(nil), "chef.automate.domain.event.api.Target")
	proto.RegisterType((*EventMsg)(nil), "chef.automate.domain.event.api.EventMsg")
	proto.RegisterType((*EventResponse)(nil), "chef.automate.domain.event.api.EventResponse")
	proto.RegisterType((*PublishRequest)(nil), "chef.automate.domain.event.api.PublishRequest")
	proto.RegisterType((*PublishResponse)(nil), "chef.automate.domain.event.api.PublishResponse")
	proto.RegisterType((*SubscribeRequest)(nil), "chef.automate.domain.event.api.SubscribeRequest")
	proto.RegisterType((*SubscribeResponse)(nil), "chef.automate.domain.event.api.SubscribeResponse")
	proto.RegisterType((*StartRequest)(nil), "chef.automate.domain.event.api.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "chef.automate.domain.event.api.StartResponse")
	proto.RegisterType((*StopRequest)(nil), "chef.automate.domain.event.api.StopRequest")
	proto.RegisterType((*StopResponse)(nil), "chef.automate.domain.event.api.StopResponse")
}

func init() {
	proto.RegisterFile("interservice/event/event.proto", fileDescriptor_7331b547e63846c1)
}

var fileDescriptor_7331b547e63846c1 = []byte{
	// 631 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0x56, 0x1b, 0xb7, 0x4d, 0x4e, 0x7a, 0xf9, 0xff, 0x61, 0x81, 0x65, 0xa1, 0xb6, 0xb2, 0x04,
	0x6a, 0x95, 0x62, 0xd3, 0xb0, 0x41, 0x20, 0x2a, 0x81, 0xc2, 0xa2, 0x12, 0x85, 0x6a, 0x1c, 0x21,
	0xd1, 0xdd, 0xd8, 0x99, 0x3a, 0xae, 0xe2, 0xd8, 0x78, 0xc6, 0x91, 0xfa, 0x18, 0x3c, 0x0c, 0xef,
	0x87, 0x7c, 0x66, 0xc6, 0x49, 0x1a, 0x84, 0xc3, 0xa2, 0x9b, 0x68, 0x2e, 0xdf, 0xe5, 0xf8, 0x9c,
	0xcf, 0x31, 0x1c, 0x26, 0x53, 0xc9, 0x0b, 0xc1, 0x8b, 0x59, 0x12, 0x71, 0x9f, 0xcf, 0xf8, 0x54,
	0xaa, 0x5f, 0x2f, 0x2f, 0x32, 0x99, 0x91, 0xc3, 0x68, 0xcc, 0x6f, 0x3d, 0x56, 0xca, 0x2c, 0x65,
	0x92, 0x7b, 0xa3, 0x2c, 0x65, 0xc9, 0xd4, 0x53, 0x08, 0x96, 0x27, 0xce, 0x51, 0x9c, 0x65, 0xf1,
	0x84, 0xfb, 0x88, 0x0e, 0xcb, 0x5b, 0x5f, 0x26, 0x29, 0x17, 0x92, 0xa5, 0xb9, 0x12, 0x70, 0x9e,
	0x3d, 0x04, 0x08, 0x59, 0x94, 0x91, 0x96, 0x77, 0x8f, 0xa0, 0xf3, 0xa9, 0xd2, 0x1a, 0xde, 0xe7,
	0x9c, 0x10, 0xb0, 0xbe, 0xb0, 0x94, 0xdb, 0x1b, 0xc7, 0x1b, 0x27, 0x1d, 0x8a, 0x6b, 0x77, 0x06,
	0xed, 0xeb, 0x22, 0x1b, 0x95, 0x11, 0x2f, 0xc8, 0x3e, 0x6c, 0x5e, 0x0e, 0xf4, 0xed, 0xe6, 0xe5,
	0x80, 0xb8, 0xb0, 0x6b, 0xee, 0x90, 0xb7, 0x89, 0x37, 0x4b, 0x67, 0x8b, 0x98, 0xca, 0xc3, 0x6e,
	0x2d, 0x63, 0x8c, 0xef, 0x90, 0xc5, 0xc2, 0xb6, 0x8e, 0x5b, 0x95, 0x6f, 0xb5, 0x76, 0xbf, 0xc3,
	0xd6, 0x87, 0x48, 0x66, 0xab, 0xa6, 0x87, 0x00, 0x5f, 0xc3, 0x3b, 0x1e, 0x61, 0xc9, 0xda, 0x72,
	0xe1, 0x84, 0x1c, 0x43, 0x77, 0x90, 0x88, 0x7c, 0xc2, 0xee, 0xb1, 0x26, 0xe5, 0xb7, 0x78, 0xe4,
	0xde, 0xc0, 0xb6, 0xc2, 0x3f, 0x8e, 0xf6, 0x90, 0x15, 0x31, 0x7f, 0x0c, 0xed, 0x9f, 0x16, 0xb4,
	0x71, 0x58, 0x57, 0x22, 0x26, 0x36, 0xec, 0xe0, 0xba, 0xf6, 0x30, 0x5b, 0xf2, 0x1e, 0xac, 0xda,
	0xa2, 0xdb, 0x3f, 0xf5, 0xfe, 0x1e, 0x20, 0xaf, 0x1e, 0x3f, 0x45, 0x1a, 0x19, 0xcc, 0x07, 0x8e,
	0x45, 0x74, 0xfb, 0x27, 0x4d, 0x12, 0x06, 0x4f, 0xe7, 0x51, 0xf9, 0xc3, 0x48, 0xc9, 0x1b, 0xe8,
	0x5c, 0x97, 0xe1, 0x24, 0x11, 0x63, 0x3e, 0xb2, 0xb7, 0x50, 0xda, 0xf1, 0x54, 0x3a, 0x3d, 0x93,
	0x4e, 0x6f, 0x68, 0xe2, 0x4b, 0xe7, 0x60, 0xf2, 0x4e, 0x87, 0xc1, 0xde, 0x46, 0xd6, 0xf3, 0xa6,
	0x82, 0x10, 0x4c, 0x75, 0x80, 0x08, 0x58, 0xdf, 0x78, 0x11, 0xda, 0x3b, 0x2a, 0xd5, 0xd5, 0x9a,
	0x5c, 0x98, 0x08, 0xd8, 0x6d, 0x54, 0x7c, 0xd1, 0xa4, 0xa8, 0xd0, 0xd4, 0x04, 0xe7, 0xc2, 0x8c,
	0xd9, 0xee, 0xac, 0xc7, 0x57, 0x68, 0x6a, 0xc2, 0xd1, 0x03, 0x6b, 0xc4, 0x24, 0xb3, 0x01, 0xd9,
	0x4f, 0x57, 0xba, 0x10, 0xe0, 0x3b, 0x4a, 0x11, 0x44, 0x1c, 0x9c, 0x48, 0xe5, 0x2b, 0xec, 0x2e,
	0xf6, 0xb3, 0xde, 0xbb, 0xa7, 0xb0, 0x87, 0x03, 0xa4, 0x5c, 0xe4, 0xd9, 0x54, 0xf0, 0x2a, 0x17,
	0x41, 0x19, 0x45, 0x5c, 0x08, 0xcc, 0x45, 0x9b, 0x9a, 0xad, 0xfb, 0x19, 0xf6, 0x75, 0x47, 0x29,
	0xff, 0x51, 0x72, 0x21, 0xc9, 0x5b, 0x68, 0x5d, 0x89, 0x18, 0x71, 0x6b, 0x4c, 0xd9, 0x44, 0x8f,
	0x56, 0x24, 0xb7, 0x07, 0x07, 0xb5, 0x5a, 0xa3, 0x35, 0x81, 0xff, 0x82, 0x32, 0x14, 0x51, 0x91,
	0x84, 0x5c, 0x9b, 0xbb, 0x4f, 0xe0, 0xff, 0x85, 0x33, 0x25, 0xe1, 0xee, 0xc3, 0x6e, 0x20, 0x59,
	0x21, 0x0d, 0xe8, 0x00, 0xf6, 0xf4, 0x5e, 0x03, 0xf6, 0xa0, 0x1b, 0xc8, 0x2c, 0x37, 0xf7, 0x88,
	0xaf, 0xb6, 0xea, 0xba, 0xff, 0xab, 0x05, 0xbb, 0x58, 0x67, 0xa0, 0xfe, 0x50, 0xc9, 0x1d, 0xec,
	0xe8, 0x32, 0x89, 0xd7, 0x18, 0xe3, 0xa5, 0xee, 0x38, 0xfe, 0xda, 0x78, 0xfd, 0xfc, 0x39, 0x74,
	0xea, 0x27, 0x22, 0xaf, 0x9a, 0xd8, 0x0f, 0x1b, 0xe2, 0x9c, 0xff, 0x03, 0x43, 0x3b, 0x8e, 0x60,
	0x0b, 0xdb, 0x43, 0xce, 0x1a, 0xb9, 0x0b, 0x5d, 0x75, 0x5e, 0xae, 0x89, 0xd6, 0x2e, 0x0c, 0xac,
	0xaa, 0xc9, 0xa4, 0xd7, 0x4c, 0xab, 0x27, 0xe3, 0x9c, 0xad, 0x07, 0x56, 0x16, 0x1f, 0xcf, 0x6f,
	0xfc, 0x38, 0x91, 0xe3, 0x32, 0xf4, 0xa2, 0x2c, 0xf5, 0x2b, 0xa6, 0x6f, 0x98, 0x3e, 0xcb, 0x13,
	0x7f, 0xf5, 0x23, 0x19, 0x6e, 0xe3, 0xcb, 0xf2, 0xfa, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb2,
	0xed, 0xef, 0x3a, 0x41, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventServiceClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error)
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.event.api.EventService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error) {
	out := new(SubscribeResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.event.api.EventService/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.event.api.EventService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.event.api.EventService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
type EventServiceServer interface {
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error)
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Stop(context.Context, *StopRequest) (*StopResponse, error)
}

// UnimplementedEventServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (*UnimplementedEventServiceServer) Publish(ctx context.Context, req *PublishRequest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (*UnimplementedEventServiceServer) Subscribe(ctx context.Context, req *SubscribeRequest) (*SubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (*UnimplementedEventServiceServer) Start(ctx context.Context, req *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedEventServiceServer) Stop(ctx context.Context, req *StopRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.event.api.EventService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.event.api.EventService/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.event.api.EventService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.event.api.EventService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.event.api.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _EventService_Publish_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _EventService_Subscribe_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _EventService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _EventService_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interservice/event/event.proto",
}
