package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
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
		gr := convertkegameresponse(j)
		penampung = append(penampung, gr)

	}
	c.JSON(http.StatusOK, gin.H{
		"data": penampung,
	})
}

func (h *gameHandler) GetGame(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr) //convert jadi id
	gim, err := h.gs.S_FindById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors:": err,
		})
	}

	gimResponse := convertkegameresponse(gim)

	c.JSON(http.StatusOK, gin.H{
		"data": gimResponse,
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
func (h *gameHandler) UpdateGame(c *gin.Context) {
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

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr) //convert jadi id
	gim, err := h.gs.S_Update(id, gr)
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

func convertkegameresponse(g game.Game) game.GameResponse {

	return game.GameResponse{
		ID:         g.ID,
		Judul:      g.Judul,
		TahunRilis: g.TahunRilis,
		Harga:      g.Harga,
		Genre:      g.Genre,
	}
}
