package handler

import (
	"context"
	"fmt"
	"user/service"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (*AuthService) Register(ctx context.Context, req *service.UserRequest) (resp *service.UserResponse, err error) {
	fmt.Println("register succss", req)

	// var user model.User

	// user.Create(req)
	return nil, nil
}

func (*AuthService) Login(ctx context.Context, req *service.UserRequest) (resp *service.UserResponse, err error) {
	return nil, nil
}

func (*AuthService) GetUser(ctx context.Context, req *service.UserRequest) (resp *service.UserResponse, err error) {
	return nil, nil
}
