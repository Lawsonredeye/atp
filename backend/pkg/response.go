package pkg

import "github.com/labstack/echo/v4"

func ErrorResponse(c echo.Context, err error, status int) error {
	return c.JSON(status, map[string]interface{}{
		"success": false,
		"error":   err.Error(),
		"status":  status,
	})
}

func SuccessResponse(c echo.Context, data interface{}, status int) error {
	return c.JSON(status, map[string]interface{}{
		"success": true,
		"data":    data,
		"status":  status,
	})
}
