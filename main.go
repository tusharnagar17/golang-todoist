package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// fmt.Println("App is listening to prot 4000")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "Successfully at / endpoint"})
	})

	log.Fatal(app.Listen(":4000"))
}
