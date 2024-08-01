package services

import (
	"context"
	"fast_food_auth/db"
	"fast_food_auth/internals/server/models"
	"fast_food_auth/internals/server/repositories"
	"fast_food_auth/pkg/exceptions"
	"fast_food_auth/pkg/validation"
	"fast_food_auth/pkg/encrypt"

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

	if valid, field := validation.GetEmptyField(user); !valid {
		return exceptions.NewErrorWithMessage(ctx, exceptions.EMPTY_REQUIRED_FIELD, field)
	}

	createdUser, _ := us.userRepository.GetUserByEmail(ctx, user.Email)

	if createdUser.Email != "" && createdUser.Username != "" {
		return exceptions.NewError(ctx, exceptions.USER_ALREADY_EXISTS)
	}

	id := uuid.NewString()

	passwordText := encrypt.EncryptPassword(user.Password)

	userDb := &db.CreateUserParams{
		ID:       id,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Password: passwordText,
	}

	return us.userRepository.CreateUser(ctx, *userDb)
}

