package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()

	if err != nil {
		return err
	}

	files := form.File["image"]
	filename := ""

	for _, file := range files {
		filename = file.Filename
		fmt.Println("file->", filename)

		if err := c.SaveFile(file, "./uploads/"+filename); err != nil {
			return err
		}
	}
	fmt.Println("file", filename)
	return c.JSON(fiber.Map{
		"url": "http://localhost:8000/api/uploads/" + filename,
	})
}
