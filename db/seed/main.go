package main

import (
	"simple-user-inventory/db"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}

	conn, err := db.NewConnection()
	if err != nil {
		panic(err)
	}

	if err = conn.User().Seed(); err != nil {
		panic(err)
	}
}
