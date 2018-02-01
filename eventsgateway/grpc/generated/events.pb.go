// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eventsgateway/grpc/protobuf/events.proto

/*
Package eventforwarder is a generated protocol buffer package.

It is generated from these files:
	eventsgateway/grpc/protobuf/events.proto

It has these top-level messages:
	Event
	Response
*/
package eventforwarder

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Event struct {
	Id        string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name      string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Topic     string            `protobuf:"bytes,3,opt,name=topic" json:"topic,omitempty"`
	Props     map[string]string `protobuf:"bytes,4,rep,name=props" json:"props,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Timestamp int64             `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Event) GetProps() map[string]string {
	if m != nil {
		return m.Props
	}
	return nil
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Response struct {
	Code    int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "eventforwarder.Event")
	proto.RegisterType((*Response)(nil), "eventforwarder.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GRPCForwarder service

type GRPCForwarderClient interface {
	SendEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error)
}

type gRPCForwarderClient struct {
	cc *grpc.ClientConn
}

func NewGRPCForwarderClient(cc *grpc.ClientConn) GRPCForwarderClient {
	return &gRPCForwarderClient{cc}
}

func (c *gRPCForwarderClient) SendEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/eventforwarder.GRPCForwarder/SendEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GRPCForwarder service

type GRPCForwarderServer interface {
	SendEvent(context.Context, *Event) (*Response, error)
}

func RegisterGRPCForwarderServer(s *grpc.Server, srv GRPCForwarderServer) {
	s.RegisterService(&_GRPCForwarder_serviceDesc, srv)
}

func _GRPCForwarder_SendEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCForwarderServer).SendEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventforwarder.GRPCForwarder/SendEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCForwarderServer).SendEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

var _GRPCForwarder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eventforwarder.GRPCForwarder",
	HandlerType: (*GRPCForwarderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEvent",
			Handler:    _GRPCForwarder_SendEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eventsgateway/grpc/protobuf/events.proto",
}

func init() { proto.RegisterFile("eventsgateway/grpc/protobuf/events.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0xd2, 0xa8, 0x19, 0xb1, 0xc8, 0xa0, 0xb0, 0x14, 0x0f, 0x21, 0xa7, 0x9c, 0x12,
	0xa8, 0x20, 0xc1, 0x83, 0x17, 0xa9, 0x1e, 0x2d, 0xeb, 0x13, 0x6c, 0x93, 0x69, 0x08, 0x9a, 0xec,
	0xb2, 0xbb, 0x6d, 0xc9, 0x63, 0xfa, 0x46, 0x92, 0x8d, 0xa1, 0x28, 0xbd, 0xcd, 0xff, 0xed, 0xfc,
	0x3b, 0xff, 0x0c, 0xa4, 0xb4, 0xa7, 0xce, 0x9a, 0x5a, 0x58, 0x3a, 0x88, 0x3e, 0xaf, 0xb5, 0x2a,
	0x73, 0xa5, 0xa5, 0x95, 0x9b, 0xdd, 0x36, 0x1f, 0xdf, 0x32, 0xa7, 0x71, 0xee, 0xd4, 0x56, 0xea,
	0x83, 0xd0, 0x15, 0xe9, 0xe4, 0xdb, 0x83, 0x70, 0x35, 0x20, 0x9c, 0x83, 0xdf, 0x54, 0xcc, 0x8b,
	0xbd, 0x34, 0xe2, 0x7e, 0x53, 0x21, 0xc2, 0xac, 0x13, 0x2d, 0x31, 0xdf, 0x11, 0x57, 0xe3, 0x2d,
	0x84, 0x56, 0xaa, 0xa6, 0x64, 0x81, 0x83, 0xa3, 0xc0, 0x47, 0x08, 0x95, 0x96, 0xca, 0xb0, 0x59,
	0x1c, 0xa4, 0x57, 0xcb, 0x38, 0xfb, 0x3b, 0x23, 0x73, 0xff, 0x67, 0xeb, 0xa1, 0x65, 0xd5, 0x59,
	0xdd, 0xf3, 0xb1, 0x1d, 0xef, 0x21, 0xb2, 0x4d, 0x4b, 0xc6, 0x8a, 0x56, 0xb1, 0x30, 0xf6, 0xd2,
	0x80, 0x1f, 0xc1, 0xa2, 0x00, 0x38, 0x5a, 0xf0, 0x06, 0x82, 0x4f, 0xea, 0x7f, 0xe3, 0x0d, 0xe5,
	0x90, 0x65, 0x2f, 0xbe, 0x76, 0x53, 0xc0, 0x51, 0x3c, 0xf9, 0x85, 0x97, 0x14, 0x70, 0xc9, 0xc9,
	0x28, 0xd9, 0x19, 0x1a, 0xb6, 0x28, 0x65, 0x45, 0xce, 0x18, 0x72, 0x57, 0x23, 0x83, 0x8b, 0x96,
	0x8c, 0x11, 0xf5, 0xe4, 0x9d, 0xe4, 0xf2, 0x1d, 0xae, 0xdf, 0xf8, 0xfa, 0xe5, 0x75, 0x8a, 0x8e,
	0xcf, 0x10, 0x7d, 0x50, 0x57, 0x8d, 0x17, 0xba, 0x3b, 0xb9, 0xd8, 0x82, 0xfd, 0xc7, 0xd3, 0xf0,
	0xe4, 0x6c, 0x73, 0xee, 0xae, 0xfe, 0xf0, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xc4, 0xf6, 0xdf,
	0xa1, 0x01, 0x00, 0x00,
}