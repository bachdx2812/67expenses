package models

type Expense struct {
	ID            int32
	UserId        int32
	User          User
	ExpenseTypeId int32
	ExpenseType   ExpenseType
	Date          string
	Amount        int32
	Content       string
}
