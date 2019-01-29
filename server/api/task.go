package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// Greeting - ok handler
func Greeting(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}
