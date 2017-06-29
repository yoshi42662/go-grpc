package main

import (
	"context"
	"fmt"

	pb "go-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	serverAddr := "127.0.0.1:8080"

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	client := pb.NewPingpongClient(conn)
	feature, err := client.Ping(context.Background(), &pb.PingRequest{
		Ping: "ping",
		Pong: "",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(feature)
}
