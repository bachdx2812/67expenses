package gqlTypes

import "server/app/models"

type ExpenseTypeType struct {
	ExpenseType models.ExpenseType
}

func (t *ExpenseTypeType) ID() int32 {
	return t.ExpenseType.ID
}

func (t *ExpenseTypeType) Name() string {
	return t.ExpenseType.Name
}
