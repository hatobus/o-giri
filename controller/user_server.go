package controller

import (
	"context"

	pb "github.com/hatobus/o-giri/protobuf"
)

type userServer struct {

}

func (u *userServer) UserSignUp(c context.Context, signUpReq *pb.SignUpRequest) (*pb.Empty, error) {
	panic("not impl")
}

func (u *userServer) UserLogin(c context.Context, loginReq *pb.LoginRequest) (*pb.LoginRes, error) {
	panic("not impl")
}
