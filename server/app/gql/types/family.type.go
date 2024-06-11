package gqlTypes

import "server/app/models"

type FamilyType struct {
	Family *models.Family
}

func (f *FamilyType) ID() int32 {
	return f.Family.ID
}

func (f *FamilyType) Name() string {
	return f.Family.Name
}

func (f *FamilyType) Users() []*UserType {
	users := make([]*UserType, len(f.Family.Users))

	for i, user := range f.Family.Users {
		users[i] = &UserType{
			User: &user,
		}
	}

	return users
}
