package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func main() {
	fmt.Println("cek")
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/", rootHandler)
	v1.GET("/player", playerHandler)
	v1.GET("/game/:id", gameHandler)
	v1.GET("/query", queryHandler)
	v1.POST("/game", postGameHandler)

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
	Judul string `json:"judul" binding:"required"`
	Harga int    `json:"harga" binding:"required"`
}

func postGameHandler(c *gin.Context) {
	var gi gameInput
	err := c.ShouldBindJSON(&gi)
	if err != nil {
		pesanErrors := []string{}
		for _, j := range err.(validator.ValidationErrors) {
			pesanError := fmt.Sprintf("Error di %s, kondisi: %s", j.Field(), j.ActualTag())
			pesanErrors = append(pesanErrors, pesanError)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": pesanErrors,
		}) //biar kalo error servernya ga mati
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"judul": gi.Judul,
		"harga": gi.Harga,
	})

}
func queryHandler(context *gin.Context) {
	tangkapQuery := context.Query("judul")
	context.JSON(http.StatusOK, gin.H{
		"judul": tangkapQuery,
	})
}
