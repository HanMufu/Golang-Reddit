package logic

import (
	"go-web-app/dao/mysql"
	"go-web-app/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}
