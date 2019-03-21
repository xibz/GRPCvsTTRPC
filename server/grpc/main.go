package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/xibz/GRPCvsTTRPC/models"
)

const port = 8080

type server struct {
	memcache map[string]models.Data
}

func (s *server) GetData(context.Context, *models.GetDataRequest) (*models.GetDataResponse, error) {
	return nil, nil
}

func (s *server) PutData(context.Context, *models.PutDataRequest) (*models.PutDataResponse, error) {
	return nil, nil
}

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := &server{}
	server := grpc.NewServer()
	models.RegisterRouteServer(server, s)
	if err := server.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
