package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	fmt.Println("cek")
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/player", playerHandler)
	router.GET("/game/:id", gameHandler)
	router.GET("/query", queryHandler)
	router.POST("/game", postGameHandler)

	router.Run()
}

func rootHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"nama": "ilham adikusuma",
		"nim":  "202410103034",
	})
}

func playerHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"id":    "20",
		"uname": "hamadi",
	})
}

func gameHandler(context *gin.Context) {
	tangkapId := context.Param("id")
	context.JSON(http.StatusOK, gin.H{
		"id": tangkapId,
	})
}

type gameInput struct {
	Judul    string
	Harga    int
	SubJudul string `json:"sub_judul"`
}

func postGameHandler(c *gin.Context) {
	var gi gameInput
	err := c.ShouldBindJSON(&gi)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"judul":     gi.Judul,
		"harga":     gi.Harga,
		"sub_judul": gi.SubJudul,
	})

}
func queryHandler(context *gin.Context) {
	tangkapQuery := context.Query("judul")
	context.JSON(http.StatusOK, gin.H{
		"judul": tangkapQuery,
	})
}
