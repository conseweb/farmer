// Code generated by protoc-gen-go.
// source: teller.proto
// DO NOT EDIT!

package protos

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

// NextLotteryInfoReq
type NextLotteryInfoReq struct {
}

func (m *NextLotteryInfoReq) Reset()                    { *m = NextLotteryInfoReq{} }
func (m *NextLotteryInfoReq) String() string            { return proto.CompactTextString(m) }
func (*NextLotteryInfoReq) ProtoMessage()               {}
func (*NextLotteryInfoReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

// NextLotteryInfoRsp
type NextLotteryInfoRsp struct {
	Error     *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	StartTime int64  `protobuf:"varint,2,opt,name=startTime" json:"startTime,omitempty"`
	EndTime   int64  `protobuf:"varint,3,opt,name=endTime" json:"endTime,omitempty"`
}

func (m *NextLotteryInfoRsp) Reset()                    { *m = NextLotteryInfoRsp{} }
func (m *NextLotteryInfoRsp) String() string            { return proto.CompactTextString(m) }
func (*NextLotteryInfoRsp) ProtoMessage()               {}
func (*NextLotteryInfoRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *NextLotteryInfoRsp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// storage object of farmer lottery
type LotteryFx struct {
	Fid   string `protobuf:"bytes,1,opt,name=fid" json:"fid,omitempty"`
	Value uint64 `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	// middle R, when teller receive farmer's lottery, so farmer's lottery relate to call queue.
	Mr uint64 `protobuf:"varint,3,opt,name=mr" json:"mr,omitempty"`
	// when handle lottery, calculate the distence between fx and lx, first xx framers will be selected as the ledger's voter
	// when selected, candidate will be ledger's ids, otherwise nothing
	Ledgers []string `protobuf:"bytes,4,rep,name=ledgers" json:"ledgers,omitempty"`
	Dist    uint64   `protobuf:"varint,5,opt,name=dist" json:"dist,omitempty"`
}

func (m *LotteryFx) Reset()                    { *m = LotteryFx{} }
func (m *LotteryFx) String() string            { return proto.CompactTextString(m) }
func (*LotteryFx) ProtoMessage()               {}
func (*LotteryFx) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

// storage object of ledger lottery
type LotteryLx struct {
	Lid   string `protobuf:"bytes,1,opt,name=lid" json:"lid,omitempty"`
	Value uint64 `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	// the distence of value and end R
	Dist uint64 `protobuf:"varint,3,opt,name=dist" json:"dist,omitempty"`
	// win a seat for ledger?
	Won bool `protobuf:"varint,4,opt,name=won" json:"won,omitempty"`
	// farmers
	Farmers []string `protobuf:"bytes,5,rep,name=farmers" json:"farmers,omitempty"`
}

func (m *LotteryLx) Reset()                    { *m = LotteryLx{} }
func (m *LotteryLx) String() string            { return proto.CompactTextString(m) }
func (*LotteryLx) ProtoMessage()               {}
func (*LotteryLx) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

// LotteryFxTicket farmer only
type LotteryFxTicket struct {
	Fid         string `protobuf:"bytes,1,opt,name=fid" json:"fid,omitempty"`
	Fx          uint64 `protobuf:"varint,2,opt,name=fx" json:"fx,omitempty"`
	Mr          uint64 `protobuf:"varint,3,opt,name=mr" json:"mr,omitempty"`
	Idx         int64  `protobuf:"varint,4,opt,name=idx" json:"idx,omitempty"`
	LotteryName string `protobuf:"bytes,5,opt,name=lotteryName" json:"lotteryName,omitempty"`
}

func (m *LotteryFxTicket) Reset()                    { *m = LotteryFxTicket{} }
func (m *LotteryFxTicket) String() string            { return proto.CompactTextString(m) }
func (*LotteryFxTicket) ProtoMessage()               {}
func (*LotteryFxTicket) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

// SendLotteryFxReq
type SendLotteryFxReq struct {
	Fid string `protobuf:"bytes,1,opt,name=fid" json:"fid,omitempty"`
	Fx  uint64 `protobuf:"varint,2,opt,name=fx" json:"fx,omitempty"`
}

func (m *SendLotteryFxReq) Reset()                    { *m = SendLotteryFxReq{} }
func (m *SendLotteryFxReq) String() string            { return proto.CompactTextString(m) }
func (*SendLotteryFxReq) ProtoMessage()               {}
func (*SendLotteryFxReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

// SendLotteryFxRsp
type SendLotteryFxRsp struct {
	Error  *Error           `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Ticket *LotteryFxTicket `protobuf:"bytes,2,opt,name=ticket" json:"ticket,omitempty"`
}

func (m *SendLotteryFxRsp) Reset()                    { *m = SendLotteryFxRsp{} }
func (m *SendLotteryFxRsp) String() string            { return proto.CompactTextString(m) }
func (*SendLotteryFxRsp) ProtoMessage()               {}
func (*SendLotteryFxRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *SendLotteryFxRsp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *SendLotteryFxRsp) GetTicket() *LotteryFxTicket {
	if m != nil {
		return m.Ticket
	}
	return nil
}

// LotteryLxTicket ledger only
type LotteryLxTicket struct {
	Lid         string `protobuf:"bytes,1,opt,name=lid" json:"lid,omitempty"`
	Lx          uint64 `protobuf:"varint,2,opt,name=lx" json:"lx,omitempty"`
	LotteryName string `protobuf:"bytes,3,opt,name=lotteryName" json:"lotteryName,omitempty"`
}

func (m *LotteryLxTicket) Reset()                    { *m = LotteryLxTicket{} }
func (m *LotteryLxTicket) String() string            { return proto.CompactTextString(m) }
func (*LotteryLxTicket) ProtoMessage()               {}
func (*LotteryLxTicket) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

// SendLotteryLxReq
type SendLotteryLxReq struct {
	Lid string `protobuf:"bytes,1,opt,name=lid" json:"lid,omitempty"`
	Lx  uint64 `protobuf:"varint,2,opt,name=lx" json:"lx,omitempty"`
}

func (m *SendLotteryLxReq) Reset()                    { *m = SendLotteryLxReq{} }
func (m *SendLotteryLxReq) String() string            { return proto.CompactTextString(m) }
func (*SendLotteryLxReq) ProtoMessage()               {}
func (*SendLotteryLxReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

// SendLotteryLxRsp
type SendLotteryLxRsp struct {
	Error  *Error           `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Ticket *LotteryLxTicket `protobuf:"bytes,2,opt,name=ticket" json:"ticket,omitempty"`
}

func (m *SendLotteryLxRsp) Reset()                    { *m = SendLotteryLxRsp{} }
func (m *SendLotteryLxRsp) String() string            { return proto.CompactTextString(m) }
func (*SendLotteryLxRsp) ProtoMessage()               {}
func (*SendLotteryLxRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *SendLotteryLxRsp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *SendLotteryLxRsp) GetTicket() *LotteryLxTicket {
	if m != nil {
		return m.Ticket
	}
	return nil
}

// StartLotteryReq
type StartLotteryReq struct {
	// when to start a new round of lottery, is a time utc timestamp, if smaller than NOW more than 1m, using now
	StartUTC int64 `protobuf:"varint,1,opt,name=startUTC" json:"startUTC,omitempty"`
	// how long the round of lottery will last, using ms,s,m,h words, such as 30m means 30 minutes
	LastInterval string `protobuf:"bytes,2,opt,name=lastInterval" json:"lastInterval,omitempty"`
}

func (m *StartLotteryReq) Reset()                    { *m = StartLotteryReq{} }
func (m *StartLotteryReq) String() string            { return proto.CompactTextString(m) }
func (*StartLotteryReq) ProtoMessage()               {}
func (*StartLotteryReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

// StartLotteryRsp
type StartLotteryRsp struct {
	Error *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *StartLotteryRsp) Reset()                    { *m = StartLotteryRsp{} }
func (m *StartLotteryRsp) String() string            { return proto.CompactTextString(m) }
func (*StartLotteryRsp) ProtoMessage()               {}
func (*StartLotteryRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

func (m *StartLotteryRsp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// GetLotteryResultReq
type GetLotteryResultReq struct {
	LotteryName string `protobuf:"bytes,1,opt,name=lotteryName" json:"lotteryName,omitempty"`
}

func (m *GetLotteryResultReq) Reset()                    { *m = GetLotteryResultReq{} }
func (m *GetLotteryResultReq) String() string            { return proto.CompactTextString(m) }
func (*GetLotteryResultReq) ProtoMessage()               {}
func (*GetLotteryResultReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{12} }

// LotteryResult
type LotteryResult struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// farmer lottery list
	Fxs []*LotteryFx `protobuf:"bytes,2,rep,name=fxs" json:"fxs,omitempty"`
	// ledger lottery list
	Lxs []*LotteryLx `protobuf:"bytes,3,rep,name=lxs" json:"lxs,omitempty"`
}

func (m *LotteryResult) Reset()                    { *m = LotteryResult{} }
func (m *LotteryResult) String() string            { return proto.CompactTextString(m) }
func (*LotteryResult) ProtoMessage()               {}
func (*LotteryResult) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{13} }

func (m *LotteryResult) GetFxs() []*LotteryFx {
	if m != nil {
		return m.Fxs
	}
	return nil
}

func (m *LotteryResult) GetLxs() []*LotteryLx {
	if m != nil {
		return m.Lxs
	}
	return nil
}

// GetLotteryResultRsp
type GetLotteryResultRsp struct {
	Error  *Error         `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Result *LotteryResult `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
}

func (m *GetLotteryResultRsp) Reset()                    { *m = GetLotteryResultRsp{} }
func (m *GetLotteryResultRsp) String() string            { return proto.CompactTextString(m) }
func (*GetLotteryResultRsp) ProtoMessage()               {}
func (*GetLotteryResultRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{14} }

func (m *GetLotteryResultRsp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *GetLotteryResultRsp) GetResult() *LotteryResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*NextLotteryInfoReq)(nil), "protos.NextLotteryInfoReq")
	proto.RegisterType((*NextLotteryInfoRsp)(nil), "protos.NextLotteryInfoRsp")
	proto.RegisterType((*LotteryFx)(nil), "protos.LotteryFx")
	proto.RegisterType((*LotteryLx)(nil), "protos.LotteryLx")
	proto.RegisterType((*LotteryFxTicket)(nil), "protos.LotteryFxTicket")
	proto.RegisterType((*SendLotteryFxReq)(nil), "protos.SendLotteryFxReq")
	proto.RegisterType((*SendLotteryFxRsp)(nil), "protos.SendLotteryFxRsp")
	proto.RegisterType((*LotteryLxTicket)(nil), "protos.LotteryLxTicket")
	proto.RegisterType((*SendLotteryLxReq)(nil), "protos.SendLotteryLxReq")
	proto.RegisterType((*SendLotteryLxRsp)(nil), "protos.SendLotteryLxRsp")
	proto.RegisterType((*StartLotteryReq)(nil), "protos.StartLotteryReq")
	proto.RegisterType((*StartLotteryRsp)(nil), "protos.StartLotteryRsp")
	proto.RegisterType((*GetLotteryResultReq)(nil), "protos.GetLotteryResultReq")
	proto.RegisterType((*LotteryResult)(nil), "protos.LotteryResult")
	proto.RegisterType((*GetLotteryResultRsp)(nil), "protos.GetLotteryResultRsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for LotteryAPI service

type LotteryAPIClient interface {
	// returns next lottery info, something about time begin, end etc...
	NextLotteryInfo(ctx context.Context, in *NextLotteryInfoReq, opts ...grpc.CallOption) (*NextLotteryInfoRsp, error)
	// receive lottery number form farmer
	SendLotteryFx(ctx context.Context, in *SendLotteryFxReq, opts ...grpc.CallOption) (*SendLotteryFxRsp, error)
	// receive lottery number form ledger
	SendLotteryLx(ctx context.Context, in *SendLotteryLxReq, opts ...grpc.CallOption) (*SendLotteryLxRsp, error)
	// send a command to start new round of lottery immediately
	StartLottery(ctx context.Context, in *StartLotteryReq, opts ...grpc.CallOption) (*StartLotteryRsp, error)
	// get lottery result to verify or someelse use
	GetLotteryResult(ctx context.Context, in *GetLotteryResultReq, opts ...grpc.CallOption) (*GetLotteryResultRsp, error)
}

type lotteryAPIClient struct {
	cc *grpc.ClientConn
}

func NewLotteryAPIClient(cc *grpc.ClientConn) LotteryAPIClient {
	return &lotteryAPIClient{cc}
}

func (c *lotteryAPIClient) NextLotteryInfo(ctx context.Context, in *NextLotteryInfoReq, opts ...grpc.CallOption) (*NextLotteryInfoRsp, error) {
	out := new(NextLotteryInfoRsp)
	err := grpc.Invoke(ctx, "/protos.LotteryAPI/NextLotteryInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lotteryAPIClient) SendLotteryFx(ctx context.Context, in *SendLotteryFxReq, opts ...grpc.CallOption) (*SendLotteryFxRsp, error) {
	out := new(SendLotteryFxRsp)
	err := grpc.Invoke(ctx, "/protos.LotteryAPI/SendLotteryFx", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lotteryAPIClient) SendLotteryLx(ctx context.Context, in *SendLotteryLxReq, opts ...grpc.CallOption) (*SendLotteryLxRsp, error) {
	out := new(SendLotteryLxRsp)
	err := grpc.Invoke(ctx, "/protos.LotteryAPI/SendLotteryLx", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lotteryAPIClient) StartLottery(ctx context.Context, in *StartLotteryReq, opts ...grpc.CallOption) (*StartLotteryRsp, error) {
	out := new(StartLotteryRsp)
	err := grpc.Invoke(ctx, "/protos.LotteryAPI/StartLottery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lotteryAPIClient) GetLotteryResult(ctx context.Context, in *GetLotteryResultReq, opts ...grpc.CallOption) (*GetLotteryResultRsp, error) {
	out := new(GetLotteryResultRsp)
	err := grpc.Invoke(ctx, "/protos.LotteryAPI/GetLotteryResult", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LotteryAPI service

type LotteryAPIServer interface {
	// returns next lottery info, something about time begin, end etc...
	NextLotteryInfo(context.Context, *NextLotteryInfoReq) (*NextLotteryInfoRsp, error)
	// receive lottery number form farmer
	SendLotteryFx(context.Context, *SendLotteryFxReq) (*SendLotteryFxRsp, error)
	// receive lottery number form ledger
	SendLotteryLx(context.Context, *SendLotteryLxReq) (*SendLotteryLxRsp, error)
	// send a command to start new round of lottery immediately
	StartLottery(context.Context, *StartLotteryReq) (*StartLotteryRsp, error)
	// get lottery result to verify or someelse use
	GetLotteryResult(context.Context, *GetLotteryResultReq) (*GetLotteryResultRsp, error)
}

func RegisterLotteryAPIServer(s *grpc.Server, srv LotteryAPIServer) {
	s.RegisterService(&_LotteryAPI_serviceDesc, srv)
}

func _LotteryAPI_NextLotteryInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextLotteryInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LotteryAPIServer).NextLotteryInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.LotteryAPI/NextLotteryInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LotteryAPIServer).NextLotteryInfo(ctx, req.(*NextLotteryInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LotteryAPI_SendLotteryFx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendLotteryFxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LotteryAPIServer).SendLotteryFx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.LotteryAPI/SendLotteryFx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LotteryAPIServer).SendLotteryFx(ctx, req.(*SendLotteryFxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LotteryAPI_SendLotteryLx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendLotteryLxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LotteryAPIServer).SendLotteryLx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.LotteryAPI/SendLotteryLx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LotteryAPIServer).SendLotteryLx(ctx, req.(*SendLotteryLxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LotteryAPI_StartLottery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartLotteryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LotteryAPIServer).StartLottery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.LotteryAPI/StartLottery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LotteryAPIServer).StartLottery(ctx, req.(*StartLotteryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LotteryAPI_GetLotteryResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLotteryResultReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LotteryAPIServer).GetLotteryResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.LotteryAPI/GetLotteryResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LotteryAPIServer).GetLotteryResult(ctx, req.(*GetLotteryResultReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _LotteryAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.LotteryAPI",
	HandlerType: (*LotteryAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NextLotteryInfo",
			Handler:    _LotteryAPI_NextLotteryInfo_Handler,
		},
		{
			MethodName: "SendLotteryFx",
			Handler:    _LotteryAPI_SendLotteryFx_Handler,
		},
		{
			MethodName: "SendLotteryLx",
			Handler:    _LotteryAPI_SendLotteryLx_Handler,
		},
		{
			MethodName: "StartLottery",
			Handler:    _LotteryAPI_StartLottery_Handler,
		},
		{
			MethodName: "GetLotteryResult",
			Handler:    _LotteryAPI_GetLotteryResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor4,
}

func init() { proto.RegisterFile("teller.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x6d, 0xd3, 0xd6, 0xcd, 0x4b, 0x4b, 0xbb, 0xb3, 0x2b, 0x86, 0x28, 0xb2, 0x0c, 0x88,
	0x8b, 0xc2, 0x0a, 0xf5, 0xe4, 0x49, 0x54, 0x56, 0x29, 0x0e, 0x45, 0xb6, 0x55, 0x54, 0xf0, 0x10,
	0xcd, 0x54, 0x82, 0xd3, 0xa4, 0xce, 0xcc, 0x6a, 0xfc, 0xdb, 0xbd, 0x38, 0x33, 0xf9, 0x61, 0x3b,
	0x9d, 0xd0, 0x8b, 0xa7, 0xc2, 0xfb, 0xf1, 0x79, 0xdf, 0xef, 0x9b, 0xd7, 0xc0, 0x50, 0x52, 0xc6,
	0x28, 0xbf, 0xd8, 0xf0, 0x5c, 0xe6, 0x68, 0x60, 0x7e, 0x44, 0x14, 0x50, 0xce, 0xf3, 0x2a, 0x88,
	0x4f, 0x01, 0xcd, 0x69, 0x21, 0x49, 0x2e, 0x25, 0xe5, 0xbf, 0x67, 0xd9, 0x2a, 0xbf, 0xa2, 0x3f,
	0xf0, 0xfb, 0xfd, 0xa8, 0xd8, 0xa0, 0xbb, 0xd0, 0x37, 0xad, 0x61, 0xe7, 0xac, 0x73, 0x1e, 0x4c,
	0x47, 0x25, 0x42, 0x5c, 0x5c, 0xea, 0x20, 0x3a, 0x06, 0x5f, 0xc8, 0x98, 0xcb, 0x65, 0xba, 0xa6,
	0x61, 0x57, 0x55, 0x78, 0x68, 0x0c, 0x37, 0x69, 0x96, 0x98, 0x80, 0xa7, 0x03, 0xf8, 0x0a, 0xfc,
	0x8a, 0xf9, 0xaa, 0x40, 0x01, 0x78, 0xab, 0x34, 0x31, 0x30, 0x1f, 0x8d, 0xa0, 0xff, 0x33, 0x66,
	0xd7, 0x65, 0x67, 0x0f, 0x01, 0x74, 0xd7, 0xdc, 0x34, 0xf5, 0x34, 0x85, 0xd1, 0xe4, 0x1b, 0xe5,
	0x22, 0xec, 0x9d, 0x79, 0xaa, 0x76, 0x08, 0xbd, 0x24, 0x15, 0x32, 0xec, 0xeb, 0x34, 0x5e, 0x34,
	0x4c, 0x62, 0x98, 0xac, 0x8d, 0x59, 0xb7, 0x95, 0x54, 0x55, 0xf9, 0x2b, 0xcf, 0x14, 0xb1, 0x73,
	0x7e, 0xa4, 0x47, 0xac, 0x62, 0xbe, 0xd6, 0x23, 0xfa, 0x7a, 0x04, 0xfe, 0x00, 0xe3, 0x46, 0xe8,
	0x32, 0xfd, 0xfa, 0x9d, 0xca, 0x5d, 0xb9, 0x4a, 0xdf, 0xaa, 0x70, 0x68, 0x55, 0x45, 0x69, 0x52,
	0x18, 0xaa, 0x87, 0x4e, 0x20, 0x60, 0x25, 0x64, 0x1e, 0xab, 0x15, 0x68, 0xb9, 0x3e, 0x7e, 0x04,
	0x93, 0x85, 0x5a, 0x4a, 0x43, 0x57, 0xeb, 0x6e, 0x45, 0xe3, 0x8f, 0x76, 0xf1, 0xc1, 0x57, 0x78,
	0x00, 0x03, 0x69, 0xf4, 0x1a, 0x42, 0x30, 0xbd, 0x5d, 0xa7, 0x2d, 0x3b, 0xf8, 0x59, 0xe3, 0x90,
	0x6c, 0x39, 0x64, 0xdb, 0x32, 0x58, 0xed, 0xd0, 0x32, 0xe2, 0x39, 0x8c, 0x90, 0xda, 0x88, 0x93,
	0x60, 0x19, 0x21, 0xff, 0xc1, 0x48, 0xad, 0x1a, 0x3f, 0x85, 0xf1, 0x42, 0xdf, 0x5d, 0x15, 0xd7,
	0x32, 0x26, 0x70, 0x64, 0x4e, 0xf1, 0xdd, 0xf2, 0xa5, 0x81, 0x7b, 0xe8, 0x14, 0x86, 0x2c, 0x16,
	0x72, 0x96, 0xa9, 0x12, 0x75, 0x13, 0x86, 0xe9, 0xe3, 0xc7, 0x56, 0xeb, 0x21, 0x51, 0xf8, 0x21,
	0x9c, 0xbc, 0xa6, 0xff, 0x26, 0x89, 0x6b, 0x26, 0xf5, 0x3c, 0x6b, 0x3f, 0xc6, 0x3e, 0xfe, 0x0c,
	0xa3, 0x9d, 0x42, 0x7d, 0x7f, 0x59, 0x93, 0x46, 0xf7, 0xd4, 0x9b, 0x17, 0x42, 0x09, 0xf1, 0xd4,
	0x98, 0xe3, 0xbd, 0x57, 0xd2, 0x79, 0xa6, 0xf2, 0x9e, 0x33, 0x4f, 0x0a, 0xfc, 0xc9, 0x21, 0xe5,
	0xe0, 0x52, 0xef, 0xc3, 0x80, 0x9b, 0xd2, 0x6a, 0xa9, 0xb7, 0x2c, 0x6e, 0xc9, 0x99, 0xfe, 0xe9,
	0x02, 0x54, 0x91, 0xe7, 0x6f, 0x67, 0xe8, 0x0d, 0x8c, 0xad, 0xaf, 0x01, 0x8a, 0xea, 0xc6, 0xfd,
	0x8f, 0x47, 0xd4, 0x9a, 0x13, 0x1b, 0x7c, 0x03, 0x5d, 0xc2, 0x68, 0xe7, 0xa4, 0x51, 0x58, 0x97,
	0xdb, 0x7f, 0x8b, 0xa8, 0x25, 0xe3, 0xc0, 0x10, 0x37, 0x86, 0xb4, 0x62, 0x48, 0x85, 0x79, 0x01,
	0xc3, 0xed, 0x0b, 0x40, 0xcd, 0x95, 0x59, 0x27, 0x15, 0xb9, 0x13, 0x86, 0x31, 0x87, 0x89, 0xfd,
	0x12, 0xe8, 0x4e, 0x5d, 0xee, 0x38, 0x97, 0xa8, 0x3d, 0xa9, 0x79, 0x5f, 0xca, 0xef, 0xf4, 0x93,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae, 0xdd, 0x2e, 0x28, 0xbe, 0x05, 0x00, 0x00,
}
