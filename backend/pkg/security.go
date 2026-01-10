package pkg

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/lawson/otterprep/domain"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ObfuscateDetail obfuscates a string by replacing all but the first and last characters with a dot
func ObfuscateDetail(msg string, dataType string) string {
	switch dataType {
	case "email":
		return msg[:1] + strings.Repeat(".", len(msg)-2) + msg[len(msg)-1:]
	case "phone":
		return msg[:3] + strings.Repeat(".", len(msg)-6) + msg[len(msg)-4:]
	case "name":
		return msg[:1] + strings.Repeat(".", len(msg)-2) + msg[len(msg)-1:]
	case "password":
		return strings.Repeat("*", len(msg))
	default:
		return msg
	}
}

// GenerateAccessToken generates a JWT token for a user.
func GenerateAccessToken(userId int64, userRole string, accessTokenExpiry time.Duration, secret string) (string, error) {
	claims := &domain.JWTClaims{
		UserID: userId,
		Role:   userRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func GenerateRefreshToken(userId int64, userRole string, refreshTokenExpiry time.Duration, secret string) (string, error) {
	claims := &domain.JWTClaims{
		UserID: userId,
		Role:   userRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseToken parses and validates a JWT token, returning the claims
func ParseToken(tokenString string, secret string) (*domain.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, ErrInvalidToken
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*domain.JWTClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
