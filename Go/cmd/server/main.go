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
	pb.UnimplementedHelloServiceServer
}

func (s *Server) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {	
	
	res := &pb.HelloResponse{Msg: "Hello, " + in.GetName()}
	
	return res, nil
}

func main() {
	fmt.Println("Running Golang gRPC server")
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}
	s  := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
