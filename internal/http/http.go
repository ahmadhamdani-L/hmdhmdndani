package http

import (
	"fmt"
	"lion-super-app/configs"
	"lion-super-app/internal/app/auth"
	"lion-super-app/internal/factory"
	"net/http"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, f *factory.Factory) {
	// index
	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", configs.App().Name(), configs.App().Version())
		return c.String(http.StatusOK, message)
	})

	// e.GET("/swagger/*", echoSwagger.WrapHandler)

	// routes
	auth.NewHandler(f).Route(e.Group("/auth"))
}
