package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "todo!",
	})
}

func getTodoDetail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "todo_detail!",
	})
}

func main() {
	router := gin.Default()

	router.GET("/todo", getTodo)
	router.GET("/todo/detail", getTodoDetail)

	router.Run()
}

// package controller

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func getTodo(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "todo",
// 	})
// }

// func getTodoDetail(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "todo_detail",
// 	})
// }
