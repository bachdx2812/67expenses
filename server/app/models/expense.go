package models

type Expense struct {
	ID      int32
	UserId  int32
	User    User
	Date    string
	Content string
}
