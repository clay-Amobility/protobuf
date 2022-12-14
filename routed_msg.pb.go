routed_msg.proto                                                                                    0000664 0001751 0001751 00000007657 14335525512 014235  0                                                                                                    ustar   tfisher                         tfisher                                                                                                                                                                                                                syntax = "proto3";

import "google/protobuf/timestamp.proto";

package routedmsgpb;

message Position {
  int32 latitude = 1; // in 1/10th microdegrees [-900000000,900000001]
  int32 longitude = 2; // in 1/10th microdegrees [-1799999999,1800000001]
}

message ClientIdentity {
  uint32 clientId = 1;
  string entityType = 2;
}

message RoutedMsg {
  bytes msgBytes = 1;
  google.protobuf.Timestamp time = 2; // in UTC time
  Position position = 3;
  // optional...only applies to some types of messages
  uint32 customRadiusOpt = 4;
  bool clientCanOverrideRadiusOpt = 5; // only applies if customRadiusOpt is set
  repeated ClientIdentity clientIdsOpt = 6;  // if specified, these are the
  //potential clients to receive the message, rather than those within a geofence
}

message AutoPublishAddMsg {
 bytes msgBytes = 1;
 string msgType = 2; // such as "TIM", "RSA", etc
 string id = 3; // must just be unique for the publishing client for the given
               // msgType
 string description = 11;
 repeated string entityTypes = 12; // optional, to publish to certain
                                                  // EntityTypes and not others; if
                                                  // empty it implies send to everyone
 // the next 4 fields pertain to messages that should be delivered to everyone
 // within a radius at some periodicity
 Position positionOpt = 4;
 uint32 customRadiusOpt = 5; // optional;
 bool clientCanOverrideRadiusOpt = 6; // only applies if customRadiusOpt is set
 uint32 frequencySecOpt = 7;
 google.protobuf.Timestamp startTime = 8; // just set to current time if it???s
                                         // not in the future
 google.protobuf.Timestamp endTime = 9; // this should be set within 24 hours of
                                       // startTime
 // if LineCrossing is defined then deliver messages to anyone crossing that
 // line
 LineCrossing lineCrossingOpt = 10;
}

message LineCrossing {
 Position endpoint1 = 1;
 Position endpoint2 = 2;
 string roadIdentifier = 3;
 enum Direction {
   Undefined = 0;
   NB = 1;
   EB = 2;
   SB = 3;
   WB = 4;
 }
 Direction direction = 4;
}

// this is the message a client should publish to remove an Autopublished Message
// that it added
message AutoPublishDeleteMsg {
  string id = 1;
  string msgType = 2; // such as "TIM", "RSA", etc
}

// acknowledgment for the AutoPublishAddMsg
message AutoPublishAddAck {
 string msgType = 1; // such as "TIM", "RSA", etc
 string id = 2;
 string msgDescription = 5; // matches 'description' from original AutoPublishAddMsg
 bool success = 3;
 string failMsg = 4; // only set in case of failure
}

// acknowledgment for the AutoPublishDeleteMsg
message AutoPublishDeleteAck {
 string msgType = 1; // such as "TIM", "RSA", etc
 string id = 2;
 string msgDescription = 5; // matches 'description' from original AutoPublishAddMsg
 bool success = 3;
 string failMsg = 4; // only set in case of failure
}

// acknowledgment for the request to delete all messages for the client (which has no associated Protobuf)
message AutoPublishDeleteAllAck {
 bool success = 1;
 string failMsg = 2; // only set in case of failure
 repeated string ids = 3;
}

// response to client that a message started (applies both to messages that were scheduled and to messages that were started immediately)
message AutoPublishStarted {
 string msgType = 1; // such as "TIM", "RSA", etc
 string id = 2;
 string msgDescription = 3; // matches 'description' from original AutoPublishAddMsg
}

// response to client that a message stopped (applies both to messages that expired and to messages that were stopped immediately)
message AutoPublishStopped {
 string msgType = 1; // such as "TIM", "RSA", etc
 string id = 2;
 string msgDescription = 3; // matches 'description' from original AutoPublishAddMsg
}

message OutboundMsg {
 ClientIdentity clientIdentity = 1; // identifies the publisher
 google.protobuf.Timestamp time = 2; // in UTC time.SECS ONLY nanos is 0
 bytes msgBytes = 3;
}
                                                                                 static_msg.pb.go                                                                                    0000664 0001751 0001751 00000052211 14335525601 014045  0                                                                                                    ustar   tfisher                         tfisher                                                                                                                                                                                                                // Code generated by protoc-gen-go. DO NOT EDIT.
