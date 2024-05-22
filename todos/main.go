package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID          int    `json:"id"`
	IsDone      bool   `json:"isDone"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var todos = []Todo{}
var nextTodoID = 1

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", createTodo)
	router.Run(":8080")
}

func getTodos(c *gin.Context) {
	// returns todos slice
	c.IndentedJSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var newTodo Todo

	// catch error when creating newtodo
	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newTodo.ID = nextTodoID
	nextTodoID++
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}
