package service

import (
	"log"
	"os"
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
	at := os.Getenv("SERVER_LISTEN_AT")
	if len(at) == 0 {
		log.Fatal("env param SERVER_LISTEN_AT is empty")
	}

	conn, err := db.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(conn)
	s.Run(at)
}
