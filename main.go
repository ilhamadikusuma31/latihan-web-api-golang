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
	gimObj := game.NewRepository(db)
	//gim, err := gimObj.FindAll()
	//if err != nil {
	//	fmt.Println("gagal")
	//}
	//for _, g := range gim {
	//	fmt.Println(g.Harga)
	//}

	gim := game.Game{
		Judul:      "watch dog",
		TahunRilis: 2015,
		Harga:      3000,
		Genre:      "",
	}
	gimObj.Create(gim)

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/", handler.RootHandler)
	v1.GET("/developer", handler.DeveloperHandler)
	v1.GET("/game/:id", handler.GameHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/game", handler.PostGameHandler)

	router.Run()
}
