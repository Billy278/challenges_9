package server

import (
	"challenges_9/module/router"

	"github.com/gin-gonic/gin"
)

func NewServer() {
	bookCtrl := InitControllers()
	rt := gin.Default()
	//init middleware
	rt.Use(gin.Recovery(), gin.Logger())
	router.NewRouter(rt, bookCtrl.BookCtr)
	rt.Run(":8080")
}
