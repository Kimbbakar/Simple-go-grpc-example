package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kimbbakar/simple-go-grpc-example/api"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer lis.Close()

	log.Print("Creating api server")
	s := api.Server{}
	server := grpc.NewServer()
	log.Print("binding grpc server with api server")
	api.RegisterPingServer(server, &s)
	log.Println("starting gRPC server...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
