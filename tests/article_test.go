package tests

import (
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/markbest/go-grpc-example/protos"
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
	r, err := c.GetArticleInfo(context.Background(), &pb.AQueryRequest{ArticleId: 10})
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
	t.Log(len(r.Data))
}

func Benchmark_GetArticleInfo(b *testing.B) {
	b.StopTimer()

	conn, _ := grpc.Dial(address, grpc.WithInsecure())
	defer conn.Close()
	c := pb.NewArticleClient(conn)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.GetArticleInfo(context.Background(), &pb.AQueryRequest{ArticleId: 10})
	}
}

func Benchmark_GetArticleList(b *testing.B) {
	b.StopTimer()

	conn, _ := grpc.Dial(address, grpc.WithInsecure())
	defer conn.Close()
	c := pb.NewArticleClient(conn)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.GetArticleList(context.Background(), &pb.AQueryRequest{Page: 1, Limit: 20})
	}
}
