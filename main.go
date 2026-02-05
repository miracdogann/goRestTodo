package main

import (
	"fiber_rest/dal"
	"fiber_rest/database"
	"fiber_rest/services"

	"github.com/gofiber/fiber/v3"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&dal.Todo{}) // veritabanı migrationları yapar
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		// return c.SendString("hello merhaba")
		return c.JSON(fiber.Map{
			"message": "İstek Başarılı",
		})
	})

	app.Post("/todo", services.CreateTodo)
	app.Get("/todos", services.GetTodos)
	app.Get("/todos/:todoID", services.GetTodo)
	app.Put("/todos/:todoID", services.UpdateTodo)
	app.Delete("/todos/:todoID", services.DeleteTodo)
	app.Listen(":3000")
}
