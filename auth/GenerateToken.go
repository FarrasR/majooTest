package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	jwt.StandardClaims
}


func GenerateToken(username string, password string) (string, bool) {


	if(username != "admin" && password != "admin") {
		return "nil" ,false;
	}

	claims := &jwtCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Secret key should have been on env
	t, err := token.SignedString([]byte("secret key"))
	if err != nil {
		panic(err)
	}
	return t , true
}


