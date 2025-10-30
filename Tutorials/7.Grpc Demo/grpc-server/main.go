package main

import (
	"log"
	"net"

	pb "grpc-demo/protos/greet"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("监听端口失败: %v", err)
	}

	s := grpc.NewServer()
	sr := &server{}
	pb.RegisterGreeterServer(s, sr)

	log.Println("gRPC 服务端启动，监听 :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}
