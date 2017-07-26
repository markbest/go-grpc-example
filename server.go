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

type articleServer struct{}

func (a *articleServer) GetArticleInfo(ctx context.Context, in *pb.AQueryRequest) (*pb.ArticleInfo, error) {
	var article Article
	var rs pb.ArticleInfo

	article = GetArticleInfo(in.Id)
	json_content, _ := json.Marshal(article)
	json.Unmarshal(json_content, &rs)
	return &rs, nil
}

func (a *articleServer) GetArticleList(ctx context.Context, in *pb.AQueryRequest) (*pb.ArticleList, error) {
	var articles []Article
	var rs []*pb.ArticleInfo
	var nrs pb.ArticleList

	articles = GetArticleList(in.Page, in.Limit)
	json_content, _ := json.Marshal(articles)
	json.Unmarshal(json_content, &rs)
	nrs.List = rs
	return &nrs, nil
}

type categoryServer struct{}

func (c *categoryServer) GetCategoryInfo(ctx context.Context, in *pb.CQueryRequest) (*pb.CategoryInfo, error) {
	var category Category
	var rs pb.CategoryInfo

	category = GetCategoryInfo(in.Id)
	json_content, _ := json.Marshal(category)
	json.Unmarshal(json_content, &rs)
	return &rs, nil
}

func (c *categoryServer) GetCategoryList(ctx context.Context, in *pb.CQueryRequest) (*pb.CategoryList, error) {
	var categories []Category
	var rs []*pb.CategoryInfo
	var nrs pb.CategoryList

	categories = GetCategoryList(in.Page, in.Limit)
	json_content, _ := json.Marshal(categories)
	json.Unmarshal(json_content, &rs)
	nrs.List = rs
	return &nrs, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterArticleServer(s, &articleServer{})
	pb.RegisterCategoryServer(s, &categoryServer{})
	s.Serve(lis)
}