// source: static_msg.proto

package staticmsgpb

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

type Position struct {
	Latitude             int32    `protobuf:"varint,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            int32    `protobuf:"varint,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4ae77d93c5d83be, []int{0}
}

func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetLatitude() int32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Position) GetLongitude() int32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type StaticAddMsg struct {
	MsgBytes             []byte               `protobuf:"bytes,1,opt,name=msgBytes,proto3" json:"msgBytes,omitempty"`
	MsgType              string               `protobuf:"bytes,2,opt,name=msgType,proto3" json:"msgType,omitempty"`
	Id                   string               `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Description          string               `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Position             *Position            `protobuf:"bytes,6,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StaticAddMsg) Reset()         { *m = StaticAddMsg{} }
func (m *StaticAddMsg) String() string { return proto.CompactTextString(m) }
func (*StaticAddMsg) ProtoMessage()    {}
func (*StaticAddMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4ae77d93c5d83be, []int{1}
}

func (m *StaticAddMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticAddMsg.Unmarshal(m, b)
}
func (m *StaticAddMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticAddMsg.Marshal(b, m, deterministic)
}
func (m *StaticAddMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticAddMsg.Merge(m, src)
}
func (m *StaticAddMsg) XXX_Size() int {
	return xxx_messageInfo_StaticAddMsg.Size(m)
}
func (m *StaticAddMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticAddMsg.DiscardUnknown(m)
}

var xxx_messageInfo_StaticAddMsg proto.InternalMessageInfo

func (m *StaticAddMsg) GetMsgBytes() []byte {
	if m != nil {
		return m.MsgBytes
	}
	return nil
}

func (m *StaticAddMsg) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *StaticAddMsg) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StaticAddMsg) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *StaticAddMsg) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *StaticAddMsg) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *StaticAddMsg) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

type StaticDeleteMsg struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StaticDeleteMsg) Reset()         { *m = StaticDeleteMsg{} }
func (m *StaticDeleteMsg) String() string { return proto.CompactTextString(m) }
func (*StaticDeleteMsg) ProtoMessage()    {}
func (*StaticDeleteMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4ae77d93c5d83be, []int{2}
}

func (m *StaticDeleteMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticDeleteMsg.Unmarshal(m, b)
}
func (m *StaticDeleteMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticDeleteMsg.Marshal(b, m, deterministic)
}
func (m *StaticDeleteMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticDeleteMsg.Merge(m, src)
}
func (m *StaticDeleteMsg) XXX_Size() int {
	return xxx_messageInfo_StaticDeleteMsg.Size(m)
}
func (m *StaticDeleteMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticDeleteMsg.DiscardUnknown(m)
}

var xxx_messageInfo_StaticDeleteMsg proto.InternalMessageInfo

func (m *StaticDeleteMsg) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StaticMsgList struct {
	Messages             []*StaticMsg `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StaticMsgList) Reset()         { *m = StaticMsgList{} }
func (m *StaticMsgList) String() string { return proto.CompactTextString(m) }
func (*StaticMsgList) ProtoMessage()    {}
func (*StaticMsgList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4ae77d93c5d83be, []int{3}
}

func (m *StaticMsgList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticMsgList.Unmarshal(m, b)
}
func (m *StaticMsgList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticMsgList.Marshal(b, m, deterministic)
}
func (m *StaticMsgList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticMsgList.Merge(m, src)
}
func (m *StaticMsgList) XXX_Size() int {
	return xxx_messageInfo_StaticMsgList.Size(m)
}
func (m *StaticMsgList) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticMsgList.DiscardUnknown(m)
}

var xxx_messageInfo_StaticMsgList proto.InternalMessageInfo

func (m *StaticMsgList) GetMessages() []*StaticMsg {
	if m != nil {
		return m.Messages
	}
	return nil
}

type StaticMsg struct {
	MsgType              string    `protobuf:"bytes,1,opt,name=msgType,proto3" json:"msgType,omitempty"`
	Id                   string    `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Position             *Position `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *StaticMsg) Reset()         { *m = StaticMsg{} }
func (m *StaticMsg) String() string { return proto.CompactTextString(m) }
func (*StaticMsg) ProtoMessage()    {}
func (*StaticMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4ae77d93c5d83be, []int{4}
}

