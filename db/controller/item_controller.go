package controller

import (
	"simple-user-inventory/db/models"

	"gorm.io/gorm"
)

type ItemController struct {
	db *gorm.DB
}

func NewItemController(db *gorm.DB) ItemController {
	return ItemController{db}
}

func (c ItemController) Create(
	name string,
	description string,
	price uint64,
) error {
	item := &models.Item{
		Name:        name,
		Description: description,
		Price:       price,
	}
	result := c.db.Create(item)
	return result.Error
}
