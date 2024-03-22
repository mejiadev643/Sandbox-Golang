package jwtpersonal

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(w http.ResponseWriter, r *http.Request) {
	// Create the Claims// equal to the payload
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user"] = "MejiaDev"
	claims["permisions"] = []string{"read", "write", "delete"}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the secret

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}
	// Print the token
	println(tokenString)
	w.Write([]byte(tokenString))
}
