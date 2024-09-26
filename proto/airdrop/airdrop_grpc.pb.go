// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: savour_rpc/airdrop.proto

package airdrop

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AirdropService_SubmitDppLinkPoints_FullMethodName = "/services.savour_rpc.airdrop.AirdropService/submitDppLinkPoints"
)

// AirdropServiceClient is the client API for AirdropService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AirdropServiceClient interface {
	SubmitDppLinkPoints(ctx context.Context, in *DppLinkPointsReq, opts ...grpc.CallOption) (*DppLinkPointsRep, error)
}

type airdropServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAirdropServiceClient(cc grpc.ClientConnInterface) AirdropServiceClient {
	return &airdropServiceClient{cc}
}

func (c *airdropServiceClient) SubmitDppLinkPoints(ctx context.Context, in *DppLinkPointsReq, opts ...grpc.CallOption) (*DppLinkPointsRep, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DppLinkPointsRep)
	err := c.cc.Invoke(ctx, AirdropService_SubmitDppLinkPoints_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AirdropServiceServer is the server API for AirdropService service.
// All implementations should embed UnimplementedAirdropServiceServer
// for forward compatibility.
type AirdropServiceServer interface {
	SubmitDppLinkPoints(context.Context, *DppLinkPointsReq) (*DppLinkPointsRep, error)
}

// UnimplementedAirdropServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAirdropServiceServer struct{}

func (UnimplementedAirdropServiceServer) SubmitDppLinkPoints(context.Context, *DppLinkPointsReq) (*DppLinkPointsRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitDppLinkPoints not implemented")
}
func (UnimplementedAirdropServiceServer) testEmbeddedByValue() {}

// UnsafeAirdropServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AirdropServiceServer will
// result in compilation errors.
type UnsafeAirdropServiceServer interface {
	mustEmbedUnimplementedAirdropServiceServer()
}

func RegisterAirdropServiceServer(s grpc.ServiceRegistrar, srv AirdropServiceServer) {
	// If the following call pancis, it indicates UnimplementedAirdropServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AirdropService_ServiceDesc, srv)
}

func _AirdropService_SubmitDppLinkPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DppLinkPointsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AirdropServiceServer).SubmitDppLinkPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AirdropService_SubmitDppLinkPoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AirdropServiceServer).SubmitDppLinkPoints(ctx, req.(*DppLinkPointsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AirdropService_ServiceDesc is the grpc.ServiceDesc for AirdropService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AirdropService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.savour_rpc.airdrop.AirdropService",
	HandlerType: (*AirdropServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "submitDppLinkPoints",
			Handler:    _AirdropService_SubmitDppLinkPoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "savour_rpc/airdrop.proto",
}
