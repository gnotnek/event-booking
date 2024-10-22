package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type middleware struct {
	secretKey string
	jwtKey    string
}

type jwtClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

func NewMiddleware(secretKey string, jwtjwtKey string) *middleware {
	return &middleware{
		secretKey: secretKey,
		jwtKey:    jwtjwtKey,
	}
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

func (m *middleware) CreateJWTToken(c *fiber.Ctx, userID uuid.UUID) (string, error) {
	claims := jwtClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(m.jwtKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (m *middleware) ValidateJWTToken(c *fiber.Ctx) (*jwtClaims, error) {
	tokenStr := c.Get("User-Token")[7:]
	claims := &jwtClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.jwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
