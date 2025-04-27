package jwt

import (
	"fmt"
	"time"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken генерирует jwt-токен
func GenerateToken(claims *model.UserClaims, secretKey []byte, duration time.Duration) (string, error) {
	if claims == nil {
		return "", fmt.Errorf("info is nil")
	}
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

// VerifyToken валидирует jwt-токен и возвращает его claims
func VerifyToken(tokenStr string, secretKey []byte) (*model.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&model.UserClaims{},
		func(token *jwt.Token) (any, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("invalid token: %v", "unexpected token signing method")
			}

			return secretKey, nil
		},
	)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors == jwt.ValidationErrorExpired {
			return nil, sys.AccessTokenExpiredError
		}
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(*model.UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
