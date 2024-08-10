package handler

import (
	"fast_food_order/internals/queue"

	"github.com/gofiber/fiber/v2"
)

type orderHandler struct {
	orderQueuePublish queue.QueuePublisher
}

func NewOrderHandler() *orderHandler {
	orderQueuePublish := queue.NewOrderQueuePublisher()

	return &orderHandler{
		orderQueuePublish,
	}
}

func (oh *orderHandler) SendMessage(ctx *fiber.Ctx) error {

	oh.orderQueuePublish.SendMessage([]byte("Deu bom"))

	return ctx.SendStatus(fiber.StatusAccepted)
}
