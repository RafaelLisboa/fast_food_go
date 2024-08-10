package config

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type rabbitCredentials struct {
	user     string
	password string
}

func getCredentials() *rabbitCredentials {
	err := godotenv.Load("build/.env")

	if err != nil {
		log.Info("Error loading .env file " + err.Error())
	}

	return &rabbitCredentials{
		user:     os.Getenv("RABBITMQ_USER"),
		password: os.Getenv("RABBITMQ_PASSWORD"),
	}

}

const QUEUE_NAME = "ORDERS"

var (
	rabbitInstance *amqp.Connection
	rabbitOnce     sync.Once
)

func Connect() *amqp.Connection {

	if rabbitInstance != nil {
		return rabbitInstance
	}

	rabbitOnce.Do(func() {
		credentials := getCredentials()

		connStr := fmt.Sprintf("amqp://%s:%s@%s/",
			credentials.user, credentials.password, "localhost")

		var err error

		rabbitInstance, err = amqp.Dial(connStr)

		if err != nil {
			panic(errors.New("error connecting on rabbit " + err.Error()))
		}

		ch, err := rabbitInstance.Channel()

		if err != nil {
			panic(errors.New("error connecting on rabbit " + err.Error()))
		}

		defer ch.Close()

		_, err = ch.QueueDeclare(QUEUE_NAME, false, false, false, false, nil)

		if err != nil {
			panic(errors.New("error connecting on rabbit " + err.Error()))
		}

	})

	return rabbitInstance
}
