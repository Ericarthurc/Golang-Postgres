package middleware

import (
	"eric/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func AuthenticateMiddleware(c *fiber.Ctx) error {
	cookieToken := c.Cookies("authToken")
	if cookieToken == "" {
		return errors.New("authentication failed, no token")
	}

	err := utils.VerifyJWTToken(c, cookieToken)
	if err != nil {
		return err
	}

	return c.Next()
}
