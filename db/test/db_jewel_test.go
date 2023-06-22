package test

import (
	"simple-user-inventory/db"
	"simple-user-inventory/db/controller"
	"simple-user-inventory/db/test/common"
	"testing"
)

func TestGain(t *testing.T) {
	common.SetupUser()
	orm, err := db.NewOrm()
	if err != nil {
		t.Fatal(err)
	}
	// Ginji's id is 1
	newJewel := controller.JewelData{
		Red:    5,
		Blue:   4,
		Green:  3,
		Yellow: 2,
		Black:  1,
	}
	if j, err := orm.Jewel().Gain(1, &newJewel); err != nil || *j != newJewel {
		t.Fatal(err)
	}
}

func TestJewels(t *testing.T) {
	orm, err := db.NewOrm()
	if err != nil {
		t.Fatal(err)
	}
	// Ginji's id is 1
	jewel, err := orm.Jewel().Jewels(1)
	if err != nil {
		t.Fatal(err)
	}

	if jewel.Red != 5 {
		t.Fatalf("red is not 5, but %d", jewel.Red)
	}
	if jewel.Blue != 4 {
		t.Fatalf("blue is not 4, but %d", jewel.Blue)
	}
	if jewel.Green != 3 {
		t.Fatalf("green is not 3, but %d", jewel.Green)
	}
	if jewel.Yellow != 2 {
		t.Fatalf("yellow is not 2, but %d", jewel.Yellow)
	}
	if jewel.Black != 1 {
		t.Fatalf("black is not 1, but %d", jewel.Black)
	}
}

func TestConsumeJewel(t *testing.T) {
	orm, err := db.NewOrm()
	if err != nil {
		t.Fatal(err)
	}
	// Ginji's id is 1
	j, err := orm.Jewel().Consume(1, &controller.JewelData{
		Red:    5,
		Blue:   4,
		Green:  3,
		Yellow: 2,
		Black:  1,
	})
	if err != nil {
		t.Fatal(err)
	}

	if j.Red != 0 {
		t.Fatalf("red is not 0, but %d", j.Red)
	}
	if j.Blue != 0 {
		t.Fatalf("blue is not 0, but %d", j.Blue)
	}
	if j.Green != 0 {
		t.Fatalf("green is not 0, but %d", j.Green)
	}
	if j.Yellow != 0 {
		t.Fatalf("yellow is not 0, but %d", j.Yellow)
	}
	if j.Black != 0 {
		t.Fatalf("black is not 0, but %d", j.Black)
	}
}
