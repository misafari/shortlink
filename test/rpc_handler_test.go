package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"ir.safari.shortlink/api/gen"
	rpc2 "ir.safari.shortlink/api/rpc"
	"log"
	"net"
	"testing"
)

func TestPing(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "buffnet", grpc.WithContextDialer(dialer()), grpc.WithInsecure())
	assert.Nil(t, err)
	defer conn.Close()

	client := gen.NewShortLinkRpcServiceClient(conn)

	res, rpcErr := client.Ping(ctx, &gen.PingRequest{
		Income: 2,
	})
	assert.Nil(t, rpcErr)
	assert.Equal(t, int32(4), res.GetOutcome())
}

// helper
func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	rpcHandler := rpc2.NewServiceRpcImpl(nil)

	gen.RegisterShortLinkRpcServiceServer(server, rpcHandler)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}
