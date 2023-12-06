package middleware

import (
	"errors"
	"fmt"
	"log"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/config"
)

type MyCustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func ParseJWT(tokenString string, jwtConfig config.JWTConfig) (*MyCustomClaims, error) {
	hmacSampleSecret := []byte(jwtConfig.Secret)

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err_msg := fmt.Sprintf("unexpected signing method: %v", token.Header["alg"])
			log.Printf("[IDM] %s", err_msg)
			return nil, fmt.Errorf(err_msg)
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		log.Println("[IDM] That's not even a token")
		return nil, fmt.Errorf("token is malformed: %w", err)
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		log.Println("[IDM] Invalid signature")
		return nil, fmt.Errorf("invalid token signature: %w", err)
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		log.Println("[IDM] Timing is everything")
		return nil, fmt.Errorf("token is either expired or not active yet: %w", err)
	} else {
		err_msg := fmt.Sprintf("error getting claims: %v", err)
		log.Printf("[IDM] %s", err_msg)
		return nil, fmt.Errorf(err_msg)
	}
}
