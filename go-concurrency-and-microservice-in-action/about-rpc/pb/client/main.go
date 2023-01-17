package main

import (
	"context"
	"fmt"

	"github.com/lvdbing/books/go-concurrency-and-microservice-in-action/about-rpc/pb"
	"google.golang.org/grpc"
)

func main() {
	serviceAddress := "127.0.0.1:1234"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()

	stringClient := pb.NewStringServiceClient(conn)
	stringReq := &pb.StringRequst{A: "A", B: "B"}
	reply, _ := stringClient.Concat(context.Background(), stringReq)
	fmt.Printf("StringService Concat: %s concat %s = %s\n", stringReq.A, stringReq.B, reply.Ret)

	stringReq = &pb.StringRequst{A: "ABCDE", B: "CDEFG"}
	reply, _ = stringClient.Diff(context.Background(), stringReq)
	fmt.Printf("StringService Diff: %s diff %s = %s\n", stringReq.A, stringReq.B, reply.Ret)
}
