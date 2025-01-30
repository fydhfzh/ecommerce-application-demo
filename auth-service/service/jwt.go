package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

const (
	JWT_SIGNING_METHOD = "HS256"
	JWT_SECRET_KEY     = "aeb99a7128a9c0fb285df917585d574bd5d4a9fbc769b1513fe7d2a3a2ced8bf"
)

type JwtService struct{}

func NewJwtService() JwtService {
	return JwtService{}
}

func (j *JwtService) GenerateToken(email string) (string, error) {
	claims := UserClaims{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "ecommerce-application",
			Subject:   email,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(JWT_SIGNING_METHOD), claims)

	byteSecret := []byte(JWT_SECRET_KEY)

	strToken, err := token.SignedString(byteSecret)
	if err != nil {
		return "", err
	}

	return strToken, nil
}
