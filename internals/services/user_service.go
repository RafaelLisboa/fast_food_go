package services

import (
	"context"
	"crypto/aes"
	"fast_food_auth/db"
	"fast_food_auth/internals/models"
	"fast_food_auth/internals/repositories"
	"fmt"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, user models.User) error	
}


type userService struct {
	userRepository repositories.UserRepository;
}


func NewUserService() UserService {
	repository := repositories.NewUserRepository();
	
	return &userService{
		userRepository: repository,
	}
}


func (us *userService) CreateUser(ctx context.Context, user models.User) error {

	id := uuid.NewString()

	cypher, err := aes.NewCipher([]byte("hedqieufqworh9238yu3jieiu4fhedw"))

	if err != nil {
		ctx.Done()
	}

	fmt.Println(user);

	passwordTextByteBuffer := make([]byte, len(user.Password))

	cypher.Encrypt(passwordTextByteBuffer, []byte(passwordTextByteBuffer))

	passwordText := string(passwordTextByteBuffer)

	userDb := &db.CreateUserParams{
		ID: id,
		Name: user.Name,
		Email: user.Email,
		Username: user.Username,
		Password: passwordText,
	}

	return us.userRepository.CreateUser(ctx, *userDb)
}
