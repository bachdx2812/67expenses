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
	input *gqlInputs.NewExpenseInputForm
	models.Expense
	repo repositories.ExpenseRepository
}

func NewCreateExpenseForm(
	input *gqlInputs.NewExpenseInputForm,
	expense *models.Expense,
) CreateExpenseForm {
	form := CreateExpenseForm{
		Form:  Form{},
		input: input,
		repo:  *repositories.NewExpenseRepository(nil, database.Db),
	}

	form.assignAttributes(input)

	return form
}

func (form *CreateExpenseForm) assignAttributes(input *gqlInputs.NewExpenseInputForm) {
	form.AddAttributes(
		&formAttributes.StringAttribute{
			FieldAttribute: formAttributes.FieldAttribute{
				Code: "Content",
			},
			Value: helpers.GetStringOrDefault(input.Content),
		},
		&formAttributes.TimeAttribute{
			FieldAttribute: formAttributes.FieldAttribute{
				Code: "Date",
			},
			Value: helpers.GetStringOrDefault(input.Date),
		},
	)
}

func (form *CreateExpenseForm) Save() error {
	if err := form.Validate(); err != nil {
		return err
	}

	return form.repo.Create(&form.Expense)
}

func (form *CreateExpenseForm) Validate() error {
	form.ValidateContent().ValidateDate().summaryErrors()

	if form.IsValid() {
		return nil
	}

	return exceptions.NewUnprocessableContentError("Please check your input", form.Errors)
}

func (form *CreateExpenseForm) ValidateContent() *CreateExpenseForm {
	field := form.GetAttribute("Content")

	field.ValidateRequired()

	if field.IsClean() {
		form.Expense.Content = form.Content
	}

	return form
}

func (form *CreateExpenseForm) ValidateDate() *CreateExpenseForm {
	field := form.GetAttribute("Date")

	field.ValidateRequired()

	if field.IsClean() {
		form.Expense.Date = form.Date
	}

	return form
}
