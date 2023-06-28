package models

import (
	"fmt"
	"math"

	"gorm.io/gorm"
)

type JewelData struct {
	Red    uint64 `json:"red" gorm:"not null"`
	Blue   uint64 `json:"blue" gorm:"not null"`
	Green  uint64 `json:"green" gorm:"not null"`
	Yellow uint64 `json:"yellow" gorm:"not null"`
	Black  uint64 `json:"black" gorm:"not null"`
}

type Jewel struct {
	gorm.Model
	UserID uint `gorm:"unique"`

	JewelData `gorm:"not null"`
}

// Add JewelData to jewel. JewelData will contain result.
func (j *JewelData) AddTo(jewel *Jewel) error {
	if j.Red > math.MaxUint64-jewel.Red {
		return fmt.Errorf("red overflow %d + %d", jewel.Red, j.Red)
	}
	if j.Blue > math.MaxUint64-jewel.Blue {
		return fmt.Errorf("blue overflow %d + %d", jewel.Blue, j.Blue)
	}
	if j.Green > math.MaxUint64-jewel.Green {
		return fmt.Errorf("green overflow %d + %d", jewel.Green, j.Green)
	}
	if j.Yellow > math.MaxUint64-jewel.Yellow {
		return fmt.Errorf("yellow overflow %d + %d", jewel.Yellow, j.Yellow)
	}
	if j.Black > math.MaxUint64-jewel.Black {
		return fmt.Errorf("black overflow %d + %d", jewel.Black, j.Black)
	}

	j.Red += jewel.Red
	j.Blue += jewel.Blue
	j.Green += jewel.Green
	j.Yellow += jewel.Yellow
	j.Black += jewel.Black
	return nil
}

// Sub JewelData from jewel. JewelData will contain result.
func (j *JewelData) SubFrom(jewel *Jewel) error {
	if j.Red > jewel.Red {
		return fmt.Errorf("red underflow %d - %d", jewel.Red, j.Red)
	}
	if j.Blue > jewel.Blue {
		return fmt.Errorf("blue underflow %d - %d", jewel.Blue, j.Blue)
	}
	if j.Green > jewel.Green {
		return fmt.Errorf("green underflow %d - %d", jewel.Green, j.Green)
	}
	if j.Yellow > jewel.Yellow {
		return fmt.Errorf("yellow underflow %d - %d", jewel.Yellow, j.Yellow)
	}
	if j.Black > jewel.Black {
		return fmt.Errorf("black underflow %d - %d", jewel.Black, j.Black)
	}

	j.Red = jewel.Red - j.Red
	j.Blue = jewel.Blue - j.Blue
	j.Green = jewel.Green - j.Green
	j.Yellow = jewel.Yellow - j.Yellow
	j.Black = jewel.Black - j.Black
	return nil
}
