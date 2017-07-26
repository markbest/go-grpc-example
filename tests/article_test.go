package tests

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc/server/protos"
	"testing"
)

const (
	address = "localhost:50051"
)

func Test_GetArticleInfo(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	c := pb.NewArticleClient(conn)
	r, err := c.GetArticleInfo(context.Background(), &pb.AQueryRequest{Id: 10})
	if err != nil {
		t.Error(err)
	}
	t.Log(r.Title)
}

func Test_GetArticleList(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	c := pb.NewArticleClient(conn)
	r, err := c.GetArticleList(context.Background(), &pb.AQueryRequest{Page: 1, Limit: 20})
	if err != nil {
		t.Error(err)
	}
	t.Log(len(r.List))
}
