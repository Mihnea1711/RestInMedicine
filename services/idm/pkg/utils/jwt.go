package utils

import (
	"errors"
	"fmt"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/config"
)

type MyCustomClaims struct {
	jwt.RegisteredClaims
}

func CreateJWT(userRole string, jwtConfig config.JWTConfig) (string, error) {
	hmacSampleSecret := []byte(jwtConfig.Secret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		// A usual scenario is to set the expiration time relative to the current time
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 30 * time.Hour)), // expiry date set to 30 days from now
		IssuedAt:  jwt.NewNumericDate(time.Now()),                          // issue data
		NotBefore: jwt.NewNumericDate(time.Now()),                          // not valid before
		Subject:   userRole,                                                // data to be included in the jwt body
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		err_msg := fmt.Sprintf("error signing token: %v", err)
		log.Printf("[IDM] %s", err_msg)
		return "", fmt.Errorf(err_msg)
	}

	log.Printf("[IDM] JWT token created successfully.")
	return tokenString, nil
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
		/*
			// sub, err := claims.GetSubject()
			// if err != nil {
			// 	err_msg := fmt.Sprintf("error getting subject: %v", err)
			// 	log.Printf("[IDM] %s", err_msg)
			// }
			// exp, err := claims.GetExpirationTime()
			// if err != nil {
			// 	err_msg := fmt.Sprintf("error getting expiry date: %v", err)
			// 	log.Printf("[IDM] %s", err_msg)
			// }
			// date := exp.Format("2006-01-02")
			// log.Printf("[IDM] Subject: %v, Expiry Date: %v", sub, date)
		*/
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
