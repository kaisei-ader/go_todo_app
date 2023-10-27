package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET TODO!",
	})
}

func postTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "POST TODO!",
	})
}

func main() {
	router := gin.Default()

	todo := router.Group("/todo")
	{
		todo.GET("", getTodo)
		todo.POST("", postTodo)
	}

	router.Run()
}
