package server

import (
	"errors"
	"os"
	"simple-user-inventory/db"

	"github.com/labstack/echo/v4"
)

type Server struct {
	listenAt string
	db.Connection
}

func NewServer(conn db.Connection) (Server, error) {
	s := Server{}
	at := os.Getenv("SERVER_LISTEN_AT")
	if len(at) == 0 {
		return s, errors.New("env param SERVER_LISTEN_AT is empty")
	}
	s.Connection = conn
	s.listenAt = at
	return s, nil
}

func (s Server) Run() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		panic("not implemented")
	})
	e.Logger.Fatal(e.Start(s.listenAt))
}
