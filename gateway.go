package main

import (
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	. "github.com/markbest/go-grpc-example/conf"
	pb "github.com/markbest/go-grpc-example/protos"
)

func run(grpcEndpoint, httpEndpoint string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	// article route
	err := pb.RegisterArticleHandlerFromEndpoint(ctx, mux, grpcEndpoint, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		return err
	}
	// category route
	err = pb.RegisterCategoryHandlerFromEndpoint(ctx, mux, grpcEndpoint, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		return err
	}
	return http.ListenAndServe(httpEndpoint, mux)
}

func main() {
	if err := InitConfig(); err != nil {
		log.Fatal(err)
	}

	if err := run(Conf.App.GrpcPort, Conf.App.HttpPort); err != nil {
		log.Fatal(err)
	}
}
