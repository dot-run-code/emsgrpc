// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// TopicClient is the client API for Topic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TopicClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	PublishStream(ctx context.Context, opts ...grpc.CallOption) (Topic_PublishStreamClient, error)
	StreamFile(ctx context.Context, opts ...grpc.CallOption) (Topic_StreamFileClient, error)
}

type topicClient struct {
	cc grpc.ClientConnInterface
}

func NewTopicClient(cc grpc.ClientConnInterface) TopicClient {
	return &topicClient{cc}
}

func (c *topicClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, "/proto.Topic/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicClient) PublishStream(ctx context.Context, opts ...grpc.CallOption) (Topic_PublishStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Topic_ServiceDesc.Streams[0], "/proto.Topic/PublishStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &topicPublishStreamClient{stream}
	return x, nil
}

type Topic_PublishStreamClient interface {
	Send(*PublishStreamRequest) error
	CloseAndRecv() (*AcknowledgementResponse, error)
	grpc.ClientStream
}

type topicPublishStreamClient struct {
	grpc.ClientStream
}

func (x *topicPublishStreamClient) Send(m *PublishStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *topicPublishStreamClient) CloseAndRecv() (*AcknowledgementResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AcknowledgementResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *topicClient) StreamFile(ctx context.Context, opts ...grpc.CallOption) (Topic_StreamFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &Topic_ServiceDesc.Streams[1], "/proto.Topic/StreamFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &topicStreamFileClient{stream}
	return x, nil
}

type Topic_StreamFileClient interface {
	Send(*PublishStreamRequest) error
	CloseAndRecv() (*AcknowledgementResponse, error)
	grpc.ClientStream
}

type topicStreamFileClient struct {
	grpc.ClientStream
}

func (x *topicStreamFileClient) Send(m *PublishStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *topicStreamFileClient) CloseAndRecv() (*AcknowledgementResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AcknowledgementResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TopicServer is the server API for Topic service.
// All implementations should embed UnimplementedTopicServer
// for forward compatibility
type TopicServer interface {
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	PublishStream(Topic_PublishStreamServer) error
	StreamFile(Topic_StreamFileServer) error
}

// UnimplementedTopicServer should be embedded to have forward compatible implementations.
type UnimplementedTopicServer struct {
}

func (UnimplementedTopicServer) Publish(context.Context, *PublishRequest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedTopicServer) PublishStream(Topic_PublishStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method PublishStream not implemented")
}
func (UnimplementedTopicServer) StreamFile(Topic_StreamFileServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamFile not implemented")
}

// UnsafeTopicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TopicServer will
// result in compilation errors.
type UnsafeTopicServer interface {
	mustEmbedUnimplementedTopicServer()
}

func RegisterTopicServer(s grpc.ServiceRegistrar, srv TopicServer) {
	s.RegisterService(&Topic_ServiceDesc, srv)
}

func _Topic_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Topic/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Topic_PublishStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TopicServer).PublishStream(&topicPublishStreamServer{stream})
}

type Topic_PublishStreamServer interface {
	SendAndClose(*AcknowledgementResponse) error
	Recv() (*PublishStreamRequest, error)
	grpc.ServerStream
}

type topicPublishStreamServer struct {
	grpc.ServerStream
}

func (x *topicPublishStreamServer) SendAndClose(m *AcknowledgementResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *topicPublishStreamServer) Recv() (*PublishStreamRequest, error) {
	m := new(PublishStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Topic_StreamFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TopicServer).StreamFile(&topicStreamFileServer{stream})
}

type Topic_StreamFileServer interface {
	SendAndClose(*AcknowledgementResponse) error
	Recv() (*PublishStreamRequest, error)
	grpc.ServerStream
}

type topicStreamFileServer struct {
	grpc.ServerStream
}

func (x *topicStreamFileServer) SendAndClose(m *AcknowledgementResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *topicStreamFileServer) Recv() (*PublishStreamRequest, error) {
	m := new(PublishStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Topic_ServiceDesc is the grpc.ServiceDesc for Topic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Topic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Topic",
	HandlerType: (*TopicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Topic_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PublishStream",
			Handler:       _Topic_PublishStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamFile",
			Handler:       _Topic_StreamFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "topic.proto",
}