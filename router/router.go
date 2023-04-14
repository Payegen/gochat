package router

import (
	"gochat/service"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// log
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// r.Static("../public", "../public")
	r.GET("/ping", service.GetIndex)
	r.GET("/index", service.GetHtml)
	r.POST("/upload", service.UpLoadFile)
	r.GET("/sendmsg", service.WsInit)
	return r
}
