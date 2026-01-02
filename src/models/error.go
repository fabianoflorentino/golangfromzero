package models

import "errors"

var (
	ErrNameBlank     error = errors.New("name can not be blank")
	ErrEmailBlank    error = errors.New("email can not be blank")
	ErrPasswordBlank error = errors.New("password can not be blank")
)
