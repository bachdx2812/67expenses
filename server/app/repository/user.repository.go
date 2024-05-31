package repository

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
