package main

import (
	"context"
	"log"
	"net"

	"go-grpc/configs"
	pb "go-grpc/libs"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloGRPCServer
}

func (s *server) ServerCheck(ctx context.Context, in *pb.Ping) (*pb.Pong, error) {
	clientHello := in.GetHello()
	log.Printf("Received: %s", clientHello)
	return &pb.Pong{Message: "Server is alive", Alive: true}, nil
}

func main() {
	configs.StartInfo(true)
	listenAddress, err := net.Listen("tcp", configs.BindAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloGRPCServer(s, &server{})
	if err := s.Serve(listenAddress); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
