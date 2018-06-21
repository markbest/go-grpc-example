package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	. "github.com/markbest/go-grpc-example/conf"
	pb "github.com/markbest/go-grpc-example/protos"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(Conf.App.GrpcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// article client
	articleClient := pb.NewArticleClient(conn)
	ars, err := articleClient.GetArticleInfo(context.Background(), &pb.AQueryRequest{ArticleId: 10})
	if err != nil {
		log.Fatal(err)
	}
	log.Print(ars.Title)

	ars1, err := articleClient.GetArticleList(context.Background(), &pb.AQueryRequest{Page: 1, Limit: 20})
	if err != nil {
		log.Fatal(err)
	}
	log.Print(len(ars1.Data))

	// category client
	categoryClient := pb.NewCategoryClient(conn)
	crs, err := categoryClient.GetCategoryInfo(context.Background(), &pb.CQueryRequest{CatId: 2})
	if err != nil {
		log.Fatal(err)
	}
	log.Print(crs.Title)

	crs1, err := categoryClient.GetCategoryList(context.Background(), &pb.CQueryRequest{Page: 1, Limit: 20})
	if err != nil {
		log.Fatal(err)
	}
	log.Print(len(crs1.Data))
}
