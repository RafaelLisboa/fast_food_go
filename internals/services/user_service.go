package services

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fast_food_auth/db"
	"fast_food_auth/internals/models"
	"fast_food_auth/internals/repositories"
	"io"
	"log"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, user models.User) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService() UserService {
	repository := repositories.NewUserRepository()

	return &userService{
		userRepository: repository,
	}
}

func (us *userService) CreateUser(ctx context.Context, user models.User) error {

	id := uuid.NewString()

	passwordText := encryptPassword(user.Password)

	userDb := &db.CreateUserParams{
		ID:       id,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Password: passwordText,
	}

	return us.userRepository.CreateUser(ctx, *userDb)
}
