package handler

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

func primeNumber(n int) bool {
	if n < 2 {
		return false
	}
	sq_root := int(math.Sqrt(float64(n)))
	for i := 2; i <= sq_root; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Release(c *fiber.Ctx) error {

	var number = rand.Intn(100)
	var result = ""
	if primeNumber(number) {
		fmt.Println("you're success releasing pokemon", number)
		result = "success"
	} else {
		fmt.Println("you're failed release pokemon", number)
		result = "failed"
	}

	return c.Status(200).JSON(fiber.Map{
		"Message": "done",
		"data":    result,
	})
}
