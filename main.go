package main

import "github.com/gofiber/fiber/v2"

type PostID int

type Data struct {
	ID    PostID `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var posts = map[PostID]Data{
	1: {
		ID:    001,
		Name:  "John Doe",
		Email: "John@Email.com",
	},
	2: {
		ID:    002,
		Name:  "Phoebe",
		Email: "Phoebe@Email.com",
	},
	3: {
		ID:    003,
		Name:  "Dyon",
		Email: "Dyon@Email.com",
	},
}

func main() {
	app := fiber.New()

	app.Get("/posts", func(c *fiber.Ctx) error {
		var result []Data
		for _, post := range posts {
			result = append(result, post)
		}
		return c.JSON(result)
	})
	app.Listen(":3000")
}
