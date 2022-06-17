package routes

import (
	"mini-project-pikachu/handler"

	"github.com/gofiber/fiber/v2"
)

func MainRouter(app fiber.Router) {

	// app.Get("/list", handler.List)
	// app.Get("/mylist",handler.MyList)
	// app.Get("/detail",handler.Detail)

	app.Post("/catch", handler.Catch)
	app.Post("/release", handler.Release)
	app.Put("/rename", handler.Rename)

}
