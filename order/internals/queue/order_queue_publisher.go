package queue

import (
	"fast_food_order/config"

	"github.com/streadway/amqp"
)

type orderQueuePublisher struct {
	queueConnection *amqp.Connection
}

func NewOrderQueuePublisher() QueuePublisher {

	queueConnection := config.Connect()

	return &orderQueuePublisher{
		queueConnection,
	}
}

func (oqp *orderQueuePublisher) SendMessage(messageBody []byte ) error {

	ch, err := oqp.queueConnection.Channel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	return ch.Publish("", config.QUEUE_NAME, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body: messageBody,
	})
}
