package test

import (
	"simple-user-inventory/db"
	"simple-user-inventory/db/test/common"
	"testing"
)

func TestCreate(t *testing.T) {
	common.SetupEnv()
	orm, err := db.NewOrm()
	if err != nil {
		t.Fatal(err)
	}
	if err = orm.User().Create(
		"Ginji",
		"ginji@user.moe",
		"ginjikyunmoemoe",
	); err != nil {
		t.Fatal(err)
	}
}

func TestRead(t *testing.T) {
	orm, err := db.NewOrm()
	if err != nil {
		t.Fatal(err)
	}
	user, err := orm.User().Read("ginji@user.moe")
	if err != nil {
		t.Fatal(err)
	}
	if user.Name != "Ginji" {
		t.Fatalf("the name of the user is not Ginji, instead %s", user.Name)
	}
}

func TestReadByUuid(t *testing.T) {
	orm, err := db.NewOrm()
	if err != nil {
		t.Fatal(err)
	}
	user, err := orm.User().Read("ginji@user.moe")
	if err != nil {
		t.Fatal(err)
	}

	user, err = orm.User().ReadByUuid(user.Uuid)
	if err != nil {
		t.Fatal(err)
	}
	if user.Name != "Ginji" {
		t.Fatalf("the name of the user is not Ginji, instead %s", user.Name)

	}
}

func TestReadId(t *testing.T) {
	orm, err := db.NewOrm()
	if err != nil {
		t.Fatal(err)
	}
	id, err := orm.User().ReadId("ginji@user.moe")
	if err != nil {
		t.Fatal(err)
	}
	if id != 1 {
		t.Fatalf("the user id is not 1 instead %d", id)
	}
}

func TestUuidToId(t *testing.T) {
	orm, err := db.NewOrm()
	if err != nil {
		t.Fatal(err)
	}
	user, err := orm.User().Read("ginji@user.moe")
	if err != nil {
		t.Fatal(err)
	}

	id, err := orm.User().UuidToId(user.Uuid)
	if err != nil {
		t.Fatal(err)
	}
	// ginji's id is 1
	if id != 1 {
		t.Fatalf("the user id is not 1 instead %d", id)
	}
}

func TestVerifyPassword(t *testing.T) {
	orm, err := db.NewOrm()
	if err != nil {
		t.Fatal(err)
	}
	ok, err := orm.User().VerifyPassword("ginji@user.moe", "ginjikyunmoemoe")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("ginjikyunmoemoe was not verified")
	}
}

func TestUpdate(t *testing.T) {
	orm, err := db.NewOrm()
	if err != nil {
		t.Fatal(err)
	}
	// Ginji's id is 1
	if err = orm.User().Update(1, "Ginjiro", "ginjiro@user.moe"); err != nil {
		t.Fatal(err)
	}
	user, err := orm.User().Read("ginjiro@user.moe")
	if err != nil {
		t.Fatal(err)
	}
	//log.Printf("%#v\n", user)
	if user.Name != "Ginjiro" {
		t.Fatal("the name of the user is not Ginjiro")
	}
}

func TestUpdatePassword(t *testing.T) {
	orm, err := db.NewOrm()
	if err != nil {
		t.Fatal(err)
	}
	// Ginji's id is 1
	if err = orm.User().UpdatePassword(1, "ginjikyunkyunmoe"); err != nil {
		t.Fatal(err)
	}
}
