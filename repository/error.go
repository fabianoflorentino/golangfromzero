package repository

import "errors"

var (
	ErrEmailAlreadyExist error = errors.New("email already used")
	ErrPgCode                  = "23505"
)
