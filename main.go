package main

import (
	"fmt"
	"os"

	"mini-project-pikachu/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/subosito/gotenv"
)

func init() {
	// loading the environment
	envFile := ".env"
	_, err := os.Stat(envFile)
	if !os.IsNotExist(err) {
		gotenv.Load(envFile)
	}

	fmt.Println("reading .env")
}

func main() {

	app := fiber.New()

	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("yay!!"))
	})
	routes.MainRouter(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
		// => 404 "Not Found"
	})

	port := ":" + os.Getenv("WEB_PORT")

	app.Listen(port)
}
