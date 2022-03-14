package grpc

import (
	"fmt"
	"net"

	"github.com/404th/portfolio_app/internal/adapters/framework/left/grpc/pb/user_service"
	"github.com/404th/portfolio_app/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPorts
}

func NewAdapter(api ports.APIPorts) *Adapter {
	return &Adapter{api}
}

func (s *Adapter) Run(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	user_service.RegisterUserServiceServer(grpcServer, s)

	if err = grpcServer.Serve(lis); err != nil {
		return err
	}

	fmt.Printf("Server is working over port%v", port)
	return nil
}
