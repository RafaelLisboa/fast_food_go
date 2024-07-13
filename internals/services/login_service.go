package services

import (
	"context"
	"errors"
	"fast_food_auth/internals/models"
	"fast_food_auth/internals/repositories"
)

type loginService struct {
	tokenService TokenService
	userRepository repositories.UserRepository
}

type LoginService interface {
	Login(user models.User) (string, error)
}

func NewLoginService() LoginService {
	tokenService := NewTokenService();

	return &loginService{tokenService: tokenService}
}


func (ls *loginService) Login(user models.User) (string, error) {
	if user.Email == "" {
		return "", errors.New("email is required when login");
	}

	userRecord, err := ls.userRepository.GetUserByEmail(context.Background(), user.Email);

	if err != nil {
		return "", err;
	}

	token , err := ls.tokenService.createTokenByUserId(userRecord.ID.String())

	if err != nil {
		return "", err
	}

	tokenStr := string(token)

	return tokenStr, nil
}