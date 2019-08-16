package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Logger returns middleware which wraps "github.com/labstack/echo/v4/middleware".Logger()
// see https://echo.labstack.com/middleware/logger
func Logger() echo.MiddlewareFunc {
	return middleware.Logger()
}
