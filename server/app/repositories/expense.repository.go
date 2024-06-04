package repositories

import (
	"context"
	"server/app/models"

	"gorm.io/gorm"
)

type ExpenseRepository struct {
	Repository
}

func NewExpenseRepository(c *context.Context, db *gorm.DB) *ExpenseRepository {
	return &ExpenseRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (repo *ExpenseRepository) Create(expense *models.Expense) error {
	return repo.db.Create(expense).Error
}
