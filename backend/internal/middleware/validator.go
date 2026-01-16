package middleware

import (
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Email regex pattern - RFC 5322 compliant with common restrictions
// This pattern validates:
// - Local part: alphanumeric, dots, hyphens, underscores, plus signs
// - Domain: alphanumeric with hyphens, must have at least one dot
// - TLD: 2-10 alphabetic characters
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,10}$`)

// Password requirements
const (
	MinPasswordLength = 8
	MaxPasswordLength = 72 // bcrypt max length
)

// CustomValidator wraps the validator package for Echo
type CustomValidator struct {
	Validator *validator.Validate
}

// NewValidator creates a new custom validator instance
func NewValidator() *CustomValidator {
	v := validator.New()

	// Register function to get json tag name for error messages
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Register custom email validation with regex
	v.RegisterValidation("email", validateEmail)

	// Register custom password validation
	v.RegisterValidation("password", validatePassword)

	return &CustomValidator{Validator: v}
}

// validateEmail validates email format using regex
func validateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	if email == "" {
		return false
	}
	return emailRegex.MatchString(email)
}

// validatePassword validates password meets requirements
func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return len(password) >= MinPasswordLength && len(password) <= MaxPasswordLength
}

// Validate validates a struct or slice based on validation tags
func (cv *CustomValidator) Validate(i interface{}) error {
	val := reflect.ValueOf(i)

	// Handle pointer
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// If it's a slice, validate each element
	if val.Kind() == reflect.Slice {
		for j := 0; j < val.Len(); j++ {
			if err := cv.Validator.Struct(val.Index(j).Interface()); err != nil {
				return err
			}
		}
		return nil
	}

	// Otherwise validate as struct
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}

// ValidationError represents a single validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrorResponse is the response for validation errors
type ValidationErrorResponse struct {
	Success bool              `json:"success"`
	Status  int               `json:"status"`
	Error   string            `json:"error"`
	Details []ValidationError `json:"details"`
}

// FormatValidationErrors formats validator errors into a user-friendly response
func FormatValidationErrors(err error) ValidationErrorResponse {
	var errors []ValidationError

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors = append(errors, ValidationError{
				Field:   e.Field(),
				Message: getErrorMessage(e),
			})
		}
	}

	return ValidationErrorResponse{
		Success: false,
		Status:  http.StatusBadRequest,
		Error:   "validation failed",
		Details: errors,
	}
}

// getErrorMessage returns a human-readable error message for each validation tag
func getErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + " is required"
	case "email":
		return e.Field() + " must be a valid email address (e.g., user@example.com)"
	case "password":
		return e.Field() + " must be between 8 and 72 characters"
	case "min":
		return e.Field() + " must be at least " + e.Param() + " characters"
	case "max":
		return e.Field() + " must be at most " + e.Param() + " characters"
	case "gte":
		return e.Field() + " must be greater than or equal to " + e.Param()
	case "gt":
		return e.Field() + " must be greater than " + e.Param()
	case "lte":
		return e.Field() + " must be less than or equal to " + e.Param()
	case "lt":
		return e.Field() + " must be less than " + e.Param()
	case "len":
		return e.Field() + " must be exactly " + e.Param() + " characters"
	case "oneof":
		return e.Field() + " must be one of: " + e.Param()
	case "alphanum":
		return e.Field() + " must contain only alphanumeric characters"
	case "alpha":
		return e.Field() + " must contain only alphabetic characters"
	case "numeric":
		return e.Field() + " must be a valid number"
	case "url":
		return e.Field() + " must be a valid URL"
	case "dive":
		return e.Field() + " contains invalid items"
	default:
		return e.Field() + " is invalid"
	}
}

// BindAndValidate binds the request body and validates it
func BindAndValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"status":  http.StatusBadRequest,
			"error":   "invalid request body: " + err.Error(),
		})
	}

	if err := c.Validate(i); err != nil {
		return c.JSON(http.StatusBadRequest, FormatValidationErrors(err))
	}

	return nil
}
