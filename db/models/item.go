package models

import "gorm.io/gorm"

type ItemData struct {
	Name        string `json:"name" gorm:"unique;not null;size:128"`
	Description string `json:"description" gorm:"not null;size:512"`
	Price       uint64 `json:"price" gorm:"not null"`
}

type Item struct {
	gorm.Model

	ItemData `gorm:"not null"`

	// need somewhere else to get data,params,etc...

	Users []*User `gorm:"not null;many2many:user_items"`
}
