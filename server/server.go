package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"

	pb "go-grpc/pb"
	"google.golang.org/grpc"
)

type pingpongServer struct {
}

func (p *pingpongServer) Ping(c context.Context, ping *pb.PingRequest) (*pb.PingResponse, error) {
  fmt.Println("Request recieved.")
	fmt.Println(ping.Ping)

	pong := pb.PingResponse{
		Ping: ping.Ping,
		Pong: "pong",
	}

	return &pong, nil
}

func main() {
	log.Println("Server is starting...")

	port := 8080
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPingpongServer(grpcServer, &pingpongServer{})
	grpcServer.Serve(lis)
}
