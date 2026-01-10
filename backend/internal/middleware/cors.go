package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CORSConfig holds the configuration for CORS middleware
type CORSConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge           int // in seconds
}

// DefaultCORSConfig returns a default CORS configuration
func DefaultCORSConfig() CORSConfig {
	return CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-Requested-With",
			"X-CSRF-Token",
		},
		ExposeHeaders: []string{
			"Link",
			"X-Total-Count",
		},
		AllowCredentials: true,
		MaxAge:           86400, // 24 hours
	}
}

// ProductionCORSConfig returns a stricter CORS configuration for production
// Pass in the allowed frontend origins
func ProductionCORSConfig(allowedOrigins []string) CORSConfig {
	return CORSConfig{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-Requested-With",
		},
		ExposeHeaders: []string{
			"Link",
			"X-Total-Count",
		},
		AllowCredentials: true,
		MaxAge:           86400,
	}
}

// CORSMiddleware returns a CORS middleware with the given configuration
func CORSMiddleware(config CORSConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			origin := req.Header.Get("Origin")

			// Check if origin is allowed
			allowedOrigin := ""
			for _, o := range config.AllowOrigins {
				if o == "*" || o == origin {
					allowedOrigin = origin
					break
				}
			}

			// If no origin matched and we don't allow all, don't set CORS headers
			if allowedOrigin == "" && len(config.AllowOrigins) > 0 && config.AllowOrigins[0] != "*" {
				return next(c)
			}

			// If we allow all origins
			if len(config.AllowOrigins) > 0 && config.AllowOrigins[0] == "*" {
				allowedOrigin = "*"
			}

			// Set CORS headers
			res.Header().Set("Access-Control-Allow-Origin", allowedOrigin)

			if config.AllowCredentials {
				res.Header().Set("Access-Control-Allow-Credentials", "true")
			}

			// Handle preflight request
			if req.Method == http.MethodOptions {
				res.Header().Set("Access-Control-Allow-Methods", joinStrings(config.AllowMethods, ", "))
				res.Header().Set("Access-Control-Allow-Headers", joinStrings(config.AllowHeaders, ", "))
				if config.MaxAge > 0 {
					res.Header().Set("Access-Control-Max-Age", intToString(config.MaxAge))
				}
				return c.NoContent(http.StatusNoContent)
			}

			// Set expose headers for actual requests
			if len(config.ExposeHeaders) > 0 {
				res.Header().Set("Access-Control-Expose-Headers", joinStrings(config.ExposeHeaders, ", "))
			}

			return next(c)
		}
	}
}

// Helper functions
func joinStrings(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result += sep + strs[i]
	}
	return result
}

func intToString(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		result = string(rune('0'+n%10)) + result
		n /= 10
	}
	return result
}
