package server

import (
	"context"
	pb "grpcdemo/proto"
)

type HelloServer struct {
	pb.UnimplementedHelloServer
}

func NewHelloServer() *HelloServer {
	return &HelloServer{}
}
func (h *HelloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Msg: "hello !"}, nil
}
