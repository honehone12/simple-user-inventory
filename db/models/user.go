package models

import (
	"simple-user-inventory/db/utils"
	"simple-user-inventory/operation/role"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserData struct {
	Name  string `json:"name" gorm:"not null;size:128"`
	Uuid  string `json:"uuid" gorm:"unique;not null;size:64"`
	Email string `json:"email" gorm:"unique;not null;size:128"`
}

type User struct {
	gorm.Model

	UserData `gorm:"not null"`

	Salt         []byte `gorm:"not null;size:64"`
	PasswordHash []byte `gorm:"not null;size:64"`

	Role uint8 `gorm:"not null"`

	Balance *Balance `gorm:"not null"`
	Jewel   *Jewel   `gorm:"not null"`
	Items   []*Item  `gorm:"not null;many2many:user_items"`
}

func NewUser(
	name string,
	email string,
	password string,
) (*User, error) {
	return newUserInternl(name, email, password, role.Consumer)
}

func NewAdmin(
	name string,
	email string,
	password string,
) (*User, error) {
	return newUserInternl(name, email, password, role.Admin)
}

func newUserInternl(
	name string,
	email string,
	password string,
	role uint8,
) (*User, error) {
	hasher := utils.NewPasswordHasher(password)
	hashed, err := hasher.Hash()
	if err != nil {
		return nil, err
	}

	uuid := uuid.NewString()
	return &User{
		UserData: UserData{
			Name:  name,
			Uuid:  uuid,
			Email: email,
		},

		Salt:         hashed.Salt,
		PasswordHash: hashed.DK,

		Role: role,

		Balance: &Balance{
			BalanceData: BalanceData{Coin: 0},
		},
		Jewel: &Jewel{
			JewelData: JewelData{
				Red:    0,
				Blue:   0,
				Green:  0,
				Yellow: 0,
				Black:  0,
			},
		},
	}, nil
}
