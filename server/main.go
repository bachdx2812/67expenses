package main

import (
	"os"
	"server/database"
	"server/initializers"

	"github.com/gin-gonic/gin"
)

func main() {
	os.Setenv("TZ", os.Getenv("TIME_ZONE"))
	initializers.LoadEnv()
	initializers.LoadDb()

	r := gin.Default()

	r.Use(initializers.CorsConfig())

	r.POST("/gql", initializers.GqlHandler(database.Db))

	r.Run()
}
