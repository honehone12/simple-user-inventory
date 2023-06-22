package controller

import (
	"fmt"
	"math"
	"simple-user-inventory/db/models"

	"gorm.io/gorm"
)

type JewelData struct {
	Red    uint64
	Blue   uint64
	Green  uint64
	Yellow uint64
	Black  uint64
}

func FromJewel(jewel *models.Jewel) *JewelData {
	return &JewelData{
		Red:    jewel.Red,
		Blue:   jewel.Blue,
		Green:  jewel.Green,
		Yellow: jewel.Yellow,
		Black:  jewel.Black,
	}
}

func (j *JewelData) Add(jewel *models.Jewel) (*models.Jewel, error) {
	if j.Red > math.MaxUint64-jewel.Red {
		return nil, fmt.Errorf("red overflow %d + %d", jewel.Red, j.Red)
	}
	if j.Blue > math.MaxUint64-jewel.Blue {
		return nil, fmt.Errorf("blue overflow %d + %d", jewel.Blue, j.Blue)
	}
	if j.Green > math.MaxUint64-jewel.Green {
		return nil, fmt.Errorf("green overflow %d + %d", jewel.Green, j.Green)
	}
	if j.Yellow > math.MaxUint64-jewel.Yellow {
		return nil, fmt.Errorf("yellow overflow %d + %d", jewel.Yellow, j.Yellow)
	}
	if j.Black > math.MaxUint64-jewel.Black {
		return nil, fmt.Errorf("black overflow %d + %d", jewel.Black, j.Black)
	}
	return &models.Jewel{
		Red:    jewel.Red + j.Red,
		Blue:   jewel.Blue + j.Blue,
		Green:  jewel.Green + j.Green,
		Yellow: jewel.Yellow + j.Yellow,
		Black:  jewel.Black + j.Black,
	}, nil
}

func (j *JewelData) Sub(jewel *models.Jewel) (*models.Jewel, error) {
	if j.Red > jewel.Red {
		return nil, fmt.Errorf("red underflow %d - %d", jewel.Red, j.Red)
	}
	if j.Blue > jewel.Blue {
		return nil, fmt.Errorf("blue underflow %d - %d", jewel.Blue, j.Blue)
	}
	if j.Green > jewel.Green {
		return nil, fmt.Errorf("green underflow %d - %d", jewel.Green, j.Green)
	}
	if j.Yellow > jewel.Yellow {
		return nil, fmt.Errorf("yellow underflow %d - %d", jewel.Yellow, j.Yellow)
	}
	if j.Black > jewel.Black {
		return nil, fmt.Errorf("black underflow %d - %d", jewel.Black, j.Black)
	}
	return &models.Jewel{
		Red:    jewel.Red - j.Red,
		Blue:   jewel.Blue - j.Blue,
		Green:  jewel.Green - j.Green,
		Yellow: jewel.Yellow - j.Yellow,
		Black:  jewel.Black - j.Black,
	}, nil
}

type JewelController struct {
	db *gorm.DB
}

func NewJewelController(db *gorm.DB) JewelController {
	return JewelController{db}
}

func (c JewelController) Jewels(id uint) (*JewelData, error) {
	jewel := &models.Jewel{}
	result := c.db.Select(
		"Red",
		"Blue",
		"Green",
		"Yellow",
		"Black",
	).Where("user_id = ?", id).Take(jewel)
	if result.Error != nil {
		return nil, result.Error
	}
	return FromJewel(jewel), nil
}

func (c JewelController) Gain(id uint, jewelData *JewelData) (*JewelData, error) {
	jewel := &models.Jewel{}
	err := c.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Select(
			"ID",
			"Red",
			"Blue",
			"Green",
			"Yellow",
			"Black",
		).Where("user_id = ?", id).Take(jewel)
		if result.Error != nil {
			return result.Error
		}

		add, err := jewelData.Add(jewel)
		if err != nil {
			return err
		}

		result = tx.Model(jewel).Select(
			"Red",
			"Blue",
			"Green",
			"Yellow",
			"Black",
		).Updates(add)
		return result.Error
	})
	if err != nil {
		return nil, err
	}
	return FromJewel(jewel), nil
}

func (c JewelController) Consume(id uint, jewelData *JewelData) (*JewelData, error) {
	jewel := &models.Jewel{}
	err := c.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Select(
			"ID",
			"Red",
			"Blue",
			"Green",
			"Yellow",
			"Black",
		).Where("user_id = ?", id).Take(jewel)
		if result.Error != nil {
			return result.Error
		}

		sub, err := jewelData.Sub(jewel)
		if err != nil {
			return err
		}

		result = tx.Model(jewel).Select(
			"Red",
			"Blue",
			"Green",
			"Yellow",
			"Black",
		).Updates(sub)
		return result.Error
	})
	if err != nil {
		return nil, err
	}
	return FromJewel(jewel), nil
}
