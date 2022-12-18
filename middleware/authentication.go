package middleware

import (
	"errors"
	"time"
	"final_project/config"
	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	Gmail string `json:"email"`
	jwt.RegisteredClaims
}

var env config.Config = config.Get_config()

func GenerateJWT(email string) (string, error) {
	var jwtKey = []byte(env.Secret_key)
	var iatTime = time.Now()
	var expirationTime = time.Now().Add(10 * time.Minute)
	claims := &MyClaims{
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

func DecodeJWT(jwt_token string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(jwt_token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")	 
		}
		return []byte(env.Secret_key), nil
	})
	if err != nil {
		return nil, err
	}

	myClaims := token.Claims.(*MyClaims)
	return myClaims, nil
}