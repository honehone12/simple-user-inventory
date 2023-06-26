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

type Orm struct {
	db *gorm.DB
}

func (orm Orm) User() controller.UserController {
	return controller.NewUserController(orm.db)
}

func (orm Orm) Balance() controller.BalanceController {
	return controller.NewBalanceController(orm.db)
}

func (orm Orm) Jewel() controller.JewelController {
	return controller.NewJewelController(orm.db)
}

func (orm Orm) Item() controller.ItemController {
	return controller.NewItemController(orm.db)
}

func NewOrm() (Orm, error) {
	orm := Orm{}
	var err error

	dsn := os.Getenv("POSTGRES_DSN")
	if len(dsn) == 0 {
		return orm, errors.New("env param POSTGRES_DSN is empty")
	}

	orm.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return orm, err
	}

	if err = orm.migrate(); err != nil {
		return orm, err
	}

	log.Println("new database connection is done")
	return orm, nil
}

func (conn Orm) migrate() error {
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
