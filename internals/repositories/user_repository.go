package repositories

import (
	"context"
	"fast_food_auth/db"
	"fmt"
)

type UserRepository interface {
	CreateUser(ctx context.Context, params db.CreateUserParams) error
	GetUserByID(ctx context.Context, id string) (db.User, error)
	GetUserByEmail(ctx context.Context, email string) (db.User, error)
	
}

type userRepository struct {
	db *db.Queries
}

func NewUserRepository() UserRepository {
	conn, err := db.GetDBInstance()

	if err != nil {
		fmt.Errorf("%v", err)
		panic("Error connecting with database")
	}

	return &userRepository{
		db: conn,
	}

}

func (ur *userRepository) CreateUser(ctx context.Context, params db.CreateUserParams)  error {
	err := ur.db.CreateUser(ctx, params);

	return err;
}


func (ur *userRepository) GetUserByID(ctx context.Context, id string) (db.User, error) {
	return ur.db.GetUserByID(ctx, id);
}

func (ur *userRepository) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	return ur.db.GetUserByEmail(ctx, email);
}
