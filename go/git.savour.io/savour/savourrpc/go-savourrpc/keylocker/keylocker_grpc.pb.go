// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: savourrpc/keylocker.proto

package keylocker

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

// LeyLockerServiceClient is the client API for LeyLockerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LeyLockerServiceClient interface {
	GetSupportChain(ctx context.Context, in *SupportChainReq, opts ...grpc.CallOption) (*SupportChainRep, error)
	SetSocialKey(ctx context.Context, in *SetSocialKeyReq, opts ...grpc.CallOption) (*SetSocialKeyRep, error)
	GetSocialKey(ctx context.Context, in *GetSocialKeyReq, opts ...grpc.CallOption) (*GetSocialKeyRep, error)
}

type leyLockerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLeyLockerServiceClient(cc grpc.ClientConnInterface) LeyLockerServiceClient {
	return &leyLockerServiceClient{cc}
}

func (c *leyLockerServiceClient) GetSupportChain(ctx context.Context, in *SupportChainReq, opts ...grpc.CallOption) (*SupportChainRep, error) {
	out := new(SupportChainRep)
	err := c.cc.Invoke(ctx, "/savourrpc.keylocker.LeyLockerService/getSupportChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leyLockerServiceClient) SetSocialKey(ctx context.Context, in *SetSocialKeyReq, opts ...grpc.CallOption) (*SetSocialKeyRep, error) {
	out := new(SetSocialKeyRep)
	err := c.cc.Invoke(ctx, "/savourrpc.keylocker.LeyLockerService/setSocialKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leyLockerServiceClient) GetSocialKey(ctx context.Context, in *GetSocialKeyReq, opts ...grpc.CallOption) (*GetSocialKeyRep, error) {
	out := new(GetSocialKeyRep)
	err := c.cc.Invoke(ctx, "/savourrpc.keylocker.LeyLockerService/getSocialKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeyLockerServiceServer is the server API for LeyLockerService service.
// All implementations must embed UnimplementedLeyLockerServiceServer
// for forward compatibility
type LeyLockerServiceServer interface {
	GetSupportChain(context.Context, *SupportChainReq) (*SupportChainRep, error)
	SetSocialKey(context.Context, *SetSocialKeyReq) (*SetSocialKeyRep, error)
	GetSocialKey(context.Context, *GetSocialKeyReq) (*GetSocialKeyRep, error)
	mustEmbedUnimplementedLeyLockerServiceServer()
}

// UnimplementedLeyLockerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLeyLockerServiceServer struct {
}

func (UnimplementedLeyLockerServiceServer) GetSupportChain(context.Context, *SupportChainReq) (*SupportChainRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupportChain not implemented")
}
func (UnimplementedLeyLockerServiceServer) SetSocialKey(context.Context, *SetSocialKeyReq) (*SetSocialKeyRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSocialKey not implemented")
}
func (UnimplementedLeyLockerServiceServer) GetSocialKey(context.Context, *GetSocialKeyReq) (*GetSocialKeyRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSocialKey not implemented")
}
func (UnimplementedLeyLockerServiceServer) mustEmbedUnimplementedLeyLockerServiceServer() {}

// UnsafeLeyLockerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LeyLockerServiceServer will
// result in compilation errors.
type UnsafeLeyLockerServiceServer interface {
	mustEmbedUnimplementedLeyLockerServiceServer()
}

func RegisterLeyLockerServiceServer(s grpc.ServiceRegistrar, srv LeyLockerServiceServer) {
	s.RegisterService(&LeyLockerService_ServiceDesc, srv)
}

func _LeyLockerService_GetSupportChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportChainReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeyLockerServiceServer).GetSupportChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/savourrpc.keylocker.LeyLockerService/getSupportChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeyLockerServiceServer).GetSupportChain(ctx, req.(*SupportChainReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeyLockerService_SetSocialKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSocialKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeyLockerServiceServer).SetSocialKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/savourrpc.keylocker.LeyLockerService/setSocialKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeyLockerServiceServer).SetSocialKey(ctx, req.(*SetSocialKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeyLockerService_GetSocialKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSocialKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeyLockerServiceServer).GetSocialKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/savourrpc.keylocker.LeyLockerService/getSocialKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeyLockerServiceServer).GetSocialKey(ctx, req.(*GetSocialKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LeyLockerService_ServiceDesc is the grpc.ServiceDesc for LeyLockerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LeyLockerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "savourrpc.keylocker.LeyLockerService",
	HandlerType: (*LeyLockerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getSupportChain",
			Handler:    _LeyLockerService_GetSupportChain_Handler,
		},
		{
			MethodName: "setSocialKey",
			Handler:    _LeyLockerService_SetSocialKey_Handler,
		},
		{
			MethodName: "getSocialKey",
			Handler:    _LeyLockerService_GetSocialKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "savourrpc/keylocker.proto",
}
