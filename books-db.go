package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func connectToDb() {
	address := GetenvOrElse("DB_HOST", "127.0.0.1") + ":" + GetenvOrElse("DB_PORT", "3306")
	cfg := mysql.Config{
		User:                 GetenvOrElse("DB_USER", "root"),
		Passwd:               GetenvOrElse("DB_PASS", "password"),
		Net:                  "tcp",
		Addr:                 address,
		DBName:               "booksdb",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func createBook(book Book) (*Book, error) {
	_, err := db.Exec("INSERT INTO books (title, author, quantity) VALUES (?, ?, ?)", book.Title, book.Author, book.Quantity)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func updateBook(book Book) (*Book, error) {
	_, err := db.Exec("UPDATE books SET title = ?, author = ?, quantity = ? WHERE id = ?", book.Title, book.Author, book.Quantity, book.Id)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func listBooks() ([]Book, error) {
	var books []Book
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, fmt.Errorf("listBooks: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Quantity); err != nil {
			return nil, fmt.Errorf("listBooks: %v", err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("listBooks: %v", err)
	}
	return books, nil
}

func bookById(id string) (*Book, error) {
	var book Book
	row := db.QueryRow("SELECT * FROM books WHERE id = ?", id)
	if row != nil {
		if err := row.Scan(&book.Id, &book.Title, &book.Author, &book.Quantity); err != nil {
			return nil, fmt.Errorf("bookById %q: %v", id, err)
		}
		return &book, nil
	}

	return nil, fmt.Errorf("bookById %q not found", id)
}
