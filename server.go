package main

import (
	"encoding/json"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	. "grpc/model"
	pb "grpc/server/protos"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) GetArticleInfo(ctx context.Context, in *pb.Request) (*pb.ArticleInfoResponse, error) {
	var a Article
	var rs pb.ArticleInfoResponse

	a = GetArticleInfo(in.Id)
	json_content, _ := json.Marshal(a)
	json.Unmarshal(json_content, &rs)
	return &rs, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterArticleServer(s, &server{})
	s.Serve(lis)
}
