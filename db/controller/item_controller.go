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

func (c ItemController) Seed() error {
	apple := models.Item{
		Name:        "Apple",
		Description: "Red sweet ball",
		Price:       100,
	}
	banana := models.Item{
		Name:        "Banana",
		Description: "Yellow strong stick",
		Price:       100,
	}
	chocolate := models.Item{
		Name:        "Chocolate",
		Description: "Too much sugar",
		Price:       200,
	}
	dinosaur := models.Item{
		Name:        "DenoSaur",
		Description: "Delicious!!",
		Price:       500,
	}
	elvis := models.Item{
		Name:        "Elvis",
		Description: "?:",
		Price:       1000,
	}
	f := models.Item{
		Name:        "f",
		Description: "Words start from 'f'",
		Price:       3000,
	}

	result := c.db.Create([]*models.Item{
		&apple, &banana, &chocolate, &dinosaur, &elvis, &f,
	})
	return result.Error
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
