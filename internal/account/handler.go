package account

import (
	"event-booking/internal/auth"
	"event-booking/internal/entity"
	"time"

	"github.com/gofiber/fiber/v2"
)

type httpHandler struct {
	svc *Service
	jwt auth.Jwt
}

func NewHttpHandler(svc *Service, jwt auth.Jwt) *httpHandler {
	return &httpHandler{
		svc: svc,
		jwt: jwt,
	}
}

type SignUpPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role" default:"user"`
}

func (h *httpHandler) SignUpUserHandler(c *fiber.Ctx) error {
	user := new(SignUpPayload)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	newUser := &entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	if user.Role != "" {
		newUser.Role = user.Role
	}

	err := h.svc.SignUpUserService(newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    &newUser,
	})
}

type SignInPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *httpHandler) SignInUserHandler(c *fiber.Ctx) error {
	user := new(SignInPayload)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	userEntity := &entity.User{
		Email:    user.Email,
		Password: user.Password,
	}

	authenticatedUser, err := h.svc.SignInUserService(userEntity)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Unauthorized",
		})
	}

	token, err := h.jwt.CreateToken(authenticatedUser.ID, authenticatedUser.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User signed in successfully",
		"user":    authenticatedUser,
	})
}
