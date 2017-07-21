package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc/server/protos"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewArticleClient(conn)
	r, err := c.GetArticleInfo(context.Background(), &pb.Request{Id: 9})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Print(r.Title)
}
