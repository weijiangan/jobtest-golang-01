// Code generated by protoc-gen-go. DO NOT EDIT.
// source: brunotest.proto

/*
Package brunotest is a generated protocol buffer package.

It is generated from these files:
	brunotest.proto

It has these top-level messages:
	AuditEvent
	QueryParam
	Response
*/
package brunotest

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

type AuditEvent struct {
	ClientIp string            `protobuf:"bytes,1,opt,name=client_ip,json=clientIp" json:"client_ip,omitempty"`
	ServerIp string            `protobuf:"bytes,2,opt,name=server_ip,json=serverIp" json:"server_ip,omitempty"`
	Tag      map[string]string `protobuf:"bytes,3,rep,name=tag" json:"tag,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Message  string            `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
}

func (m *AuditEvent) Reset()                    { *m = AuditEvent{} }
func (m *AuditEvent) String() string            { return proto.CompactTextString(m) }
func (*AuditEvent) ProtoMessage()               {}
func (*AuditEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuditEvent) GetClientIp() string {
	if m != nil {
		return m.ClientIp
	}
	return ""
}

func (m *AuditEvent) GetServerIp() string {
	if m != nil {
		return m.ServerIp
	}
	return ""
}

func (m *AuditEvent) GetTag() map[string]string {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *AuditEvent) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type QueryParam struct {
	ClientIp string            `protobuf:"bytes,1,opt,name=client_ip,json=clientIp" json:"client_ip,omitempty"`
	ServerIp string            `protobuf:"bytes,2,opt,name=server_ip,json=serverIp" json:"server_ip,omitempty"`
	Tag      map[string]string `protobuf:"bytes,3,rep,name=tag" json:"tag,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *QueryParam) Reset()                    { *m = QueryParam{} }
func (m *QueryParam) String() string            { return proto.CompactTextString(m) }
func (*QueryParam) ProtoMessage()               {}
func (*QueryParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *QueryParam) GetClientIp() string {
	if m != nil {
		return m.ClientIp
	}
	return ""
}

func (m *QueryParam) GetServerIp() string {
	if m != nil {
		return m.ServerIp
	}
	return ""
}

func (m *QueryParam) GetTag() map[string]string {
	if m != nil {
		return m.Tag
	}
	return nil
}

type Response struct {
	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
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
	proto.RegisterType((*AuditEvent)(nil), "brunotest.AuditEvent")
	proto.RegisterType((*QueryParam)(nil), "brunotest.QueryParam")
	proto.RegisterType((*Response)(nil), "brunotest.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for App service

type AppClient interface {
	Send(ctx context.Context, in *AuditEvent, opts ...grpc.CallOption) (*Response, error)
	Query(ctx context.Context, in *QueryParam, opts ...grpc.CallOption) (App_QueryClient, error)
}

type appClient struct {
	cc *grpc.ClientConn
}

func NewAppClient(cc *grpc.ClientConn) AppClient {
	return &appClient{cc}
}

func (c *appClient) Send(ctx context.Context, in *AuditEvent, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/brunotest.App/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Query(ctx context.Context, in *QueryParam, opts ...grpc.CallOption) (App_QueryClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_App_serviceDesc.Streams[0], c.cc, "/brunotest.App/Query", opts...)
	if err != nil {
		return nil, err
	}
	x := &appQueryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type App_QueryClient interface {
	Recv() (*AuditEvent, error)
	grpc.ClientStream
}

type appQueryClient struct {
	grpc.ClientStream
}

func (x *appQueryClient) Recv() (*AuditEvent, error) {
	m := new(AuditEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for App service

type AppServer interface {
	Send(context.Context, *AuditEvent) (*Response, error)
	Query(*QueryParam, App_QueryServer) error
}

func RegisterAppServer(s *grpc.Server, srv AppServer) {
	s.RegisterService(&_App_serviceDesc, srv)
}

func _App_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brunotest.App/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Send(ctx, req.(*AuditEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Query_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryParam)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AppServer).Query(m, &appQueryServer{stream})
}

type App_QueryServer interface {
	Send(*AuditEvent) error
	grpc.ServerStream
}

type appQueryServer struct {
	grpc.ServerStream
}

func (x *appQueryServer) Send(m *AuditEvent) error {
	return x.ServerStream.SendMsg(m)
}

var _App_serviceDesc = grpc.ServiceDesc{
	ServiceName: "brunotest.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _App_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Query",
			Handler:       _App_Query_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "brunotest.proto",
}

func init() { proto.RegisterFile("brunotest.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x65, 0x29, 0x28, 0x0c, 0x89, 0x9a, 0x55, 0x93, 0x06, 0x13, 0x21, 0x3d, 0x71, 0x6a, 0x08,
	0x18, 0xa3, 0xde, 0xc0, 0x70, 0xe0, 0x86, 0x95, 0x3b, 0x59, 0xe8, 0xa4, 0x56, 0x61, 0x77, 0xb3,
	0xbb, 0xad, 0xe1, 0xaf, 0xfc, 0x0e, 0xbf, 0xca, 0x6c, 0x0b, 0x96, 0x26, 0x9c, 0x4c, 0xbc, 0xcd,
	0xbc, 0x99, 0x79, 0x33, 0xef, 0x65, 0xe0, 0x7c, 0xa9, 0x12, 0x2e, 0x0c, 0x6a, 0xe3, 0x4b, 0x25,
	0x8c, 0xa0, 0xcd, 0x5f, 0xc0, 0xfb, 0x26, 0x00, 0xa3, 0x24, 0x8c, 0xcd, 0x24, 0x45, 0x6e, 0xe8,
	0x0d, 0x34, 0x57, 0xeb, 0x18, 0xb9, 0x59, 0xc4, 0xd2, 0x25, 0x5d, 0xd2, 0x6b, 0x06, 0x8d, 0x1c,
	0x98, 0x4a, 0x5b, 0xd4, 0xa8, 0x52, 0x54, 0xb6, 0x58, 0xcd, 0x8b, 0x39, 0x30, 0x95, 0xb4, 0x0f,
	0x8e, 0x61, 0x91, 0xeb, 0x74, 0x9d, 0x5e, 0x6b, 0x70, 0xeb, 0x17, 0x2b, 0x0b, 0x76, 0x7f, 0xce,
	0xa2, 0x09, 0x37, 0x6a, 0x1b, 0xd8, 0x56, 0xea, 0xc2, 0xe9, 0x06, 0xb5, 0x66, 0x11, 0xba, 0xb5,
	0x8c, 0x6c, 0x9f, 0xb6, 0xef, 0xa1, 0xb1, 0x6f, 0xa5, 0x17, 0xe0, 0x7c, 0xe0, 0x76, 0x77, 0x8b,
	0x0d, 0xe9, 0x15, 0xd4, 0x53, 0xb6, 0x4e, 0x70, 0x77, 0x42, 0x9e, 0x3c, 0x55, 0x1f, 0x88, 0xf7,
	0x45, 0x00, 0x5e, 0x12, 0x54, 0xdb, 0x19, 0x53, 0x6c, 0xf3, 0x1f, 0x62, 0x0a, 0xf6, 0xb2, 0x98,
	0x3f, 0x9f, 0x3c, 0x81, 0x46, 0x80, 0x5a, 0x0a, 0xae, 0x91, 0x76, 0xa0, 0xa5, 0x0d, 0x33, 0x89,
	0x5e, 0xac, 0x44, 0x88, 0xd9, 0x7c, 0x3d, 0x80, 0x1c, 0x7a, 0x16, 0x21, 0x1e, 0x3a, 0x56, 0x2d,
	0x39, 0x36, 0x48, 0xc1, 0x19, 0x49, 0x49, 0xef, 0xa0, 0xf6, 0x8a, 0x3c, 0xa4, 0xd7, 0x47, 0xfd,
	0x6f, 0x5f, 0x1e, 0xc0, 0xfb, 0xad, 0x5e, 0x85, 0x3e, 0x42, 0x3d, 0xd3, 0x55, 0x1a, 0x2b, 0x94,
	0xb6, 0x8f, 0xb3, 0x79, 0x95, 0x3e, 0x19, 0x0f, 0xa1, 0xb3, 0x12, 0x1b, 0x3f, 0x8a, 0xcd, 0x5b,
	0xb2, 0xf4, 0x3f, 0x31, 0x7e, 0x8f, 0x19, 0x8f, 0x18, 0x2f, 0x26, 0xc6, 0x67, 0x63, 0x1b, 0xce,
	0x51, 0x9b, 0x99, 0x7d, 0xbe, 0x19, 0x59, 0x9e, 0x64, 0x5f, 0x38, 0xfc, 0x09, 0x00, 0x00, 0xff,
	0xff, 0xef, 0x19, 0x8f, 0x63, 0x98, 0x02, 0x00, 0x00,
}
