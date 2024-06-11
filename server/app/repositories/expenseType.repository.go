package repositories

import (
	"context"
	"server/app/models"

	"gorm.io/gorm"
)

type ExpenseTypeRepository struct {
	Repository
}

func NewExpenseTypeRepository(c *context.Context, db *gorm.DB) *ExpenseTypeRepository {
	return &ExpenseTypeRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (repo *ExpenseTypeRepository) List(expenseTypes *[]*models.ExpenseType) error {
	return repo.db.Table("expense_types").Find(expenseTypes).Error
}
