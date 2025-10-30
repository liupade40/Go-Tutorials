package main

import (
	"context"
	pb "grpc-demo/protos/greet"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "Hello " + req.Name + " from Go gRPC server",
	}, nil
}