func (m *StaticMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticMsg.Unmarshal(m, b)
}
func (m *StaticMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticMsg.Marshal(b, m, deterministic)
}
func (m *StaticMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticMsg.Merge(m, src)
}
func (m *StaticMsg) XXX_Size() int {
	return xxx_messageInfo_StaticMsg.Size(m)
}
func (m *StaticMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticMsg.DiscardUnknown(m)
}

var xxx_messageInfo_StaticMsg proto.InternalMessageInfo

func (m *StaticMsg) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *StaticMsg) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StaticMsg) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

type StaticAddAck struct {
	MsgType              string   `protobuf:"bytes,1,opt,name=msgType,proto3" json:"msgType,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	MsgDescription       string   `protobuf:"bytes,5,opt,name=msgDescription,proto3" json:"msgDescription,omitempty"`
	Success              bool     `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	FailMsg              string   `protobuf:"bytes,4,opt,name=failMsg,proto3" json:"failMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StaticAddAck) Reset()         { *m = StaticAddAck{} }
func (m *StaticAddAck) String() string { return proto.CompactTextString(m) }
func (*StaticAddAck) ProtoMessage()    {}
func (*StaticAddAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4ae77d93c5d83be, []int{5}
}

func (m *StaticAddAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticAddAck.Unmarshal(m, b)
}
func (m *StaticAddAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticAddAck.Marshal(b, m, deterministic)
}
func (m *StaticAddAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticAddAck.Merge(m, src)
}
func (m *StaticAddAck) XXX_Size() int {
	return xxx_messageInfo_StaticAddAck.Size(m)
}
func (m *StaticAddAck) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticAddAck.DiscardUnknown(m)
}

var xxx_messageInfo_StaticAddAck proto.InternalMessageInfo

func (m *StaticAddAck) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *StaticAddAck) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StaticAddAck) GetMsgDescription() string {
	if m != nil {
		return m.MsgDescription
	}
	return ""
}

func (m *StaticAddAck) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *StaticAddAck) GetFailMsg() string {
	if m != nil {
		return m.FailMsg
	}
	return ""
}

type StaticDeleteAck struct {
	MsgType              string   `protobuf:"bytes,1,opt,name=msgType,proto3" json:"msgType,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	MsgDescription       string   `protobuf:"bytes,5,opt,name=msgDescription,proto3" json:"msgDescription,omitempty"`
	Success              bool     `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	FailMsg              string   `protobuf:"bytes,4,opt,name=failMsg,proto3" json:"failMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StaticDeleteAck) Reset()         { *m = StaticDeleteAck{} }
func (m *StaticDeleteAck) String() string { return proto.CompactTextString(m) }
func (*StaticDeleteAck) ProtoMessage()    {}
func (*StaticDeleteAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4ae77d93c5d83be, []int{6}
}

func (m *StaticDeleteAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticDeleteAck.Unmarshal(m, b)
}
func (m *StaticDeleteAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticDeleteAck.Marshal(b, m, deterministic)
}
func (m *StaticDeleteAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticDeleteAck.Merge(m, src)
}
func (m *StaticDeleteAck) XXX_Size() int {
	return xxx_messageInfo_StaticDeleteAck.Size(m)
}
func (m *StaticDeleteAck) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticDeleteAck.DiscardUnknown(m)
}

var xxx_messageInfo_StaticDeleteAck proto.InternalMessageInfo

func (m *StaticDeleteAck) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *StaticDeleteAck) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StaticDeleteAck) GetMsgDescription() string {
	if m != nil {
		return m.MsgDescription
	}
	return ""
}

func (m *StaticDeleteAck) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *StaticDeleteAck) GetFailMsg() string {
	if m != nil {
		return m.FailMsg
	}
	return ""
}

