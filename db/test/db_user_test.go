package test

import (
	"simple-user-inventory/db"
	"simple-user-inventory/db/test/common"
	"testing"
)

func TestCreate(t *testing.T) {
	common.SetupEnv()
	conn, err := db.NewOrm()
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
	conn, err := db.NewOrm()
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
	conn, err := db.NewOrm()
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
	conn, err := db.NewOrm()
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
	conn, err := db.NewOrm()
	if err != nil {
		t.Fatal(err)
	}
	// Ginji's id is 1
	if err = conn.User().UpdatePassword(1, "ginjikyunkyunmoe"); err != nil {
		t.Fatal(err)
	}
}
