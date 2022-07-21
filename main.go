package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("cek")
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/player", playerHandler)
	router.Run()
}

func rootHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"nama": "ilham adikuusuma",
		"nim":  "202410103034",
	})
}

func playerHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"id":    "20",
		"uname": "hamadi",
	})
}
