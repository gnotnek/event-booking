package account

import (
	"event-booking/internal/api/responses"
	"event-booking/internal/auth"
	"event-booking/internal/entity"
	"time"

	"github.com/gofiber/fiber/v2"
)

type httpHandler struct {
	svc *Service
	jwt auth.Auth
}

func NewHttpHandler(svc *Service, jwt auth.Auth) *httpHandler {
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
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse("Bad Request"))
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
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
	}

	return c.Status(fiber.StatusCreated).JSON(responses.NewSuccessResponse("User created successfully"))
}

type SignInPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *httpHandler) SignInUserHandler(c *fiber.Ctx) error {
	user := new(SignInPayload)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse("Bad Request"))
	}

	userEntity := &entity.User{
		Email:    user.Email,
		Password: user.Password,
	}

	authenticatedUser, err := h.svc.SignInUserService(userEntity)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.NewErrorResponse("Unauthorized"))
	}

	token, err := h.jwt.CreateToken(authenticatedUser.ID, authenticatedUser.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
	}

	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
	})

	userDTO := responses.UserResponseObject{ID: authenticatedUser.ID, Name: authenticatedUser.Name, Email: authenticatedUser.Email, Role: authenticatedUser.Role}

	return c.Status(fiber.StatusOK).JSON(responses.DataResponse{
		Message: "Log in successful",
		Data:    userDTO,
	})
}

func (h *httpHandler) SignOutUserHandler(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})

	return c.Status(fiber.StatusOK).JSON(responses.NewSuccessResponse("User signed out successfully"))
}

func (h *httpHandler) GetUserByIDHandler(c *fiber.Ctx) error {
	userID := c.Params("id")

	user, err := h.svc.FindByIDService(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.NewErrorResponse("User not found"))
	}

	userDTO := responses.UserResponseObject{ID: user.ID, Email: user.Email, Role: user.Role}

	return c.Status(fiber.StatusOK).JSON(responses.DataResponse{
		Message: "User found",
		Data:    userDTO,
	})
}

func (h *httpHandler) RefreshTokenHandler(c *fiber.Ctx) error {
	token := c.Cookies("jwt")

	newToken, err := h.jwt.RefreshToken(token)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse("Bad Request"))
	}

	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   newToken,
		Expires: time.Now().Add(time.Hour * 24),
	})

	return c.Status(fiber.StatusOK).JSON(responses.NewSuccessResponse("Token refreshed successfully"))
}

type UpdateUserPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (h *httpHandler) UpdateUserHandler(c *fiber.Ctx) error {
	user := new(UpdateUserPayload)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse("Bad Request"))
	}

	newUser := &entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}

	err := h.svc.UpdateUserService(newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewSuccessResponse("User updated successfully"))
}
