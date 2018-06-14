package controller

import (
	"encoding/json"

	"golang.org/x/net/context"

	. "github.com/markbest/go-grpc-example/model"
	pb "github.com/markbest/go-grpc-example/protos"
)

type ArticleServer struct{}

// get article info
func (a *ArticleServer) GetArticleInfo(ctx context.Context, in *pb.AQueryRequest) (*pb.ArticleInfo, error) {
	var article Article
	var rs pb.ArticleInfo

	article = GetArticleInfo(in.Id)
	jsonContent, _ := json.Marshal(article)
	json.Unmarshal(jsonContent, &rs)
	return &rs, nil
}

// get article list
func (a *ArticleServer) GetArticleList(ctx context.Context, in *pb.AQueryRequest) (*pb.ArticleList, error) {
	var articles []Article
	var rs []*pb.ArticleInfo
	var nrs pb.ArticleList

	articles = GetArticleList(in.Page, in.Limit)
	jsonContent, _ := json.Marshal(articles)
	json.Unmarshal(jsonContent, &rs)
	nrs.List = rs
	return &nrs, nil
}
