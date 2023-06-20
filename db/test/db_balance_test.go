package test

import (
	"simple-user-inventory/db"
	"simple-user-inventory/db/test/common"
	"testing"
)

func TestFund(t *testing.T) {
	common.SetupUser()
	orm, err := db.NewOrm()
	if err != nil {
		t.Fatal(err)
	}
	// Ginji's id is 1
	if err = orm.Balance().Fund(1, 1000); err != nil {
		t.Fatal(err)
	}
}

func TestCoin(t *testing.T) {
	orm, err := db.NewOrm()
	if err != nil {
		t.Fatal(err)
	}
	// Ginji's id is 1
	coin, err := orm.Balance().Coin(1)
	if err != nil {
		t.Fatal(err)
	}
	if coin != 1000 {
		t.Fatalf("balance is not 1000 but, %d", coin)
	}
}

func TestConsumeCoin(t *testing.T) {
	orm, err := db.NewOrm()
	if err != nil {
		t.Fatal(err)
	}
	// Ginji's id is 1
	err = orm.Balance().Consume(1, 1000)
	if err != nil {
		t.Fatal(err)
	}
	coin, err := orm.Balance().Coin(1)
	if err != nil {
		t.Fatal(err)
	}
	if coin != 0 {
		t.Fatalf("balance is not 0, but %d", coin)
	}
}
