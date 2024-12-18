package auth

import (
	"event-booking/internal/api/responses"

	"github.com/gofiber/fiber/v2"
)

type Middleware struct {
	jwtService *JwtService
}

func NewMiddleware(jwtService *JwtService) *Middleware {
	return &Middleware{jwtService: jwtService}
}

func (m *Middleware) AuthRequired(c *fiber.Ctx) error {
	tokenString := c.Cookies("jwt")
	claims, err := m.jwtService.ValidateToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.NewErrorResponse("Unauthorized"))
	}

	c.Locals("userID", claims.UserID)
	c.Locals("role", claims.Role)
	return c.Next()
}

func (m *Middleware) AdminRequired(c *fiber.Ctx) error {
	if err := m.AuthRequired(c); err != nil {
		return err
	}

	role := c.Locals("role").(string)
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(responses.NewErrorResponse("Admin access only"))
	}

	return c.Next()
}
