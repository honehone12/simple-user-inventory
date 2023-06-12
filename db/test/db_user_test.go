package test

import (
	"simple-user-inventory/db"
	"testing"

	"github.com/joho/godotenv"
)

func setupEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}
}

func TestCreate(t *testing.T) {
	setupEnv()
	conn, err := db.NewConnection()
	if err != nil {
		t.Fatal(err)
	}
	if err = conn.User().Create(
		"Ginji",
		"ginji@user.moe",
		"ginjikyunmoemoe",
	); err != nil {
		t.Fatal(err)
	}
}

func TestRead(t *testing.T) {
	conn, err := db.NewConnection()
	if err != nil {
		t.Fatal(err)
	}
	user, err := conn.User().Read("ginji@user.moe")
	if err != nil {
		t.Fatal(err)
	}
	//log.Printf("%#v\n", user)
	if user.Name != "Ginji" {
		t.Fatal("the name of the user is not Ginji")
	}
}

func TestReadId(t *testing.T) {
	conn, err := db.NewConnection()
	if err != nil {
		t.Fatal(err)
	}
	id, err := conn.User().ReadId("ginji@user.moe")
	if err != nil {
		t.Fatal(err)
	}
	if id != 1 {
		t.Fatalf("the user id is not 1 instead %d", id)
	}
}

func TestUpdate(t *testing.T) {
	conn, err := db.NewConnection()
	if err != nil {
		t.Fatal(err)
	}
	// Ginji's id is 1
	if err = conn.User().Update(1, "Ginjiro", "ginjiro@user.moe"); err != nil {
		t.Fatal(err)
	}
	user, err := conn.User().Read("ginjiro@user.moe")
	if err != nil {
		t.Fatal(err)
	}
	//log.Printf("%#v\n", user)
	if user.Name != "Ginjiro" {
		t.Fatal("the name of the user is not Ginjiro")
	}
}

func TestUpdatePassword(t *testing.T) {
	conn, err := db.NewConnection()
	if err != nil {
		t.Fatal(err)
	}
	// Ginji's id is 1
	if err = conn.User().UpdatePassword(1, "ginjikyunkyunmoe"); err != nil {
		t.Fatal(err)
	}
}
