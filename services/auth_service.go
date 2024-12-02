package services

import (
	"errors"
	"fmt"
	"user_auth/internals"
	"user_auth/models"
	"user_auth/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Register(username, password string) error
	Login(username, password string) (string, error)
	UpdateUserDetails(user models.User) error
}

type AuthService struct {
	repo       repositories.IUserRepository
	jwtManager *internals.JWTManager
}

func NewAuthService(repo repositories.IUserRepository, jwtManager *internals.JWTManager) IAuthService {
	return &AuthService{
		repo:       repo,
		jwtManager: jwtManager,
	}
}

func (s *AuthService) Register(username, password string) error {
	user, err := s.repo.FindByUsername(username)

	if user != nil {
		return fmt.Errorf("User already registered!")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user = &models.User{Username: username, Password: string(hashedPassword)}

	s.repo.Create(*user)

	return nil
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("invalid credentials")
	}
	return s.jwtManager.Generate(int64(user.Id), user.Username, user.Name, user.Lastname, user.DNI, user.Address, user.PostalCode, user.Phone)
}

func (s *AuthService) UpdateUserDetails(user models.User) error {
	return s.repo.UpdateUserDetails(user)
}
