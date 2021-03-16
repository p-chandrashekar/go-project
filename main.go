package main

import "github.com/gofiber/fiber/v2"

// Here is the repo name go-with-jwt-auth

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8000")
}
