package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("User existed")
	ErrorUserNotExist    = errors.New("User is not existed")
	ErrorInvalidPassword = errors.New("Wrong password")
	ErrorInvalidID       = errors.New("Invalid ID")
)
