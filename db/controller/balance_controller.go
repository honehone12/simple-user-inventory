package controller

import (
	"fmt"
	"math"
	"simple-user-inventory/db/models"

	"gorm.io/gorm"
)

var (
	ErrorBalanceTooMuch   = fmt.Errorf("balance overflow")
	ErrorBalanceNotEnough = fmt.Errorf("balance underflow")
)

type BalanceController struct {
	db *gorm.DB
}

func NewBalanceController(db *gorm.DB) BalanceController {
	return BalanceController{db}
}

func (c BalanceController) Coin(id uint) (*models.BalanceData, error) {
	balance := &models.Balance{}
	result := c.db.Select("Coin").Where("user_id = ?", id).Take(balance)
	if result.Error != nil {
		return nil, result.Error
	}
	return &balance.BalanceData, nil
}

func (c BalanceController) Fund(id uint, value uint64) (*models.BalanceData, error) {
	balance := &models.Balance{}
	err := c.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Select("ID", "Coin").Where("user_id = ?", id).Take(balance)
		if result.Error != nil {
			return result.Error
		}
		if value > (math.MaxUint64 - balance.Coin) {
			return ErrorBalanceTooMuch
		}

		result = tx.Model(balance).Update("Coin", balance.Coin+value)
		return result.Error
	})
	if err != nil {
		return nil, err
	}
	return &balance.BalanceData, nil
}

func (c BalanceController) Consume(id uint, value uint64) (*models.BalanceData, error) {
	balance := &models.Balance{}
	err := c.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Select("ID", "Coin").Where("user_id = ?", id).Take(balance)
		if result.Error != nil {
			return result.Error
		}
		if value > balance.Coin {
			return ErrorBalanceNotEnough
		}

		result = tx.Model(balance).Update("Coin", balance.Coin-value)
		return result.Error
	})
	if err != nil {
		return nil, err
	}
	return &balance.BalanceData, nil
}
