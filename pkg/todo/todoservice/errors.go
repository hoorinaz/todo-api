package todoservice

import "errors"

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrNotFound     = errors.New("Not Found Todo")
)
