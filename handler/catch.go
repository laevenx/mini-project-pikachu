package handler

import (
	"mini-project-pikachu/models"
	"mini-project-pikachu/usecase"

	"github.com/gofiber/fiber/v2"
)

func Catch(c *fiber.Ctx) error {
	var option models.Option
	var status bool
	if err := c.BodyParser(&option); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	m, status, err := usecase.Catch(option)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"Message": "done",
		"data":    m,
		"status":  status,
	})
}
