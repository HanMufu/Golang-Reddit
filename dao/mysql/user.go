package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"go-web-app/models"
)

const secret = "hanmufu.com"

func CheckUserExist(username string) (err error) {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("User existed")
	}
	return
}

func InsertUser(user *models.User) (err error) {
	// encrypt password
	password := encryptPassword(user.Password)
	sqlStr := "insert into user(user_id, username, password) values (?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := "select user_id, username, password from user where username=?"
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return errors.New("User is not existed")
	}
	if err != nil {
		return
	}
	password := encryptPassword(oPassword)
	if password != user.Password {
		return errors.New("Wrong password")
	}
	return
}
