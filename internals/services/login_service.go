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
	Login(ctx context.Context, user models.LoginRequest) (*models.Token, error)
}

func NewLoginService() LoginService {
	tokenService := NewTokenService()
	userRepository := repositories.NewUserRepository()

	return &loginService{tokenService: tokenService,
		userRepository: userRepository}

}

func (ls *loginService) Login(ctx context.Context, user models.LoginRequest) (*models.Token, error) {

	if valid, field := GetEmptyField(user); !valid {
		return nil, exceptions.NewErrorWithMessage(ctx, exceptions.EMPTY_REQUIRED_FIELD, field)
	}

	userRecord, err := ls.userRepository.GetUserByEmail(ctx, user.Email)

	encryptedPassword := encryptPassword(user.Password);


	if encryptedPassword != userRecord.Password {
		return nil, exceptions.NewError(ctx, exceptions.LOGIN_FAILED)
	}

	if err != nil {
		return nil, err
	}

	token, err := ls.tokenService.createTokenByUserId(userRecord.ID.String())

	if err != nil {
		return nil, err
	}



	return token, nil
}
