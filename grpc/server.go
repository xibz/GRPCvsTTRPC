package grpc

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/xibz/GRPCvsTTRPC/models/grpcmodels"
)

// Server .
type Server struct {
	memcache map[string]*grpcmodels.Data
}

// NewGRPCServer returns a new GRPC server
func NewGRPCServer() *Server {
	return &Server{
		memcache: map[string]*grpcmodels.Data{},
	}
}

// Start .
func (s *Server) Start(socketpath string) error {
	l, err := net.Listen("unix", socketpath)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	grpcmodels.RegisterRouteServer(server, s)
	if err := server.Serve(l); err != nil {
		return err
	}

	return nil
}

var errNotFound = fmt.Errorf("NotFound")

// GetData .
func (s *Server) GetData(ctx context.Context, req *grpcmodels.GetDataRequest) (*grpcmodels.GetDataResponse, error) {
	v, ok := s.memcache[req.Key]
	if !ok {
		return nil, errNotFound
	}

	resp := &grpcmodels.GetDataResponse{
		Data: v,
	}

	return resp, nil
}

// PutData .
func (s *Server) PutData(ctx context.Context, req *grpcmodels.PutDataRequest) (*grpcmodels.PutDataResponse, error) {
	s.memcache[req.Key] = req.Data
	return &grpcmodels.PutDataResponse{}, nil
}
