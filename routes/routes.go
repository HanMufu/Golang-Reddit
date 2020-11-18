package routes

import (
	"go-web-app/controller"
	"go-web-app/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// user register
	r.POST("/signup", controller.SignUpHandler)

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})

	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"msg": "404",
		})
	})
	return r
}
