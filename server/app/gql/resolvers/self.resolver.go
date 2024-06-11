package resolvers

import (
	"context"
	gqlTypes "server/app/gql/types"
	"server/app/repositories"
	"server/database"
)

func (r *Resolver) Self(ctx context.Context) (*gqlTypes.UserType, error) {
	user, err := r.AuthUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	repo := repositories.NewUserRepository(
		&ctx,
		database.Db.Preload("Family.Users"),
	)

	if err := repo.Find(user); err != nil {
		return nil, err
	}

	return &gqlTypes.UserType{
		User: user,
	}, nil
}
