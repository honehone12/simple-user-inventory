package db

import (
	"errors"
	"log"
	"os"
	"simple-user-inventory/db/controller"
	"simple-user-inventory/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnection struct {
	user *controller.UserController
}

func (conn *DbConnection) User() *controller.UserController {
	return conn.user
}

func NewConnection() (*DbConnection, error) {
	dsn := os.Getenv("POSTGRES_DSN")
	if len(dsn) == 0 {
		return nil, errors.New("env param POSTGRES_DSN is empty")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = migrate(db); err != nil {
		return nil, err
	}

	return &DbConnection{
		user: controller.NewUserController(db),
	}, nil
}

func migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.User{},
		&models.Balance{},
		&models.Jewel{},
		&models.Item{},
	); err != nil {
		return err
	}
	log.Println("migration is done")
	return nil
}