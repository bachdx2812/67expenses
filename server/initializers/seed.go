package initializers

import (
	"server/app/models"
	"server/database"
)

func Seed() {
	db := database.Db

	if db.Table("families").Find(&models.Family{}).RowsAffected == 0 {
		db.Table("families").Create(&models.Family{
			Name: "Bach Phuong",
			ID:   1,
		})
		db.Table("families").Create(&models.Family{
			Name: "Thai An",
			ID:   2,
		})
	}

	if db.Table("users").Find(&models.User{}).RowsAffected == 0 {
		name := "Bach"
		db.Table("users").Create(&models.User{
			FamilyId:          1,
			Name:              &name,
			Phone:             "0865882991",
			EncryptedPassword: "$2a$10$8080LRcosHOEqh6kpHbQie/vgubdGQkXIvPQUkAaCzGZCs6FMsmc6",
		})

		name = "Thai"
		db.Table("users").Create(&models.User{
			FamilyId: 2,
			Name:     &name,
			Phone:    "$2a$10$8080LRcosHOEqh6kpHbQie/vgubdGQkXIvPQUkAaCzGZCs6FMsmc6",
		})
	}

	if db.Table("expense_types").Find(&models.ExpenseType{}).RowsAffected == 0 {
		db.Table("expense_types").Create(&models.ExpenseType{
			Name: "Food",
			ID:   1,
		})

		db.Table("expense_types").Create(&models.ExpenseType{
			Name: "Transportation",
			ID:   2,
		})
	}
}
