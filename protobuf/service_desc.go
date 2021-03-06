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

func loginHandler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
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

func signUpHandler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserSignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ogiri.Public/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserSignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var UserServicedesc = grpc.ServiceDesc{
	ServiceName: "ogiri.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    loginHandler,
		},
		{
			MethodName: "SignUp",
			Handler: signUpHandler,

		},
	},
	Streams: []grpc.StreamDesc{},
	Metadata: "api.proto",
}
