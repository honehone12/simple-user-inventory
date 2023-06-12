package test

import (
	"simple-user-inventory/db"
	"testing"

	"github.com/joho/godotenv"
)

func setupUser() {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}
	conn, err := db.NewConnection()
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

func TestFund(t *testing.T) {
	setupUser()
	conn, err := db.NewConnection()
	if err != nil {
		t.Fatal(err)
	}
	// Ginji's id is 1
	if err = conn.Balance().Fund(1, 1000); err != nil {
		t.Fatal(err)
	}
}

func TestCoin(t *testing.T) {
	conn, err := db.NewConnection()
	if err != nil {
		t.Fatal(err)
	}
	// Ginji's id is 1
	coin, err := conn.Balance().Coin(1)
	if err != nil {
		t.Fatal(err)
	}
	if coin != 1000 {
		t.Fatal("balance is not 1000")
	}
}

func TestConsume(t *testing.T) {
	conn, err := db.NewConnection()
	if err != nil {
		t.Fatal(err)
	}
	// Ginji's id is 1
	err = conn.Balance().Consume(1, 1000)
	if err != nil {
		t.Fatal(err)
	}
	coin, err := conn.Balance().Coin(1)
	if err != nil {
		t.Fatal(err)
	}
	if coin != 0 {
		t.Fatal("balance is not 0")
	}
}
