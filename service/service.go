package service

import (
	"log"
	"os"
	"simple-user-inventory/db"
	"simple-user-inventory/server"
	"simple-user-inventory/server/metadata"

	"github.com/joho/godotenv"
)

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	name := os.Getenv("SERVICE_NAME")
	if len(name) == 0 {
		log.Fatal("env param SERVICE_NAME is empty")
	}
	ver := os.Getenv("VERSION")
	if len(ver) == 0 {
		log.Fatal("env param VERSION is empty")
	}
	at := os.Getenv("SERVER_LISTEN_AT")
	if len(at) == 0 {
		log.Fatal("env param SERVER_LISTEN_AT is empty")
	}

	orm, err := db.NewOrm()
	if err != nil {
		log.Fatal(err)
	}

	server.Run(metadata.Metadata{
		Name:    name,
		Version: ver,
	}, at, orm)
}
