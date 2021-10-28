package repo

import (
	"errors"
)

var (
	BadRequest error = errors.New("bad request")
	InvalidId  error = errors.New("invalid id")
	NotFound   error = errors.New("card not found")
	CardEmpty  error = errors.New("card is empty")
)
