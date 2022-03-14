package grpc

import (
	"context"
	"fmt"

	"github.com/404th/portfolio_app/internal/adapters/framework/left/grpc/pb/user_service"
	"github.com/google/uuid"
)

func (s *Adapter) SignUp(ctx context.Context, in *user_service.CreateUser) (*user_service.CreateUser, error) {
	var em_user *user_service.CreateUser
	if in.GetEmail() == "" || in.GetPassword() == 0 || in.GetUsername() == "" {
		return em_user, fmt.Errorf("invalid input")
	}

	_, err := s.api.GetSignUp(in.GetEmail(), in.GetUsername(), in.GetPassword())
	if err != nil {
		return em_user, err
	}

	em_user = &user_service.CreateUser{
		Username: in.GetUsername(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}

	return em_user, nil
}

func (s *Adapter) SignIn(ctx context.Context, in *user_service.SignInUser) (*user_service.CreateUser, error) {
	var em_user *user_service.CreateUser
	if in.GetPassword() == 0 || in.GetUsername() == "" {
		return em_user, fmt.Errorf("invalid input")
	}

	username, email, err := s.api.GetSignIn(in.GetUsername(), in.GetPassword())
	if err != nil {
		return em_user, err
	}

	em_user = &user_service.CreateUser{
		Username: username,
		Email:    email,
		Password: in.GetPassword(),
	}

	return em_user, nil
}

func (s *Adapter) DeleteUser(ctx context.Context, in *user_service.UUIDTracker) (*user_service.AnswerForDelete, error) {
	var em_user *user_service.AnswerForDelete
	if in.GetId() == "" {
		return em_user, fmt.Errorf("invalid input")
	}

	id := uuid.MustParse(in.GetId())
	is_deleted, err := s.api.GetDeleteUser(id)
	if err != nil {
		return em_user, err
	}

	em_user = &user_service.AnswerForDelete{
		IsDeleted: is_deleted,
	}

	return em_user, nil
}
