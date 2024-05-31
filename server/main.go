package main

import (
	"os"
	"server/app/models"
	"server/app/repository"
	"server/database"
	"server/initializers"

	"github.com/gin-gonic/gin"
)

func main() {
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")
	initializers.LoadEnv()
	initializers.LoadDb()

	r := gin.Default()

	r.Use(initializers.CorsConfig())

	r.POST("/gql", initializers.GqlHandler(database.Db))

	userRepo := repository.NewUserRepository(nil, database.Db)
	user := models.User{
		Phone: "0865882991",
	}

	userRepo.Find(&user)

	r.Run()
}
