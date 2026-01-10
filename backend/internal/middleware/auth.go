package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ContextKey string

const (
	UserIDKey   ContextKey = "user_id"
	UserRoleKey ContextKey = "role"
)

type JWTClaims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func JWTAuthMiddleware(jwtSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"success": false,
					"error":   "missing authorization header",
					"status":  http.StatusUnauthorized,
				})
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"success": false,
					"error":   "invalid authorization header format",
					"status":  http.StatusUnauthorized,
				})
			}

			tokenString := parts[1]

			token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "invalid signing method")
				}
				return []byte(jwtSecret), nil
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"success": false,
					"error":   "invalid or expired token",
					"status":  http.StatusUnauthorized,
				})
			}

			claims, ok := token.Claims.(*JWTClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"success": false,
					"error":   "invalid token claims",
					"status":  http.StatusUnauthorized,
				})
			}

			// Set user info in context
			c.Set(string(UserIDKey), claims.UserID)
			c.Set(string(UserRoleKey), claims.Role)

			return next(c)
		}
	}
}

// Helper functions to extract values from context
func GetUserID(c echo.Context) (int64, bool) {
	id, ok := c.Get(string(UserIDKey)).(int64)
	return id, ok
}

func GetUserRole(c echo.Context) (string, bool) {
	role, ok := c.Get(string(UserRoleKey)).(string)
	return role, ok
}
