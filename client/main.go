package main

import (
	"context"
	"flag"
	"log"
	"time"

	"go-grpc/configs"
	pb "go-grpc/libs"

	"google.golang.org/grpc"
)

var (
	check = flag.Bool("check", false, "Server Testing")
)

func main() {
	flag.Parse()
	configs.StartInfo(false)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, configs.BindAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloGRPCClient(conn)

	if *check {
		r, err := c.ServerCheck(ctx, &pb.Ping{Hello: "Hello"})
		if err != nil {
			log.Fatalf("could not connect: %v", err)
		}
		if r.GetAlive() {
			log.Println(r.GetMessage())
		} else {
			log.Println("Server is Down")
		}
	}
}
