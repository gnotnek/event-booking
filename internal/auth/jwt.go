package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Jwt interface {
	CreateToken(userID uuid.UUID) (string, error)
}

type JwtService struct {
	jwtKey string
}

func NewJwtService(jwtKey string) *JwtService {
	return &JwtService{
		jwtKey: jwtKey,
	}
}

type claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func (j *JwtService) CreateToken(userID uuid.UUID) (string, error) {
	claims := claims{
		UserID: userID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.jwtKey))
}

func (j *JwtService) ValidateToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.jwtKey), nil
	})
	if err != nil {
		return nil
	}

	_, ok := token.Claims.(*claims)
	if !ok || !token.Valid {
		return err
	}

	return nil
}

func (j *JwtService) AuthRequired(c *fiber.Ctx) error {
	err := j.ValidateToken(c.Cookies("jwt"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	return nil
}
