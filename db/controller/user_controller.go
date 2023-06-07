package controller

import (
	"simple-user-inventory/db/models"

	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db}
}

func (c *UserController) Seed() error {
	alice, err := models.NewUser("Alice", "alice@user.moe", "alicekyunmoemoe")
	if err != nil {
		return err
	}
	bob, err := models.NewUser("Bob", "bob@user.moe", "bobkyunmoemoe")
	if err != nil {
		return err
	}
	charlie, err := models.NewUser("Charlie", "charlie@user.moe", "charliekyunmoemoe")
	if err != nil {
		return err
	}
	dave, err := models.NewUser("Dave", "dave@user.moe", "davekyunmoemoe")
	if err != nil {
		return err
	}
	eve, err := models.NewUser("Eve", "eve@user.moe", "evekyunmoemoe")
	if err != nil {
		return err
	}
	fergie, err := models.NewUser("Fergie", "fergie@user.moe", "fergiekyunmoemoe")
	if err != nil {
		return err
	}

	result := c.db.Create([]*models.User{
		alice, bob, charlie, dave, eve, fergie,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *UserController) Create(
	name string,
	email string,
	password string,
) error {
	user, err := models.NewUser(name, email, password)
	if err != nil {
		return err
	}

	result := c.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *UserController) Read(email string) (*models.User, error) {
	user := &models.User{}
	result := c.db.Select(
		"ID", "CreatedAt", "UpdatedAt",
		"Name", "Email",
	).Where(&models.User{Email: email}).Take(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
