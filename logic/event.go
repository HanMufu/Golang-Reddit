package logic

import (
	"go-web-app/dao/mongodb"
	"go-web-app/models"

	"go.uber.org/zap"
)

func GetEventByID(eid int64) (event *models.Event, err error) {
	event, err = mongodb.GetEventById(eid)
	if err != nil {
		zap.L().Error("mysql.GetEventById(pid) failed", zap.Int64("Event ID", eid), zap.Error(err))
		return
	}
	return
}
