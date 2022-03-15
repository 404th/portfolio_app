package client

import (
	"fmt"
	"log"
	"os"

	"github.com/404th/portfolio_app/internal/adapters/framework/left/grpc/pb/user_service"
	"google.golang.org/grpc"
)

type GrpcClientI interface {
	UserService() user_service.UserServiceClient
}

type GrpcClient struct {
	connections map[string]interface{}
}

func New() (*GrpcClient, error) {
	var target string = fmt.Sprintf("%s:%s", os.Getenv("PG_HOST"), os.Getenv("PG_PORT"))
	client, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to create client:%s", err.Error())
		return nil, err
	}
	defer client.Close()

	return &GrpcClient{
		connections: map[string]interface{}{
			"user_service": user_service.NewUserServiceClient(client),
		},
	}, nil
}

func (g *GrpcClient) UserService() user_service.UserServiceClient {
	return g.connections["user_service"].(user_service.UserServiceClient)
}
