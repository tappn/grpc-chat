package main

import (
	"log"
	"net"

	"github.com/tappn/grpc-chat/gen/go/room/v1"
	"github.com/tappn/grpc-chat/handler"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	// todo server option
	grpcServer := grpc.NewServer(opts...)
	room.RegisterRoomServiceServer(grpcServer, handler.NewRoomHandler())
	log.Println("------ start grpc server ------")
	grpcServer.Serve(listen)
}
