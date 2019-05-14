package api

import (
	fmt "fmt"
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message %s", in.Msg)
	return &PingMessage{Msg: fmt.Sprintf("Hi %s", in.Msg)}, nil
}
