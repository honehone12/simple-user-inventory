package main

import (
	"simple-user-inventory/db"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}

	orm, err := db.NewOrm()
	if err != nil {
		panic(err)
	}

	if err = orm.User().Seed(); err != nil {
		panic(err)
	}
}
