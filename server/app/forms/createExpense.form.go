package forms

import (
	"server/app/exceptions"
	"server/app/formAttributes"
	gqlInputs "server/app/gql/inputs"
	"server/app/helpers"
	"server/app/models"
	"server/app/repositories"
	"server/database"
)

type CreateExpenseForm struct {
	Form
	Input   *gqlInputs.NewExpenseInputForm
	Expense models.Expense
	repo    repositories.ExpenseRepository
}

func NewCreateExpenseForm(
	input *gqlInputs.NewExpenseInputForm,
	expense *models.Expense,
) CreateExpenseForm {
	form := CreateExpenseForm{
		Form:    Form{},
		Input:   input,
		Expense: *expense,
		repo:    *repositories.NewExpenseRepository(nil, database.Db),
	}

	form.assignAttributes()

	return form
}

func (form *CreateExpenseForm) assignAttributes() {
	form.AddAttributes(
		&formAttributes.StringAttribute{
			FieldAttribute: formAttributes.FieldAttribute{
				Code: "Content",
			},
			Value: helpers.GetStringOrDefault(form.Input.Content),
		},
		&formAttributes.TimeAttribute{
			FieldAttribute: formAttributes.FieldAttribute{
				Code: "Date",
			},
			Value: helpers.GetStringOrDefault(form.Input.Date),
		},
	)
}

func (form *CreateExpenseForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	return form.repo.Create(&form.Expense)
}

func (form *CreateExpenseForm) validate() error {
	form.validateContent().validateDate().summaryErrors()

	if form.IsValid() {
		return nil
	}

	return exceptions.NewUnprocessableContentError("Please check your input", form.Errors)
}

func (form *CreateExpenseForm) validateContent() *CreateExpenseForm {
	field := form.GetAttribute("Content")

	field.ValidateRequired()

	if field.IsClean() {
		form.Expense.Content = *form.Input.Content
	}

	return form
}

func (form *CreateExpenseForm) validateDate() *CreateExpenseForm {
	field := form.GetAttribute("Date")

	field.ValidateRequired()

	if field.IsClean() {
		form.Expense.Date = *form.Input.Date
	}

	return form
}
