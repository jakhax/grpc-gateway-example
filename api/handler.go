package api

import (
	context "context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {

	log.Printf("Message %s \n", in.Greeting)

	return &PingMessage{Greeting: "hello world"}, nil

}