type StaticDeleteAllAck struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	FailMsg              string   `protobuf:"bytes,2,opt,name=failMsg,proto3" json:"failMsg,omitempty"`
	Ids                  []string `protobuf:"bytes,3,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StaticDeleteAllAck) Reset()         { *m = StaticDeleteAllAck{} }
func (m *StaticDeleteAllAck) String() string { return proto.CompactTextString(m) }
func (*StaticDeleteAllAck) ProtoMessage()    {}
func (*StaticDeleteAllAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4ae77d93c5d83be, []int{7}
}

func (m *StaticDeleteAllAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticDeleteAllAck.Unmarshal(m, b)
}
func (m *StaticDeleteAllAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticDeleteAllAck.Marshal(b, m, deterministic)
}
func (m *StaticDeleteAllAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticDeleteAllAck.Merge(m, src)
}
func (m *StaticDeleteAllAck) XXX_Size() int {
	return xxx_messageInfo_StaticDeleteAllAck.Size(m)
}
func (m *StaticDeleteAllAck) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticDeleteAllAck.DiscardUnknown(m)
}

var xxx_messageInfo_StaticDeleteAllAck proto.InternalMessageInfo

func (m *StaticDeleteAllAck) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *StaticDeleteAllAck) GetFailMsg() string {
	if m != nil {
		return m.FailMsg
	}
	return ""
}

func (m *StaticDeleteAllAck) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type StaticStarted struct {
	MsgType              string   `protobuf:"bytes,1,opt,name=msgType,proto3" json:"msgType,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	MsgDescription       string   `protobuf:"bytes,3,opt,name=msgDescription,proto3" json:"msgDescription,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StaticStarted) Reset()         { *m = StaticStarted{} }
func (m *StaticStarted) String() string { return proto.CompactTextString(m) }
func (*StaticStarted) ProtoMessage()    {}
func (*StaticStarted) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4ae77d93c5d83be, []int{8}
}

func (m *StaticStarted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticStarted.Unmarshal(m, b)
}
func (m *StaticStarted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticStarted.Marshal(b, m, deterministic)
}
func (m *StaticStarted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticStarted.Merge(m, src)
}
func (m *StaticStarted) XXX_Size() int {
	return xxx_messageInfo_StaticStarted.Size(m)
}
func (m *StaticStarted) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticStarted.DiscardUnknown(m)
}

var xxx_messageInfo_StaticStarted proto.InternalMessageInfo

func (m *StaticStarted) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *StaticStarted) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StaticStarted) GetMsgDescription() string {
	if m != nil {
		return m.MsgDescription
	}
	return ""
}

type StaticStopped struct {
	MsgType              string   `protobuf:"bytes,1,opt,name=msgType,proto3" json:"msgType,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	MsgDescription       string   `protobuf:"bytes,3,opt,name=msgDescription,proto3" json:"msgDescription,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StaticStopped) Reset()         { *m = StaticStopped{} }
func (m *StaticStopped) String() string { return proto.CompactTextString(m) }
func (*StaticStopped) ProtoMessage()    {}
func (*StaticStopped) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4ae77d93c5d83be, []int{9}
}

func (m *StaticStopped) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticStopped.Unmarshal(m, b)
}
func (m *StaticStopped) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticStopped.Marshal(b, m, deterministic)
}
func (m *StaticStopped) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticStopped.Merge(m, src)
}
func (m *StaticStopped) XXX_Size() int {
	return xxx_messageInfo_StaticStopped.Size(m)
}
func (m *StaticStopped) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticStopped.DiscardUnknown(m)
}

var xxx_messageInfo_StaticStopped proto.InternalMessageInfo

func (m *StaticStopped) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *StaticStopped) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StaticStopped) GetMsgDescription() string {
	if m != nil {
		return m.MsgDescription
	}
	return ""
}

func init() {
	proto.RegisterType((*Position)(nil), "staticmsgpb.Position")
	proto.RegisterType((*StaticAddMsg)(nil), "staticmsgpb.StaticAddMsg")
	proto.RegisterType((*StaticDeleteMsg)(nil), "staticmsgpb.StaticDeleteMsg")
	proto.RegisterType((*StaticMsgList)(nil), "staticmsgpb.StaticMsgList")
	proto.RegisterType((*StaticMsg)(nil), "staticmsgpb.StaticMsg")
	proto.RegisterType((*StaticAddAck)(nil), "staticmsgpb.StaticAddAck")
	proto.RegisterType((*StaticDeleteAck)(nil), "staticmsgpb.StaticDeleteAck")
	proto.RegisterType((*StaticDeleteAllAck)(nil), "staticmsgpb.StaticDeleteAllAck")
	proto.RegisterType((*StaticStarted)(nil), "staticmsgpb.StaticStarted")
	proto.RegisterType((*StaticStopped)(nil), "staticmsgpb.StaticStopped")
}

func init() { proto.RegisterFile("static_msg.proto", fileDescriptor_c4ae77d93c5d83be) }

