package ttrpc

import (
	"context"
	"fmt"
	"net"

	"github.com/containerd/ttrpc"

	"github.com/xibz/GRPCvsTTRPC/models/ttrpcmodels"
)

// Server .
type Server struct {
	memcache map[string]*ttrpcmodels.Data
}

// NewTTRPCServer returns a new GRPC server
func NewTTRPCServer() *Server {
	return &Server{
		memcache: map[string]*ttrpcmodels.Data{},
	}
}

// Start .
func (s *Server) Start(socketpath string) error {
	l, err := net.Listen("unix", socketpath)
	if err != nil {
		return err
	}

	server, err := ttrpc.NewServer()
	if err != nil {
		return err
	}

	ttrpcmodels.RegisterRouteService(server, s)
	if err := server.Serve(context.Background(), l); err != nil {
		return err
	}

	return nil
}

var errNotFound = fmt.Errorf("NotFound")

// GetData .
func (s *Server) GetData(ctx context.Context, req *ttrpcmodels.GetDataRequest) (*ttrpcmodels.GetDataResponse, error) {
	v, ok := s.memcache[req.Key]
	if !ok {
		return nil, errNotFound
	}

	resp := &ttrpcmodels.GetDataResponse{
		Data: v,
	}

	return resp, nil
}

// PutData .
func (s *Server) PutData(ctx context.Context, req *ttrpcmodels.PutDataRequest) (*ttrpcmodels.PutDataResponse, error) {
	s.memcache[req.Key] = req.Data
	return &ttrpcmodels.PutDataResponse{}, nil
}
