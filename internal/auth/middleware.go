package auth

import (
	"github.com/gofiber/fiber/v2"
)

type middleware struct {
	secretKey string
}

func NewMiddleware(secretKey string) *middleware {
	return &middleware{secretKey}
}

func (m *middleware) StaticToken(c *fiber.Ctx) error {
	tokenStr := c.Get("Authorization")
	if tokenStr != m.secretKey {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Unauthorized",
		})
	}

	return c.Next()
}
