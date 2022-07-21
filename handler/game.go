package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"web-api-golang/game"
)

//nama func jadiin kapital biar public
func RootHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"nama": "ilham adikusuma",
		"nim":  "202410103034",
	})
}

func DeveloperHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"id":    "20",
		"uname": "konami",
	})
}

func GameHandler(context *gin.Context) {
	tangkapId := context.Param("id")
	context.JSON(http.StatusOK, gin.H{
		"id": tangkapId,
	})
}

func QueryHandler(context *gin.Context) {
	tangkapQuery := context.Query("judul")
	context.JSON(http.StatusOK, gin.H{
		"judul": tangkapQuery,
	})
}

func PostGameHandler(c *gin.Context) {
	var gi game.GameInput
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
