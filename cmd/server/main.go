package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/danielbh/hello-grpc-go/pb"
	"github.com/danielbh/hello-grpc-go/service"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	cacheSize := flag.Int("size", 128, "set cache size")
	flag.Parse()

	cacheServer, err := service.NewCache(*cacheSize)

	if err != nil {
		log.Fatal("cache failed to initialize: ", err)
	}

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = startGRPCServer(cacheServer, *enableTLS, listener)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}

func startGRPCServer(
	cacheService pb.CacheServiceServer,
	enableTLS bool,
	listener net.Listener,
) error {
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(loggerInterceptor),
	}

	grpcServer := grpc.NewServer(serverOptions...)

	pb.RegisterCacheServiceServer(grpcServer, cacheService)
	reflection.Register(grpcServer)

	log.Printf("Start GRPC server at %s, TLS = %t", listener.Addr().String(), enableTLS)
	return grpcServer.Serve(listener)
}

func loggerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Println("Request: ", info.FullMethod)
	return handler(ctx, req)
}
