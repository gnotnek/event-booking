package health

import "github.com/gofiber/fiber/v2"

type httpHandler struct {
	service *Service
}

func NewHttpHandler(service *Service) *httpHandler {
	return &httpHandler{
		service: service,
	}
}

func (h *httpHandler) HealthCheck(c *fiber.Ctx) error {
	healthComponent, isHealthy := h.service.Check()
	if isHealthy {
		return c.JSON(fiber.Map{
			"status":     "ok",
			"components": healthComponent,
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status":     "fail",
		"components": healthComponent,
	})
}
