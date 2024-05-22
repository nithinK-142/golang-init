package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "here we GO!!")
	})
	router.GET("/todos", getTodos)
	router.POST("/todos", createTodo)
	router.PUT("/todos/:id", editTodo)
	router.DELETE("/todos/:id", deleteTodo)
	router.Run(":" + port)
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

func editTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Todo ID"})
		return
	}

	var editedTodo Todo
	// catch error when creating newtodo
	if err := c.BindJSON(&editedTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = editedTodo
			todos[i].ID = id
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Todo edited successfully"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func deleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Todo ID"})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			fmt.Println(todos[:i])
			fmt.Println(todos[i+1:])
			todos = append(todos[:i], todos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}
