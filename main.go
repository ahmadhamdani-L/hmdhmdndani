package main

import (
	db "lion-super-app/internal/database"
	"lion-super-app/internal/factory"
	"lion-super-app/internal/http"
	"lion-super-app/internal/middleware"
	// "lion-super-app/pkg/redis"
	"github.com/labstack/echo/v4"
	"lion-super-app/internal/database/migration"
)

func main() {
	db.Init()
	migration.Init()

	e := echo.New()
	middleware.Init(e)

	f := factory.NewFactory()
	http.Init(e, f)
	// redis.Init()
	e.Logger.Fatal(e.Start(":3030"))
}
