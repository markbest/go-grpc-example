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

func (s *server) GetArticleInfo(ctx context.Context, in *pb.QueryRequest) (*pb.ArticleInfo, error) {
	var a Article
	var rs pb.ArticleInfo

	a = GetArticleInfo(in.Id)
	json_content, _ := json.Marshal(a)
	json.Unmarshal(json_content, &rs)
	return &rs, nil
}

func (s *server) GetArticleList(ctx context.Context, in *pb.QueryRequest) (*pb.ArticleList, error) {
	var a []Article
	var rs []*pb.ArticleInfo
	var nrs pb.ArticleList

	a = GetArticleList(in.Page, in.Limit)
	json_content, _ := json.Marshal(a)
	json.Unmarshal(json_content, &rs)
	nrs.List = rs
	return &nrs, nil
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
