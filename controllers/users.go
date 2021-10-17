package controllers

import (
	"eric/database"
	"eric/models"
	"eric/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func getUserByCredentials(c *fiber.Ctx, username string, password string) (*models.User, error) {
	user, err := models.GetUserByUsername(database.DBPool, username)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("password incorrect")
	}
	return user, nil
}

func GetUsersHandler(c *fiber.Ctx) error {
	var err error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "data": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "data": nil})
}

func GetUserHandler(c *fiber.Ctx) error {
	var err error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "data": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "data": nil})
}

func LoginUserHandler(c *fiber.Ctx) error {
	var err error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "data": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "data": nil})
}

func LogoutUserHandler(c *fiber.Ctx) error {
	var err error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "data": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "data": nil})
}

func CreateUserHandler(c *fiber.Ctx) error {
	var err error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "data": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "data": nil})
}

func UpdateUserHandler(c *fiber.Ctx) error {
	var err error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "data": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "data": nil})
}

func DeleteUserHandler(c *fiber.Ctx) error {
	var err error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "data": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "data": nil})
}
