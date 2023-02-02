package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Book struct (Model)

type Book struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Pages  uint    `json:"pages"`
	Author *Author `json:"author"`
}

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

var books []Book

/**
Routes Handlers
**/
// Get all books
func getBooksController(c *fiber.Ctx) error {
	return c.JSON(books)
}

// Get single book
func getBookController(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, book := range books {
		if book.ID == id {
			return c.JSON(book)
		}
	}

	return c.Status(404).JSON(fiber.Map{"message": "Book not found"})
}

// Add new book
func addBookController(c *fiber.Ctx) error {
	var book Book
	if err := c.BodyParser(&book); err != nil {
		return err
	}

	books = append(books, book)
	return c.Status(201).JSON(book)
}

// Delete book
func deleteBookController(c *fiber.Ctx) error {
	id := c.Params("id")
	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			return c.Status(200).JSON(fiber.Map{"message": "Book deleted"})
		}
	}
	return c.Status(404).JSON(fiber.Map{"message": "Book not found"})
}

func main() {
	// Create new Fiber instance
	app := fiber.New()

	app.Use(logger.New())

	books = append(books, Book{ID: "1", Name: "The Alchemist", Pages: 200,
		Author: &Author{ID: "1", Name: "Paulo Coelho", Age: 50}})

	books = append(books, Book{ID: "2", Name: "The Monk Who Sold His Ferrari", Pages: 300,
		Author: &Author{ID: "2", Name: "Robin Sharma", Age: 40}})

	// Routes
	v1 := app.Group("/api/v1/books")
	v1.Get("/", getBooksController)
	v1.Get("/:id<min(0)>", getBookController)
	v1.Post("/", addBookController)
	v1.Delete("/book", deleteBookController)
	// Start server
	log.Fatal(app.Listen(":8000"))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
