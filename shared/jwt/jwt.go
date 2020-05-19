package jwt

import "net/http"

type jwtProvider struct {
}

type JwtProvider interface {
	GenerateToken() (string, error)
	DecodeToken(w http.ResponseWriter)
}
