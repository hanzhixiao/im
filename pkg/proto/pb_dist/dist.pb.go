// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v4.22.2
// source: pb_dist/dist.proto

package pb_dist

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "im/pkg/proto/pb_chat"
	pb_enum "im/pkg/proto/pb_enum"
	_ "im/pkg/proto/pb_invite"
	pb_msg "im/pkg/proto/pb_msg"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DistMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic    pb_enum.TOPIC          `protobuf:"varint,1,opt,name=topic,proto3,enum=pb_enum.TOPIC" json:"topic,omitempty"`
	SubTopic pb_enum.SUB_TOPIC      `protobuf:"varint,2,opt,name=sub_topic,json=subTopic,proto3,enum=pb_enum.SUB_TOPIC" json:"sub_topic,omitempty"`
	Msg      *pb_msg.SrvChatMessage `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DistMessageReq) Reset() {
	*x = DistMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_dist_dist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistMessageReq) ProtoMessage() {}

func (x *DistMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_dist_dist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistMessageReq.ProtoReflect.Descriptor instead.
func (*DistMessageReq) Descriptor() ([]byte, []int) {
	return file_pb_dist_dist_proto_rawDescGZIP(), []int{0}
}

func (x *DistMessageReq) GetTopic() pb_enum.TOPIC {
	if x != nil {
		return x.Topic
	}
	return pb_enum.TOPIC(0)
}

func (x *DistMessageReq) GetSubTopic() pb_enum.SUB_TOPIC {
	if x != nil {
		return x.SubTopic
	}
	return pb_enum.SUB_TOPIC(0)
}

func (x *DistMessageReq) GetMsg() *pb_msg.SrvChatMessage {
	if x != nil {
		return x.Msg
	}
	return nil
}

type DistMessageResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DistMessageResp) Reset() {
	*x = DistMessageResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_dist_dist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistMessageResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistMessageResp) ProtoMessage() {}

func (x *DistMessageResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_dist_dist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistMessageResp.ProtoReflect.Descriptor instead.
func (*DistMessageResp) Descriptor() ([]byte, []int) {
	return file_pb_dist_dist_proto_rawDescGZIP(), []int{1}
}

func (x *DistMessageResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DistMessageResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *DistMessageResp) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ChatInviteNotificationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId       int64                     `protobuf:"varint,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`                                              // 发送者uid
	SenderPlatform pb_enum.PLATFORM_TYPE     `protobuf:"varint,2,opt,name=sender_platform,json=senderPlatform,proto3,enum=pb_enum.PLATFORM_TYPE" json:"sender_platform,omitempty"` // 发送者平台
	Notifications  []*ChatInviteNotification `protobuf:"bytes,3,rep,name=notifications,proto3" json:"notifications,omitempty"`
}

func (x *ChatInviteNotificationReq) Reset() {
	*x = ChatInviteNotificationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_dist_dist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatInviteNotificationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatInviteNotificationReq) ProtoMessage() {}

func (x *ChatInviteNotificationReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_dist_dist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatInviteNotificationReq.ProtoReflect.Descriptor instead.
func (*ChatInviteNotificationReq) Descriptor() ([]byte, []int) {
	return file_pb_dist_dist_proto_rawDescGZIP(), []int{2}
}

func (x *ChatInviteNotificationReq) GetSenderId() int64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *ChatInviteNotificationReq) GetSenderPlatform() pb_enum.PLATFORM_TYPE {
	if x != nil {
		return x.SenderPlatform
	}
	return pb_enum.PLATFORM_TYPE(0)
}

func (x *ChatInviteNotificationReq) GetNotifications() []*ChatInviteNotification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

type ChatInviteNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReceiverId       int64              `protobuf:"varint,1,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
	ReceiverServerId int64              `protobuf:"varint,2,opt,name=receiver_server_id,json=receiverServerId,proto3" json:"receiver_server_id,omitempty"`
	Invite           *pb_msg.ChatInvite `protobuf:"bytes,3,opt,name=invite,proto3" json:"invite,omitempty"`
}

func (x *ChatInviteNotification) Reset() {
	*x = ChatInviteNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_dist_dist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatInviteNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatInviteNotification) ProtoMessage() {}

func (x *ChatInviteNotification) ProtoReflect() protoreflect.Message {
	mi := &file_pb_dist_dist_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatInviteNotification.ProtoReflect.Descriptor instead.
func (*ChatInviteNotification) Descriptor() ([]byte, []int) {
	return file_pb_dist_dist_proto_rawDescGZIP(), []int{3}
}

func (x *ChatInviteNotification) GetReceiverId() int64 {
	if x != nil {
		return x.ReceiverId
	}
	return 0
}

func (x *ChatInviteNotification) GetReceiverServerId() int64 {
	if x != nil {
		return x.ReceiverServerId
	}
	return 0
}

func (x *ChatInviteNotification) GetInvite() *pb_msg.ChatInvite {
	if x != nil {
		return x.Invite
	}
	return nil
}

type ChatInviteNotificationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ChatInviteNotificationResp) Reset() {
	*x = ChatInviteNotificationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_dist_dist_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatInviteNotificationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatInviteNotificationResp) ProtoMessage() {}

func (x *ChatInviteNotificationResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_dist_dist_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatInviteNotificationResp.ProtoReflect.Descriptor instead.
func (*ChatInviteNotificationResp) Descriptor() ([]byte, []int) {
	return file_pb_dist_dist_proto_rawDescGZIP(), []int{4}
}

func (x *ChatInviteNotificationResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ChatInviteNotificationResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_pb_dist_dist_proto protoreflect.FileDescriptor

var file_pb_dist_dist_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x62, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x1a, 0x12, 0x70,
	0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x10, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x62, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x62, 0x5f, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x91, 0x01, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x54, 0x4f, 0x50, 0x49,
	0x43, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x2f, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x62,
	0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x55, 0x42, 0x5f, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x52,
	0x08, 0x73, 0x75, 0x62, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x28, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2e,
	0x53, 0x72, 0x76, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x4b, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0xc0, 0x01, 0x0a, 0x19, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0f, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x50,
	0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x0e, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x45, 0x0a, 0x0d,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x62, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x2c, 0x0a, 0x12, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x06, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x62, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x52, 0x06, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x22, 0x42, 0x0a, 0x1a, 0x43, 0x68, 0x61,
	0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xab, 0x01,
	0x0a, 0x04, 0x44, 0x69, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x2e,
	0x44, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x18,
	0x2e, 0x70, 0x62, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x61, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x74,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x70, 0x62, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x70, 0x62, 0x5f, 0x64, 0x69, 0x73, 0x74,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x13, 0x5a, 0x11, 0x2e,
	0x2f, 0x70, 0x62, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x3b, 0x70, 0x62, 0x5f, 0x64, 0x69, 0x73, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_dist_dist_proto_rawDescOnce sync.Once
	file_pb_dist_dist_proto_rawDescData = file_pb_dist_dist_proto_rawDesc
)

func file_pb_dist_dist_proto_rawDescGZIP() []byte {
	file_pb_dist_dist_proto_rawDescOnce.Do(func() {
		file_pb_dist_dist_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_dist_dist_proto_rawDescData)
	})
	return file_pb_dist_dist_proto_rawDescData
}

var file_pb_dist_dist_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pb_dist_dist_proto_goTypes = []interface{}{
	(*DistMessageReq)(nil),             // 0: pb_dist.DistMessageReq
	(*DistMessageResp)(nil),            // 1: pb_dist.DistMessageResp
	(*ChatInviteNotificationReq)(nil),  // 2: pb_dist.ChatInviteNotificationReq
	(*ChatInviteNotification)(nil),     // 3: pb_dist.ChatInviteNotification
	(*ChatInviteNotificationResp)(nil), // 4: pb_dist.ChatInviteNotificationResp
	(pb_enum.TOPIC)(0),                 // 5: pb_enum.TOPIC
	(pb_enum.SUB_TOPIC)(0),             // 6: pb_enum.SUB_TOPIC
	(*pb_msg.SrvChatMessage)(nil),      // 7: pb_msg.SrvChatMessage
	(pb_enum.PLATFORM_TYPE)(0),         // 8: pb_enum.PLATFORM_TYPE
	(*pb_msg.ChatInvite)(nil),          // 9: pb_msg.ChatInvite
}
var file_pb_dist_dist_proto_depIdxs = []int32{
	5, // 0: pb_dist.DistMessageReq.topic:type_name -> pb_enum.TOPIC
	6, // 1: pb_dist.DistMessageReq.sub_topic:type_name -> pb_enum.SUB_TOPIC
	7, // 2: pb_dist.DistMessageReq.msg:type_name -> pb_msg.SrvChatMessage
	8, // 3: pb_dist.ChatInviteNotificationReq.sender_platform:type_name -> pb_enum.PLATFORM_TYPE
	3, // 4: pb_dist.ChatInviteNotificationReq.notifications:type_name -> pb_dist.ChatInviteNotification
	9, // 5: pb_dist.ChatInviteNotification.invite:type_name -> pb_msg.ChatInvite
	0, // 6: pb_dist.Dist.DistMessage:input_type -> pb_dist.DistMessageReq
	2, // 7: pb_dist.Dist.ChatInviteNotification:input_type -> pb_dist.ChatInviteNotificationReq
	1, // 8: pb_dist.Dist.DistMessage:output_type -> pb_dist.DistMessageResp
	4, // 9: pb_dist.Dist.ChatInviteNotification:output_type -> pb_dist.ChatInviteNotificationResp
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pb_dist_dist_proto_init() }
func file_pb_dist_dist_proto_init() {
	if File_pb_dist_dist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_dist_dist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistMessageReq); i {
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
		file_pb_dist_dist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistMessageResp); i {
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
		file_pb_dist_dist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatInviteNotificationReq); i {
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
		file_pb_dist_dist_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatInviteNotification); i {
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
		file_pb_dist_dist_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatInviteNotificationResp); i {
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
			RawDescriptor: file_pb_dist_dist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_dist_dist_proto_goTypes,
		DependencyIndexes: file_pb_dist_dist_proto_depIdxs,
		MessageInfos:      file_pb_dist_dist_proto_msgTypes,
	}.Build()
	File_pb_dist_dist_proto = out.File
	file_pb_dist_dist_proto_rawDesc = nil
	file_pb_dist_dist_proto_goTypes = nil
	file_pb_dist_dist_proto_depIdxs = nil
}
