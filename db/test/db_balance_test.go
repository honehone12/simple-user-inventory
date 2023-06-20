package test

import (
	"simple-user-inventory/db"
	"simple-user-inventory/db/test/common"
	"testing"
)

func TestFund(t *testing.T) {
	common.SetupUser()
	conn, err := db.NewOrm()
	if err != nil {
		t.Fatal(err)
	}
	// Ginji's id is 1
	if err = conn.Balance().Fund(1, 1000); err != nil {
		t.Fatal(err)
	}
}

func TestCoin(t *testing.T) {
	conn, err := db.NewOrm()
	if err != nil {
		t.Fatal(err)
	}
	// Ginji's id is 1
	coin, err := conn.Balance().Coin(1)
	if err != nil {
		t.Fatal(err)
	}
	if coin != 1000 {
		t.Fatalf("balance is not 1000 but, %d", coin)
	}
}

func TestConsumeCoin(t *testing.T) {
	conn, err := db.NewOrm()
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
		t.Fatalf("balance is not 0, but %d", coin)
	}
}
