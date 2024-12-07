package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Auth interface {
	CreateToken(userID uuid.UUID, role string) (string, error)
	ValidateToken(tokenString string) (*claims, error)
	AuthRequired(c *fiber.Ctx) error
	AdminRequired(c *fiber.Ctx) error
}

type AuthService struct {
	jwtKey string
}

func NewAuthService(jwtKey string) *AuthService {
	return &AuthService{jwtKey: jwtKey}
}

type claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"` // Added Role field
	jwt.RegisteredClaims
}

func (j *AuthService) CreateToken(userID uuid.UUID, role string) (string, error) {
	claims := claims{
		UserID: userID.String(),
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.jwtKey))
}

func (j *AuthService) ValidateToken(tokenString string) (*claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.jwtKey), nil
	})
	if err != nil || !token.Valid {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	claims, ok := token.Claims.(*claims)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid claims")
	}

	return claims, nil
}

func (j *AuthService) AuthRequired(c *fiber.Ctx) error {
	tokenString := c.Cookies("jwt")
	claims, err := j.ValidateToken(tokenString)
	if err != nil {
		log.Error().Err(err).Msg("Failed to validate token")
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	log.Info().
		Str("userID", claims.UserID).
		Str("role", claims.Role).
		Msg("Token validated successfully")

	c.Locals("userID", claims.UserID)
	c.Locals("role", claims.Role)
	return c.Next()
}

func (j *AuthService) AdminRequired(c *fiber.Ctx) error {
	// Extract JWT token from cookies
	tokenString := c.Cookies("jwt")
	if tokenString == "" {
		log.Error().Msg("JWT cookie is missing")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	// Validate token
	claims, err := j.ValidateToken(tokenString)
	if err != nil {
		log.Error().Err(err).Msg("Failed to validate token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	log.Info().
		Str("userID", claims.UserID).
		Str("role", claims.Role).
		Msg("Token validated successfully")

	// Check if the user has the admin role
	if claims.Role != "admin" {
		log.Warn().
			Str("userID", claims.UserID).
			Str("role", claims.Role).
			Msg("Access denied: Admin access only")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Admin access only"})
	}

	// Store user data in locals for later use
	c.Locals("userID", claims.UserID)
	c.Locals("role", claims.Role)

	log.Info().Msg("Admin access granted")
	return c.Next()
}
