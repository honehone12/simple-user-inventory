package models

import "gorm.io/gorm"

type Balance struct {
	gorm.Model
	UserID uint `gorm:"unique"`

	Coin uint64 `gorm:"not null"`
}
