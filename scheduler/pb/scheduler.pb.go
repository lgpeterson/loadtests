// Code generated by protoc-gen-go.
// source: pb/scheduler.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/scheduler.proto

It has these top-level messages:
	LoadTestReq
	LoadTestResp
	RegisterExecutorReq
	RegisterExecutorResp
*/
package pb

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

type LoadTestReq struct {
	Url                       string  `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Script                    string  `protobuf:"bytes,2,opt,name=script" json:"script,omitempty"`
	ScriptName                string  `protobuf:"bytes,3,opt,name=script_name" json:"script_name,omitempty"`
	RunTime                   int32   `protobuf:"varint,4,opt,name=run_time" json:"run_time,omitempty"`
	GrowthFactor              float64 `protobuf:"fixed64,8,opt,name=growth_factor" json:"growth_factor,omitempty"`
	TimeBetweenGrowth         float64 `protobuf:"fixed64,9,opt,name=time_between_growth" json:"time_between_growth,omitempty"`
	StartingRequestsPerSecond int32   `protobuf:"varint,10,opt,name=starting_requests_per_second" json:"starting_requests_per_second,omitempty"`
	MaxRequestsPerSecond      int32   `protobuf:"varint,11,opt,name=max_requests_per_second" json:"max_requests_per_second,omitempty"`
	ScriptConfig              string  `protobuf:"bytes,12,opt,name=script_config" json:"script_config,omitempty"`
}

func (m *LoadTestReq) Reset()                    { *m = LoadTestReq{} }
func (m *LoadTestReq) String() string            { return proto.CompactTextString(m) }
func (*LoadTestReq) ProtoMessage()               {}
func (*LoadTestReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LoadTestResp struct {
	// Types that are valid to be assigned to Phase:
	//	*LoadTestResp_Preparing_
	//	*LoadTestResp_Start
	//	*LoadTestResp_Finish
	//	*LoadTestResp_Error
	Phase isLoadTestResp_Phase `protobuf_oneof:"phase"`
}

func (m *LoadTestResp) Reset()                    { *m = LoadTestResp{} }
func (m *LoadTestResp) String() string            { return proto.CompactTextString(m) }
func (*LoadTestResp) ProtoMessage()               {}
func (*LoadTestResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isLoadTestResp_Phase interface {
	isLoadTestResp_Phase()
}

type LoadTestResp_Preparing_ struct {
	Preparing *LoadTestResp_Preparing `protobuf:"bytes,1,opt,name=preparing,oneof"`
}
type LoadTestResp_Start struct {
	Start *LoadTestResp_Started `protobuf:"bytes,2,opt,name=start,oneof"`
}
type LoadTestResp_Finish struct {
	Finish *LoadTestResp_Finished `protobuf:"bytes,3,opt,name=finish,oneof"`
}
type LoadTestResp_Error struct {
	Error *LoadTestResp_Errored `protobuf:"bytes,4,opt,name=error,oneof"`
}

func (*LoadTestResp_Preparing_) isLoadTestResp_Phase() {}
func (*LoadTestResp_Start) isLoadTestResp_Phase()      {}
func (*LoadTestResp_Finish) isLoadTestResp_Phase()     {}
func (*LoadTestResp_Error) isLoadTestResp_Phase()      {}

func (m *LoadTestResp) GetPhase() isLoadTestResp_Phase {
	if m != nil {
		return m.Phase
	}
	return nil
}

func (m *LoadTestResp) GetPreparing() *LoadTestResp_Preparing {
	if x, ok := m.GetPhase().(*LoadTestResp_Preparing_); ok {
		return x.Preparing
	}
	return nil
}

func (m *LoadTestResp) GetStart() *LoadTestResp_Started {
	if x, ok := m.GetPhase().(*LoadTestResp_Start); ok {
		return x.Start
	}
	return nil
}

func (m *LoadTestResp) GetFinish() *LoadTestResp_Finished {
	if x, ok := m.GetPhase().(*LoadTestResp_Finish); ok {
		return x.Finish
	}
	return nil
}

func (m *LoadTestResp) GetError() *LoadTestResp_Errored {
	if x, ok := m.GetPhase().(*LoadTestResp_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LoadTestResp) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _LoadTestResp_OneofMarshaler, _LoadTestResp_OneofUnmarshaler, []interface{}{
		(*LoadTestResp_Preparing_)(nil),
		(*LoadTestResp_Start)(nil),
		(*LoadTestResp_Finish)(nil),
		(*LoadTestResp_Error)(nil),
	}
}

func _LoadTestResp_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LoadTestResp)
	// phase
	switch x := m.Phase.(type) {
	case *LoadTestResp_Preparing_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Preparing); err != nil {
			return err
		}
	case *LoadTestResp_Start:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Start); err != nil {
			return err
		}
	case *LoadTestResp_Finish:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Finish); err != nil {
			return err
		}
	case *LoadTestResp_Error:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LoadTestResp.Phase has unexpected type %T", x)
	}
	return nil
}

func _LoadTestResp_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LoadTestResp)
	switch tag {
	case 1: // phase.preparing
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LoadTestResp_Preparing)
		err := b.DecodeMessage(msg)
		m.Phase = &LoadTestResp_Preparing_{msg}
		return true, err
	case 2: // phase.start
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LoadTestResp_Started)
		err := b.DecodeMessage(msg)
		m.Phase = &LoadTestResp_Start{msg}
		return true, err
	case 3: // phase.finish
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LoadTestResp_Finished)
		err := b.DecodeMessage(msg)
		m.Phase = &LoadTestResp_Finish{msg}
		return true, err
	case 4: // phase.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LoadTestResp_Errored)
		err := b.DecodeMessage(msg)
		m.Phase = &LoadTestResp_Error{msg}
		return true, err
	default:
		return false, nil
	}
}

type LoadTestResp_Preparing struct {
	Count int32 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *LoadTestResp_Preparing) Reset()                    { *m = LoadTestResp_Preparing{} }
func (m *LoadTestResp_Preparing) String() string            { return proto.CompactTextString(m) }
func (*LoadTestResp_Preparing) ProtoMessage()               {}
func (*LoadTestResp_Preparing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type LoadTestResp_Started struct {
}

func (m *LoadTestResp_Started) Reset()                    { *m = LoadTestResp_Started{} }
func (m *LoadTestResp_Started) String() string            { return proto.CompactTextString(m) }
func (*LoadTestResp_Started) ProtoMessage()               {}
func (*LoadTestResp_Started) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

type LoadTestResp_Finished struct {
}

func (m *LoadTestResp_Finished) Reset()                    { *m = LoadTestResp_Finished{} }
func (m *LoadTestResp_Finished) String() string            { return proto.CompactTextString(m) }
func (*LoadTestResp_Finished) ProtoMessage()               {}
func (*LoadTestResp_Finished) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 2} }

type LoadTestResp_Errored struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *LoadTestResp_Errored) Reset()                    { *m = LoadTestResp_Errored{} }
func (m *LoadTestResp_Errored) String() string            { return proto.CompactTextString(m) }
func (*LoadTestResp_Errored) ProtoMessage()               {}
func (*LoadTestResp_Errored) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 3} }

type RegisterExecutorReq struct {
	DropletId int64 `protobuf:"varint,1,opt,name=droplet_id" json:"droplet_id,omitempty"`
	Port      int64 `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
}

