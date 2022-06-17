package handler

import (
	"mini-project-pikachu/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MyPoke struct {
	Id       int
	Name     string
	Nickname string
}

func Fibonacci(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]

}
func Rename(c *fiber.Ctx) error {
	var num = 0
	var name = c.Query("nickname")
	var loop = c.Query("num")
	num, err := strconv.Atoi(loop)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	nickname, err := usecase.Rename(name, num)

	return c.Status(200).JSON(fiber.Map{
		"Message": "done",
		"data":    nickname,
	})
}
