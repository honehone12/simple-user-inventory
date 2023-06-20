package server

import (
	"net/http"
	"simple-user-inventory/db"
	"simple-user-inventory/server/context"
	"simple-user-inventory/server/handlers"
	"simple-user-inventory/server/metadata"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(metadata metadata.Metadata, listenAt string, db db.Connection) {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &context.Context{
				Context:    c,
				Connection: db,
			}
			return next(ctx)
		}
	})
	e.Validator = context.NewValidator()
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, metadata)
	})
	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)

	e.Logger.Fatal(e.Start(listenAt))
}
