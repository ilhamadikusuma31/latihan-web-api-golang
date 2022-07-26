package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"web-api-golang/game"
	"web-api-golang/handler"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/game-pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	//bikin table + fitur crud
	gimRepo := game.NewRepository(db)

	//bikin fitur crud
	gimService := game.NewService(gimRepo)

	//bikin handler (route) untuk game
	obj := handler.NewGameHandler(gimService)

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/game", obj.GetGames)
	v1.GET("/game/:id", obj.GetGame)
	v1.POST("/game", obj.PostGameHandler)

	router.Run()

	//main
	//handler
	//service
	//repository
	//db
	//mysql
}
