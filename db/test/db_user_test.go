package test

import (
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
	//log.Printf("%#v\n", user)
	if user.Name != "Ginji" {
		t.Fatal("the name of the user is not Ginji")
	}
}

func TestReadId(t *testing.T) {
	setupTest()
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
	setupTest()
	conn, err := db.NewConnection()
	if err != nil {
		t.Fatal(err)
	}
	// Ginji's id is 1
	err = conn.User().Update(1, "Ginjiro", "ginjiro@user.moe")
	if err != nil {
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
	setupTest()
	conn, err := db.NewConnection()
	if err != nil {
		t.Fatal(err)
	}
	// Ginji's id is 1
	err = conn.User().UpdatePassword(1, "ginjikyunkyunmoe")
	if err != nil {
		t.Fatal(err)
	}
}
