package logic

import (
	"go-web-app/dao/mysql"
	"go-web-app/models"
	"go-web-app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) {
	// check if user existed
	mysql.QueryUserByUsername()
	// generate UID
	snowflake.GenID()
	// write into database
	mysql.InsertUser()
}
