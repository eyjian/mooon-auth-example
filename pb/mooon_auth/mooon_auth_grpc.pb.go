// Written by yijian on 2024/01/21
// 注意本文中的 auth 为 authentication 的缩写，"鉴权"之意，而非 authorization 的缩写

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: mooon_auth.proto

package mooon_auth

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

const (
	MooonAuth_Authenticate_FullMethodName = "/mooon_auth.MooonAuth/Authenticate"
)

// MooonAuthClient is the client API for MooonAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MooonAuthClient interface {
	Authenticate(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthResp, error)
}

type mooonAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewMooonAuthClient(cc grpc.ClientConnInterface) MooonAuthClient {
	return &mooonAuthClient{cc}
}

func (c *mooonAuthClient) Authenticate(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthResp, error) {
	out := new(AuthResp)
	err := c.cc.Invoke(ctx, MooonAuth_Authenticate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MooonAuthServer is the server API for MooonAuth service.
// All implementations must embed UnimplementedMooonAuthServer
// for forward compatibility
type MooonAuthServer interface {
	Authenticate(context.Context, *AuthReq) (*AuthResp, error)
	mustEmbedUnimplementedMooonAuthServer()
}

// UnimplementedMooonAuthServer must be embedded to have forward compatible implementations.
type UnimplementedMooonAuthServer struct {
}

func (UnimplementedMooonAuthServer) Authenticate(context.Context, *AuthReq) (*AuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedMooonAuthServer) mustEmbedUnimplementedMooonAuthServer() {}

// UnsafeMooonAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MooonAuthServer will
// result in compilation errors.
type UnsafeMooonAuthServer interface {
	mustEmbedUnimplementedMooonAuthServer()
}

func RegisterMooonAuthServer(s grpc.ServiceRegistrar, srv MooonAuthServer) {
	s.RegisterService(&MooonAuth_ServiceDesc, srv)
}

func _MooonAuth_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MooonAuthServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MooonAuth_Authenticate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MooonAuthServer).Authenticate(ctx, req.(*AuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MooonAuth_ServiceDesc is the grpc.ServiceDesc for MooonAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MooonAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mooon_auth.MooonAuth",
	HandlerType: (*MooonAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _MooonAuth_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mooon_auth.proto",
}
