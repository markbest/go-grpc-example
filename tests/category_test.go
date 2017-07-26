package tests

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc/server/protos"
	"testing"
)

func Test_GetCategoryInfo(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	c := pb.NewCategoryClient(conn)
	r, err := c.GetCategoryInfo(context.Background(), &pb.CQueryRequest{Id: 2})
	if err != nil {
		t.Error(err)
	}
	t.Log(r.Title)
}

func Test_GetCategoryList(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	c := pb.NewCategoryClient(conn)
	r, err := c.GetCategoryList(context.Background(), &pb.CQueryRequest{Page: 1, Limit: 20})
	if err != nil {
		t.Error(err)
	}
	t.Log(len(r.List))
}
