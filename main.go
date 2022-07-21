package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("cek")
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"nama": "ilham adikuusuma",
			"nim":  "202410103034",
		})
	})

	router.GET("/player", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"id":    "20",
			"uname": "hamadi",
		})
	})

	router.Run()
}
