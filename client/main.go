package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpcdemo/proto"
	"log"
)

func main() {
	conn, _ := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	c := pb.NewHelloClient(conn)
	ctx := context.Background()
	resp, err := c.SayHello(ctx, &pb.HelloRequest{Msg: "123"})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp.Msg)
}
