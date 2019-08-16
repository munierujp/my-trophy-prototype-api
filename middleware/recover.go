package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Recover returns middleware which wraps "github.com/labstack/echo/v4/middleware".Recover()
// see https://echo.labstack.com/middleware/recover
func Recover() echo.MiddlewareFunc {
	return middleware.Recover()
}
