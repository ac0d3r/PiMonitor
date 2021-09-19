package router

import (
	"pimonitor/device/raspivid"
	"pimonitor/handler"
	"pimonitor/handler/middleware"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func AddRoute(engine *gin.Engine) {
	engine.Use(static.Serve("/", static.LocalFile("frontend", false)))
	engine.POST("/login", handler.UserLogin)

	addAPIRoute(engine.Group("api"))
}

func AddWebSocketStream(engine *gin.Engine, opts raspivid.Options) {
	streamGroup := engine.Group("/vi")
	streamGroup.Use(middleware.TokenAuth())
	{
		ws := raspivid.NewWebSocketHandler(opts)
		streamGroup.GET("/stream", ws.Handler)
	}
}

func addAPIRoute(group *gin.RouterGroup) {
	userGroup := group.Group("/user")
	userGroup.Use(middleware.TokenAuth())
	{
		userGroup.POST("/logout", handler.UserLogout)
		userGroup.POST("/update", handler.UserUpdateInfomation)
	}
}
