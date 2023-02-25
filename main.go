package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	connectToDb()

	router := gin.Default()
	bookAddRoutes(router)
	router.Run(":" + port)

}
