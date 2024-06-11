package gqlTypes

import "server/app/models"

type ExpenseType struct {
	Expense *models.Expense
}

func (e *ExpenseType) ID() int32 {
	return e.Expense.ID
}

func (e *ExpenseType) UserId() int32 {
	return e.Expense.UserId
}

func (e *ExpenseType) Date() string {
	return e.Expense.Date
}

func (e *ExpenseType) Content() string {
	return e.Expense.Content
}
