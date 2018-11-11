package main

import (
	"context"
	"log"

	"github.com/jakhax/grpc-golang-example/api"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Unable to connect to server: %v", err)
	}
	defer conn.Close()

	c := api.NewPingClient(conn)

	res, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "Hello world Server"})
	if err != nil {
		log.Fatal("Error occurred when sending message: %v", err)
	}
	log.Printf("Response: %s", res.Greeting)
}
