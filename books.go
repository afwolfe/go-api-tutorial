package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BookAddRoutes(router *gin.Engine) {
	router.GET("/books", GetBooks)
	router.GET("/books/:id", GetBookById)
	router.POST("/books", CreateBook)
	router.PATCH("/books/:id/checkout", PatchCheckoutBookById)
	router.PATCH("/books/:id/return", PatchReturnBookById)
}

func BookById(id string) (*book, error) {
	for i, b := range books {
		if b.Id == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func PatchCheckoutBookById(c *gin.Context) {
	id := c.Param("id")

	book, err := BookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	} else if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "No copies of book available to borrow."})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func PatchReturnBookById(c *gin.Context) {
	id := c.Param("id")

	book, err := BookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	book, err := BookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
