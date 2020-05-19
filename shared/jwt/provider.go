package jwt

import (
	jwt2 "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

var Key = []byte("The_Secret_Key")
type Claims struct {
	Username string
	Email    string
	jwt2.StandardClaims
}

func (cl Claims) GenerateToken()(string , error){

	token := jwt2.NewWithClaims(jwt2.SigningMethodHS256, cl)
	stringToken, err := token.SignedString(Key)
	return stringToken, err
}
func(cl Claims) SetToken(w http.ResponseWriter)  {
	stringToken, err:= cl.GenerateToken()
	if err!=nil{
		log.Println("Generate Token Failed! error: ", err.Error())
		return
	}
	w.Header().Set("Authorization", stringToken)

}


func NewJwtProvider(username string, email string) JwtProvider{
	expTime:= time.Now().Add(6 *time.Hour)
		return Claims{
		Username: username,
		Email:    email,
		StandardClaims : jwt2.StandardClaims{
			ExpiresAt: expTime.Unix(),

		},
	}
}