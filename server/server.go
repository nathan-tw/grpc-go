package main

import (
	"log"
	"net"

	"github.com/nathan-tw/grpc-go/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("Failde to listen on port 9000:", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	reflection.Register(grpcServer)
	log.Println("gRPC server is running.")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failde to serve gRPC server over port 9000:", err)
	} 
}
