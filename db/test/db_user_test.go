package test

import (
	"log"
	"simple-user-inventory/db"
	"testing"

	"github.com/joho/godotenv"
)

func setupTest() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}
}

func TestSeed(t *testing.T) {
	setupTest()
	conn, err := db.NewConnection()
	if err != nil {
		t.Fatal(err)
	}
	err = conn.User().Seed()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreate(t *testing.T) {
	setupTest()
	conn, err := db.NewConnection()
	if err != nil {
		t.Fatal(err)
	}
	err = conn.User().Create(
		"Ginji",
		"ginji@user.moe",
		"ginjikyunmoemoe",
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRead(t *testing.T) {
	setupTest()
	conn, err := db.NewConnection()
	if err != nil {
		t.Fatal(err)
	}
	user, err := conn.User().Read("ginji@user.moe")
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v\n", user)
	if user.Name != "Ginji" {
		t.Fatal("the name of the user is not Ginji")
	}
}
