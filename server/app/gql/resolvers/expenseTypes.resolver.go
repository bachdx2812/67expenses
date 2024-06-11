package resolvers

import (
	"context"
	gqlTypes "server/app/gql/types"
	"server/app/models"
	"server/app/repositories"
	"server/database"
)

func (r *Resolver) ExpenseTypes(ctx context.Context) ([]*gqlTypes.ExpenseTypeType, error) {
	var expenseTypes []*models.ExpenseType

	repo := repositories.NewExpenseTypeRepository(&ctx, database.Db)

	if err := repo.List(&expenseTypes); err != nil {
		return nil, err
	}

	expenseTypeTypes := make([]*gqlTypes.ExpenseTypeType, len(expenseTypes))

	for i, expenseType := range expenseTypes {
		expenseTypeTypes[i] = &gqlTypes.ExpenseTypeType{
			ExpenseType: *expenseType,
		}
	}

	return expenseTypeTypes, nil
}
