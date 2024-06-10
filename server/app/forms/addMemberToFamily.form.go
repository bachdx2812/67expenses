package forms

import (
	"server/app/exceptions"
	"server/app/formAttributes"
	gqlInputs "server/app/gql/inputs"
	"server/app/helpers"
	"server/app/models"
	"server/app/repositories"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AddMemberToFamilyForm struct {
	Form
	Input  *gqlInputs.AddMemberToFamilyInputForm
	Family models.Family
	User   models.User
	repo   repositories.UserRepository
}

func NewAddMemberToFamilyForm(
	input *gqlInputs.AddMemberToFamilyInputForm,
	family *models.Family,
	db *gorm.DB,
) AddMemberToFamilyForm {
	form := AddMemberToFamilyForm{
		Form:   Form{},
		Input:  input,
		Family: *family,
		repo:   *repositories.NewUserRepository(nil, db),
	}

	form.assignAttributes()

	return form
}

func (form *AddMemberToFamilyForm) assignAttributes() {
	form.AddAttributes(
		&formAttributes.StringAttribute{
			FieldAttribute: formAttributes.FieldAttribute{
				Code: "Phone",
			},
			Value: helpers.GetStringOrDefault(form.Input.Phone),
		},
		&formAttributes.StringAttribute{
			FieldAttribute: formAttributes.FieldAttribute{
				Code: "Password",
			},
			Value: helpers.GetStringOrDefault(form.Input.Password),
		},
		&formAttributes.StringAttribute{
			FieldAttribute: formAttributes.FieldAttribute{
				Code: "Name",
			},
			Value: helpers.GetStringOrDefault(form.Input.Name),
		},
	)
}

func (form *AddMemberToFamilyForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	return form.repo.AddMemberToFamily(form.User, form.Family)
}

func (form *AddMemberToFamilyForm) validate() error {
	form.ValidatePhone().ValidatePassword().ValidateName().summaryErrors()

	if form.IsValid() {
		return nil
	}

	return exceptions.NewUnprocessableContentError("Please check your input", form.Errors)
}

func (form *AddMemberToFamilyForm) ValidatePhone() *AddMemberToFamilyForm {
	field := form.GetAttribute("Phone")
	field.ValidateRequired()

	if field.IsClean() {
		form.User.Phone = *form.Input.Phone
	}

	return form
}

func (form *AddMemberToFamilyForm) ValidatePassword() *AddMemberToFamilyForm {
	field := form.GetAttribute("Password")
	field.ValidateRequired()

	if field.IsClean() {
		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(*form.Input.Password), bcrypt.DefaultCost)
		if err != nil {
			field.AddError(err)
		} else {
			form.User.EncryptedPassword = string(encryptedPassword)
		}
	}

	return form
}

func (form *AddMemberToFamilyForm) ValidateName() *AddMemberToFamilyForm {
	field := form.GetAttribute("Name")
	field.ValidateRequired()

	if field.IsClean() {
		form.User.Name = form.Input.Name
	}

	return form
}
