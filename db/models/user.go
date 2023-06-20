package models

import (
	"simple-user-inventory/db/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `gorm:"not null;size:256"`
	Uuid         string `gorm:"not null;size:64"`
	Email        string `gorm:"unique;not null;size:256"`
	Salt         []byte `gorm:"not null;size:64"`
	PasswordHash []byte `gorm:"not null;size:64"`

	Balance *Balance `gorm:"not null"`
	Jewel   *Jewel   `gorm:"not null"`
	Items   []Item   `gorm:"not null;many2many:user_items"`
}

func NewUser(
	name string,
	email string,
	password string,
) (*User, error) {
	hasher := utils.NewPasswordHasher(password)
	hashed, err := hasher.Hash()
	if err != nil {
		return nil, err
	}

	uuid := uuid.NewString()
	return &User{
		Name:         name,
		Uuid:         uuid,
		Email:        email,
		Salt:         hashed.Salt,
		PasswordHash: hashed.DK,

		Balance: &Balance{},
		Jewel:   &Jewel{},
	}, nil
}
