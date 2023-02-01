package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)


// Book struct (Model)

type Book struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Pages  uint `json:"pages"`
    Author *Author `json:"author"`
}

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  uint `json:"age"`
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
 }}

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


func main() {
	// Create new Fiber instance
	app := fiber.New()

   app.Use(logger.New())

    books = append(books, Book{ID: "1", Name: "The Alchemist", Pages: 200, Author: 
	&Author{ID: "1", Name: "Paulo Coelho", Age: 50}})

    books = append(books,Book{ID:"2",Name:"The Monk Who Sold His Ferrari",Pages:300,Author: &Author{ID:"2",Name:"Robin Sharma",Age:40}})


	// Routes
   app.Route("/api/v1/books",func(router fiber.Router) {
	router.Get("/",getBooksController)
	router.Get("/:id<min(0)>",getBookController)
	router.Post("/addBook",addBookController)
	// router.put("/books/:id",updateBookController)
	// router.delete("/books/:id",deleteBookController)

   })

	// app.Get("/books",getBooksController)
	// app.Get("/books/:id<min(0)>",getBookController)
	// app.Post("/books",addBookController)
	// // app.put("/books/:id",updateBookController)
	// // app.delete("/books/:id",deleteBookController)

	// Start server
	log.Fatal(app.Listen(":8000"))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
