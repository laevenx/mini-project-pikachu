package main

import (
	"fmt"
	"os"
	"strconv"

	"mini-project-pikachu/routes"

	"mini-project-pikachu/config/mysql"

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
func init() {
	// loading databases
	fmt.Println("connecting mysql bni")

	port, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		fmt.Println(err)
	}
	// setup mysql databases
	database := mysql.Config{
		Host: os.Getenv("MYSQL_HOST"),
		DB:   os.Getenv("MYSQL_DATABASE"),
		Port: port,
		User: os.Getenv("MYSQL_USER"),
		Pass: os.Getenv("MYSQL_PASS"),
	}

	err = database.Connect()
	if err != nil {
		panic(err)
	}
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
