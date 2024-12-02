package handlers

import (
	"context"
	"user_auth/models"
	"user_auth/pb"
	"user_auth/services"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	authService services.IAuthService
}

func NewAuthHandler(authService services.IAuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (a *AuthHandler) Register(ctx context.Context, p *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if err := a.authService.Register(p.Username, p.Password); err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{Message: "Registered successfully"}, nil
}

func (a *AuthHandler) Login(ctx context.Context, p *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := a.authService.Login(p.Username, p.Password)

	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{Token: token}, nil
}

func (a *AuthHandler) UpdateUserDetails(ctx context.Context, p *pb.UpdateUserDetailsRequest) (*pb.UpdateUserDetailsResponse, error) {
	user := models.User{
		Id:         int(p.Id),
		Username:   p.Username,
		Name:       p.Name,
		Lastname:   p.Lastname,
		DNI:        int(p.Dni),
		Address:    p.Address,
		PostalCode: int(p.PostalCode),
		Phone:      int(p.Phone),
	}

	if err := a.authService.UpdateUserDetails(user); err != nil {
		return nil, err
	}
	return &pb.UpdateUserDetailsResponse{Message: "User details updated successfully"}, nil
}
