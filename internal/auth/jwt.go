package auth

import (
	"time"

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
