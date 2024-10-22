package account

import (
	"event-booking/internal/auth"
	"event-booking/internal/entity"

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

func (h *httpHandler) SignUpUserHandler(c *fiber.Ctx) error {
	user := new(entity.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	user, err := h.svc.SignUpUserService(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *httpHandler) SignInUserHandler(c *fiber.Ctx) error {
	user := new(entity.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	user, err := h.svc.SignInUserService(user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Unauthorized",
		})
	}

	token, err := h.jwt.CreateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
		"user":  user,
	})
}
