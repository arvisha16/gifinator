package main

import (
	"errors"
	"log"
	"net"

	pb "github.com/GoogleCloudPlatform/k8s-render-demo/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (server) RenderFrame(ctx context.Context, req *pb.RenderRequest) (*pb.RenderResponse, error) {
	return nil, errors.New("not implemented")
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("listen failed: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterRenderServer(srv, server{})
	srv.Serve(l)
}