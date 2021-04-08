package operation

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const (
	constUsername string = "test.api"
	constPassword string = "02062021"
)

var (
	AppName          = "Mahasiswa Service"
	JWTLoginDuration = time.Duration(30) * time.Minute
	JWTSignMethod    = jwt.SigningMethodHS256
	JWTSignatureKey  = []byte("test.mhs.key")
)

// ConnectClaims is open id connect jwt claims
type ConnectClaims struct {
	jwt.StandardClaims
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	UserName string `json:"user,omitempty"`
}

type LoginRequest struct {
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token,omitempty"`
}

// Authenticate
func Authenticate(user, pass string) ([]byte, error) {
	hash, _ := hashGenerator(constUsername + constPassword)
	userPassword := user + pass
	ok := compareHash(userPassword, hash)
	if !ok {
		return nil, fmt.Errorf("Username or password not match")
	}
	return generateJwtToken(user)
}

// hashgenerator
func hashGenerator(userPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(userPassword), 10)
	return string(bytes), err
}

// compareHash
func compareHash(userPassword, hash string) bool {
	// Hash compare
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(userPassword))
	return err == nil
}

func generateJwtToken(username string) ([]byte, error) {
	claims := &ConnectClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    AppName,
			ExpiresAt: time.Now().Add(JWTLoginDuration).Unix(),
		},
		Name:     "Test API - Diah",
		Email:    "test.email@gmail.com",
		UserName: username,
	}

	token := jwt.NewWithClaims(JWTSignMethod, claims)

	signToken, err := token.SignedString(JWTSignatureKey)
	if err != nil {
		return nil, err
	}

	return json.Marshal(LoginResponse{Token: signToken})
}
