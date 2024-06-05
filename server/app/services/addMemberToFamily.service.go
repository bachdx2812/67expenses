package services

import (
	"server/app/forms"
	gqlInputs "server/app/gql/inputs"
	"server/app/models"

	"gorm.io/gorm"
)

type AddMemberToFamilyService struct {
	Db     *gorm.DB
	Args   gqlInputs.AddMemberToFamilyInput
	Family *models.Family
}

func (service *AddMemberToFamilyService) Execute() error {
	form := forms.NewAddMemberToFamilyForm(
		&service.Args.Input,
		service.Family,
		service.Db,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
