package models

import "gorm.io/gorm"

type Jewel struct {
	gorm.Model
	UserID uint `gorm:"unique"`

	Red    uint64 `gorm:"not null"`
	Blue   uint64 `gorm:"not null"`
	Green  uint64 `gorm:"not null"`
	Yellow uint64 `gorm:"not null"`
	Black  uint64 `gorm:"not null"`
}
