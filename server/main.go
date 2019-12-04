package main

import (
	"github.com/qreasio/playgrpc/pkg/hrd"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	address = "localhost:4040"
)

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	server := &hrd.Server{}
	hrd.RegisterHumanResourceServer(grpcServer, server)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
}
