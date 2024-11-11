package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Jwt interface {
	CreateToken(userID uuid.UUID, role string) (string, error)
	ValidateToken(tokenString string) (*claims, error)
	AuthRequired(c *fiber.Ctx) error
	AdminOnly(c *fiber.Ctx) error
}

type JwtService struct {
	jwtKey string
}

func NewJwtService(jwtKey string) *JwtService {
	return &JwtService{jwtKey: jwtKey}
}

type claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"` // Added Role field
	jwt.RegisteredClaims
}

func (j *JwtService) CreateToken(userID uuid.UUID, role string) (string, error) {
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

func (j *JwtService) ValidateToken(tokenString string) (*claims, error) {
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

func (j *JwtService) AuthRequired(c *fiber.Ctx) error {
	tokenString := c.Cookies("jwt")
	claims, err := j.ValidateToken(tokenString)
	if err != nil {
		return err
	}

	c.Locals("userID", claims.UserID)
	c.Locals("role", claims.Role)
	return c.Next()
}

func (j *JwtService) AdminOnly(c *fiber.Ctx) error {
	// Run AuthRequired to validate token and set user role
	if err := j.AuthRequired(c); err != nil {
		return err
	}

	// Assert role as string and check if the role is "admin"
	role, ok := c.Locals("role").(string)
	if !ok || role != "admin" {
		log.Error().Msg("Admin access only")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Admin access only"})
	}

	return c.Next()
}
