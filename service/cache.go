package service

import (
	"context"

	lru "github.com/hashicorp/golang-lru"

	"github.com/danielbh/hello-grpc-go/pb"
)

// Cache Implements cache interface
type Cache interface {
	Set(context.Context, *pb.SetRequest) (*pb.SetResponse, error)
	Get(context.Context, *pb.GetRequest) (*pb.GetResponse, error)
}

type cache struct {
	lru *lru.Cache
}

// NewCache creates new cache instance
func NewCache(size int) (Cache, error) {

	lru, err := lru.New(size)

	if err != nil {
		return nil, err
	}

	return &cache{
		lru: lru,
	}, nil
}

func (c *cache) Set(ctx context.Context, request *pb.SetRequest) (*pb.SetResponse, error) {
	evicted := c.lru.Add(request.Key, request.Value)
	return &pb.SetResponse{Evicted: evicted}, nil
}

func (c *cache) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	value, ok := c.lru.Get(request.Key)

	if ok {
		return &pb.GetResponse{Value: value.(*pb.Any)}, nil
	}

	return &pb.GetResponse{}, nil
}
