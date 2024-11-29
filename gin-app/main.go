package main

import (
	"gowebcompare/todo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var TodoList *todo.TodoList

func main() {
	TodoList = todo.New()
	router := gin.Default()

	api := router.Group("/api")
	api.GET("/get_list", getList)
	api.GET("/get_todo/:id", getTodo)
	api.POST("/create_todo", createTodo)
	api.PUT("/update_todo", updateTodo)
	api.DELETE("/delete_todo/:id", deleteTodo)

	router.Run(":8080")
}

func getList(c *gin.Context) {
	list := TodoList.GetList()
	c.JSON(http.StatusOK, list)
}

func getTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	todo, err := TodoList.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func createTodo(c *gin.Context) {
	todo := new(todo.Todo)
	if err := c.ShouldBindJSON(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := TodoList.CreateTodo(*todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func updateTodo(c *gin.Context) {
	todo := new(todo.Todo)
	if err := c.ShouldBindJSON(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := TodoList.UpdateTodo(*todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func deleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := TodoList.DeleteTodo(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
