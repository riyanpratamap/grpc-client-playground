package main

import (
	"context"
	"log"

	"github.com/riyanpratamap/grpc-client-playground/chat"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Erro dial grpc", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from client!",
	}
	response, err2 := c.SayHello(context.Background(), &message)
	if err2 != nil {
		log.Printf("Error when calling SayHello: ", err2)
	}

	log.Printf("Response from server: %s", response.Body)

}
