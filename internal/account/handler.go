package account

import (
	"event-booking/internal/api/responses"
	"event-booking/internal/api/validator"
	"event-booking/internal/auth"
	"event-booking/internal/entity"
	"time"

	"github.com/gofiber/fiber/v2"
)

type httpHandler struct {
	svc       *Service
	jwt       *auth.JwtService
	validator *validator.Validator
}

func NewHttpHandler(svc *Service, jwt *auth.JwtService, validator *validator.Validator) *httpHandler {
	return &httpHandler{
		svc:       svc,
		jwt:       jwt,
		validator: validator,
	}
}

type SignUpPayload struct {
	Name     string `json:"name" validate:"required,min=3,max=50,name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=50"`
	Role     string `json:"role" validate:"omitempty,oneof=admin user" default:"user"`
}

func (h *httpHandler) SignUpUserHandler(c *fiber.Ctx) error {
	user := new(SignUpPayload)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse("Bad Request"))
	}

	if err := h.validator.ValidateStruct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse(err.Error()))
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
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

func (h *httpHandler) SignInUserHandler(c *fiber.Ctx) error {
	user := new(SignInPayload)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse("Bad Request"))
	}

	if err := h.validator.ValidateStruct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse(err.Error()))
	}

	userEntity := &entity.User{
		Email:    user.Email,
		Password: user.Password,
	}

	authenticatedUser, err := h.svc.SignInUserService(userEntity)
	if err != nil {
		if err.Error() == "user is not verified" {
			return c.Status(fiber.StatusUnauthorized).JSON(responses.NewErrorResponse("Please verify your email"))
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(responses.NewErrorResponse("Invalid email or password"))
		}
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
	Name     string `json:"name" validate:"required,min=3,max=50,name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

func (h *httpHandler) UpdateUserHandler(c *fiber.Ctx) error {
	user := new(UpdateUserPayload)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse("Bad Request"))
	}

	if err := h.validator.ValidateStruct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse(err.Error()))
	}

	newUser := &entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	err := h.svc.UpdateUserService(newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewSuccessResponse("User updated successfully"))
}

type RequestVerificationCodePayload struct {
	Email string `json:"email" validate:"required,email"`
}

func (h *httpHandler) RequestVerificationCodeHandler(c *fiber.Ctx) error {
	email := new(RequestVerificationCodePayload)
	if err := c.BodyParser(email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse("Bad Request"))
	}

	if err := h.validator.ValidateStruct(email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse(err.Error()))
	}

	err := h.svc.GenerateVerificationCode(email.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewSuccessResponse("Verification code sent to your email successfully"))
}

type ValidateVerificationCodePayload struct {
	Code  string `json:"code" validate:"required,len=6"`
	Email string `json:"email" validate:"required,email"`
}

func (h *httpHandler) ValidateVerificationCodeHandler(c *fiber.Ctx) error {
	var payload ValidateVerificationCodePayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse("Bad Request"))
	}

	err := h.svc.ValidateVerificationCode(payload.Email, payload.Code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewSuccessResponse("Verification code validated successfully"))
}
