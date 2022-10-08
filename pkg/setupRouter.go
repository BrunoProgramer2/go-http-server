package pkg

import (
	"BrunoProgramer2.github.io/go-http-server/routes"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", routes.Home)
	r.GET("/app.js", routes.JsRender)
	return r
}
