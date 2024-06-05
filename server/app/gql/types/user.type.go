package gqlTypes

import "server/app/models"

type UserType struct {
	User models.User
}

func (u *UserType) ID() int32 {
	return u.User.ID
}

func (u *UserType) Name() string {
	return *u.User.Name
}

func (u *UserType) Phone() string {
	return u.User.Phone
}

func (u *UserType) Family() *FamilyType {
	return &FamilyType{
		Family: u.User.Family,
	}
}

func (u *UserType) FamilyId() int32 {
	return u.User.FamilyId
}