func (m *RegisterExecutorReq) Reset()                    { *m = RegisterExecutorReq{} }
func (m *RegisterExecutorReq) String() string            { return proto.CompactTextString(m) }
func (*RegisterExecutorReq) ProtoMessage()               {}
func (*RegisterExecutorReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type RegisterExecutorResp struct {
	InfluxAddr     string `protobuf:"bytes,1,opt,name=influx_addr" json:"influx_addr,omitempty"`
	InfluxUsername string `protobuf:"bytes,2,opt,name=influx_username" json:"influx_username,omitempty"`
	InfluxPassword string `protobuf:"bytes,3,opt,name=influx_password" json:"influx_password,omitempty"`
	InfluxDb       string `protobuf:"bytes,4,opt,name=influx_db" json:"influx_db,omitempty"`
	InfluxSsl      bool   `protobuf:"varint,5,opt,name=influx_ssl" json:"influx_ssl,omitempty"`
}

func (m *RegisterExecutorResp) Reset()                    { *m = RegisterExecutorResp{} }
func (m *RegisterExecutorResp) String() string            { return proto.CompactTextString(m) }
func (*RegisterExecutorResp) ProtoMessage()               {}
func (*RegisterExecutorResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*LoadTestReq)(nil), "loadtests.LoadTestReq")
	proto.RegisterType((*LoadTestResp)(nil), "loadtests.LoadTestResp")
	proto.RegisterType((*LoadTestResp_Preparing)(nil), "loadtests.LoadTestResp.Preparing")
	proto.RegisterType((*LoadTestResp_Started)(nil), "loadtests.LoadTestResp.Started")
	proto.RegisterType((*LoadTestResp_Finished)(nil), "loadtests.LoadTestResp.Finished")
	proto.RegisterType((*LoadTestResp_Errored)(nil), "loadtests.LoadTestResp.Errored")
	proto.RegisterType((*RegisterExecutorReq)(nil), "loadtests.RegisterExecutorReq")
	proto.RegisterType((*RegisterExecutorResp)(nil), "loadtests.RegisterExecutorResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Scheduler service

type SchedulerClient interface {
	LoadTest(ctx context.Context, in *LoadTestReq, opts ...grpc.CallOption) (Scheduler_LoadTestClient, error)
	RegisterExecutor(ctx context.Context, in *RegisterExecutorReq, opts ...grpc.CallOption) (*RegisterExecutorResp, error)
}

type schedulerClient struct {
	cc *grpc.ClientConn
}

func NewSchedulerClient(cc *grpc.ClientConn) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) LoadTest(ctx context.Context, in *LoadTestReq, opts ...grpc.CallOption) (Scheduler_LoadTestClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Scheduler_serviceDesc.Streams[0], c.cc, "/loadtests.Scheduler/LoadTest", opts...)
	if err != nil {
		return nil, err
	}
	x := &schedulerLoadTestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Scheduler_LoadTestClient interface {
	Recv() (*LoadTestResp, error)
	grpc.ClientStream
}

