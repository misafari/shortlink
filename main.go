package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"ir.safari.shortlink/api/gen"
	"ir.safari.shortlink/api/rpc"
	"ir.safari.shortlink/repository"
	"log"
	"net"
)

const (
	rpcPort = ":50551"
)

var ctx = context.Background()

func main() {
	// init redis
	redisInit := redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%s:%d", "localhost", 20156),
		PoolSize:   5,
		DB:         0,
		MaxRetries: 10,
	})

	redisManager := repository.NewRedisManager(redisInit)
	_ = repository.NewCachedUrlRedisRepository(redisManager)

	// init mongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	urlRepository := repository.NewOriginalUrlRepository(client, ctx)

	// init rpc server
	lis, err := net.Listen("tcp", rpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	rpcImpl := rpc.NewServiceRpcImpl(urlRepository)

	gen.RegisterShortLinkRpcServiceServer(grpcServer, rpcImpl)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
