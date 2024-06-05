package resolvers

import (
	"context"
	"server/app/exceptions"
	gqlInputs "server/app/gql/inputs"
	gqlTypes "server/app/gql/types"
	"server/app/models"
	"server/app/repositories"
	"server/app/services"
	"server/database"
)

func (r *Resolver) AddMemberToFamily(ctx context.Context, args gqlInputs.AddMemberToFamilyInput) (*gqlTypes.FamilyType, error) {
	user, err := r.AuthUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	family := models.Family{ID: user.FamilyId}
	familyRepo := repositories.NewFamilyRepository(nil, database.Db.Preload("Users"))

	if err := familyRepo.Find(&family); err != nil {
		return nil, exceptions.NewBadRequestError("Invalid family")
	}

	service := services.AddMemberToFamilyService{
		Db:     database.Db,
		Args:   args,
		Family: &family,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	}

	if err := familyRepo.Find(&family); err != nil {
		return nil, exceptions.NewBadRequestError("Invalid family")
	}

	return &gqlTypes.FamilyType{
		Family: family,
	}, nil
}
