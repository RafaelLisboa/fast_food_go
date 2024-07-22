package services

import (
	"context"
	"fast_food_auth/internals/models"
	"fast_food_auth/internals/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	secretKey          = "4b8e8d2b24f83e945aee647ed75a1e2167c6e7a2b9b3f913f5b9b3f9135c6e8d"
	refreshTokenSecret = "8a9e4f8b7d2a4e5b8e3d1f2c7a1b9c5f6e7a9d2b4e5c8a9f3d1b2e4c5a7b9d2f"
)

type tokenService struct {
	tokenRepository repositories.TokenRepository
	userService     UserService
}

type TokenService interface {
	createTokenByUserId(id string) (*models.Token, error)
	isTokenValid(token string) bool
}

func NewTokenService() TokenService {
	tokenRepository := repositories.NewTokenRepository()

	return &tokenService{
		tokenRepository: tokenRepository,
	}
}

func (ts *tokenService) createTokenByUserId(id string) (*models.Token, error) {

	expTime := time.Now().Add(time.Minute * 10).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  id,
			"exp": expTime,
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	
	if err != nil {
		//TODO: implement
	}

	refreshTokenString, err := ts.createRefreshToken(id)

	if err != nil {
		return nil, err
	}

	return &models.Token{
		AcessToken: tokenString,
		ExpiresIn:  uint32(expTime),
		RefreshToken: refreshTokenString,
	}, nil

}

func (ts *tokenService) isTokenValid(token string) bool {
	if token == "" {
		return false
	}

	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return false
	}

	return tokenParsed.Valid
}

func (ts *tokenService) createRefreshToken(userId string) (string, error) {
	refreshTokenExpTime := time.Now().Add(time.Hour * 1).Unix()

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  userId,
			"exp": refreshTokenExpTime,
		})

	refreshTokenStringJwt, err := refreshToken.SignedString([]byte(refreshTokenSecret))

	if err != nil {
		panic(err)
	}

	refreshTokenParam := &models.RefreshToken{
		UserId: userId,
		ExpiresIn: uint32(refreshTokenExpTime),
		Token: refreshTokenStringJwt,
	}

	err = ts.tokenRepository.CreateRefreshTokenByUserId(context.Background(), refreshTokenParam);

	if err != nil {
		return "", err
	}

	return refreshTokenStringJwt, nil

}
