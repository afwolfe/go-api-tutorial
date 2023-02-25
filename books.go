package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

func bookAddRoutes(router *gin.Engine) {
	router.GET("/books", GetBooks)
	router.GET("/books/:id", GetBookById)
	router.POST("/books", PostCreateBook)
	router.PATCH("/books/:id/checkout", PatchCheckoutBookById)
	router.PATCH("/books/:id/return", PatchReturnBookById)
}

func PatchCheckoutBookById(c *gin.Context) {
	id := c.Param("id")

	book, err := bookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	} else if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "No copies of book available to borrow."})
		return
	}
	book.Quantity -= 1
	updatedBook, err2 := updateBook(*book)
	if err2 != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Unknown error occurred borrowing book."})
		return
	}
	c.IndentedJSON(http.StatusOK, updatedBook)
}

func PatchReturnBookById(c *gin.Context) {
	id := c.Param("id")

	book, err := bookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	book.Quantity += 1
	updatedBook, err2 := updateBook(*book)
	if err2 != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Unknown error occurred returning book."})
		return
	}
	c.IndentedJSON(http.StatusOK, updatedBook)
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	book, err := bookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func GetBooks(c *gin.Context) {
	books, err := listBooks()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

func PostCreateBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	createdBook, err := createBook(newBook)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "An unknown error occurred while creating the book."})
		return
	}
	c.IndentedJSON(http.StatusCreated, createdBook)
}
