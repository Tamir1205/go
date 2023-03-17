package auth

import (
	"context"
	"github.com/Tamir1205/midterm1/internal/storage/users"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userRepository users.Repository
}

type Service interface {
	CreateUser(ctx context.Context, dto CreateUserDto) (*int64, error)
	SignIn(ctx context.Context, request SignInRequest) error
}

func NewService(userRepository users.Repository) Service {
	return &service{userRepository: userRepository}
}

func (s *service) CreateUser(ctx context.Context, dto CreateUserDto) (*int64, error) {
	password, err := s.hashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	createUser, err := s.userRepository.CreateUser(ctx, users.User{
		Email:     dto.Email,
		Password:  password,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	})
	if err != nil {
		return nil, err
	}

	return &createUser, nil
}

func (s *service) SignIn(ctx context.Context, request SignInRequest) error {
	user, err := s.userRepository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return err
	}

	if err := s.comparePasswords(user.Password, request.Password); err != nil {
		return err
	}

	return nil
}

func (s *service) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (s *service) comparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
