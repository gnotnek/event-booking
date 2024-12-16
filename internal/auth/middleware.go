package auth

import (
	"event-booking/internal/api/responses"
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
	RefreshToken(tokenString string) (string, error)
}

type AuthService struct {
	jwtKey string
}

func NewAuthService(jwtKey string) *AuthService {
	return &AuthService{jwtKey: jwtKey}
}

type claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
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
		log.Error().Err(err).Msg("Failed to validate token")
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	claims, ok := token.Claims.(*claims)
	if !ok {
		log.Error().Msg("Invalid claims")
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

	c.Locals("userID", claims.UserID)
	c.Locals("role", claims.Role)
	return c.Next()
}

func (j *AuthService) AdminRequired(c *fiber.Ctx) error {
	// Extract JWT token from cookies
	tokenString := c.Cookies("jwt")
	if tokenString == "" {
		log.Error().Msg("JWT cookie is missing")
		return c.Status(fiber.StatusUnauthorized).JSON(responses.NewErrorResponse("Unauthorized"))
	}

	// Validate token
	claims, err := j.ValidateToken(tokenString)
	if err != nil {
		log.Error().Err(err).Msg("Failed to validate token")
		return c.Status(fiber.StatusUnauthorized).JSON(responses.NewErrorResponse("Unauthorized"))
	}

	// Check if the user has the admin role
	if claims.Role != "admin" {
		log.Warn().
			Str("userID", claims.UserID).
			Str("role", claims.Role).
			Msg("Access denied: Admin access only")
		return c.Status(fiber.StatusForbidden).JSON(responses.NewErrorResponse("Access denied: Admin access only"))
	}

	// Store user data in locals for later use
	c.Locals("userID", claims.UserID)
	c.Locals("role", claims.Role)

	return c.Next()
}

func (j *AuthService) RefreshToken(tokenString string) (string, error) {
	claims, err := j.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	return j.CreateToken(uuid.MustParse(claims.UserID), claims.Role)
}
