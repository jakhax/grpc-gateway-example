package main

import (
	"fmt"
	"log"
)

func main() {

	grpcAddress := fmt.Sprintf("%s:%d", "localhost", 7777)

	restAddress := fmt.Sprintf("%s:%d", "localhost", 7778)
	certFile := "cert/server.crt"

	keyFile := "cert/server.key"

	// fire the gRPC server in a goroutine
	go func() {
		err := startGRPCServer(grpcAddress, certFile, keyFile)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	// fire the REST proxy server in a goroutine
	go func() {
		err := startProxyRESTServer(restAddress, grpcAddress, certFile)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	log.Printf("Entering infinite loop")
	select {}

}
