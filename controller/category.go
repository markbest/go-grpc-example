package controller

import (
	"encoding/json"

	"golang.org/x/net/context"

	. "github.com/markbest/go-grpc-example/model"
	pb "github.com/markbest/go-grpc-example/protos"
)

type CategoryServer struct{}

// get category info
func (c *CategoryServer) GetCategoryInfo(ctx context.Context, in *pb.CQueryRequest) (*pb.CategoryInfo, error) {
	var category Category
	var rs pb.CategoryInfo

	category = GetCategoryInfo(in.CatId)
	jsonContent, _ := json.Marshal(category)
	json.Unmarshal(jsonContent, &rs)
	return &rs, nil
}

// get category list
func (c *CategoryServer) GetCategoryList(ctx context.Context, in *pb.CQueryRequest) (*pb.CategoryList, error) {
	var categories []Category
	var rs []*pb.CategoryInfo
	var nrs pb.CategoryList

	categories = GetCategoryList(in.Page, in.Limit)
	jsonContent, _ := json.Marshal(categories)
	json.Unmarshal(jsonContent, &rs)
	nrs.Data = rs
	return &nrs, nil
}
