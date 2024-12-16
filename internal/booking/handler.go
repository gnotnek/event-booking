package booking

import (
	"event-booking/internal/api/responses"
	"event-booking/internal/api/validator"
	"event-booking/internal/entity"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type httpHandler struct {
	svc       *Service
	validator *validator.Validator
}

func NewHttpHandler(svc *Service, validator *validator.Validator) *httpHandler {
	return &httpHandler{
		svc:       svc,
		validator: validator,
	}
}

type BookingInputPayload struct {
	UserID   uuid.UUID `json:"user_id" validate:"required"`
	EventID  uuid.UUID `json:"event_id" validate:"required"`
	Quantity int       `json:"quantity" validate:"required"`
}

func (h *httpHandler) BookEventHandler(c *fiber.Ctx) error {
	book := new(BookingInputPayload)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse("Bad Request"))
	}

	if err := h.validator.ValidateStruct(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse(err.Error()))
	}

	newBook := &entity.Booking{
		UserID:   book.UserID,
		EventID:  book.EventID,
		Quantity: book.Quantity,
	}

	newBook, err := h.svc.CreateBookingService(newBook)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
	}

	return c.Status(fiber.StatusCreated).JSON(responses.NewDataResponse("Booking created successfully", responses.BookingResponseObject{
		ID:         newBook.ID,
		UserID:     newBook.UserID,
		EventID:    newBook.EventID,
		Quantity:   newBook.Quantity,
		TotalPrice: newBook.TotalPrice,
		CreatedAt:  newBook.CreatedAt,
		UpdatedAt:  newBook.UpdatedAt,
	}))
}

func (h *httpHandler) GetBookedEventsHandler(c *fiber.Ctx) error {
	bookings, err := h.svc.FindAllBookingService()
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return c.Status(fiber.StatusNotFound).JSON(responses.NewErrorResponse("Booking not found"))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
		}
	}

	var bookedEvents []responses.BookingResponseObject
	for _, book := range bookings {
		bookedEvents = append(bookedEvents, responses.BookingResponseObject{
			ID:         book.ID,
			UserID:     book.UserID,
			EventID:    book.EventID,
			Quantity:   book.Quantity,
			TotalPrice: book.TotalPrice,
			CreatedAt:  book.CreatedAt,
			UpdatedAt:  book.UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewDataResponse("Bookings found", bookedEvents))
}

func (h *httpHandler) GetBookedEventByIDHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	book, err := h.svc.FindBookingService(id)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return c.Status(fiber.StatusNotFound).JSON(responses.NewErrorResponse("Booking not found"))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
		}
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewDataResponse("Booking found", responses.BookingResponseObject{
		ID:         book.ID,
		UserID:     book.UserID,
		EventID:    book.EventID,
		Quantity:   book.Quantity,
		TotalPrice: book.TotalPrice,
		CreatedAt:  book.CreatedAt,
		UpdatedAt:  book.UpdatedAt,
	}))
}

func (h *httpHandler) CancelBookedEventHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	book, err := h.svc.FindBookingService(id)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return c.Status(fiber.StatusNotFound).JSON(responses.NewErrorResponse("Booking not found"))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
		}
	}

	err = h.svc.DeleteBookingService(id, book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewSuccessResponse("Booking canceled successfully"))
}

func (h *httpHandler) UpdateBookedEventHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	newBook := new(BookingInputPayload)

	if err := c.BodyParser(newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse("Bad Request"))
	}

	if err := h.validator.ValidateStruct(newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse(err.Error()))
	}

	book, err := h.svc.SaveBookingService(id, *newBook)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return c.Status(fiber.StatusNotFound).JSON(responses.NewErrorResponse("Booking not found"))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
		}
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewDataResponse("Booking updated successfully", responses.BookingResponseObject{
		ID:         book.ID,
		UserID:     book.UserID,
		EventID:    book.EventID,
		Quantity:   book.Quantity,
		TotalPrice: book.TotalPrice,
		CreatedAt:  book.CreatedAt,
		UpdatedAt:  book.UpdatedAt,
	}))
}
