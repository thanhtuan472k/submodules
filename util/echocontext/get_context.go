package echocontext

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
)

// GetContext ...
func GetContext(c echo.Context) context.Context {
	return c.Request().Context()
}
