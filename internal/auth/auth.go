package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Auth interface {
	CreateToken(userID uuid.UUID, role string) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
	RefreshToken(tokenString string) (string, error)
	AuthRequired(c *fiber.Ctx) error
	AdminRequired(c *fiber.Ctx) error
}
