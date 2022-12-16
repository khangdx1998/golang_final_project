// package middleware
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Gmail string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string, secret_key string) (string, error) {
	var jwtKey = []byte("my_secret_key")
	var iatTime = time.Now()
	var expirationTime = time.Now().Add(10 * time.Minute)
	claims := &Claims{
		Gmail: email,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt: jwt.NewNumericDate(iatTime),
		},
	}
	// Token with signing algorithm and payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// JWT Token include signing algorithm, payload, signing secret key
	jwt, _ := token.SignedString(jwtKey)
	
	return jwt, nil
}

func DecodeJWT(jwt_token string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(jwt_token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")	 
		}
		return []byte("my_secret_key"), nil
	})
	if err != nil {
		return nil, err
	}

	myClaims := token.Claims.(*Claims)
	return myClaims, nil
}


func main(){
	jwt, _ := GenerateJWT("khangdx@gmail.com", "khangngocmy")
	fmt.Println(jwt)
	var jwt_token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImtoYW5nZHhAZ21haWwuY29tIiwiZXhwIjoxNjcxMTI5NDc1LCJpYXQiOjE2NzExMjg4NzV9.Ir-kPFUTJOjmaZsWq9ZC93LRDrPsTWD7EPvxkD314dM"
	claims, _ := DecodeJWT(jwt_token)
	json_string, _ := json.Marshal(claims)
	fmt.Println(string(json_string))
}
