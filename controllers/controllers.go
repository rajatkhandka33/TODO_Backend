package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Todo represents a task item
type Todo struct {
	Id        string `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

// In-memory storage for todos
var Todos []Todo

// Generates a unique ID for each todo
func generateRandomId() string {
	return uuid.New().String()
}

// Sets up the /todos routes
func TodosController(r *gin.Engine) {
	x := r.Group("/todos")

	// GET /todos - List all todos
	x.GET("/", func(c *gin.Context) {
		if len(Todos) == 0 {
			c.JSON(http.StatusOK, gin.H{"message": "No tasks yet!"})
			return
		}
		c.JSON(http.StatusOK, Todos)
	})

	// POST /todos - Create a new todo
	x.POST("/", func(c *gin.Context) {
		var todo Todo
		if err := c.BindJSON(&todo); err != nil {
			log.Println("Invalid JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		todo.Id = generateRandomId()
		Todos = append(Todos, todo)

		c.JSON(http.StatusCreated, todo)
	})

	// GET /todos/:id - Get a specific todo by ID
	x.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, todo := range Todos {
			if todo.Id == id {
				c.JSON(http.StatusOK, todo)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "No such task found"})
	})

	// PUT /todos/:id - Update a todo
	x.PUT("/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, todo := range Todos {
			if todo.Id == id {
				var updatedTodo Todo
				if err := c.BindJSON(&updatedTodo); err != nil {
					log.Println("Invalid JSON:", err)
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
					return
				}
				updatedTodo.Id = id // Keep the original ID
				Todos[i] = updatedTodo
				c.JSON(http.StatusOK, updatedTodo)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "No such task found"})
	})

	// DELETE /todos/:id - Delete a todo
	x.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, todo := range Todos {
			if todo.Id == id {
				Todos = append(Todos[:i], Todos[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "No such task found"})
	})
}
