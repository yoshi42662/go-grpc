package main

import (
	"context"
	"fmt"

	pb "client/pb"
	"google.golang.org/grpc"
	"time"
)

func main() {
	host := "server"
	port := "8080"

	serverAddr := fmt.Sprintf("%s:%s", host, port)

	for {
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

		time.Sleep(3 * time.Second)
	}
}
