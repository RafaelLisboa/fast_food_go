package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DbCredentials struct {
	Username string
	Password string
	Host     string
	Database string
}

func LoadEnv() {
	err := godotenv.Load("build/.env")
	
	if err != nil {
		panic(err)
	}
}

func GetDatabaseCredentials() *DbCredentials {
	LoadEnv()

	return &DbCredentials{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Host: os.Getenv("HOST"),
		Database: os.Getenv("DATABASE"),
	}
	
}
