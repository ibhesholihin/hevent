package jwt

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService(secretKey string) JWTService {
	return &jwtService{
		secretKey: secretKey,
		issuer:    "heraldevent",
	}
}

func (s *jwtService) GenerateToken(ctx context.Context, userID int64, isAdmin bool) (accesstoken string, refreshtoken string, err error) {
	claims := &jwtCustomClaims{
		userID,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 90).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", "", err
	}

	refreshclaims := &jwtCustomClaims{
		userID,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	ts := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims)

	reftoken, err := ts.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", "", err
	}

	return token, reftoken, nil
}

func (s *jwtService) ValidateToken(ctx context.Context, tokenString string) (token *jwt.Token, err error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
}
