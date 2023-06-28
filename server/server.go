package server

import (
	"simple-user-inventory/db"
	"simple-user-inventory/server/context"
	"simple-user-inventory/server/handlers"
	"simple-user-inventory/server/handlers/balance"
	"simple-user-inventory/server/handlers/items"
	"simple-user-inventory/server/handlers/jewel"
	"simple-user-inventory/server/handlers/user"
	"simple-user-inventory/server/utils"

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
	if !utils.IsDev() {
		e.Use(middleware.Recover())
	}
	e.Use(middleware.Logger())
	e.Use(echoS.Middleware(store))

	e.GET("/", handlers.Root)
	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)

	e.GET("/balance/coin", balance.Coin)
	e.POST("/balance/fund", balance.Fund)

	e.GET("/jewel/jewels", jewel.Jewels)
	e.POST("/jewel/gain", jewel.Gain)

	e.GET("/items/list", user.Items)
	e.POST("/items/purchase", user.Purchase)

	// want another service(admin or dev)
	e.POST("/items/create", items.Create)

	e.Logger.Fatal(e.Start(listenAt))
}
