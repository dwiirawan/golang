package main

import (
	"belajar-restful/config"
	"belajar-restful/handlers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func authHandler(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()

	fmt.Println(headers, "<<< iki headers")

	return ctx.Next()
}

func main() {
	app := fiber.New()

	config.Connect()

	private := app.Group("api", func(c *fiber.Ctx) error {
		return authHandler(c)
	})

	private.Get("private", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("success")
	})

	user := app.Group("/user")

	user.Get("/", handlers.GetUser)
	user.Post("/", handlers.AddUser)
	user.Get("/:id", handlers.GetUserById)
	user.Get("/:nama", handlers.GetNama)
	user.Delete("/:id", handlers.DeleteUserById)
	user.Put("/:id", handlers.UpdateUser)

	app.Listen(":3000")
}
