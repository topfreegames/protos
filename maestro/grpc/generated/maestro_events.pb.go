// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: maestro_events.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RoomStatus_RoomStatusType int32

const (
	RoomStatus_ready       RoomStatus_RoomStatusType = 0
	RoomStatus_occupied    RoomStatus_RoomStatusType = 1
	RoomStatus_terminating RoomStatus_RoomStatusType = 2
	RoomStatus_terminated  RoomStatus_RoomStatusType = 3
)

// Enum value maps for RoomStatus_RoomStatusType.
var (
	RoomStatus_RoomStatusType_name = map[int32]string{
		0: "ready",
		1: "occupied",
		2: "terminating",
		3: "terminated",
	}
	RoomStatus_RoomStatusType_value = map[string]int32{
		"ready":       0,
		"occupied":    1,
		"terminating": 2,
		"terminated":  3,
	}
)

func (x RoomStatus_RoomStatusType) Enum() *RoomStatus_RoomStatusType {
	p := new(RoomStatus_RoomStatusType)
	*p = x
	return p
}

func (x RoomStatus_RoomStatusType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoomStatus_RoomStatusType) Descriptor() protoreflect.EnumDescriptor {
	return file_maestro_events_proto_enumTypes[0].Descriptor()
}

func (RoomStatus_RoomStatusType) Type() protoreflect.EnumType {
	return &file_maestro_events_proto_enumTypes[0]
}

func (x RoomStatus_RoomStatusType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoomStatus_RoomStatusType.Descriptor instead.
func (RoomStatus_RoomStatusType) EnumDescriptor() ([]byte, []int) {
	return file_maestro_events_proto_rawDescGZIP(), []int{1, 0}
}

type PlayerEvent_PlayerEventType int32

const (
	PlayerEvent_PLAYER_JOINED PlayerEvent_PlayerEventType = 0
	PlayerEvent_PLAYER_LEFT   PlayerEvent_PlayerEventType = 1
)

// Enum value maps for PlayerEvent_PlayerEventType.
var (
	PlayerEvent_PlayerEventType_name = map[int32]string{
		0: "PLAYER_JOINED",
		1: "PLAYER_LEFT",
	}
	PlayerEvent_PlayerEventType_value = map[string]int32{
		"PLAYER_JOINED": 0,
		"PLAYER_LEFT":   1,
	}
)

func (x PlayerEvent_PlayerEventType) Enum() *PlayerEvent_PlayerEventType {
	p := new(PlayerEvent_PlayerEventType)
	*p = x
	return p
}

func (x PlayerEvent_PlayerEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerEvent_PlayerEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_maestro_events_proto_enumTypes[1].Descriptor()
}

func (PlayerEvent_PlayerEventType) Type() protoreflect.EnumType {
	return &file_maestro_events_proto_enumTypes[1]
}

func (x PlayerEvent_PlayerEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerEvent_PlayerEventType.Descriptor instead.
func (PlayerEvent_PlayerEventType) EnumDescriptor() ([]byte, []int) {
	return file_maestro_events_proto_rawDescGZIP(), []int{3, 0}
}

type RoomEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room      *Room             `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	EventType string            `protobuf:"bytes,2,opt,name=eventType,proto3" json:"eventType,omitempty"`
	Metadata  map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RoomEvent) Reset() {
	*x = RoomEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maestro_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomEvent) ProtoMessage() {}

func (x *RoomEvent) ProtoReflect() protoreflect.Message {
	mi := &file_maestro_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomEvent.ProtoReflect.Descriptor instead.
func (*RoomEvent) Descriptor() ([]byte, []int) {
	return file_maestro_events_proto_rawDescGZIP(), []int{0}
}

func (x *RoomEvent) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *RoomEvent) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *RoomEvent) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type RoomStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room       *Room                     `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	StatusType RoomStatus_RoomStatusType `protobuf:"varint,2,opt,name=statusType,proto3,enum=eventforwarder.RoomStatus_RoomStatusType" json:"statusType,omitempty"`
}

func (x *RoomStatus) Reset() {
	*x = RoomStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maestro_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomStatus) ProtoMessage() {}

func (x *RoomStatus) ProtoReflect() protoreflect.Message {
	mi := &file_maestro_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomStatus.ProtoReflect.Descriptor instead.
func (*RoomStatus) Descriptor() ([]byte, []int) {
	return file_maestro_events_proto_rawDescGZIP(), []int{1}
}

func (x *RoomStatus) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *RoomStatus) GetStatusType() RoomStatus_RoomStatusType {
	if x != nil {
		return x.StatusType
	}
	return RoomStatus_ready
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Game     string            `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
	RoomId   string            `protobuf:"bytes,2,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Host     string            `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	Port     int32             `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	Metadata map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RoomType string            `protobuf:"bytes,7,opt,name=roomType,proto3" json:"roomType,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maestro_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_maestro_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_maestro_events_proto_rawDescGZIP(), []int{2}
}

func (x *Room) GetGame() string {
	if x != nil {
		return x.Game
	}
	return ""
}

func (x *Room) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *Room) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Room) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Room) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Room) GetRoomType() string {
	if x != nil {
		return x.RoomType
	}
	return ""
}

type PlayerEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId  string                      `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Room      *Room                       `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	EventType PlayerEvent_PlayerEventType `protobuf:"varint,3,opt,name=eventType,proto3,enum=eventforwarder.PlayerEvent_PlayerEventType" json:"eventType,omitempty"`
	Metadata  map[string]string           `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PlayerEvent) Reset() {
	*x = PlayerEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maestro_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerEvent) ProtoMessage() {}

func (x *PlayerEvent) ProtoReflect() protoreflect.Message {
	mi := &file_maestro_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerEvent.ProtoReflect.Descriptor instead.
func (*PlayerEvent) Descriptor() ([]byte, []int) {
	return file_maestro_events_proto_rawDescGZIP(), []int{3}
}

func (x *PlayerEvent) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerEvent) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *PlayerEvent) GetEventType() PlayerEvent_PlayerEventType {
	if x != nil {
		return x.EventType
	}
	return PlayerEvent_PLAYER_JOINED
}

func (x *PlayerEvent) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type RoomInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomType               string            `protobuf:"bytes,1,opt,name=roomType,proto3" json:"roomType,omitempty"`
	Game                   string            `protobuf:"bytes,2,opt,name=game,proto3" json:"game,omitempty"`
	NumberOfTeams          int32             `protobuf:"varint,3,opt,name=numberOfTeams,proto3" json:"numberOfTeams,omitempty"`
	PlayersPerTeam         int32             `protobuf:"varint,4,opt,name=playersPerTeam,proto3" json:"playersPerTeam,omitempty"`
	MinimumNumberOfPlayers int32             `protobuf:"varint,5,opt,name=minimumNumberOfPlayers,proto3" json:"minimumNumberOfPlayers,omitempty"`
	MatchmakingScript      string            `protobuf:"bytes,6,opt,name=matchmakingScript,proto3" json:"matchmakingScript,omitempty"`
	WebhookUrl             string            `protobuf:"bytes,7,opt,name=webhookUrl,proto3" json:"webhookUrl,omitempty"`
	Metadata               map[string]string `protobuf:"bytes,8,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tags                   map[string]string `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RoomInfo) Reset() {
	*x = RoomInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maestro_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomInfo) ProtoMessage() {}

func (x *RoomInfo) ProtoReflect() protoreflect.Message {
	mi := &file_maestro_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomInfo.ProtoReflect.Descriptor instead.
func (*RoomInfo) Descriptor() ([]byte, []int) {
	return file_maestro_events_proto_rawDescGZIP(), []int{4}
}

func (x *RoomInfo) GetRoomType() string {
	if x != nil {
		return x.RoomType
	}
	return ""
}

func (x *RoomInfo) GetGame() string {
	if x != nil {
		return x.Game
	}
	return ""
}

func (x *RoomInfo) GetNumberOfTeams() int32 {
	if x != nil {
		return x.NumberOfTeams
	}
	return 0
}

func (x *RoomInfo) GetPlayersPerTeam() int32 {
	if x != nil {
		return x.PlayersPerTeam
	}
	return 0
}

func (x *RoomInfo) GetMinimumNumberOfPlayers() int32 {
	if x != nil {
		return x.MinimumNumberOfPlayers
	}
	return 0
}

func (x *RoomInfo) GetMatchmakingScript() string {
	if x != nil {
		return x.MatchmakingScript
	}
	return ""
}

func (x *RoomInfo) GetWebhookUrl() string {
	if x != nil {
		return x.WebhookUrl
	}
	return ""
}

func (x *RoomInfo) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *RoomInfo) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maestro_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_maestro_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_maestro_events_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_maestro_events_proto protoreflect.FileDescriptor

var file_maestro_events_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x61, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x22, 0xd5, 0x01, 0x0a, 0x09, 0x52, 0x6f, 0x6f, 0x6d, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x1c,
	0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x43, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xcd,
	0x01, 0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a,
	0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x49, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x4a, 0x0a, 0x0e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x69, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x0e,
	0x0a, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x10, 0x03, 0x22, 0xf3,
	0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3e, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x52,
	0x6f, 0x6f, 0x6d, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xd9, 0x02, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x49, 0x0a, 0x09, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x35, 0x0a, 0x0f, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d,
	0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x01,
	0x22, 0x80, 0x04, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x54, 0x65,
	0x61, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x50, 0x65,
	0x72, 0x54, 0x65, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x50, 0x65, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x36, 0x0a, 0x16, 0x6d,
	0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x6d, 0x69, 0x6e,
	0x69, 0x6d, 0x75, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x72,
	0x6c, 0x12, 0x42, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x61,
	0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x3b, 0x0a,
	0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61,
	0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x38, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xbf, 0x03,
	0x0a, 0x0d, 0x47, 0x52, 0x50, 0x43, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x44, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x18, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x18, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x6f, 0x6f,
	0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x1a, 0x18, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x40, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x69, 0x6e, 0x67, 0x12,
	0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x1a, 0x18, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x48, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73,
	0x79, 0x6e, 0x63, 0x12, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a,
	0x18, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_maestro_events_proto_rawDescOnce sync.Once
	file_maestro_events_proto_rawDescData = file_maestro_events_proto_rawDesc
)

