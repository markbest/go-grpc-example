package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/markbest/go-grpc-example/server/protos"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//article client
	c0 := pb.NewArticleClient(conn)
	r0, err0 := c0.GetArticleInfo(context.Background(), &pb.AQueryRequest{Id: 10})
	if err0 != nil {
		log.Fatal(err0)
	}
	log.Print(r0.Title)

	r1, err1 := c0.GetArticleList(context.Background(), &pb.AQueryRequest{Page: 1, Limit: 20})
	if err1 != nil {
		log.Fatal(err1)
	}
	log.Print(len(r1.List))

	//category client
	c1 := pb.NewCategoryClient(conn)
	r2, err2 := c1.GetCategoryInfo(context.Background(), &pb.CQueryRequest{Id: 2})
	if err2 != nil {
		log.Fatal(err2)
	}
	log.Print(r2.Title)

	r3, err3 := c1.GetCategoryList(context.Background(), &pb.CQueryRequest{Page: 1, Limit: 20})
	if err3 != nil {
		log.Fatal(err3)
	}
	log.Print(len(r3.List))
}
