package resolvers

import (
	"context"
	gqlInputs "server/app/gql/inputs"
	gqlTypes "server/app/gql/types"
	"server/app/models"
	"server/app/services"
	"server/database"
)

func (r *Resolver) CreateExpense(ctx context.Context, args gqlInputs.NewExpenseInput) (*gqlTypes.ExpenseType, error) {
	user, err := r.AuthUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	expense := models.Expense{UserId: user.ID}

	service := services.CreateExpenseService{
		Ctx:     &ctx,
		Db:      database.Db,
		Args:    args,
		Expense: &expense,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	}

	return &gqlTypes.ExpenseType{
		Expense: models.Expense{
			UserId: user.ID,
			User:   *user,
			// Date:    args.Date,
			// Content: args.Content,
		},
	}, nil
}
