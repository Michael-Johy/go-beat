package main

import (
	"context"
	"example.com/user/hello/grpcs/proto"
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

	client := proto.NewPlayerServiceClient(conn)

	player, e := client.GetBasketBallPlayer(context.Background(), &proto.PlayerRequest{
		Id: "1",
	})

	if e != nil {
		log.Fatalf("failed to get player for param %v ", 1)
	}

	log.Println(player)

}
