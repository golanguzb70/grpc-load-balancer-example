package main

import (
	"log"
	"net"

	"github.com/golanguzb70/grpc-load-balancer-example/server/config"
	pb "github.com/golanguzb70/grpc-load-balancer-example/server/genproto/post_service"
	"github.com/golanguzb70/grpc-load-balancer-example/server/service"
	"google.golang.org/grpc"
)

func main() {
	config := config.Load()

	postService := service.NewPostService()

	lis, err := net.Listen("tcp", ":"+config.GrpcPort)
	if err != nil {
		log.Printf("Error while listening: %v\n", err)
		return
	}

	server := grpc.NewServer()
	pb.RegisterPostServiceServer(server, postService)

	log.Println("gRPC server is running on port :" + config.GrpcPort)

	err = server.Serve(lis)
	if err != nil {
		log.Println("Error while running gRPC server", err)
	}
}
