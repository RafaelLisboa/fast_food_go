package services

import (
	"context"
	"fast_food_auth/db"
	"fast_food_auth/internals/exceptions"
	"fast_food_auth/internals/models"
	"fast_food_auth/internals/repositories"
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

	createdUser, err := us.userRepository.GetUserByEmail(ctx, user.Email)

	if err != nil {
		exceptions.NewError(ctx, exceptions.INTERNAL_ERROR)
		return err
	}

	if createdUser.ID.String() != "" {
		exceptions.NewError(ctx, exceptions.USER_ALREADY_EXISTS)
		return err
	}

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
