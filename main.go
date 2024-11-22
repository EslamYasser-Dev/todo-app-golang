package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Note struct {
	ID      string `json:"id"`
	Done    bool   `json:"done"`
	Content string `json:"content"`
}

func main() {
	fmt.Println("hello ")
	app := fiber.New()
	notes := []Note{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello"})
	})

	app.Post("/api/notes", func(c *fiber.Ctx) error {
		note := &Note{}

		if err := c.BodyParser(note); err != nil {
			return err
		}

		if note.Content == "" {
			return c.Status(400).JSON(fiber.Map{"error": "content is required"})
		}

		note.ID = uuid.NewString()
		notes = append(notes, *note)
		return c.Status(201).JSON(note)
	})

	app.Patch("/api/notes:id", func(c *fiber.Ctx) error {
		return nil
	})

	log.Fatal(app.Listen(":4000"))

}
