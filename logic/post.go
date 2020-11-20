package logic

import (
	"go-web-app/dao/mysql"
	"go-web-app/models"
	"go-web-app/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	p.PostID = snowflake.GenID()
	return mysql.CreatePost(p)
}
