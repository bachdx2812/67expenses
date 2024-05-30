package main

import (
	"os"
	"server/initializers"
)

func main() {
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")
	initializers.LoadEnv()
	initializers.LoadDb()
}
