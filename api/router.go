package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iissy/gooa/controller"
)

func Start() {
	app := gin.New()
	app.Use(gin.Recovery())
	api := app.Group("/api")

	article := api.Group("/ding")
	{
		article.GET("/getuserinfo", controller.Getuserinfo)
	}

	app.Run(":8080")
}
