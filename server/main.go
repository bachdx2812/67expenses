package main

import (
	"os"
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
}
