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

func (j *JwtService) ValidateToken(tokenString string) (*claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.jwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
