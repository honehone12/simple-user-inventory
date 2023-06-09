package main

import (
	"simple-user-inventory/db"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	conn, err := db.NewConnection()
	if err != nil {
		panic(err)
	}

	err = conn.User().Seed()
	if err != nil {
		panic(err)
	}
}
