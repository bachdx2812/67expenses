package resolvers

import (
	"context"
	gqlInputs "server/app/gql/inputs"
	gqlTypes "server/app/gql/types"
	"server/app/services"
	"server/database"
)

func (r *Resolver) SignIn(ctx context.Context, args gqlInputs.SignInInput) (*gqlTypes.AccessToken, error) {
	service := services.AuthService{
		SignInInputForm: args.Input,
		Db:              database.Db,
	}

	token, err := service.Execute()
	if err != nil {
		return nil, err
	}

	return &gqlTypes.AccessToken{Token: *token}, nil
}
