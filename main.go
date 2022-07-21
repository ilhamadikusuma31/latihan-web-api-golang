package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web-api-golang/handler"
)

func main() {
	fmt.Println("cek")
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/", handler.RootHandler)
	v1.GET("/developer", handler.DeveloperHandler)
	v1.GET("/game/:id", handler.GameHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/game", handler.PostGameHandler)

	router.Run()
}
