package handler

import (
	"context"
	"strconv"
	"user/model"
	"user/service"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (*AuthService) Register(ctx context.Context, req *service.UserRequest) (resp *service.UserResponse, err error) {
	var user model.User
	u, err := user.Create(req)

	if err != nil {
		return nil, err
	}

	return &service.UserResponse{
		Code: 200,
		Data: &service.UserModel{
			UserID:   strconv.Itoa(int(u.ID)),
			UserName: u.UserName,
		},
	}, nil
}

func (*AuthService) Login(ctx context.Context, req *service.UserRequest) (resp *service.UserResponse, err error) {
	return nil, nil
}

func (*AuthService) GetUser(ctx context.Context, req *service.UserRequest) (resp *service.UserResponse, err error) {
	return nil, nil
}
