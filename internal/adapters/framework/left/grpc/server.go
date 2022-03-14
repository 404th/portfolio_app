package grpc

import (
	"net"

	"github.com/404th/portfolio_app/internal/adapters/framework/right/db"
	"google.golang.org/grpc"
)

type Server struct {
	db *db.DB
}

func NewServer(db *db.DB) *Server {
	return &Server{db}
}

func (s *Server) Run(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	if err = grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
