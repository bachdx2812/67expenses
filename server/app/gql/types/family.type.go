package gqlTypes

import "server/app/models"

type FamilyType struct {
	Family models.Family
}

func (f *FamilyType) ID() int32 {
	return f.Family.ID
}

func (f *FamilyType) Name() string {
	return f.Family.Name
}

func (f *FamilyType) Users() []*UserType {
	return nil
}
