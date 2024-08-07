package main

import (
	"context"
	"fmt"
	"github.com/raelzeraxd/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {	
	
	res := &pb.HelloReply{Message: "Hello from Golang grpc server to " + in.GetName()}
	
	return res, nil
}

func main() {
	fmt.Println("Running Golang gRPC server")
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}
	s  := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
