package main

import (
	"context"
	"fmt"

	"github.com/raelzeraxd/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := pb.NewGreeterClient(conn)

	req := &pb.HelloRequest{
		Name: "John",
	}
	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Message)
}
