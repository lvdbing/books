package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	aboutrpc "github.com/lvdbing/books/go-concurrency-and-microservice-in-action/about-rpc"
)

func main() {
	stringService := new(aboutrpc.StringService)
	rpc.Register(stringService)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(l, nil)
}
