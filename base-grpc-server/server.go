package main

import (
	"context"
	basketBallPlayer "example.com/user/hello/base-grpc-proto"
	"fmt"
	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct{}

var c *cache.Cache

func (*server) GetBasketBallPlayer(ctx context.Context, in *basketBallPlayer.PlayerRequest) (*basketBallPlayer.PlayerResponse, error) {
	id := in.GetId()
	player, found := c.Get(id)
	if found {
		player := player.(basketBallPlayer.Player)
		return &basketBallPlayer.PlayerResponse{
			Result: &player,
		}, nil
	}
	return nil, fmt.Errorf("could not find player with id: %v", id)
}

func main() {
	fmt.Println("Starting gRPC micro-service...")
	c = cache.New(time.Minute*60, time.Minute*70)
	c.Set("1", basketBallPlayer.Player{
		Id:        "1",
		FirstName: "f1",
		LastName:  "l1",
		Age:       12,
		PhotoUrl:  "www.photo.url",
	}, cache.DefaultExpiration)

	l, e := net.Listen("tcp", ":50051")
	if e != nil {
		log.Fatalf("Failed to start listener %v", e)
	}

	s := grpc.NewServer()
	basketBallPlayer.RegisterPlayerServiceServer(s, &server{})

	if e := s.Serve(l); e != nil {
		log.Fatalf("failed to serve %v", e)
	}
}
