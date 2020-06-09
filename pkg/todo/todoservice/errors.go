package todoservice

import "errors"

// type(
// 	ErrUnauthorized struct {
// 		text: "Unauthorized",
// 		status: 401
// 	}
// 	ErrNotFound struct {
// 		text: "Not Found Todo",
// 		status : 404
// 	}

// )

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrNotFound     = errors.New("Not Found Todo")
)
