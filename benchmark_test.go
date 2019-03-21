package grpc_test

import (
	"bytes"
	"context"
	"net"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/containerd/ttrpc"
	"google.golang.org/grpc"

	grpcserver "github.com/xibz/GRPCvsTTRPC/grpc"
	"github.com/xibz/GRPCvsTTRPC/models/grpcmodels"
	"github.com/xibz/GRPCvsTTRPC/models/ttrpcmodels"
	ttrpcserver "github.com/xibz/GRPCvsTTRPC/ttrpc"
)

func BenchmarkGRPCSmallPayload(b *testing.B) {
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
	defer func() {
		conn.Close()
		os.Remove(socketPath)
	}()

	client := grpcmodels.NewRouteClient(conn)
	getKeys := []string{
		"foo",
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
			},
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

func BenchmarkGRPCLargeArrayPayload(b *testing.B) {
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
	defer func() {
		conn.Close()
		os.Remove(socketPath)
	}()

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
				StringShape: strings.Repeat("foobarbaz", 1024*10),
				IntShape:    1,
				DoubleShape: 3.14,
				BoolShape:   true,
				BytesShape:  bytes.Repeat([]byte("a"), 1024*10),
				Shapes: []*grpcmodels.NestedShape{
					{
						NestedInt:   2,
						NestedBytes: bytes.Repeat([]byte("b"), 1024*10),
					},
					{
						NestedInt:   3,
						NestedBytes: bytes.Repeat([]byte("c"), 1024*10),
					},
					{
						NestedInt:   4,
						NestedBytes: bytes.Repeat([]byte("d"), 1024*10),
					},
					{
						NestedInt:   5,
						NestedBytes: bytes.Repeat([]byte("e"), 1024*10),
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

func BenchmarkGRPCLargeNArrayPayload(b *testing.B) {
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
	defer func() {
		conn.Close()
		os.Remove(socketPath)
	}()

	client := grpcmodels.NewRouteClient(conn)
	getKeys := []string{
		"bar",
		"not-exist",
	}

	putData := []grpcmodels.PutDataRequest{
		{
			Key: "bar",
			Data: &grpcmodels.Data{
				StringShape: strings.Repeat("foobarbaz", 1024*10),
				IntShape:    1,
				DoubleShape: 3.14,
				BoolShape:   true,
				BytesShape:  bytes.Repeat([]byte("a"), 1024*10),
				Shapes: []*grpcmodels.NestedShape{
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
				},
			},
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

func BenchmarkGRPCLargePayload(b *testing.B) {
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
	defer func() {
		conn.Close()
		os.Remove(socketPath)
	}()

	client := grpcmodels.NewRouteClient(conn)
	getKeys := []string{
		"foo",
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
				BytesShape:  bytes.Repeat([]byte{'a'}, 1024*1024*10),
			},
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

func BenchmarkTTRPCSmallPayload(b *testing.B) {
	socketPath := "/tmp/server.sock"
	ctx := context.Background()

	s := ttrpcserver.NewTTRPCServer()
	go func() {
		err := s.Start(socketPath)
		if err != nil {
			b.Fatalf("Failed to start server: %v", err)
		}
	}()

	time.Sleep(250 * time.Millisecond)
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		b.Fatalf("Failed to dial: %v", err)
	}
	defer func() {
		conn.Close()
		os.Remove(socketPath)
	}()

	client := ttrpcmodels.NewRouteClient(ttrpc.NewClient(conn))
	getKeys := []string{
		"foo",
		"not-exist",
	}

	putData := []ttrpcmodels.PutDataRequest{
		{
			Key: "foo",
			Data: &ttrpcmodels.Data{
				StringShape: "a",
				IntShape:    1,
				DoubleShape: 3.14,
				BoolShape:   true,
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := putData[i%len(putData)]
		client.PutData(ctx, &req)

		key := getKeys[i%len(getKeys)]
		client.GetData(ctx, &ttrpcmodels.GetDataRequest{
			Key: key,
		})
	}
}

func BenchmarkTTRPCLargeArrayPayload(b *testing.B) {
	socketPath := "/tmp/server.sock"
	ctx := context.Background()

	s := ttrpcserver.NewTTRPCServer()
	go func() {
		err := s.Start(socketPath)
		if err != nil {
			b.Fatalf("Failed to start server: %v", err)
		}
	}()

	time.Sleep(250 * time.Millisecond)
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		b.Fatalf("Failed to dial: %v", err)
	}
	defer func() {
		conn.Close()
		os.Remove(socketPath)
	}()

	client := ttrpcmodels.NewRouteClient(ttrpc.NewClient(conn))
	getKeys := []string{
		"foo",
		"bar",
		"baz",
		"not-exist",
	}

	putData := []ttrpcmodels.PutDataRequest{
		{
			Key: "foo",
			Data: &ttrpcmodels.Data{
				StringShape: "a",
				IntShape:    1,
				DoubleShape: 3.14,
				BoolShape:   true,
				BytesShape:  []byte{'a'},
				Shapes: []*ttrpcmodels.NestedShape{
					{
						NestedInt:   2,
						NestedBytes: []byte{'a'},
					},
				},
			},
		},
		{
			Key: "bar",
			Data: &ttrpcmodels.Data{
				StringShape: strings.Repeat("foobarbaz", 1024*10),
				IntShape:    1,
				DoubleShape: 3.14,
				BoolShape:   true,
				BytesShape:  bytes.Repeat([]byte("a"), 1024*10),
				Shapes: []*ttrpcmodels.NestedShape{
					{
						NestedInt:   2,
						NestedBytes: bytes.Repeat([]byte("b"), 1024*10),
					},
					{
						NestedInt:   3,
						NestedBytes: bytes.Repeat([]byte("c"), 1024*10),
					},
					{
						NestedInt:   4,
						NestedBytes: bytes.Repeat([]byte("d"), 1024*10),
					},
					{
						NestedInt:   5,
						NestedBytes: bytes.Repeat([]byte("e"), 1024*10),
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
		client.GetData(ctx, &ttrpcmodels.GetDataRequest{
			Key: key,
		})
	}
}

func BenchmarkTTRPCLargeNArrayPayload(b *testing.B) {
	socketPath := "/tmp/server.sock"
	ctx := context.Background()

	s := ttrpcserver.NewTTRPCServer()
	go func() {
		err := s.Start(socketPath)
		if err != nil {
			b.Fatalf("Failed to start server: %v", err)
		}
	}()

	time.Sleep(250 * time.Millisecond)
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		b.Fatalf("Failed to dial: %v", err)
	}
	defer func() {
		conn.Close()
		os.Remove(socketPath)
	}()

	client := ttrpcmodels.NewRouteClient(ttrpc.NewClient(conn))
	getKeys := []string{
		"bar",
		"not-exist",
	}

	putData := []ttrpcmodels.PutDataRequest{
		{
			Key: "bar",
			Data: &ttrpcmodels.Data{
				StringShape: strings.Repeat("foobarbaz", 1024*10),
				IntShape:    1,
				DoubleShape: 3.14,
				BoolShape:   true,
				BytesShape:  bytes.Repeat([]byte("a"), 1024*10),
				Shapes: []*ttrpcmodels.NestedShape{
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
					{NestedInt: 2},
					{NestedInt: 3},
					{NestedInt: 4},
					{NestedInt: 5},
				},
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := putData[i%len(putData)]
		client.PutData(ctx, &req)

		key := getKeys[i%len(getKeys)]
		client.GetData(ctx, &ttrpcmodels.GetDataRequest{
			Key: key,
		})
	}
}

func BenchmarkTTRPCLargePayload(b *testing.B) {
	socketPath := "/tmp/server.sock"
	ctx := context.Background()

	s := ttrpcserver.NewTTRPCServer()
	go func() {
		err := s.Start(socketPath)
		if err != nil {
			b.Fatalf("Failed to start server: %v", err)
		}
	}()

	time.Sleep(250 * time.Millisecond)
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		b.Fatalf("Failed to dial: %v", err)
	}
	defer func() {
		conn.Close()
		os.Remove(socketPath)
	}()

	client := ttrpcmodels.NewRouteClient(ttrpc.NewClient(conn))
	getKeys := []string{
		"foo",
		"not-exist",
	}

	putData := []ttrpcmodels.PutDataRequest{
		{
			Key: "foo",
			Data: &ttrpcmodels.Data{
				StringShape: "a",
				IntShape:    1,
				DoubleShape: 3.14,
				BoolShape:   true,
				BytesShape:  bytes.Repeat([]byte{'a'}, 1024*1024*10),
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := putData[i%len(putData)]
		client.PutData(ctx, &req)

		key := getKeys[i%len(getKeys)]
		client.GetData(ctx, &ttrpcmodels.GetDataRequest{
			Key: key,
		})
	}
}
