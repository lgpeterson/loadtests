// Code generated by protoc-gen-go.
// source: pb/executor.proto
// DO NOT EDIT!

/*
Package executorGRPC is a generated protocol buffer package.

It is generated from these files:
	pb/executor.proto

It has these top-level messages:
	StatusMessage
	CommandMessage
	ScriptParams
*/
package executorGRPC

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

type StatusMessage struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *StatusMessage) Reset()                    { *m = StatusMessage{} }
func (m *StatusMessage) String() string            { return proto.CompactTextString(m) }
func (*StatusMessage) ProtoMessage()               {}
func (*StatusMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CommandMessage struct {
	Command      string        `protobuf:"bytes,1,opt,name=command" json:"command,omitempty"`
	ScriptParams *ScriptParams `protobuf:"bytes,2,opt,name=script_params" json:"script_params,omitempty"`
	ScriptConfig string        `protobuf:"bytes,3,opt,name=script_config" json:"script_config,omitempty"`
}

func (m *CommandMessage) Reset()                    { *m = CommandMessage{} }
func (m *CommandMessage) String() string            { return proto.CompactTextString(m) }
func (*CommandMessage) ProtoMessage()               {}
func (*CommandMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CommandMessage) GetScriptParams() *ScriptParams {
	if m != nil {
		return m.ScriptParams
	}
	return nil
}

type ScriptParams struct {
	Url                       string  `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Script                    string  `protobuf:"bytes,2,opt,name=script" json:"script,omitempty"`
	ScriptId                  string  `protobuf:"bytes,3,opt,name=script_id" json:"script_id,omitempty"`
	RunTime                   int32   `protobuf:"varint,4,opt,name=run_time" json:"run_time,omitempty"`
	MaxWorkers                int32   `protobuf:"varint,6,opt,name=max_workers" json:"max_workers,omitempty"`
	GrowthFactor              float64 `protobuf:"fixed64,8,opt,name=growth_factor" json:"growth_factor,omitempty"`
	TimeBetweenGrowth         float64 `protobuf:"fixed64,9,opt,name=time_between_growth" json:"time_between_growth,omitempty"`
	StartingRequestsPerSecond int32   `protobuf:"varint,10,opt,name=starting_requests_per_second" json:"starting_requests_per_second,omitempty"`
	MaxRequestsPerSecond      int32   `protobuf:"varint,11,opt,name=max_requests_per_second" json:"max_requests_per_second,omitempty"`
}

func (m *ScriptParams) Reset()                    { *m = ScriptParams{} }
func (m *ScriptParams) String() string            { return proto.CompactTextString(m) }
func (*ScriptParams) ProtoMessage()               {}
func (*ScriptParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*StatusMessage)(nil), "executorGRPC.StatusMessage")
	proto.RegisterType((*CommandMessage)(nil), "executorGRPC.CommandMessage")
	proto.RegisterType((*ScriptParams)(nil), "executorGRPC.ScriptParams")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Commander service

type CommanderClient interface {
	ExecuteCommand(ctx context.Context, opts ...grpc.CallOption) (Commander_ExecuteCommandClient, error)
}

type commanderClient struct {
	cc *grpc.ClientConn
}

func NewCommanderClient(cc *grpc.ClientConn) CommanderClient {
	return &commanderClient{cc}
}

func (c *commanderClient) ExecuteCommand(ctx context.Context, opts ...grpc.CallOption) (Commander_ExecuteCommandClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Commander_serviceDesc.Streams[0], c.cc, "/executorGRPC.Commander/ExecuteCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &commanderExecuteCommandClient{stream}
	return x, nil
}

type Commander_ExecuteCommandClient interface {
	Send(*CommandMessage) error
	Recv() (*StatusMessage, error)
	grpc.ClientStream
}

type commanderExecuteCommandClient struct {
	grpc.ClientStream
}

func (x *commanderExecuteCommandClient) Send(m *CommandMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *commanderExecuteCommandClient) Recv() (*StatusMessage, error) {
	m := new(StatusMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Commander service

type CommanderServer interface {
	ExecuteCommand(Commander_ExecuteCommandServer) error
}

func RegisterCommanderServer(s *grpc.Server, srv CommanderServer) {
	s.RegisterService(&_Commander_serviceDesc, srv)
}

func _Commander_ExecuteCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CommanderServer).ExecuteCommand(&commanderExecuteCommandServer{stream})
}

type Commander_ExecuteCommandServer interface {
	Send(*StatusMessage) error
	Recv() (*CommandMessage, error)
	grpc.ServerStream
}

type commanderExecuteCommandServer struct {
	grpc.ServerStream
}

func (x *commanderExecuteCommandServer) Send(m *StatusMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *commanderExecuteCommandServer) Recv() (*CommandMessage, error) {
	m := new(CommandMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Commander_serviceDesc = grpc.ServiceDesc{
	ServiceName: "executorGRPC.Commander",
	HandlerType: (*CommanderServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecuteCommand",
			Handler:       _Commander_ExecuteCommand_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4e, 0xf2, 0x40,
	0x10, 0xc7, 0xbf, 0xfd, 0x50, 0xa4, 0x53, 0x40, 0x59, 0x62, 0xdc, 0x00, 0x09, 0xa6, 0xf1, 0xc0,
	0x09, 0x15, 0x1f, 0x81, 0x18, 0x4f, 0x26, 0x28, 0x9e, 0xdd, 0x2c, 0x65, 0xc4, 0x46, 0xdb, 0xad,
	0xb3, 0xdb, 0xc0, 0xe3, 0xfa, 0x28, 0x6e, 0x97, 0x92, 0x80, 0xf1, 0xd8, 0xdf, 0xfc, 0x3a, 0x33,
	0xfb, 0x1f, 0xe8, 0xe4, 0x8b, 0x6b, 0xdc, 0x60, 0x5c, 0x58, 0x4d, 0xe3, 0x9c, 0xb4, 0xd5, 0xbc,
	0xb9, 0xfb, 0x7e, 0x78, 0x9e, 0x4d, 0xa3, 0x21, 0xb4, 0xe6, 0x56, 0xd9, 0xc2, 0x3c, 0xa2, 0x31,
	0x6a, 0x85, 0xbc, 0x0d, 0x75, 0xe3, 0x81, 0x60, 0x97, 0x6c, 0x14, 0x44, 0x2f, 0xd0, 0x9e, 0xea,
	0x34, 0x55, 0xd9, 0x72, 0x67, 0x9c, 0xc2, 0x49, 0xbc, 0x25, 0x5b, 0x85, 0xdf, 0x42, 0xcb, 0xc4,
	0x94, 0xe4, 0x56, 0xe6, 0x8a, 0x54, 0x6a, 0xc4, 0x7f, 0x87, 0xc3, 0x49, 0x6f, 0xbc, 0x3f, 0x69,
	0x3c, 0xf7, 0xca, 0xcc, 0x1b, 0xd1, 0x37, 0x83, 0xe6, 0x3e, 0xe0, 0x21, 0xd4, 0x0a, 0xfa, 0xac,
	0x1a, 0x96, 0x3b, 0xf8, 0xa2, 0xef, 0x14, 0xf0, 0x0e, 0x04, 0xd5, 0x80, 0x64, 0x29, 0x6a, 0x1e,
	0x9d, 0x41, 0x83, 0x8a, 0x4c, 0xda, 0x24, 0x45, 0x71, 0xe4, 0xc8, 0x31, 0xef, 0x42, 0x98, 0xaa,
	0x8d, 0x5c, 0x6b, 0xfa, 0x40, 0x32, 0xa2, 0xee, 0xe1, 0x39, 0xb4, 0x56, 0xa4, 0xd7, 0xf6, 0x5d,
	0xbe, 0xa9, 0xd8, 0x6d, 0x22, 0x1a, 0x0e, 0x33, 0xde, 0x87, 0x6e, 0xf9, 0xa7, 0x5c, 0xa0, 0x5d,
	0x23, 0x66, 0x72, 0xeb, 0x88, 0xc0, 0x17, 0xaf, 0x60, 0xe0, 0x12, 0x20, 0x9b, 0x64, 0x2b, 0x49,
	0xf8, 0x55, 0xa0, 0xb1, 0x46, 0xe6, 0x48, 0xd2, 0x60, 0xac, 0xdd, 0xa3, 0xc1, 0x77, 0x1e, 0xc2,
	0x45, 0x39, 0xee, 0x2f, 0x21, 0x2c, 0x85, 0xc9, 0x2b, 0x04, 0x55, 0x70, 0x48, 0xfc, 0x09, 0xda,
	0xf7, 0x3e, 0x0c, 0xac, 0x18, 0x1f, 0x1c, 0xa6, 0x73, 0x98, 0x71, 0xaf, 0xff, 0x2b, 0xbb, 0xfd,
	0x13, 0x45, 0xff, 0x46, 0xec, 0x86, 0x2d, 0xea, 0xfe, 0x9c, 0x77, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x0a, 0x26, 0x26, 0xa4, 0xe3, 0x01, 0x00, 0x00,
}
