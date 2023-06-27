package controller

import (
	"errors"
	"simple-user-inventory/db/models"
	"simple-user-inventory/db/utils"

	"gorm.io/gorm"
)

var (
	ErrorInvalidPassword = errors.New("invalid password")
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) UserController {
	return UserController{db}
}

func (c UserController) Seed() error {
	alice, err := models.NewUser("Alice", "alice@user.moe", "alicekyunmoemoe")
	if err != nil {
		return err
	}
	bob, err := models.NewUser("Bob", "bob@user.moe", "bobkyunmoemoe")
	if err != nil {
		return err
	}
	charlie, err := models.NewUser("Charlie", "charlie@user.moe", "charliekyunmoemoe")
	if err != nil {
		return err
	}
	dave, err := models.NewUser("Dave", "dave@user.moe", "davekyunmoemoe")
	if err != nil {
		return err
	}
	eve, err := models.NewUser("Eve", "eve@user.moe", "evekyunmoemoe")
	if err != nil {
		return err
	}
	fergie, err := models.NewUser("Fergie", "fergie@user.moe", "fergiekyunmoemoe")
	if err != nil {
		return err
	}

	result := c.db.Create([]*models.User{
		alice, bob, charlie, dave, eve, fergie,
	})
	return result.Error
}

func (c UserController) Create(
	name string,
	email string,
	password string,
) error {
	user, err := models.NewUser(name, email, password)
	if err != nil {
		return err
	}

	result := c.db.Create(user)
	return result.Error
}

func (c UserController) CreateAdmin(
	name string,
	email string,
	password string,
) error {
	user, err := models.NewAdmin(name, email, password)
	if err != nil {
		return err
	}

	result := c.db.Create(user)
	return result.Error
}

func (c UserController) Read(email string) (*models.User, error) {
	user := &models.User{}
	result := c.db.Select(
		"ID", "CreatedAt", "UpdatedAt",
		"Name", "Uuid", "Email", "Role",
	).Where("email = ?", email).Take(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (c UserController) ReadByUuid(uuid string) (*models.User, error) {
	user := &models.User{}
	result := c.db.Select(
		"ID", "CreatedAt", "UpdatedAt",
		"Name", "Uuid", "Email", "Role",
	).Where("uuid = ?", uuid).Take(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (c UserController) ReadId(email string) (uint, error) {
	user := &models.User{}
	result := c.db.Select("ID").Where("email = ?", email).Take(user)
	// gorm does not have zero id
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (c UserController) UuidToId(uuid string) (uint, error) {
	user := &models.User{}
	result := c.db.Select("ID").Where("uuid = ?", uuid).Take(user)
	// gorm does not have zero id
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (c UserController) VerifyPassword(email string, password string) (string, error) {
	user := &models.User{}
	result := c.db.Select("PasswordHash", "Salt", "Uuid").Where("email = ?", email).Take(user)
	if result.Error != nil {
		return "", result.Error
	}

	hasher := utils.NewPasswordHasher(password)
	ok, err := hasher.Verify(user.PasswordHash, user.Salt)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", ErrorInvalidPassword
	}

	return user.Uuid, nil
}

func (c UserController) Update(
	id uint,
	name string,
	email string,
) error {
	result := c.db.Model(&models.User{
		Model: gorm.Model{ID: id},
	}).Select("Name", "Email").Updates(&models.User{
		Model: gorm.Model{ID: id},
		UserData: &models.UserData{
			Name:  name,
			Email: email,
		},
	})
	return result.Error
}

func (c UserController) UpdatePassword(id uint, password string) error {
	hasher := utils.NewPasswordHasher(password)
	hashed, err := hasher.Hash()
	if err != nil {
		return err
	}

	result := c.db.Model(&models.User{
		Model: gorm.Model{ID: id},
	}).Select("Salt", "PasswordHash").Updates(&models.User{
		Salt:         hashed.Salt,
		PasswordHash: hashed.DK,
	})
	return result.Error
}

func (c UserController) Purchase(id uint, itemId uint) error {
	return c.db.Model(&models.User{
		Model: gorm.Model{ID: id},
	}).Association("Items").Append(&models.Item{
		Model: gorm.Model{ID: itemId},
	})
}

func (c UserController) ReadOwnedItems(id uint) ([]models.Item, error) {
	items := []models.Item{}
	err := c.db.Model(&models.User{
		Model: gorm.Model{ID: id},
	}).Select(
		"ID", "Name", "Description", "Price",
	).Association("Items").Find(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}
