package main

import (
  "belajar-restful/config"
  "belajar-restful/handlers"

  "github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New()

  config.Connect()

  app.Get("/user", handlers.GetUser)
  app.Post("/user", handlers.AddUser)
  app.Get("/user/:id", handlers.GetUserById)
  app.Delete("/user/:id", handlers.DeleteUserById)
  app.Put("/user/:id", handlers.UpdateUser)

  app.Listen(":3000")
}
