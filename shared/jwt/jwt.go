package jwt

import (
	"time"
)

type jwtProvider struct {
	signKey []byte
	expirationDate time.Time
}

type JwtProvider interface {
	GenerateToken(d Data) (string, error)
	DecodeToken(tokenstr string) (Data, error)
}
