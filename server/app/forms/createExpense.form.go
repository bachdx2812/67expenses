package forms

import (
	"server/app/constants"
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
		&formAttributes.IntAttribute[int32]{
			FieldAttribute: formAttributes.FieldAttribute{
				Code: "ExpenseTypeId",
			},
			Value: helpers.GetInt32OrDefault(form.Input.ExpenseTypeId),
		},
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
		&formAttributes.IntAttribute[int32]{
			FieldAttribute: formAttributes.FieldAttribute{
				Code: "Amount",
			},
			Value: helpers.GetInt32OrDefault(form.Input.Amount),
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
	form.validateExpenseType().
		validateContent().
		validateDate().
		validateAmount().
		summaryErrors()

	if form.IsValid() {
		return nil
	}

	return exceptions.NewUnprocessableContentError("Please check your input", form.Errors)
}

func (form *CreateExpenseForm) validateExpenseType() *CreateExpenseForm {
	field := form.GetAttribute("ExpenseTypeId")

	field.ValidateRequired()

	if field.IsClean() {
		expenseType := models.ExpenseType{ID: *form.Input.ExpenseTypeId}
		expenseTypeRepo := repositories.NewExpenseTypeRepository(nil, database.Db)

		if err := expenseTypeRepo.Find(&expenseType); err != nil {
			field.AddError("invalid expense type")
		}

		if field.IsClean() {
			form.Expense.ExpenseTypeId = expenseType.ID
		}
	}

	return form
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
	field.ValidateFormat(constants.DDMMYYYY_DateFormat, constants.HUMAN_DDMMYYYY_DateFormat)

	if field.IsClean() {
		form.Expense.Date = *form.Input.Date
	}

	return form
}

func (form *CreateExpenseForm) validateAmount() *CreateExpenseForm {
	field := form.GetAttribute("Amount")

	field.ValidateRequired()

	if field.IsClean() {
		form.Expense.Amount = *form.Input.Amount
	}

	return form
}
