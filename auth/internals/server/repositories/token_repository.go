package repositories

import (
	"context"
	"fast_food_auth/internals/db"
	"fast_food_auth/internals/server/models"
)

type TokenRepository interface {
	CreateRefreshTokenByUserId(ctx context.Context, refreshToken *models.RefreshToken)  error 
	IsRefreshTokenValid(ctx context.Context, token string) bool
}

type tokenRepository struct {
	db *db.Queries
}

func NewTokenRepository() *tokenRepository {
	conn, err := db.GetDBInstance()

	if err != nil {
		panic(err)
	}

	return &tokenRepository{
		db: conn,
	}
}


func (tr *tokenRepository) CreateRefreshTokenByUserId(ctx context.Context, refreshToken *models.RefreshToken)  error {
	params := &db.CreateRefreshTokenParams{
		
		UserID: refreshToken.UserId,
		Token: refreshToken.Token,
		ExpiresAt: int32(refreshToken.ExpiresIn),
	}
	
	return tr.db.CreateRefreshToken(ctx, *params);
}


func (tr *tokenRepository) IsRefreshTokenValid(ctx context.Context, token string) bool {
	tokenRecord, err := tr.db.GetRefreshToken(ctx, token)

	if err != nil {
		return false
	}

	if tokenRecord != "" {
		return true
	}

	return false
}
