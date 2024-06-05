package repositories

import (
	"context"
	"server/app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	Repository
}

func NewUserRepository(c *context.Context, db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (repo *UserRepository) Find(user *models.User) error {
	return repo.db.Table("users").First(&user).Error
}

func (repo *UserRepository) Auth(phone string, password string) (u *models.User, err error) {
	user := models.User{Phone: phone}

	if err := repo.db.Table("users").Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}

	if err := user.ComparePassword(password, user.EncryptedPassword); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) AddMemberToFamily(user models.User, family models.Family) error {
	return repo.db.Table("users").Create(&models.User{
		FamilyId:          family.ID,
		Name:              user.Name,
		Phone:             user.Phone,
		EncryptedPassword: user.EncryptedPassword,
	}).Error
}
