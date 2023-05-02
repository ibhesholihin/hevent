package middleware

import (
	"net/http"
	"strings"

	jwtLib "github.com/golang-jwt/jwt"
	"github.com/ibhesholihin/hevent/utils"
	"github.com/labstack/echo/v4"
)

// JWT auth for user token
func (m *Middleware) JWTAuthUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			ctx := c.Request().Context()
			authorizationHeader := c.Request().Header.Get("Authorization")
			bearerToken := strings.Split(authorizationHeader, " ")

			if len(bearerToken) != 2 {
				return c.JSON(http.StatusUnauthorized, utils.NewUnauthorizedError("invalid authorization token"))
			}

			tokenStr := bearerToken[1]
			token, err := m.jwtSvc.ValidateToken(ctx, tokenStr)
			if err != nil {
				return c.JSON(
					http.StatusUnauthorized,
					utils.NewUnauthorizedError("invalid authorization token"),
				)
			}

			if !token.Valid {
				return c.JSON(
					http.StatusUnauthorized,
					utils.NewUnauthorizedError("invalid authorization token"),
				)
			}

			claims := token.Claims.(jwtLib.MapClaims)

			//cek if isadmin true
			if claims["isadmin"].(bool) {
				return c.JSON(
					http.StatusUnauthorized,
					utils.NewUnauthorizedError("invalid authorization token"),
				)
			}

			c.Set("user_id", int64(claims["user_id"].(float64)))
			c.Set("isAdmin", bool(claims["isadmin"].(bool)))
			return next(c)
		}
	}
}

// JWT auth for admin token
func (m *Middleware) JWTAuthAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			ctx := c.Request().Context()
			authorizationHeader := c.Request().Header.Get("Authorization")
			bearerToken := strings.Split(authorizationHeader, " ")

			if len(bearerToken) != 2 {
				return c.JSON(http.StatusUnauthorized, utils.NewUnauthorizedError("invalid authorization token"))
			}

			tokenStr := bearerToken[1]
			token, err := m.jwtSvc.ValidateToken(ctx, tokenStr)
			if err != nil {
				return c.JSON(
					http.StatusUnauthorized,
					utils.NewUnauthorizedError("invalid authorization token"),
				)
			}

			if !token.Valid {
				return c.JSON(
					http.StatusUnauthorized,
					utils.NewUnauthorizedError("invalid authorization token"),
				)
			}

			claims := token.Claims.(jwtLib.MapClaims)

			//cek if isadmin false
			if !claims["isadmin"].(bool) {
				return c.JSON(
					http.StatusUnauthorized,
					utils.NewUnauthorizedError("invalid authorization token"),
				)
			}

			c.Set("user_id", int64(claims["user_id"].(float64)))
			c.Set("isAdmin", bool(claims["isadmin"].(bool)))
			return next(c)
		}
	}
}
