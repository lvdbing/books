package main

import (
	"fmt"
	"log"
	"net/rpc"

	aboutrpc "github.com/lvdbing/books/go-concurrency-and-microservice-in-action/about-rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	stringReq := &aboutrpc.StringRequest{"A", "B"}
	var reply string
	err = client.Call("StringService.Concat", stringReq, &reply)
	if err != nil {
		log.Fatal("Concat error:", err)
	}
	fmt.Printf("StringService Concat: %s concat %s =  %s\n", stringReq.A, stringReq.B, reply)

	stringReq = &aboutrpc.StringRequest{"ABCDE", "DEFG"}
	call := client.Go("StringService.Diff", stringReq, &reply, nil)
	<-call.Done
	fmt.Printf("StringService Diff: %s diff %s = %s\n", stringReq.A, stringReq.B, reply)
}
