package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model

	Name        string `gorm:"unique;not null;size:256"`
	Description string `gorm:"not null;size:512"`

	Price uint64 `gorm:"not null"`

	// need somewhere else to get data,params,etc...

	Users []*User `gorm:"not null;many2many:user_items"`
}
