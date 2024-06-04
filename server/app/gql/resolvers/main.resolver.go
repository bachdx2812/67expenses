package resolvers

import (
	"context"
	"server/app/exceptions"
	"server/app/models"
	"server/app/pkg/auths"

	"gorm.io/gorm"
)

type Resolver struct {
	Db *gorm.DB
}

func (r *Resolver) AuthUserFromCtx(ctx context.Context) (*models.User, error) {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("unauthorized")
	}

	return &user, nil
}
