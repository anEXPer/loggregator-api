// Code generated by protoc-gen-go. DO NOT EDIT.
// source: egress_query.proto

package loggregator_v2

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

type ContainerMetricRequest struct {
	SourceId string `protobuf:"bytes,1,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
}

func (m *ContainerMetricRequest) Reset()                    { *m = ContainerMetricRequest{} }
func (m *ContainerMetricRequest) String() string            { return proto.CompactTextString(m) }
func (*ContainerMetricRequest) ProtoMessage()               {}
func (*ContainerMetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ContainerMetricRequest) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

type QueryResponse struct {
	Envelopes []*Envelope `protobuf:"bytes,1,rep,name=envelopes" json:"envelopes,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *QueryResponse) GetEnvelopes() []*Envelope {
	if m != nil {
		return m.Envelopes
	}
	return nil
}

func init() {
	proto.RegisterType((*ContainerMetricRequest)(nil), "loggregator.v2.ContainerMetricRequest")
	proto.RegisterType((*QueryResponse)(nil), "loggregator.v2.QueryResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EgressQuery service

type EgressQueryClient interface {
	ContainerMetrics(ctx context.Context, in *ContainerMetricRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type egressQueryClient struct {
	cc *grpc.ClientConn
}

func NewEgressQueryClient(cc *grpc.ClientConn) EgressQueryClient {
	return &egressQueryClient{cc}
}

func (c *egressQueryClient) ContainerMetrics(ctx context.Context, in *ContainerMetricRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := grpc.Invoke(ctx, "/loggregator.v2.EgressQuery/ContainerMetrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EgressQuery service

type EgressQueryServer interface {
	ContainerMetrics(context.Context, *ContainerMetricRequest) (*QueryResponse, error)
}

func RegisterEgressQueryServer(s *grpc.Server, srv EgressQueryServer) {
	s.RegisterService(&_EgressQuery_serviceDesc, srv)
}

func _EgressQuery_ContainerMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EgressQueryServer).ContainerMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loggregator.v2.EgressQuery/ContainerMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EgressQueryServer).ContainerMetrics(ctx, req.(*ContainerMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EgressQuery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loggregator.v2.EgressQuery",
	HandlerType: (*EgressQueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ContainerMetrics",
			Handler:    _EgressQuery_ContainerMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "egress_query.proto",
}

func init() { proto.RegisterFile("egress_query.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x4d, 0x2f, 0x4a,
	0x2d, 0x2e, 0x8e, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0xcb, 0xc9, 0x4f, 0x4f, 0x2f, 0x4a, 0x4d, 0x4f, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x33, 0x92, 0xe2,
	0x4b, 0xcd, 0x2b, 0x4b, 0xcd, 0xc9, 0x2f, 0x48, 0x85, 0xc8, 0x2b, 0x99, 0x72, 0x89, 0x39, 0xe7,
	0xe7, 0x95, 0x24, 0x66, 0xe6, 0xa5, 0x16, 0xf9, 0xa6, 0x96, 0x14, 0x65, 0x26, 0x07, 0xa5, 0x16,
	0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x73, 0x71, 0x16, 0xe7, 0x97, 0x16, 0x25, 0xa7, 0xc6, 0x67,
	0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x71, 0x40, 0x04, 0x3c, 0x53, 0x94, 0xdc, 0xb9,
	0x78, 0x03, 0x41, 0xb6, 0x04, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x99, 0x71, 0x71,
	0xc2, 0x4c, 0x2e, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x92, 0xd0, 0x43, 0xb5, 0x5b, 0xcf,
	0x15, 0xaa, 0x20, 0x08, 0xa1, 0xd4, 0x28, 0x8b, 0x8b, 0xdb, 0x15, 0xec, 0x6a, 0xb0, 0x71, 0x42,
	0xd1, 0x5c, 0x02, 0x68, 0xce, 0x29, 0x16, 0x52, 0x43, 0x37, 0x07, 0xbb, 0x83, 0xa5, 0x64, 0xd1,
	0xd5, 0xa1, 0xb8, 0x50, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x65, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa4, 0xa8, 0x78, 0xc9, 0x28, 0x01, 0x00, 0x00,
}
