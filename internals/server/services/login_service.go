package services

import (
	"context"
	"fast_food_auth/internals/server/models"
	"fast_food_auth/internals/server/repositories"
	"fast_food_auth/pkg/encrypt"
	"fast_food_auth/pkg/exceptions"
	"fast_food_auth/pkg/validation"
	"fmt"
)

type loginService struct {
	tokenService   TokenService
	userRepository repositories.UserRepository
}

type LoginService interface {
	Login(ctx context.Context, user models.LoginRequest) (*models.Token, error)
	RefreshToken(ctx context.Context, token string) (*models.Token, error)
}

func NewLoginService() LoginService {
	tokenService := NewTokenService()
	userRepository := repositories.NewUserRepository()

	return &loginService{
		tokenService:   tokenService,
		userRepository: userRepository,
	}
}

func (ls *loginService) Login(ctx context.Context, user models.LoginRequest) (*models.Token, error) {
	if valid, field := validation.GetEmptyField(user); !valid {
		return nil, exceptions.NewErrorWithMessage(ctx, exceptions.EMPTY_REQUIRED_FIELD, field)
	}

	userRecord, err := ls.userRepository.GetUserByEmail(ctx, user.Email)

	encryptedPassword := encrypt.EncryptPassword(user.Password)

	if encryptedPassword != userRecord.Password {
		return nil, exceptions.NewError(ctx, exceptions.LOGIN_FAILED)
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	token, err := ls.tokenService.createTokenByUserId(ctx, userRecord.ID.String())

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return token, nil
}

func (ls *loginService) RefreshToken(ctx context.Context, token string) (*models.Token, error) {
	if isValid, userId := ls.tokenService.isTokenValid(ctx, token); isValid {
		return ls.tokenService.createTokenByUserId(ctx, userId)
	}

	return nil, exceptions.NewError(ctx, exceptions.REFRESH_TOKEN_ERROR)
}
