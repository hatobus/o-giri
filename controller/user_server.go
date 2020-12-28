package controller

import (
	"context"
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/hatobus/o-giri/infrastructure/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"

	pb "github.com/hatobus/o-giri/protobuf"
)

var (
	InvalidUserNameLength = "ユーザー名は3文字以上32文字以下である必要があります"
	InvalidPasswordLength = "パスワードは8文字以上64文字以下である必要があります"
	InvalidDuplicatedUserName = "%v はすでに登録されているユーザー名です"

	InternalServerError = "サーバー側のエラーです 運営にお問い合わせください"
)

type UserServer struct {
	DB *sql.DB
}

func NewUserServer(db *sql.DB) *UserServer {
	return &UserServer{db}
}

func (u *UserServer) UserSignUp(_ context.Context, signUpReq *pb.SignUpRequest) (*pb.Empty, error) {
	if len(signUpReq.Username) < 3 || len(signUpReq.Username) > 16 {
		return &pb.Empty{}, status.Error(codes.InvalidArgument, InvalidUserNameLength)
	}

	if len(signUpReq.Password) < 8 || len(signUpReq.Password) > 64 {
		return &pb.Empty{}, status.Error(codes.InvalidArgument, InvalidPasswordLength)
	}

	if u, err := database.UserByName(u.DB, signUpReq.Username); u != nil || err != nil {
		if u != nil {
			return &pb.Empty{}, status.Error(codes.InvalidArgument, fmt.Sprintf(InvalidDuplicatedUserName, signUpReq.Username))
		}
		if err != nil {
			log.Println(err)
			return &pb.Empty{}, status.Error(codes.Unknown, InternalServerError)
		}
	}

	encPW := sha512.Sum512([]byte(signUpReq.Password))

	newUser := &database.User{
		Name: signUpReq.Username,
		Password: hex.EncodeToString(encPW[:]),
	}

	if err := newUser.Insert(u.DB); err != nil {
		log.Println(err)
		return &pb.Empty{}, status.Error(codes.Unknown, InternalServerError)
	}

	return &pb.Empty{}, nil
}

func (u *UserServer) UserLogin(c context.Context, loginReq *pb.LoginRequest) (*pb.LoginRes, error) {
	panic("not impl")
}
