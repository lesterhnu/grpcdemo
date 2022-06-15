package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	pb "grpcdemo/proto"
	"grpcdemo/server"
	"log"
	"net/http"
	"strings"
)

const PORT = "8888"

func main() {
	err := RunServer()
	if err != nil {
		panic(err)
	}
}

func RunServer() error {
	httpMux := runHttpServer()
	grpcS := runGrpcServer()
	gatewayMux := runGatewayServer()
	httpMux.Handle("/", gatewayMux)

	return http.ListenAndServe(":"+PORT, grpcHandlerFunc(grpcS, httpMux))
}

func runGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, server.NewHelloServer())
	reflection.Register(s)
	return s
}

func runHttpServer() *http.ServeMux {
	s := http.NewServeMux()
	return s
}

func runGatewayServer() *runtime.ServeMux {
	endpoint := ":" + PORT
	gwmux := runtime.NewServeMux()
	options := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterHelloHandlerFromEndpoint(context.Background(), gwmux, endpoint, options)
	if err != nil {
		log.Fatal(err)
	}
	return gwmux
}
func grpcHandlerFunc(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
