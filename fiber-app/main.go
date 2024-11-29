package main

import (
	"gowebcompare/todo"
	"strconv"

	fiber "github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

var TodoList *todo.TodoList

func main() {
	TodoList = todo.New()
	app := fiber.New()

	myLogger := logger.New()
	app.Use(myLogger)
	api := app.Group("/api")
	api.Get("/get_list", getList)
	api.Get("/get_todo/:id", getTodo)
	api.Post("/create_todo", createTodo)
	api.Put("/update_todo", updateTodo)
	api.Delete("/delete_todo/:id", deleteTodo)

	log.Fatal(app.Listen(":8080"))

}

func getList(c fiber.Ctx) error {
	list := TodoList.GetList()
	if err := c.JSON(list); err != nil {
		c.JSON(fiber.Map{"error": err.Error()})
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}

func getTodo(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.JSON(fiber.Map{"error": err.Error()})
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	todo, err := TodoList.GetTodo(id)
	if err != nil {
		c.JSON(fiber.Map{"error": err.Error()})
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if err = c.JSON(todo); err != nil {
		c.JSON(fiber.Map{"error": err.Error()})
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}

func createTodo(c fiber.Ctx) error {
	todo := new(todo.Todo)
	if err := c.Bind().JSON(todo); err != nil {
		c.SendString(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if err := TodoList.CreateTodo(*todo); err != nil {
		c.JSON(fiber.Map{"error": err.Error()})
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.SendStatus(fiber.StatusOK)
}

func updateTodo(c fiber.Ctx) error {
	todo := new(todo.Todo)
	if err := c.Bind().JSON(todo); err != nil {
		c.JSON(fiber.Map{"error": err.Error()})
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if err := TodoList.UpdateTodo(*todo); err != nil {
		c.JSON(fiber.Map{"error": err.Error()})
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.SendStatus(fiber.StatusOK)
}

func deleteTodo(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.JSON(fiber.Map{"error": err.Error()})
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if err = TodoList.DeleteTodo(id); err != nil {
		c.JSON(fiber.Map{"error": err.Error()})
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.SendStatus(fiber.StatusOK)
}
