package services

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	secretKey = "ohfe84230_*&*(¨#812)"
)

type tokenService struct {}

type TokenService interface {
	CreateTokenByUserId(id string) (string, error)
	IsTokenValid(token string) (bool)
}


func NewTokenService() TokenService {
	return &tokenService{}
}

func (ts *tokenService) CreateTokenByUserId(id string) (string, error) {
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "id": id, 
        "exp": time.Now().Add(time.Minute * 10).Unix(), 
    })

	tokenString, err := token.SignedString(secretKey)
    
	if err != nil {
 	   return "", err
    }

	return tokenString, nil;

}

func (ts *tokenService) IsTokenValid(token string) bool {
	if token == "" {
		return false;
	}

	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return false;
	}

	return tokenParsed.Valid
}