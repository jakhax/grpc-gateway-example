package main

import (
	"context"
	"log"

	"github.com/jakhax/grpc-golang-example/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	var conn *grpc.ClientConn
	creds, err := credentials.NewClientTLSFromFile("/home/n1ght0wl/go/src/github.com/jakhax/grpc-golang-example/cert/server.crt", "")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	auth := Authentication{
		Login:    "user",
		Password: "password",
	}
	conn, err = grpc.Dial("localhost:7777", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
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
