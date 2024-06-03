package services

import (
	"context"
	"server/app/exceptions"
	"server/app/helpers"
	"server/app/repository"

	"gorm.io/gorm"
)

type AuthService struct {
	Phone    *string
	Password *string

	Ctx *context.Context
	Db  *gorm.DB
}

func (service *AuthService) Execute() (*string, error) {
	if err := service.validate(); err != nil {
		return nil, err
	}

	if token, err := service.getUserAndGenerateToken(); err != nil {
		return nil, err
	} else {
		return token, nil
	}
}

func (service *AuthService) validate() error {
	exception := exceptions.NewUnprocessableContentError("Please check your input", nil)

	if service.Phone == nil || *service.Phone == "" {
		exception.AddError("Phone", []interface{}{"is required"})
	}

	if service.Password == nil || *service.Password == "" {
		exception.AddError("Password", []interface{}{"is required"})
	}

	if len(exception.Errors) > 0 {
		return exception
	}

	return nil
}

func (service *AuthService) getUserAndGenerateToken() (*string, error) {
	repo := repository.NewUserRepository(service.Ctx, service.Db)

	if user, err := repo.Auth(*service.Phone, *service.Password); err != nil {
		return nil, exceptions.NewUnprocessableContentError("User not found", exceptions.ResourceModificationError{
			"phone": []interface{}{"is not found"},
		})
	} else {
		token, err := helpers.GenerateJwtToken(user.GenerateJwtClaims())

		if err != nil {
			return nil, exceptions.NewBadRequestError("Failed to generate JWT token")
		}

		return &token, nil
	}
}
