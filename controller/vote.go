package controller

import (
	"go-web-app/logic"
	"go-web-app/models"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func PostVoteHandler(c *gin.Context) {
	p := new(models.ParamVoteData)
	if err := c.ShouldBind(p); err != nil {
		validationError, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.L().Error("err.(validator.ValidationErrors)", zap.Error(validationError))
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(validationError.Translate(trans))
		zap.L().Error("c.ShouldBind(p) failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	userId, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
	}
	if err := logic.VoteForPost(userId, p); err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
