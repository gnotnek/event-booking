package auth

import (
	"event-booking/internal/api/responses"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
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
	tokenString := c.Cookies("jwt")
	if tokenString == "" {
		log.Error().Msg("JWT cookie is missing")
		return c.Status(fiber.StatusUnauthorized).JSON(responses.NewErrorResponse("Unauthorized"))
	}

	claims, err := m.jwtService.ValidateToken(tokenString)
	if err != nil {
		log.Error().Err(err).Msg("Failed to validate token")
		return c.Status(fiber.StatusUnauthorized).JSON(responses.NewErrorResponse("Unauthorized"))
	}

	if claims.Role != "admin" {
		log.Warn().
			Str("userID", claims.UserID).
			Str("role", claims.Role).
			Msg("Access denied: Admin access only")
		return c.Status(fiber.StatusForbidden).JSON(responses.NewErrorResponse("Access denied: Admin access only"))
	}

	c.Locals("userID", claims.UserID)
	c.Locals("role", claims.Role)

	return c.Next()
}
