package test

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/danielbh/hello-grpc-go/pb"
	"github.com/danielbh/hello-grpc-go/service"
)

func TestCacheGetNoValue(t *testing.T) {
	client := setup(t)

	response, err := client.Get(context.Background(), &pb.GetRequest{Key: "noexistKey"})

	require.NoError(t, err)
	require.Nil(t, response.Value)
}

func TestCacheSetGet(t *testing.T) {
	client := setup(t)

	setResponse, err := client.Set(context.Background(), &pb.SetRequest{Key: "key1", Value: &pb.Any{Value: []byte("value1")}})
	require.NoError(t, err)
	require.Equal(t, setResponse.Evicted, false)

	getResponse, err := client.Get(context.Background(), &pb.GetRequest{Key: "key1"})

	require.NoError(t, err)
	require.Equal(t, "value1", string(getResponse.Value.GetValue()))
}

func setup(t *testing.T) pb.CacheServiceClient {
	t.Parallel()

	cacheServer, err := service.NewCache(128)

	grpcServer := grpc.NewServer()
	pb.RegisterCacheServiceServer(grpcServer, cacheServer)

	listener, err := net.Listen("tcp", ":0") // random available port
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	addr := listener.Addr().String()

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	require.NoError(t, err)

	return pb.NewCacheServiceClient(conn)
}
