package logic

import (
	"go-web-app/dao/mysql"
	"go-web-app/models"
	"go-web-app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// check if user existed
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// generate UID
	userID := snowflake.GenID()
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// write into database
	return mysql.InsertUser(user)
}
