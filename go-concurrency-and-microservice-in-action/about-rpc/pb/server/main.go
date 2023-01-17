package main

import (
	"flag"
	"log"
	"net"

	"github.com/lvdbing/books/go-concurrency-and-microservice-in-action/about-rpc/pb"
	stringservice "github.com/lvdbing/books/go-concurrency-and-microservice-in-action/about-rpc/pb/string-service"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	stringService := new(stringservice.StringService)
	pb.RegisterStringServiceServer(grpcServer, stringService)
	grpcServer.Serve(lis)
}