type schedulerLoadTestClient struct {
	grpc.ClientStream
}

func (x *schedulerLoadTestClient) Recv() (*LoadTestResp, error) {
	m := new(LoadTestResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *schedulerClient) RegisterExecutor(ctx context.Context, in *RegisterExecutorReq, opts ...grpc.CallOption) (*RegisterExecutorResp, error) {
	out := new(RegisterExecutorResp)
	err := grpc.Invoke(ctx, "/loadtests.Scheduler/RegisterExecutor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Scheduler service

type SchedulerServer interface {
	LoadTest(*LoadTestReq, Scheduler_LoadTestServer) error
	RegisterExecutor(context.Context, *RegisterExecutorReq) (*RegisterExecutorResp, error)
}

func RegisterSchedulerServer(s *grpc.Server, srv SchedulerServer) {
	s.RegisterService(&_Scheduler_serviceDesc, srv)
}

func _Scheduler_LoadTest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LoadTestReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SchedulerServer).LoadTest(m, &schedulerLoadTestServer{stream})
}

type Scheduler_LoadTestServer interface {
	Send(*LoadTestResp) error
	grpc.ServerStream
}

type schedulerLoadTestServer struct {
	grpc.ServerStream
}

func (x *schedulerLoadTestServer) Send(m *LoadTestResp) error {
	return x.ServerStream.SendMsg(m)
}

func _Scheduler_RegisterExecutor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RegisterExecutorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SchedulerServer).RegisterExecutor(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loadtests.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterExecutor",
			Handler:    _Scheduler_RegisterExecutor_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LoadTest",
			Handler:       _Scheduler_LoadTest_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x93, 0x38, 0xb5, 0xc7, 0x2d, 0x94, 0x0d, 0x10, 0xcb, 0x20, 0x52, 0x2c, 0x0e, 0x9c,
	0x42, 0x65, 0x24, 0x38, 0x22, 0x55, 0x2a, 0xe2, 0xc0, 0x01, 0xb5, 0x70, 0xe1, 0xb2, 0xb2, 0xbd,
	0x93, 0xc4, 0xc2, 0xf5, 0x6e, 0x76, 0xd7, 0x4a, 0x3e, 0x80, 0x3f, 0xe1, 0x17, 0xf8, 0x21, 0xfe,
	0x84, 0xdd, 0x8d, 0x53, 0xaa, 0x92, 0xf4, 0x38, 0x6f, 0xde, 0xdb, 0x99, 0x79, 0x7e, 0x06, 0x22,
	0x8a, 0x37, 0xaa, 0x5c, 0x20, 0x6b, 0x6b, 0x94, 0x53, 0x21, 0xb9, 0xe6, 0x24, 0xac, 0x79, 0xce,
	0x34, 0x2a, 0xad, 0xd2, 0x3f, 0x1e, 0x44, 0x9f, 0x4d, 0xf5, 0xd5, 0x54, 0x97, 0xb8, 0x24, 0x11,
	0xf4, 0x5b, 0x59, 0xc7, 0xde, 0xa9, 0xf7, 0x3a, 0x24, 0x0f, 0x60, 0xa8, 0x4a, 0x59, 0x09, 0x1d,
	0xf7, 0x5c, 0x3d, 0x82, 0x68, 0x53, 0xd3, 0x26, 0xbf, 0xc6, 0xb8, 0xef, 0xc0, 0x13, 0x08, 0x64,
	0xdb, 0x50, 0x5d, 0x19, 0x64, 0x60, 0x10, 0xdf, 0xd2, 0xae, 0xf3, 0x35, 0x5d, 0x71, 0xf9, 0x03,
	0xa5, 0x8a, 0x87, 0x0e, 0x7c, 0x02, 0xc7, 0x73, 0xc9, 0x57, 0x7a, 0x41, 0x67, 0x79, 0xa9, 0xb9,
	0x8c, 0x03, 0x03, 0x7b, 0xe4, 0x19, 0x8c, 0xac, 0x92, 0x16, 0xa8, 0x57, 0x88, 0x0d, 0xdd, 0x70,
	0xe2, 0xd0, 0x35, 0x5f, 0xc1, 0x73, 0xa5, 0x73, 0xa9, 0xab, 0x66, 0x4e, 0x25, 0x2e, 0x5b, 0xbb,
	0x31, 0x15, 0x28, 0xa9, 0xc2, 0x92, 0x37, 0x2c, 0x06, 0xf7, 0xf2, 0x04, 0xc6, 0x76, 0xdc, 0x2e,
	0x42, 0x64, 0x09, 0xe9, 0xef, 0x1e, 0x1c, 0xfd, 0xbb, 0x51, 0x09, 0xf2, 0x0e, 0x42, 0x21, 0x51,
	0xe4, 0xd2, 0x3c, 0xec, 0x4e, 0x8d, 0xb2, 0x97, 0xd3, 0x1b, 0x4f, 0xa6, 0xb7, 0xb9, 0xd3, 0x2f,
	0x5b, 0xe2, 0xa7, 0x03, 0x72, 0x06, 0xbe, 0xdb, 0xc7, 0xd9, 0x11, 0x65, 0x93, 0x7d, 0x9a, 0x2b,
	0x4b, 0x42, 0x66, 0x14, 0x19, 0x0c, 0x67, 0x55, 0x53, 0xa9, 0x85, 0x33, 0x2b, 0xca, 0x4e, 0xf7,
	0x49, 0x3e, 0x3a, 0x96, 0xd3, 0x98, 0x29, 0x28, 0xa5, 0x71, 0x68, 0x70, 0xff, 0x94, 0x0b, 0x4b,
	0xb2, 0x8a, 0x24, 0x81, 0xf0, 0x66, 0x4d, 0x72, 0x0c, 0x7e, 0xc9, 0xdb, 0x46, 0xbb, 0xc3, 0xfc,
	0x24, 0x84, 0xc3, 0x6e, 0x9d, 0x04, 0x20, 0xd8, 0x8e, 0x49, 0x62, 0x38, 0xec, 0xf4, 0x56, 0xb0,
	0x99, 0xe7, 0x3e, 0xfa, 0xf9, 0x21, 0xf8, 0x62, 0x91, 0x2b, 0x4c, 0xdf, 0xc3, 0xe8, 0x12, 0xe7,
	0x95, 0xd2, 0x28, 0x2f, 0xd6, 0x58, 0xb6, 0xe6, 0xa3, 0xd9, 0x84, 0x10, 0x00, 0x26, 0xb9, 0xa8,
	0x51, 0xd3, 0x8a, 0x39, 0x4d, 0x9f, 0x1c, 0xc1, 0x40, 0xf0, 0xce, 0x97, 0x7e, 0xfa, 0xd3, 0x83,
	0xc7, 0xff, 0x2b, 0x8d, 0xef, 0x26, 0x18, 0x55, 0x33, 0xab, 0xdb, 0x35, 0xcd, 0x19, 0xeb, 0xe6,
	0x91, 0x31, 0x3c, 0xec, 0xc0, 0x56, 0xa1, 0x74, 0xc1, 0xea, 0xdd, 0x69, 0x88, 0x5c, 0x29, 0x93,
	0x26, 0xd6, 0x25, 0xee, 0x11, 0x84, 0x5d, 0x83, 0x15, 0xce, 0xa4, 0xd0, 0x2e, 0xd5, 0x41, 0x4a,
	0xd5, 0xb1, 0x6f, 0xb0, 0x20, 0xfb, 0xe5, 0x41, 0x78, 0xb5, 0x4d, 0x3e, 0xf9, 0x00, 0xc1, 0xd6,
	0x3d, 0xf2, 0x74, 0xa7, 0xa5, 0xcb, 0x64, 0xbc, 0xc7, 0xea, 0xf4, 0xe0, 0xcc, 0x23, 0xdf, 0xe0,
	0xe4, 0xee, 0x51, 0xe4, 0xc5, 0x2d, 0xc1, 0x0e, 0xaf, 0x92, 0xc9, 0xbd, 0x7d, 0xfb, 0xf0, 0xf9,
	0xe0, 0x7b, 0x4f, 0x14, 0xc5, 0xd0, 0xfd, 0x98, 0x6f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x8f,
	0xf8, 0xe3, 0x3a, 0xae, 0x03, 0x00, 0x00,
}
