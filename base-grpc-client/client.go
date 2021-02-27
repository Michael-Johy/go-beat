package main

import (
	"context"
	basketBallPlayer "example.com/user/hello/base-grpc-proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	sAddress := "127.0.0.1:50051"

	conn, e := grpc.Dial(sAddress, grpc.WithInsecure())

	if e != nil {
		log.Fatalf("failed to dial %v", sAddress)
	}

	defer conn.Close()

	client := basketBallPlayer.NewPlayerServiceClient(conn)

	player, e := client.GetBasketBallPlayer(context.Background(), &basketBallPlayer.PlayerRequest{
		Id: "1",
	})

	if e != nil {
		log.Fatalf("failed to get player for param %v ", 1)
	}

	log.Println(player)

}