var fileDescriptor_c4ae77d93c5d83be = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x55, 0x12, 0xba, 0x4d, 0xa6, 0xcb, 0xb2, 0xb2, 0x04, 0x8a, 0x2a, 0x24, 0x4a, 0x0e, 0xa8,
	0xa7, 0xac, 0x28, 0x1c, 0xb8, 0x2e, 0xf4, 0x48, 0x25, 0xe4, 0xed, 0x89, 0x0b, 0x4a, 0x63, 0xaf,
	0xb1, 0x88, 0x6b, 0xab, 0xe3, 0x1e, 0xf6, 0x3f, 0x90, 0x10, 0x7f, 0x8b, 0x6c, 0xaf, 0xb3, 0x29,
	0x2a, 0x02, 0x84, 0x90, 0xf6, 0x96, 0x99, 0x79, 0x7e, 0x6f, 0xde, 0xb3, 0x03, 0xe7, 0x68, 0x1b,
	0x2b, 0xdb, 0x4f, 0x0a, 0x45, 0x6d, 0x76, 0xda, 0x6a, 0x32, 0x09, 0x1d, 0x85, 0xc2, 0x6c, 0xa6,
	0xcf, 0x84, 0xd6, 0xa2, 0xe3, 0x17, 0x7e, 0xb4, 0xd9, 0x5f, 0x5f, 0x58, 0xa9, 0x38, 0xda, 0x46,
	0x99, 0x80, 0xae, 0x96, 0x90, 0x7f, 0xd0, 0x28, 0xad, 0xd4, 0x5b, 0x32, 0x85, 0xbc, 0x6b, 0xac,
	0xb4, 0x7b, 0xc6, 0xcb, 0x64, 0x96, 0xcc, 0x47, 0xb4, 0xaf, 0xc9, 0x53, 0x28, 0x3a, 0xbd, 0x15,
	0x61, 0x98, 0xfa, 0xe1, 0x5d, 0xa3, 0xfa, 0x9e, 0xc2, 0xe9, 0x95, 0x97, 0xbd, 0x64, 0x6c, 0x85,
	0xc2, 0x51, 0x29, 0x14, 0x6f, 0x6f, 0x2c, 0x47, 0x4f, 0x75, 0x4a, 0xfb, 0x9a, 0x94, 0x30, 0x56,
	0x28, 0xd6, 0x37, 0x26, 0x10, 0x15, 0x34, 0x96, 0xe4, 0x0c, 0x52, 0xc9, 0xca, 0xcc, 0x37, 0x53,
	0xc9, 0xc8, 0x0c, 0x26, 0x8c, 0x63, 0xbb, 0x93, 0xc6, 0xed, 0x57, 0x8e, 0xfd, 0x60, 0xd8, 0x22,
	0x6f, 0xa0, 0x40, 0xdb, 0xec, 0xec, 0x5a, 0x2a, 0x5e, 0x3e, 0x98, 0x25, 0xf3, 0xc9, 0x62, 0x5a,
	0x07, 0xcf, 0x75, 0xf4, 0x5c, 0xaf, 0xa3, 0x67, 0x7a, 0x07, 0x26, 0xaf, 0x61, 0xcc, 0xb7, 0xcc,
	0x9f, 0x1b, 0xfd, 0xf6, 0x5c, 0x84, 0x92, 0x97, 0x90, 0x9b, 0xdb, 0xb8, 0xca, 0x13, 0x7f, 0xec,
	0x71, 0x3d, 0xc8, 0xbb, 0x8e, 0x59, 0xd2, 0x1e, 0x56, 0x3d, 0x87, 0x47, 0x21, 0x9a, 0x25, 0xef,
	0xb8, 0xe5, 0x2e, 0x9d, 0xe0, 0x33, 0x89, 0x3e, 0xab, 0x77, 0xf0, 0x30, 0x40, 0x56, 0x28, 0xde,
	0x4b, 0xb4, 0x64, 0x01, 0xb9, 0xe2, 0x88, 0x8d, 0xf0, 0xf1, 0x65, 0xf3, 0xc9, 0xe2, 0xc9, 0x81,
	0x4c, 0x8f, 0xa6, 0x3d, 0xae, 0xfa, 0x0c, 0x45, 0xdf, 0x1e, 0x66, 0x9c, 0x1c, 0xcb, 0x38, 0xed,
	0x33, 0x1e, 0x3a, 0xca, 0xfe, 0xcc, 0xd1, 0xd7, 0x64, 0x70, 0xdb, 0x97, 0xed, 0x97, 0xbf, 0x50,
	0x7b, 0x01, 0x67, 0x0a, 0xc5, 0x72, 0x70, 0xa9, 0x23, 0x3f, 0xfb, 0xa9, 0xeb, 0x18, 0x71, 0xdf,
	0xb6, 0x1c, 0xd1, 0x2f, 0x95, 0xd3, 0x58, 0xba, 0xc9, 0x75, 0x23, 0xbb, 0x15, 0x0a, 0x7f, 0xdf,
	0x05, 0x8d, 0x65, 0xf5, 0x2d, 0x39, 0x4c, 0xfa, 0xfe, 0x6c, 0xf6, 0x11, 0xc8, 0xc1, 0x62, 0x5d,
	0x77, 0xbb, 0x5b, 0x64, 0x4a, 0x7e, 0xc9, 0x94, 0x1e, 0x30, 0x91, 0x73, 0xc8, 0x24, 0x73, 0xca,
	0xd9, 0xbc, 0xa0, 0xee, 0xb3, 0x6a, 0xe2, 0xdb, 0xb9, 0x72, 0x4f, 0x9b, 0xb3, 0x7f, 0xb2, 0x9c,
	0x1d, 0xb3, 0x3c, 0x94, 0xd0, 0xc6, 0xfc, 0x0f, 0x89, 0xcd, 0x89, 0xff, 0xe9, 0x5e, 0xfd, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x4e, 0x69, 0xb4, 0xee, 0xcf, 0x04, 0x00, 0x00,
}
                                                                                                                                                                                                                                                                                                                                                                                       static_msg.proto                                                                                    0000664 0001751 0001751 00000005611 14335525351 014207  0                                                                                                    ustar   tfisher                         tfisher                                                                                                                                                                                                                syntax = "proto3";


