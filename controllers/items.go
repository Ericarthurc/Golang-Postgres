package controllers

import (
	"eric/database"
	"eric/models"

	"github.com/gofiber/fiber/v2"
)

func GetItemsHandler(c *fiber.Ctx) error {
	var err error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "data": err})
}

func GetItemHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")

	item, err := models.GetItem(database.DBPool, idParam)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "data": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"success": true, "data": item})
}

func CreateItemHandler(c *fiber.Ctx) error {
	var err error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "data": err})
}

func UpdateItemHandler(c *fiber.Ctx) error {
	var err error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "data": err})
}

func DeleteItemHandler(c *fiber.Ctx) error {
	var err error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "data": err})
}
