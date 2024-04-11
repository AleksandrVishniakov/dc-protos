// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: orchestrator/v1/orchestrator.proto

package orchestratorv1

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

// OrchestratorClient is the client API for Orchestrator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrchestratorClient interface {
	RegisterWorker(ctx context.Context, in *WorkerRegisterRequest, opts ...grpc.CallOption) (*WorkerRegisterResponse, error)
	StartTask(ctx context.Context, in *TaskStartingRequest, opts ...grpc.CallOption) (*TaskStartingResponse, error)
	SendTakResult(ctx context.Context, in *TaskResultRequest, opts ...grpc.CallOption) (*TaskResultResponse, error)
}

type orchestratorClient struct {
	cc grpc.ClientConnInterface
}

func NewOrchestratorClient(cc grpc.ClientConnInterface) OrchestratorClient {
	return &orchestratorClient{cc}
}

func (c *orchestratorClient) RegisterWorker(ctx context.Context, in *WorkerRegisterRequest, opts ...grpc.CallOption) (*WorkerRegisterResponse, error) {
	out := new(WorkerRegisterResponse)
	err := c.cc.Invoke(ctx, "/orchestrator.Orchestrator/RegisterWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorClient) StartTask(ctx context.Context, in *TaskStartingRequest, opts ...grpc.CallOption) (*TaskStartingResponse, error) {
	out := new(TaskStartingResponse)
	err := c.cc.Invoke(ctx, "/orchestrator.Orchestrator/StartTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorClient) SendTakResult(ctx context.Context, in *TaskResultRequest, opts ...grpc.CallOption) (*TaskResultResponse, error) {
	out := new(TaskResultResponse)
	err := c.cc.Invoke(ctx, "/orchestrator.Orchestrator/SendTakResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrchestratorServer is the server API for Orchestrator service.
// All implementations must embed UnimplementedOrchestratorServer
// for forward compatibility
type OrchestratorServer interface {
	RegisterWorker(context.Context, *WorkerRegisterRequest) (*WorkerRegisterResponse, error)
	StartTask(context.Context, *TaskStartingRequest) (*TaskStartingResponse, error)
	SendTakResult(context.Context, *TaskResultRequest) (*TaskResultResponse, error)
	mustEmbedUnimplementedOrchestratorServer()
}

// UnimplementedOrchestratorServer must be embedded to have forward compatible implementations.
type UnimplementedOrchestratorServer struct {
}

func (UnimplementedOrchestratorServer) RegisterWorker(context.Context, *WorkerRegisterRequest) (*WorkerRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterWorker not implemented")
}
func (UnimplementedOrchestratorServer) StartTask(context.Context, *TaskStartingRequest) (*TaskStartingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartTask not implemented")
}
func (UnimplementedOrchestratorServer) SendTakResult(context.Context, *TaskResultRequest) (*TaskResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTakResult not implemented")
}
func (UnimplementedOrchestratorServer) mustEmbedUnimplementedOrchestratorServer() {}

// UnsafeOrchestratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrchestratorServer will
// result in compilation errors.
type UnsafeOrchestratorServer interface {
	mustEmbedUnimplementedOrchestratorServer()
}

func RegisterOrchestratorServer(s grpc.ServiceRegistrar, srv OrchestratorServer) {
	s.RegisterService(&Orchestrator_ServiceDesc, srv)
}

func _Orchestrator_RegisterWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).RegisterWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orchestrator.Orchestrator/RegisterWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).RegisterWorker(ctx, req.(*WorkerRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestrator_StartTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskStartingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).StartTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orchestrator.Orchestrator/StartTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).StartTask(ctx, req.(*TaskStartingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestrator_SendTakResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).SendTakResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orchestrator.Orchestrator/SendTakResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).SendTakResult(ctx, req.(*TaskResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Orchestrator_ServiceDesc is the grpc.ServiceDesc for Orchestrator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Orchestrator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orchestrator.Orchestrator",
	HandlerType: (*OrchestratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterWorker",
			Handler:    _Orchestrator_RegisterWorker_Handler,
		},
		{
			MethodName: "StartTask",
			Handler:    _Orchestrator_StartTask_Handler,
		},
		{
			MethodName: "SendTakResult",
			Handler:    _Orchestrator_SendTakResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orchestrator/v1/orchestrator.proto",
}
