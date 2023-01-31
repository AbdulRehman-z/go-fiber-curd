package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)


// Book struct (Model)

type Book struct {
	ID     int `json:"id"`
	Name   string `json:"name"`
	Pages  string `json:"pages"`
    Author *Author `json:"author"`
}

type Author struct {
	ID   int `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}



var books []Book

// route handlers


func main() {
	// Create new Fiber instance
	app := fiber.New()

    books = append(books, Book{ID: 1, Name: "The Alchemist", Pages: "200", Author: 
	&Author{ID: 1, Name: "Paulo Coelho", Age: "50"}})
	// Routes
	app.Get("/books",getBooksController)
	app.Get("/books/:id",getBookController)
	app.post("/books",addBookController)
	app.put("/books/:id",updateBookController)
	app.delete("/books/:id",deleteBookController)

	// Start server
	if err := app.Listen(":8000"); err != nil {
		log.Fatal(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
