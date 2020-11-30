package controller

import (
	"go-web-app/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetEventHandler(c *gin.Context) {
	eidStr := c.Param("id")
	eid, err := strconv.ParseInt(eidStr, 10, 64)
	if err != nil {
		zap.L().Error("Get event by id failed.", zap.Error(err))
		zap.L().Error("eid", zap.Int64("eid", eid))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetEventByID(eid)
	if err != nil {
		zap.L().Error("logic.GetEventByID(eid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
