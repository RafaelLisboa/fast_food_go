package services

import "fast_food_auth/internals/models"

type loginService struct {
	tokenService TokenService
}

type LoginService interface {
	Login(user models.User) (string, error)
}

func New() LoginService {
	tokenService := NewTokenService();

	return &loginService{tokenService: tokenService}
}


func (ls *loginService) Login(user models.User) (string, error) {
	
}