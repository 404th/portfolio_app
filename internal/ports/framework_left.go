package ports

import (
	"context"

	"github.com/404th/portfolio_app/internal/adapters/framework/left/grpc/pb/user_service"
)

type GRPCPorts interface {
	Run()
	SignUp(ctx context.Context, in *user_service.CreateUser) (*user_service.CreateUser, error)
	SignIn(ctx context.Context, in *user_service.SignInUser) (*user_service.CreateUser, error)
	DeleteUser(ctx context.Context, in *user_service.UUIDTracker) (*user_service.AnswerForDelete, error)
}

// GetSignUp(email, username string, password int32) (uuid.UUID, error)
// GetSignIn(username string, password int32) (string, string, error)
// GetDeleteUser(id uuid.UUID) (bool, error)
