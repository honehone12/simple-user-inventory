package controller

import (
	"fmt"
	"math"
	"simple-user-inventory/db/models"

	"gorm.io/gorm"
)

type BalanceController struct {
	db *gorm.DB
}

func NewBalanceController(db *gorm.DB) BalanceController {
	return BalanceController{db}
}

func (c BalanceController) Coin(id uint) (uint64, error) {
	balance := &models.Balance{}
	result := c.db.Select("Coin").Where("user_id = ?", id).Take(balance)
	if result.Error != nil {
		return 0, result.Error
	}
	return balance.Coin, nil
}

func (c BalanceController) Fund(id uint, value uint64) (uint64, error) {
	balance := &models.Balance{}
	err := c.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Select("ID", "Coin").Where("user_id = ?", id).Take(balance)
		if result.Error != nil {
			return result.Error
		}
		if value > (math.MaxUint64 - balance.Coin) {
			return fmt.Errorf("balance overflow %d + %d", balance.Coin, value)
		}

		result = tx.Model(balance).Update("Coin", balance.Coin+value)
		return result.Error
	})
	if err != nil {
		return 0, err
	}
	return balance.Coin, nil
}

func (c BalanceController) Consume(id uint, value uint64) (uint64, error) {
	balance := &models.Balance{}
	err := c.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Select("ID", "Coin").Where("user_id = ?", id).Take(balance)
		if result.Error != nil {
			return result.Error
		}
		if value > balance.Coin {
			return fmt.Errorf("balance underflow %d - %d", balance.Coin, value)
		}

		result = tx.Model(balance).Update("Coin", balance.Coin-value)
		return result.Error
	})
	if err != nil {
		return 0, err
	}
	return balance.Coin, nil
}
