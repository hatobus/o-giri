package main

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/hatobus/o-giri/config"
	"github.com/hatobus/o-giri/infrastructure/database"
)

func main() {
	conf, err := config.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] faild to load config. err: %v", err)
		os.Exit(1)
	}

	_, err = database.Connect(conf.MySQL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] faild to connect database. err: %v", err)
		os.Exit(1)
	}

	s := grpc.NewServer()

	lis, err := net.Listen("tcp", conf.ServerPort)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] faild to open tcp port. err: %v", err)
		os.Exit(1)
	}

	if err := s.Serve(lis); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] faild to start server. err: %v", err)
		os.Exit(1)
	}
}
