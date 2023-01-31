package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)


// route handlers
func initialController(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func main() {
	// Create new Fiber instance
	app := fiber.New()


	// Routes
	app.Get("/", initialController)

	// Start server
	if err := app.Listen(":9000"); err != nil {
		log.Fatal(err)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
