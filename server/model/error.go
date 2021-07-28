package model

import "errors"

var (
	ERROR_USER_NOTEXIST = errors.New("User not exist")
	ERROR_USER_EXISTS   = errors.New("User has exist")
	ERROR_USER_PWD      = errors.New("Wrong passwrod")
)
