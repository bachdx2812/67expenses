package forms

import (
	"server/app/exceptions"
	"server/app/formAttributes"
	gqlInputs "server/app/gql/inputs"
	"server/app/helpers"
	"server/app/models"
	"server/app/repositories"
	"server/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AddMemberToFamilyForm struct {
	Form
	*gqlInputs.AddMemberToFamilyInputForm
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
		Form:                       Form{},
		AddMemberToFamilyInputForm: input,
		Family:                     *family,
		repo:                       *repositories.NewUserRepository(nil, db),
	}

	form.assignAttributes(input)

	return form
}

func (form *AddMemberToFamilyForm) assignAttributes(input *gqlInputs.AddMemberToFamilyInputForm) {
	form.AddAttributes(
		&formAttributes.StringAttribute{
			FieldAttribute: formAttributes.FieldAttribute{
				Code: "Phone",
			},
			Value: helpers.GetStringOrDefault(input.Phone),
		},
		&formAttributes.StringAttribute{
			FieldAttribute: formAttributes.FieldAttribute{
				Code: "Password",
			},
			Value: helpers.GetStringOrDefault(input.Password),
		},
		&formAttributes.StringAttribute{
			FieldAttribute: formAttributes.FieldAttribute{
				Code: "Name",
			},
			Value: helpers.GetStringOrDefault(input.Name),
		},
	)
}

func (form *AddMemberToFamilyForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	userRepository := repositories.NewUserRepository(nil, database.Db)

	return userRepository.AddMemberToFamily(form.User, form.Family)
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
		form.User.Phone = *form.Phone
	}

	return form
}

func (form *AddMemberToFamilyForm) ValidatePassword() *AddMemberToFamilyForm {
	field := form.GetAttribute("Password")

	field.ValidateRequired()

	if field.IsClean() {
		if encryptPassword, err := bcrypt.GenerateFromPassword([]byte(*form.Password), 10); err != nil {
			field.AddError(err)
		} else {
			form.User.EncryptedPassword = string(encryptPassword)
		}
	}

	return form
}

func (form *AddMemberToFamilyForm) ValidateName() *AddMemberToFamilyForm {
	field := form.GetAttribute("Name")

	field.ValidateRequired()

	if field.IsClean() {
		form.User.Name = form.Name
	}

	return form
}
