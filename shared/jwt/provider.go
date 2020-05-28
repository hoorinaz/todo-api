package jwt

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	Data struct {
		Username string
		Email    string
	}
	claims struct {
		Data
		jwt.StandardClaims
	}
)

func (jp jwtProvider) GenerateTime() time.Time {

	return time.Now().Add(time.Hour * 6)
}

func (jp jwtProvider) GenerateToken(d Data) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		Data: d,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jp.expirationDate.Unix(),
		},
	})
	stringToken, err := token.SignedString(jp.signKey)
	return stringToken, err
}

func (jp jwtProvider) DecodeToken(tokenStr string) (Data, error) {
	c := claims{}
	token, err := jwt.ParseWithClaims(tokenStr, &c, func(token *jwt.Token) (interface{}, error) {
		return jp.signKey, nil
	})
	if err != nil {
		log.Println("decode token has error")
		return Data{}, err
	}
	if !token.Valid {
		log.Println("token is invalid")
		return Data{}, ErrTokenInvalid
	}
	return c.Data, nil
}

func NewJwtProvider() JwtProvider {
	return jwtProvider{
		signKey:        []byte("hoorieNazari"),
		expirationDate: time.Now().Add(time.Hour * 6),
	}
}
