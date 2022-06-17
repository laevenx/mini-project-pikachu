package handler

import (
	"math/rand"
	"mini-project-pikachu/usecase"

	"github.com/gofiber/fiber/v2"
)

func Release(c *fiber.Ctx) error {

	var number = rand.Intn(100)

	result := usecase.Release(number)

	return c.Status(200).JSON(fiber.Map{
		"Message": "done",
		"data":    result,
	})
}
