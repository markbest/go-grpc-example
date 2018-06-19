package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	. "github.com/markbest/go-grpc-example/conf"
	. "github.com/markbest/go-grpc-example/controller"
	pb "github.com/markbest/go-grpc-example/protos"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", Conf.App.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterArticleServer(s, &ArticleServer{})
	pb.RegisterCategoryServer(s, &CategoryServer{})
	s.Serve(lis)
}
