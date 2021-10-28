package service

import (
	"errors"
)

var (
	BadRequest  error = errors.New("bad request")
	InvalidId   error = errors.New("invalid id")
	InvalidType error = errors.New("invalid type")
)
