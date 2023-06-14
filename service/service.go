package service

import (
	"log"
	"simple-user-inventory/db"
	"simple-user-inventory/server"

	"github.com/joho/godotenv"
)

type Service struct {
	server.Server
}

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	conn, err := db.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	s, err := server.NewServer(conn)
	if err != nil {
		log.Fatal(err)
	}
	s.Run()
}
