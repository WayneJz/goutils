package middlewares

import (
	"github.com/labstack/echo/v4"
	"strings"
)

type Middlewares struct {
	ReMethodMap map[string]string
}

func (m *Middlewares) ReMethod(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		urlParams := strings.Split(c.Request().URL.Path, "/")

		return next(c)
	}
}
