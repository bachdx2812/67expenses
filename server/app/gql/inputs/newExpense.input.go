package gqlInputs

type NewExpenseInput struct {
	Input NewExpenseInputForm
}

type NewExpenseInputForm struct {
	Content *string
	Date    *string
}
