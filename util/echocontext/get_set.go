package echocontext

import "github.com/labstack/echo/v4"

// GetPayload ...
func GetPayload(c echo.Context) interface{} {
	return c.Get(KeyPayload)
}

// SetPayload ...
func SetPayload(c echo.Context, value interface{}) {
	c.Set(KeyPayload, value)
}

// GetQuery ...
func GetQuery(c echo.Context) interface{} {
	return c.Get(KeyQuery)
}

// SetQuery ...
func SetQuery(c echo.Context, value interface{}) {
	c.Set(KeyQuery, value)
}

// GetParam ...
func GetParam(c echo.Context, key string) interface{} {
	return c.Get(key)
}

// SetParam ...
func SetParam(c echo.Context, key string, value interface{}) {
	c.Set(key, value)
}
