package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thephiri/go-vue/database"
	"github.com/thephiri/go-vue/models"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
