package exceptions

import "errors"

var (
	ErrNotFound   = errors.New("item not found")
	ErrBadRequest = errors.New("bad request")
)
