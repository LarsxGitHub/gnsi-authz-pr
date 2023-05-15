// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: github.com/openconfig/gnsi/accounting/acct.proto

package accounting

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

// AccountingPullClient is the client API for AccountingPull service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountingPullClient interface {
	RecordStream(ctx context.Context, opts ...grpc.CallOption) (AccountingPull_RecordStreamClient, error)
}

type accountingPullClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountingPullClient(cc grpc.ClientConnInterface) AccountingPullClient {
	return &accountingPullClient{cc}
}

func (c *accountingPullClient) RecordStream(ctx context.Context, opts ...grpc.CallOption) (AccountingPull_RecordStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &AccountingPull_ServiceDesc.Streams[0], "/gnsi.accounting.v1.AccountingPull/RecordStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &accountingPullRecordStreamClient{stream}
	return x, nil
}

type AccountingPull_RecordStreamClient interface {
	Send(*RecordSync) error
	Recv() (*Record, error)
	grpc.ClientStream
}

type accountingPullRecordStreamClient struct {
	grpc.ClientStream
}

func (x *accountingPullRecordStreamClient) Send(m *RecordSync) error {
	return x.ClientStream.SendMsg(m)
}

func (x *accountingPullRecordStreamClient) Recv() (*Record, error) {
	m := new(Record)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AccountingPullServer is the server API for AccountingPull service.
// All implementations must embed UnimplementedAccountingPullServer
// for forward compatibility
type AccountingPullServer interface {
	RecordStream(AccountingPull_RecordStreamServer) error
	mustEmbedUnimplementedAccountingPullServer()
}

// UnimplementedAccountingPullServer must be embedded to have forward compatible implementations.
type UnimplementedAccountingPullServer struct {
}

func (UnimplementedAccountingPullServer) RecordStream(AccountingPull_RecordStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordStream not implemented")
}
func (UnimplementedAccountingPullServer) mustEmbedUnimplementedAccountingPullServer() {}

// UnsafeAccountingPullServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountingPullServer will
// result in compilation errors.
type UnsafeAccountingPullServer interface {
	mustEmbedUnimplementedAccountingPullServer()
}

func RegisterAccountingPullServer(s grpc.ServiceRegistrar, srv AccountingPullServer) {
	s.RegisterService(&AccountingPull_ServiceDesc, srv)
}

func _AccountingPull_RecordStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AccountingPullServer).RecordStream(&accountingPullRecordStreamServer{stream})
}

type AccountingPull_RecordStreamServer interface {
	Send(*Record) error
	Recv() (*RecordSync, error)
	grpc.ServerStream
}

type accountingPullRecordStreamServer struct {
	grpc.ServerStream
}

func (x *accountingPullRecordStreamServer) Send(m *Record) error {
	return x.ServerStream.SendMsg(m)
}

func (x *accountingPullRecordStreamServer) Recv() (*RecordSync, error) {
	m := new(RecordSync)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AccountingPull_ServiceDesc is the grpc.ServiceDesc for AccountingPull service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountingPull_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnsi.accounting.v1.AccountingPull",
	HandlerType: (*AccountingPullServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RecordStream",
			Handler:       _AccountingPull_RecordStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/openconfig/gnsi/accounting/acct.proto",
}

// AccountingPushClient is the client API for AccountingPush service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountingPushClient interface {
	RecordStream(ctx context.Context, opts ...grpc.CallOption) (AccountingPush_RecordStreamClient, error)
}

type accountingPushClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountingPushClient(cc grpc.ClientConnInterface) AccountingPushClient {
	return &accountingPushClient{cc}
}

func (c *accountingPushClient) RecordStream(ctx context.Context, opts ...grpc.CallOption) (AccountingPush_RecordStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &AccountingPush_ServiceDesc.Streams[0], "/gnsi.accounting.v1.AccountingPush/RecordStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &accountingPushRecordStreamClient{stream}
	return x, nil
}

type AccountingPush_RecordStreamClient interface {
	Send(*Record) error
	Recv() (*RecordSync, error)
	grpc.ClientStream
}

type accountingPushRecordStreamClient struct {
	grpc.ClientStream
}

func (x *accountingPushRecordStreamClient) Send(m *Record) error {
	return x.ClientStream.SendMsg(m)
}

func (x *accountingPushRecordStreamClient) Recv() (*RecordSync, error) {
	m := new(RecordSync)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AccountingPushServer is the server API for AccountingPush service.
// All implementations must embed UnimplementedAccountingPushServer
// for forward compatibility
type AccountingPushServer interface {
	RecordStream(AccountingPush_RecordStreamServer) error
	mustEmbedUnimplementedAccountingPushServer()
}

// UnimplementedAccountingPushServer must be embedded to have forward compatible implementations.
type UnimplementedAccountingPushServer struct {
}

func (UnimplementedAccountingPushServer) RecordStream(AccountingPush_RecordStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordStream not implemented")
}
func (UnimplementedAccountingPushServer) mustEmbedUnimplementedAccountingPushServer() {}

// UnsafeAccountingPushServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountingPushServer will
// result in compilation errors.
type UnsafeAccountingPushServer interface {
	mustEmbedUnimplementedAccountingPushServer()
}

func RegisterAccountingPushServer(s grpc.ServiceRegistrar, srv AccountingPushServer) {
	s.RegisterService(&AccountingPush_ServiceDesc, srv)
}

func _AccountingPush_RecordStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AccountingPushServer).RecordStream(&accountingPushRecordStreamServer{stream})
}

type AccountingPush_RecordStreamServer interface {
	Send(*RecordSync) error
	Recv() (*Record, error)
	grpc.ServerStream
}

type accountingPushRecordStreamServer struct {
	grpc.ServerStream
}

func (x *accountingPushRecordStreamServer) Send(m *RecordSync) error {
	return x.ServerStream.SendMsg(m)
}

func (x *accountingPushRecordStreamServer) Recv() (*Record, error) {
	m := new(Record)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AccountingPush_ServiceDesc is the grpc.ServiceDesc for AccountingPush service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountingPush_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnsi.accounting.v1.AccountingPush",
	HandlerType: (*AccountingPushServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RecordStream",
			Handler:       _AccountingPush_RecordStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/openconfig/gnsi/accounting/acct.proto",
}