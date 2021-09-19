package main

import (
	"log"
	"net/http"
	"pimonitor/database"
	"pimonitor/device/raspivid"
	"pimonitor/device/servo"
	"pimonitor/router"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.Init("pi.sqlite3"); err != nil {
		log.Fatalln(err)
	}
	// 初始化舵机
	servo.Init("16", "12", 15, 35, 10, 40)
	// 设置 raspivid 参数
	options := raspivid.Options{
		Width:          960,
		Height:         540,
		FPS:            30,
		HorizontalFlip: false,
		VerticalFlip:   false,
		Rotation:       0,
	}
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	router.AddRoute(engine)
	router.AddWebSocketStream(engine, options)
	engine.NoRoute(func(c *gin.Context) { c.Status(http.StatusNotFound) })
	engine.NoMethod(func(c *gin.Context) { c.Status(http.StatusMethodNotAllowed) })

	if err := engine.Run(":8000"); err != nil {
		log.Printf("[Main] Gin engine error: %v", err)
	}
}