func file_maestro_events_proto_rawDescGZIP() []byte {
	file_maestro_events_proto_rawDescOnce.Do(func() {
		file_maestro_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_maestro_events_proto_rawDescData)
	})
	return file_maestro_events_proto_rawDescData
}

var file_maestro_events_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_maestro_events_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_maestro_events_proto_goTypes = []interface{}{
	(RoomStatus_RoomStatusType)(0),   // 0: eventforwarder.RoomStatus.RoomStatusType
	(PlayerEvent_PlayerEventType)(0), // 1: eventforwarder.PlayerEvent.PlayerEventType
	(*RoomEvent)(nil),                // 2: eventforwarder.RoomEvent
	(*RoomStatus)(nil),               // 3: eventforwarder.RoomStatus
	(*Room)(nil),                     // 4: eventforwarder.Room
	(*PlayerEvent)(nil),              // 5: eventforwarder.PlayerEvent
	(*RoomInfo)(nil),                 // 6: eventforwarder.RoomInfo
	(*Response)(nil),                 // 7: eventforwarder.Response
	nil,                              // 8: eventforwarder.RoomEvent.MetadataEntry
	nil,                              // 9: eventforwarder.Room.MetadataEntry
	nil,                              // 10: eventforwarder.PlayerEvent.MetadataEntry
	nil,                              // 11: eventforwarder.RoomInfo.MetadataEntry
	nil,                              // 12: eventforwarder.RoomInfo.TagsEntry
}
var file_maestro_events_proto_depIdxs = []int32{
	4,  // 0: eventforwarder.RoomEvent.room:type_name -> eventforwarder.Room
	8,  // 1: eventforwarder.RoomEvent.metadata:type_name -> eventforwarder.RoomEvent.MetadataEntry
	4,  // 2: eventforwarder.RoomStatus.room:type_name -> eventforwarder.Room
	0,  // 3: eventforwarder.RoomStatus.statusType:type_name -> eventforwarder.RoomStatus.RoomStatusType
	9,  // 4: eventforwarder.Room.metadata:type_name -> eventforwarder.Room.MetadataEntry
	4,  // 5: eventforwarder.PlayerEvent.room:type_name -> eventforwarder.Room
	1,  // 6: eventforwarder.PlayerEvent.eventType:type_name -> eventforwarder.PlayerEvent.PlayerEventType
	10, // 7: eventforwarder.PlayerEvent.metadata:type_name -> eventforwarder.PlayerEvent.MetadataEntry
	11, // 8: eventforwarder.RoomInfo.metadata:type_name -> eventforwarder.RoomInfo.MetadataEntry
	12, // 9: eventforwarder.RoomInfo.tags:type_name -> eventforwarder.RoomInfo.TagsEntry
	6,  // 10: eventforwarder.GRPCForwarder.SendRoomInfo:input_type -> eventforwarder.RoomInfo
	3,  // 11: eventforwarder.GRPCForwarder.SendRoomStatus:input_type -> eventforwarder.RoomStatus
	4,  // 12: eventforwarder.GRPCForwarder.SendRoomPing:input_type -> eventforwarder.Room
	3,  // 13: eventforwarder.GRPCForwarder.SendRoomResync:input_type -> eventforwarder.RoomStatus
	2,  // 14: eventforwarder.GRPCForwarder.SendRoomEvent:input_type -> eventforwarder.RoomEvent
	5,  // 15: eventforwarder.GRPCForwarder.SendPlayerEvent:input_type -> eventforwarder.PlayerEvent
	7,  // 16: eventforwarder.GRPCForwarder.SendRoomInfo:output_type -> eventforwarder.Response
	7,  // 17: eventforwarder.GRPCForwarder.SendRoomStatus:output_type -> eventforwarder.Response
	7,  // 18: eventforwarder.GRPCForwarder.SendRoomPing:output_type -> eventforwarder.Response
	7,  // 19: eventforwarder.GRPCForwarder.SendRoomResync:output_type -> eventforwarder.Response
	7,  // 20: eventforwarder.GRPCForwarder.SendRoomEvent:output_type -> eventforwarder.Response
	7,  // 21: eventforwarder.GRPCForwarder.SendPlayerEvent:output_type -> eventforwarder.Response
	16, // [16:22] is the sub-list for method output_type
	10, // [10:16] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_maestro_events_proto_init() }
func file_maestro_events_proto_init() {
	if File_maestro_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_maestro_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_maestro_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_maestro_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_maestro_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_maestro_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_maestro_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_maestro_events_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_maestro_events_proto_goTypes,
		DependencyIndexes: file_maestro_events_proto_depIdxs,
		EnumInfos:         file_maestro_events_proto_enumTypes,
		MessageInfos:      file_maestro_events_proto_msgTypes,
	}.Build()
	File_maestro_events_proto = out.File
	file_maestro_events_proto_rawDesc = nil
	file_maestro_events_proto_goTypes = nil
	file_maestro_events_proto_depIdxs = nil
}
