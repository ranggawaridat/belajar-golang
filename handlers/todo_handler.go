package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ranggawaridat/belajar-golang/models"
)

var todos = []models.Todo{
	{ID: 1, Title: "Belajar Gin", Done: false},
}

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func GetTodoByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "todo not found",
	})
}

func CreateTodo(c *gin.Context) {
	var newTodo models.Todo

	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newTodo.ID = len(todos) + 1

	todos = append(todos, newTodo)

	c.JSON(http.StatusCreated, newTodo)
}

func UpdateTodo(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	var updatedTodo models.Todo

	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			updatedTodo.ID = id
			todos[i] = updatedTodo

			c.JSON(http.StatusOK, updatedTodo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "todo not found",
	})
}

func DeleteTodo(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)

			c.JSON(http.StatusOK, gin.H{
				"message": "todo deleted",
			})

			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "todo not found",
	})
}