import "google/protobuf/timestamp.proto";

package staticmsgpb;

message Position {

  int32 latitude = 1; // in 1/10th microdegrees

  int32 longitude = 2; // in 1/10th microdegrees

}


// this is the message a client should publish to add a new Static Message

// into the system

message StaticAddMsg {

  bytes msgBytes = 1; // the actual content of the message

  string msgType = 2; // such as "TIM", "RSA", etc

  string id = 3; // must just be unique for the publishing client for this msgType

 string description = 7;  // can be empty


  google.protobuf.Timestamp startTime = 4; // just set to current time if it???s not

                                           // in the future

  google.protobuf.Timestamp endTime = 5; //

  Position position = 6; // approximate central geoposition

}



// this is the message a client should publish to remove a Static Message

// that it added

// the reason this is a Protobuf message at all (given that it has one field)

// is in case we need to expand it in the future, this gives us flexibility

message StaticDeleteMsg {

  string id = 1;

}


// this is the list of static messages that a subscribing client has access to

message StaticMsgList {

  repeated StaticMsg messages = 1;

}


message StaticMsg {

  string msgType = 1; // such as "TIM", "RSA", etc

  string id = 2; // must just be unique for the publishing client for the given msgType  


  Position position = 3;

}


// acknowledgment for the StaticAddMsg

message StaticAddAck {

  string msgType = 1; // such as "TIM", "RSA", etc

  string id = 2;

  string msgDescription = 5; // matches ???description??? from original StaticAddMsg

  bool success = 3;

  string failMsg = 4; // only set in case of failure

}


// acknowledgment for the StaticDeleteMsg

message StaticDeleteAck {

  string msgType = 1; // such as "TIM", "RSA", etc

  string id = 2;

  string msgDescription = 5; // matches ???description??? from original StaticAddMsg

  bool success = 3;

  string failMsg = 4; // only set in case of failure

}


// acknowledgment for the request to delete all messages for the client (which has no associated Protobuf)

message StaticDeleteAllAck {

  bool success = 1;

  string failMsg = 2; // only set in case of failure

  repeated string ids = 3;

}


// response to client that a message started (applies both to messages that were scheduled and to messages that were started immediately)

message StaticStarted {

  string msgType = 1; // such as "TIM", "RSA", etc

  string id = 2;

  string msgDescription = 3; // matches ???description??? from original StaticAddMsg

}


// response to client that a message stopped (applies both to messages that expired and to messages that were stopped immediately)

message StaticStopped {

  string msgType = 1; // such as "TIM", "RSA", etc

  string id = 2;

  string msgDescription = 3; // matches ???description??? from original StaticAddMsg

}
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       