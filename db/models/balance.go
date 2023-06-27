package models

import "gorm.io/gorm"

type BalanceData struct {
	Coin uint64 `json:"coin" gorm:"not null"`
}

type Balance struct {
	gorm.Model
	UserID uint `gorm:"unique"`

	*BalanceData `gorm:"not null"`
}
