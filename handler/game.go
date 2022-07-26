package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"web-api-golang/game"
)

type gameHandler struct {
	gs game.Service
}

func NewGameHandler(param game.Service) *gameHandler {
	return &gameHandler{param}
}

//func (h *gameHandler) RootHandler(context *gin.Context) {
//	context.JSON(http.StatusOK, gin.H{
//		"nama": "ilham adikusuma",
//		"nim":  "202410103034",
//	})
//}
//
//func (h *gameHandler) DeveloperHandler(context *gin.Context) {
//	context.JSON(http.StatusOK, gin.H{
//		"id":    "20",
//		"uname": "konami",
//	})
//}
//
//func (h *gameHandler) GameHandler(context *gin.Context) {
//	tangkapId := context.Param("id")
//	context.JSON(http.StatusOK, gin.H{
//		"id": tangkapId,
//	})
//}
//
//func (h *gameHandler) QueryHandler(context *gin.Context) {
//	tangkapQuery := context.Query("judul")
//	context.JSON(http.StatusOK, gin.H{
//		"judul": tangkapQuery,
//	})
//}

func (h *gameHandler) GetGames(c *gin.Context) {
	gims, err := h.gs.S_FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors:": err,
		})
	}

	var penampung []game.GameResponse
	for _, j := range gims {
		br := game.GameResponse{
			ID:         j.ID,
			Judul:      j.Judul,
			TahunRilis: j.TahunRilis,
			Harga:      j.Harga,
			Genre:      j.Genre,
		}
		penampung = append(penampung, br)

	}
	c.JSON(http.StatusOK, gin.H{
		"data": penampung,
	})
}

func (h *gameHandler) PostGameHandler(c *gin.Context) {
	var gr game.GameRequest
	err := c.ShouldBindJSON(&gr)
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

	gim, err := h.gs.S_Create(gr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		}) //biar kalo error servernya ga mati
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gim,
	})

}
