package controller

import (
	"go-web-app/logic"
	"go-web-app/models"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	// get parameters and validate them
	//var p models.ParamSignUp
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("signup with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "Bad signup parameters",
		})
		return
	}
	// check if any parameter is empty
	if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
		zap.L().Error("signup with invalid param")
		c.JSON(http.StatusOK, gin.H{
			"msg": "Bad signup parameters",
		})
		return
	}
	zap.L().Info("User signup successfully")
	// business logic
	logic.SignUp(p)
	// return responses
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
