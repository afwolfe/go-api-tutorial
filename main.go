package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

type book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{Id: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{Id: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{Id: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	BookAddRoutes(router)
	router.Run(":" + port)

}
