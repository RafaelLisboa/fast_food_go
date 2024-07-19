package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	secretKey = "4b8e8d2b24f83e945aee647ed75a1e2167c6e7a2b9b3f913f5b9b3f9135c6e8d"
)

type tokenService struct {}

type TokenService interface {
	createTokenByUserId(id string) (string, error)
	isTokenValid(token string) (bool)
}


func NewTokenService() TokenService {
	return &tokenService{}
}

func (ts *tokenService) createTokenByUserId(id string) (string, error) {
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "id": id, 
        "exp": time.Now().Add(time.Minute * 10).Unix(), 
    })

	tokenString, err := token.SignedString([]byte(secretKey))
    
	if err != nil {
 	   return "", err
    }

	return tokenString, nil;

}

func (ts *tokenService) isTokenValid(token string) bool {
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