package middleware

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/lawson/otterprep/pkg"
)

// CustomHTTPErrorHandler handles all errors and returns a consistent JSON response
func CustomHTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	// Handle validation errors
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		_ = c.JSON(http.StatusBadRequest, FormatValidationErrors(validationErrors))
		return
	}

	code := http.StatusInternalServerError
	message := "internal server error"

	// Handle Echo HTTP errors
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if m, ok := he.Message.(string); ok {
			message = m
		} else {
			message = http.StatusText(code)
		}
	} else {
		// Map known application errors to appropriate status codes
		switch err {
		case pkg.ErrSubjectNotFound, pkg.ErrQuestionNotFound,
			pkg.ErrQuestionOptionNotFound, pkg.ErrQuizNotFound, pkg.ErrUserNotFound:
			code = http.StatusNotFound
			message = err.Error()
		case pkg.ErrInvalidName, pkg.ErrInvalidEmail, pkg.ErrInvalidUserID,
			pkg.ErrQuestionTextNotFound, pkg.ErrQuestionOptionTextNotFound,
			pkg.ErrSubjectNameNotFound, pkg.ErrInvalidPasswordLength:
			code = http.StatusBadRequest
			message = err.Error()
		case pkg.ErrInvalidPasswordHash, pkg.ErrUnauthorized, pkg.ErrInvalidRole:
			code = http.StatusUnauthorized
			message = err.Error()
		case pkg.ErrSubjectWithNameExists, pkg.ErrUserAlreadyExists:
			code = http.StatusConflict
			message = err.Error()
		case pkg.ErrInternalServerError:
			code = http.StatusInternalServerError
			message = err.Error()
		default:
			message = err.Error()
		}
	}

	// Return consistent error response
	_ = c.JSON(code, map[string]interface{}{
		"success": false,
		"error":   message,
		"status":  code,
	})
}

// RecoverMiddleware recovers from panics and returns a 500 error
func RecoverMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					c.Logger().Errorf("panic recovered: %v", r)
					_ = c.JSON(http.StatusInternalServerError, map[string]interface{}{
						"success": false,
						"error":   "internal server error",
						"status":  http.StatusInternalServerError,
					})
				}
			}()
			return next(c)
		}
	}
}

// NotFoundHandler handles 404 errors for undefined routes
func NotFoundHandler(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"success": false,
		"error":   "route not found",
		"status":  http.StatusNotFound,
	})
}

// MethodNotAllowedHandler handles 405 errors for wrong HTTP methods
func MethodNotAllowedHandler(c echo.Context) error {
	return c.JSON(http.StatusMethodNotAllowed, map[string]interface{}{
		"success": false,
		"error":   "method not allowed",
		"status":  http.StatusMethodNotAllowed,
	})
}
