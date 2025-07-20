package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret string
}

func NewJWT() *JWT {
	return &JWT{
		secret: "default-secret-key", // In production, this should come from config
	}
}

func (j *JWT) Generate(userID string) (string, error) {
	claimsMap := map[string]any{
		"user_id": userID,
	}
	return GenerateToken(claimsMap, j.secret)
}

func (j *JWT) Verify(tokenString string) (map[string]any, error) {
	return VerifyToken(tokenString, j.secret)
}

func GenerateToken(claimsMap map[string]any, secret string) (string, error) {
	claims := jwt.MapClaims{}
	for k, v := range claimsMap {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string, secret string) (map[string]any, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("invalid token claims")
	}
}
