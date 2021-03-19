package main

import (
	"github.com/gofiber/fiber"
	"github.com/jnprogrammer/go_fiber/book"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Buy Bitcoin!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.Deletebook)
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(3000)
}
