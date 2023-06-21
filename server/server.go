package server

import (
	"simple-user-inventory/db"
	"simple-user-inventory/server/context"
	"simple-user-inventory/server/handlers"

	gorillaS "github.com/gorilla/sessions"
	echoS "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(
	name string,
	version string,
	listenAt string,
	db db.Orm,
	store gorillaS.Store,
) {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &context.Context{
				Context: c,
				Orm:     db,
				Metadata: context.Metadata{
					Name:    name,
					Version: version,
				},
			}
			return next(ctx)
		}
	})
	e.Validator = context.NewValidator()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(echoS.Middleware(store))

	e.GET("/", handlers.Root)
	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)

	e.Logger.Fatal(e.Start(listenAt))
}
