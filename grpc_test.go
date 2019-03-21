package grpc_test

import (
	"bytes"
	"context"
	"net"
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc"

	grpcserver "github.com/xibz/GRPCvsTTRPC/grpc"
	"github.com/xibz/GRPCvsTTRPC/models/grpcmodels"
)

func BenchmarkGRPC(b *testing.B) {
	socketPath := "/tmp/server.sock"
	ctx := context.Background()

	s := grpcserver.NewGRPCServer()
	go s.Start(socketPath)

	conn, err := grpc.Dial(
		socketPath,
		grpc.WithInsecure(),
		grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("unix", addr, timeout)
		}),
	)
	if err != nil {
		b.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := grpcmodels.NewRouteClient(conn)
	getKeys := []string{
		"foo",
		"bar",
		"baz",
		"not-exist",
	}

	putData := []grpcmodels.PutDataRequest{
		{
			Key: "foo",
			Data: &grpcmodels.Data{
				StringShape: "a",
				IntShape:    1,
				DoubleShape: 3.14,
				BoolShape:   true,
				BytesShape:  []byte{'a'},
				Shapes: []*grpcmodels.NestedShape{
					{
						NestedInt:   2,
						NestedBytes: []byte{'a'},
					},
				},
			},
		},
		{
			Key: "bar",
			Data: &grpcmodels.Data{
				StringShape: strings.Repeat("foobarbaz", 1024*1024),
				IntShape:    1,
				DoubleShape: 3.14,
				BoolShape:   true,
				BytesShape:  bytes.Repeat([]byte("a"), 1024*1024),
				Shapes: []*grpcmodels.NestedShape{
					{
						NestedInt:   2,
						NestedBytes: bytes.Repeat([]byte("b"), 1024*1024),
					},
				},
			},
		},
		{
			Key: "baz",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := putData[i%len(putData)]
		client.PutData(ctx, &req)

		key := getKeys[i%len(getKeys)]
		client.GetData(ctx, &grpcmodels.GetDataRequest{
			Key: key,
		})
	}
}
