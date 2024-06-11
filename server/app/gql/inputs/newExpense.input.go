package gqlInputs

type NewExpenseInput struct {
	Input NewExpenseInputForm
}

type NewExpenseInputForm struct {
	ExpenseTypeId *int32
	Content       *string
	Date          *string
	Amount        *int32
}
