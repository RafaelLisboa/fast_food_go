package config

import (
	"errors"
	"os"
)

type DbCredentials struct {
	Username string
	Password string
	Host     string
	Database string
}


func GetDatabaseCredentials() *DbCredentials {

	if os.Getenv("HOST") == "" {
		panic(errors.New("error getting the env vars"))
	}

	return &DbCredentials{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Host:     os.Getenv("HOST"),
		Database: os.Getenv("DATABASE"),
	}

}
