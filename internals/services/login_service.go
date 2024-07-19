package services

import (
	"context"

	"fast_food_auth/internals/exceptions"
	"fast_food_auth/internals/models"
	"fast_food_auth/internals/repositories"
)

type loginService struct {
	tokenService   TokenService
	userRepository repositories.UserRepository
}

type LoginService interface {
	Login(ctx context.Context, user models.LoginRequest) (string, error)
}

func NewLoginService() LoginService {
	tokenService := NewTokenService()
	userRepository := repositories.NewUserRepository()

	return &loginService{tokenService: tokenService,
		userRepository: userRepository}

}

func (ls *loginService) Login(ctx context.Context, user models.LoginRequest) (string, error) {

	if valid, field := GetEmptyField(user); !valid {
		return "", exceptions.NewErrorWithMessage(ctx, exceptions.EMPTY_REQUIRED_FIELD, field)
	}

	userRecord, err := ls.userRepository.GetUserByEmail(ctx, user.Email)

	encryptedPassword := encryptPassword(user.Password);


	if encryptedPassword != userRecord.Password {
		return "", exceptions.NewError(ctx, exceptions.LOGIN_FAILED)
	}

	if err != nil {
		return "", err
	}

	token, err := ls.tokenService.createTokenByUserId(userRecord.ID.String())

	if err != nil {
		return "", err
	}

	tokenStr := string(token)

	return tokenStr, nil
}
