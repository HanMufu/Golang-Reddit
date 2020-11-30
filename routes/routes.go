package routes

import (
	"go-web-app/controller"
	"go-web-app/logger"
	"go-web-app/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := r.Group("/api/v1")

	// user register
	v1.POST("/signup", controller.SignUpHandler)

	// user login
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware())

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)
		v1.GET("/posts2", controller.GetPostListHandler2)

		v1.POST("/vote", controller.PostVoteHandler)
	}

	//r.GET("/", func(context *gin.Context) {
	//	context.String(http.StatusOK, "ok")
	//})

	//r.GET("/ping", middlewares.JWTAuthMiddleware(), func(context *gin.Context) {
	//	context.String(http.StatusOK, "pong")
	//})

	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"msg": "404",
		})
	})
	return r
}

//// JWTAuthMiddleware 基于JWT的认证中间件
//func JWTAuthMiddleware() func(c *gin.Context) {
//	return func(c *gin.Context) {
//		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
//		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
//		// 这里的具体实现方式要依据你的实际业务情况决定
//		authHeader := c.Request.Header.Get("Authorization")
//		if authHeader == "" {
//			c.JSON(http.StatusOK, gin.H{
//				"code": 2003,
//				"msg":  "请求头中auth为空",
//			})
//			c.Abort()
//			return
//		}
//		// 按空格分割
//		parts := strings.SplitN(authHeader, " ", 2)
//		if !(len(parts) == 2 && parts[0] == "Bearer") {
//			c.JSON(http.StatusOK, gin.H{
//				"code": 2004,
//				"msg":  "请求头中auth格式有误",
//			})
//			c.Abort()
//			return
//		}
//		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
//		mc, err := jwt.ParseToken(parts[1])
//		if err != nil {
//			c.JSON(http.StatusOK, gin.H{
//				"code": 2005,
//				"msg":  "无效的Token",
//			})
//			c.Abort()
//			return
//		}
//		// 将当前请求的username信息保存到请求的上下文c上
//		c.Set("userID", mc.UserID)
//		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
//	}
//}
