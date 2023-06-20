package main

import (
	"context"
	"ex00/pkg/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("FIRST ERROR")

	clinet := proto.NewTransmitterClient(cc)
	resp, err := clinet.Connection(context.Background(), &proto.Request{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
