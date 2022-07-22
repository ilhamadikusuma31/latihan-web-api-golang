package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"web-api-golang/handler"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("database berhasil konek")

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/", handler.RootHandler)
	v1.GET("/developer", handler.DeveloperHandler)
	v1.GET("/game/:id", handler.GameHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/game", handler.PostGameHandler)

	router.Run()
}
