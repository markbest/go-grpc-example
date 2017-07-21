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
	r, err := c.GetArticleList(context.Background(), &pb.QueryRequest{Page: 1, Limit:20})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Print(len(r.List))
}
