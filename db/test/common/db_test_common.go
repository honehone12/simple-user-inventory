package common

import (
	"simple-user-inventory/db"

	"github.com/joho/godotenv"
)

func SetupEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}
}

func SetupUser() {
	SetupEnv()

	conn, err := db.NewOrm()
	if err != nil {
		panic(err)
	}
	if err = conn.User().Create(
		"Ginji",
		"ginji@user.moe",
		"ginjikyunmoemoe",
	); err != nil {
		panic(err)
	}
}
