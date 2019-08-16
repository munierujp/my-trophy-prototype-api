# middleware
To avoid duplicate package names, put all [middlewares](https://echo.labstack.com/middleware) here and **don't** use Echo's middleware directly.

## Example
`middleware/logger.go`:

```go
package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logger() echo.MiddlewareFunc {
	return middleware.Logger()
}
```


`main.go`:

```go
package main

import (
	"my-app/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
}
```
