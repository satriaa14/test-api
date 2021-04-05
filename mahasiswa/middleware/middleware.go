package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/satriaa14/test-api/mahasiswa/src/github.com/dgrijalva/jwt-go"

	"github.com/satriaa14/test-api/mahasiswa/operation"
)

// Cors Setup
func MiddlewareJwtAuth(w http.ResponseWriter, r *http.Request) error {

	headerAuth := r.Header.Get("Authorization")
	if !strings.Contains(headerAuth, "Bearer") {
		return fmt.Errorf("No token parsed")
	}
	tokenString := strings.Replace(headerAuth, "Bearer ", "", -1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != operation.JWTSignMethod {
			return nil, fmt.Errorf("Signing method invalid")
		}
		return operation.JWTSignatureKey, nil
	})

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return fmt.Errorf("Token Not Valid")
	}

	ctx := context.WithValue(context.Background(), "userInfo", claims)
	r = r.WithContext(ctx)

	return nil
}
