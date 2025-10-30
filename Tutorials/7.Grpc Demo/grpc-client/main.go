package main

import (
	"context"
	"log"
	"time"

	pb "grpc-demo/protos/greet"

	"google.golang.org/grpc"
)

func main() {

	// 明文连接 gRPC 服务
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("连接 gRPC 失败: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Go Client"})
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	log.Printf("服务端返回: %s", resp.Message)
}
