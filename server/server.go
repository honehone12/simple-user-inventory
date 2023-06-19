package server

import (
	"simple-user-inventory/db"

	"github.com/labstack/echo/v4"
)

type Server struct {
	db.Connection
}

func NewServer(conn db.Connection) Server {
	return Server{conn}
}

func (s Server) Run(listenAt string) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		panic("not implemented")
	})
	e.Logger.Fatal(e.Start(listenAt))
}
