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

type Connection struct {
	db *gorm.DB
}

func (conn Connection) User() controller.UserController {
	return controller.NewUserController(conn.db)
}

func (conn Connection) Balance() controller.BalanceController {
	return controller.NewBalanceController(conn.db)
}

func (conn Connection) Jewel() controller.JewelController {
	return controller.NewJewelController(conn.db)
}

func NewConnection() (Connection, error) {
	conn := Connection{}
	var err error

	dsn := os.Getenv("POSTGRES_DSN")
	if len(dsn) == 0 {
		return conn, errors.New("env param POSTGRES_DSN is empty")
	}

	conn.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return conn, err
	}

	if err = conn.migrate(); err != nil {
		return conn, err
	}

	log.Println("new database connection is done")
	return conn, nil
}

func (conn Connection) migrate() error {
	if err := conn.db.AutoMigrate(
		&models.User{},
		&models.Balance{},
		&models.Jewel{},
		&models.Item{},
	); err != nil {
		return err
	}
	return nil
}
