package controller

import (
	"context"
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/hatobus/o-giri/infrastructure/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/hatobus/o-giri/protobuf"
)

var (
	InvalidUserNameLength = "ユーザー名は3文字以上32文字以下である必要があります"
	InvalidPasswordLength = "パスワードは8文字以上64文字以下である必要があります"
	InvalidDuplicatedUserName = "%v はすでに登録されているユーザー名です"
	InvalidUserNameNotFound = "指定されたユーザーは見つかりませんでした"
	InvalidUserNameOrPassword = "ユーザー名またはパスワードが違います"

	InternalServerError = "サーバー側のエラーです 運営にお問い合わせください"
)

type UserServer struct {
	DB *sql.DB
	Salt string
}

func NewUserServer(db *sql.DB, hashSalt string) *UserServer {
	return &UserServer{db, hashSalt}
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
		if err != sql.ErrNoRows {
			log.Println(err)
			return &pb.Empty{}, status.Error(codes.Unknown, InternalServerError)
		}
	}

	encPW := sha512.Sum512([]byte(signUpReq.Password+u.Salt))

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

func (u *UserServer) UserLogin(_ context.Context, loginReq *pb.LoginRequest) (*pb.LoginRes, error) {
	if len(loginReq.Username) < 3 || len(loginReq.Username) > 16 {
		return &pb.LoginRes{}, status.Error(codes.InvalidArgument, InvalidUserNameLength)
	}

	if len(loginReq.Password) < 8 || len(loginReq.Password) > 64 {
		return &pb.LoginRes{}, status.Error(codes.InvalidArgument, InvalidPasswordLength)
	}

	user, err := database.UserByName(u.DB, loginReq.Username)
	if user == nil {
		return &pb.LoginRes{}, status.Errorf(codes.InvalidArgument, InvalidUserNameNotFound)
	}
	if err != nil {
		log.Println(err)
		return &pb.LoginRes{}, status.Errorf(codes.Unknown, InternalServerError)
	}

	encPW := sha512.Sum512([]byte(loginReq.Password+u.Salt))

	if user.Password != hex.EncodeToString(encPW[:]) {
		return &pb.LoginRes{}, status.Errorf(codes.InvalidArgument, InvalidUserNameOrPassword)
	}

	// TODO: tokenの作成処理
	var token string
	return &pb.LoginRes{Token: token}, nil
}
