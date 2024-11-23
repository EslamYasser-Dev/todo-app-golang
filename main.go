package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type Note struct {
	ID      string `json:"id"`
	Done    bool   `json:"done"`
	Content string `json:"content"`
}

func main() {
	fmt.Println("hello ")
	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load env file")
	}

	PORT := os.Getenv("PORT")
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
		return c.Status(201).JSON(fiber.Map{"status": "success", "data": note})
	})

	app.Patch("/api/notes/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, note := range notes {
			if fmt.Sprint(note.ID) == id {
				notes[i].Done = true
				return c.Status(200).JSON(notes[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	})

	app.Delete("/api/notes/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, note := range notes {
			if fmt.Sprint(note.ID) == id {
				notes = append(notes[:i], notes[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"status": "success"})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	})

	log.Fatal(app.Listen(":" + PORT))

}
