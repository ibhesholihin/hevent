package jwt

import (
	"context"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(ctx context.Context, userID int64, isAdmin bool) (token string, refreshtoken string, err error)
	ValidateToken(ctx context.Context, tokenString string) (token *jwt.Token, err error)
}
