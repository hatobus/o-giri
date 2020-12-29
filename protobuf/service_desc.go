package ogiri

import (
	"context"
	"google.golang.org/grpc"
)

type UserServer interface {
	// SignUp
	// ユーザーの登録
	UserSignUp(context.Context, *SignUpRequest) (*Empty, error)

	// Login
	// ユーザーログイン
	UserLogin(context.Context, *LoginRequest) (*LoginRes, error)
}

func login_handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}

	info := &grpc.UnaryServerInfo{
		Server: srv,
		FullMethod: "/ogiri.Public/Login",
	}

	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserLogin(ctx, req.(*LoginRequest))
	}

	return interceptor(ctx, in, info, handler)
}

var User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ogiri.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler: login_handler,
		},
	},
	Streams: []grpc.StreamDesc{},
	Metadata: "api.proto",
}
