package services

import (
	"context"
	"fast_food_auth/db"
	"fast_food_auth/internals/models"
	"fast_food_auth/internals/repositories"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, user models.User) error	
}


type userService struct {
	userRepository repositories.UserRepository;
}


func NewUserService() *userService{
	repository := repositories.NewUserRepository();
	
	return &userService{
		userRepository: repository,
	}
}


func (us *userService) CreateUser(ctx context.Context, user models.User) error {

	id := uuid.NewString()

	userDb := &db.CreateUserParams{
		ID: id,
		Name: user.Name,
		Email: user.Email,
		Username: user.Username,
	}

	return us.userRepository.CreateUser(ctx, *userDb)
}