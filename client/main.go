package main

import (
	"bufio"
	"context"
	"log"
	"os"

	"github.com/kimbbakar/simple-go-grpc-example/api"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewPingClient(conn)

	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')

		response, err := c.SayHello(context.Background(), &api.PingMessage{Msg: text})
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}
		log.Printf("Response from server: %s", response.Msg)
	}
}
