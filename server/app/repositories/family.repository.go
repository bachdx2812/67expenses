package repositories

import (
	"context"
	"server/app/models"

	"gorm.io/gorm"
)

type FamilyRepository struct {
	Repository
}

func NewFamilyRepository(c *context.Context, db *gorm.DB) *FamilyRepository {
	return &FamilyRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (repo *FamilyRepository) Find(family *models.Family) error {
	return repo.db.Table("families").First(&family).Error
}
