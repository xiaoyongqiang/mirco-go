package main

import (
	"fmt"
	"log"
	"net"
	"swingbaby-go/config"
	"swingbaby-go/rpcs/merpaytype"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/koding/multiconfig"
)

func main() {
	var err error
	m := multiconfig.New()
	config.Config = new(config.CmdConfig)
	err = m.Load(config.Config)
	if err != nil {
		log.Fatalf("Load configuration failed. Error: %s\n", err.Error())
	}
	m.MustLoad(config.Config)

	err = config.InitializeConn()
	if err != nil {
		log.Fatalf("config.InitialzeConn() failed. Error info: %s\n", err.Error())
	}
	defer func() {
		config.DBHandle.Close()
		config.RedisHandle.Close()
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Config.MerPayTypeConf.Host, config.Config.MerPayTypeConf.Port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v\n", err)
	}

	grpcServer := grpc.NewServer()
	merpaytype.RegisterMerPayConfServer(grpcServer, new(merpaytype.Service))
	grpcServer.Serve(lis)
}
