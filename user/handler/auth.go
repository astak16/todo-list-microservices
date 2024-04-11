package handler

import (
	"context"
	"user/model"
	"user/service"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (*AuthService) Register(ctx context.Context, req *service.UserRequest) (resp *service.UserResponse, err error) {
	var user model.User
	user.Create(req)
	return nil, nil
}

func (*AuthService) Login(ctx context.Context, req *service.UserRequest) (resp *service.UserResponse, err error) {
	return nil, nil
}

func (*AuthService) GetUser(ctx context.Context, req *service.UserRequest) (resp *service.UserResponse, err error) {
	return nil, nil
}
