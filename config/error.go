package config

import "errors"

var (
	ErrLoadEnv error = errors.New("Error loading .env file")
)
