package services

import (
	"context"
	"server/app/forms"
	gqlInputs "server/app/gql/inputs"
	"server/app/models"

	"gorm.io/gorm"
)

type CreateExpenseService struct {
	Ctx     *context.Context
	Db      *gorm.DB
	Args    gqlInputs.NewExpenseInput
	Expense *models.Expense
}

func (service *CreateExpenseService) Execute() error {
	form := forms.NewCreateExpenseForm(&service.Args.Input, service.Expense)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
