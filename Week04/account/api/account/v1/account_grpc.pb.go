// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package account

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountClient interface {
	GetUserById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*UserInfo, error)
}

type accountClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountClient(cc grpc.ClientConnInterface) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) GetUserById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/account.account/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
// All implementations must embed UnimplementedAccountServer
// for forward compatibility
type AccountServer interface {
	GetUserById(context.Context, *Id) (*UserInfo, error)
	mustEmbedUnimplementedAccountServer()
}

// UnimplementedAccountServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (UnimplementedAccountServer) GetUserById(context.Context, *Id) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedAccountServer) mustEmbedUnimplementedAccountServer() {}

// UnsafeAccountServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServer will
// result in compilation errors.
type UnsafeAccountServer interface {
	mustEmbedUnimplementedAccountServer()
}

func RegisterAccountServer(s grpc.ServiceRegistrar, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.account/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetUserById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserById",
			Handler:    _Account_GetUserById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}