package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jakhax/grpc-golang-example/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// server to listen on port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "localhost", 7777))
	if err != nil {
		log.Fatalf("Error Server failed to listen:%v", err)
	}
	// create server instance from  api.Server{} struct
	s := api.Server{}

	// @todo rem to load key path from env
	creds, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	opts := []grpc.ServerOption{grpc.Creds(creds)}

	// create a grpc server object, returns a pointer
	grpcServer := grpc.NewServer(opts...)
	// attach the api.Server{} services to the grpc object
	api.RegisterPingServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
