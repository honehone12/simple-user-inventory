package controller

import (
	"simple-user-inventory/db/models"

	"gorm.io/gorm"
)

type JewelController struct {
	db *gorm.DB
}

func NewJewelController(db *gorm.DB) JewelController {
	return JewelController{db}
}

func (c JewelController) Jewels(id uint) (*models.JewelData, error) {
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
	return &jewel.JewelData, nil
}

func (c JewelController) Gain(id uint, j *models.JewelData) (*models.JewelData, error) {
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

		err := j.AddTo(jewel)
		if err != nil {
			return err
		}

		result = tx.Model(jewel).Select(
			"Red",
			"Blue",
			"Green",
			"Yellow",
			"Black",
		).Updates(j)
		return result.Error
	})
	if err != nil {
		return nil, err
	}
	return &jewel.JewelData, nil
}

func (c JewelController) Consume(id uint, j *models.JewelData) (*models.JewelData, error) {
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

		err := j.SubFrom(jewel)
		if err != nil {
			return err
		}

		result = tx.Model(jewel).Select(
			"Red",
			"Blue",
			"Green",
			"Yellow",
			"Black",
		).Updates(j)
		return result.Error
	})
	if err != nil {
		return nil, err
	}
	return &jewel.JewelData, nil
}
