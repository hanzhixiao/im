// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: pb_msg/msg.proto

package pb_msg

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Message_SendChatMessage_FullMethodName  = "/pb_msg.Message/SendChatMessage"
	Message_MessageOperation_FullMethodName = "/pb_msg.Message/MessageOperation"
)

// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageClient interface {
	SendChatMessage(ctx context.Context, in *SendChatMessageReq, opts ...grpc.CallOption) (*SendChatMessageResp, error)
	MessageOperation(ctx context.Context, in *MessageOperationReq, opts ...grpc.CallOption) (*MessageOperationResp, error)
}

type messageClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageClient(cc grpc.ClientConnInterface) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) SendChatMessage(ctx context.Context, in *SendChatMessageReq, opts ...grpc.CallOption) (*SendChatMessageResp, error) {
	out := new(SendChatMessageResp)
	err := c.cc.Invoke(ctx, Message_SendChatMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) MessageOperation(ctx context.Context, in *MessageOperationReq, opts ...grpc.CallOption) (*MessageOperationResp, error) {
	out := new(MessageOperationResp)
	err := c.cc.Invoke(ctx, Message_MessageOperation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServer is the server API for Message service.
// All implementations must embed UnimplementedMessageServer
// for forward compatibility
type MessageServer interface {
	SendChatMessage(context.Context, *SendChatMessageReq) (*SendChatMessageResp, error)
	MessageOperation(context.Context, *MessageOperationReq) (*MessageOperationResp, error)
	mustEmbedUnimplementedMessageServer()
}

// UnimplementedMessageServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServer struct {
}

func (UnimplementedMessageServer) SendChatMessage(context.Context, *SendChatMessageReq) (*SendChatMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendChatMessage not implemented")
}
func (UnimplementedMessageServer) MessageOperation(context.Context, *MessageOperationReq) (*MessageOperationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageOperation not implemented")
}
func (UnimplementedMessageServer) mustEmbedUnimplementedMessageServer() {}

// UnsafeMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServer will
// result in compilation errors.
type UnsafeMessageServer interface {
	mustEmbedUnimplementedMessageServer()
}

func RegisterMessageServer(s grpc.ServiceRegistrar, srv MessageServer) {
	s.RegisterService(&Message_ServiceDesc, srv)
}

func _Message_SendChatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendChatMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SendChatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_SendChatMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SendChatMessage(ctx, req.(*SendChatMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_MessageOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageOperationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).MessageOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_MessageOperation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).MessageOperation(ctx, req.(*MessageOperationReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Message_ServiceDesc is the grpc.ServiceDesc for Message service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Message_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb_msg.Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendChatMessage",
			Handler:    _Message_SendChatMessage_Handler,
		},
		{
			MethodName: "MessageOperation",
			Handler:    _Message_MessageOperation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_msg/msg.proto",
}
