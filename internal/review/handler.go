package review

import (
	"event-booking/internal/api/responses"
	"event-booking/internal/api/validator"
	"event-booking/internal/entity"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

type ReviewPayload struct {
	EventID uuid.UUID `json:"event_id" validate:"required"`
	UserID  uuid.UUID `json:"user_id" validate:"required"`
	Review  string    `json:"review" validate:"required"`
	Rating  int       `json:"rating" validate:"required,min=1,max=5"`
}

func (h *httpHandler) CreateReviewHandler(c *fiber.Ctx) error {
	review := new(ReviewPayload)
	if err := c.BodyParser(review); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse("Bad Request"))
	}

	if err := h.validator.ValidateStruct(review); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse(err.Error()))
	}

	newReview := &entity.Review{
		EventID: review.EventID,
		UserID:  review.UserID,
		Review:  review.Review,
		Rating:  review.Rating,
	}

	createdReview, err := h.svc.CreateReviewService(newReview)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse(err.Error()))
	}

	reviewResponse := responses.ReviewResponseObject{
		ID:        createdReview.ID,
		EventID:   createdReview.EventID,
		UserID:    createdReview.UserID,
		Review:    createdReview.Review,
		Rating:    createdReview.Rating,
		CreatedAt: createdReview.CreatedAt,
		UpdatedAt: createdReview.UpdatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(responses.NewDataResponse("Review created", reviewResponse))
}

func (h *httpHandler) UpdateReviewHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	review := new(ReviewPayload)
	if err := c.BodyParser(review); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse("Bad Request"))
	}

	if err := h.validator.ValidateStruct(review); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse(err.Error()))
	}

	newReview := &entity.Review{
		ID:      uuid.MustParse(id),
		EventID: review.EventID,
		UserID:  review.UserID,
		Review:  review.Review,
		Rating:  review.Rating,
	}

	updatedReview, err := h.svc.SaveReviewService(newReview)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse(err.Error()))
	}

	reviewResponse := responses.ReviewResponseObject{
		ID:        updatedReview.ID,
		EventID:   updatedReview.EventID,
		UserID:    updatedReview.UserID,
		Review:    updatedReview.Review,
		Rating:    updatedReview.Rating,
		CreatedAt: updatedReview.CreatedAt,
		UpdatedAt: updatedReview.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewDataResponse("Review updated", reviewResponse))
}

func (h *httpHandler) FindAllReviewHandler(c *fiber.Ctx) error {
	reviews, err := h.svc.FindAllReviewService()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse(err.Error()))
	}

	var reviewResponses []responses.ReviewResponseObject
	for _, review := range reviews {
		reviewResponses = append(reviewResponses, responses.ReviewResponseObject{
			ID:        review.ID,
			EventID:   review.EventID,
			UserID:    review.UserID,
			Review:    review.Review,
			Rating:    review.Rating,
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewDataResponse("Reviews found", reviewResponses))
}

func (h *httpHandler) FindReviewHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	review, err := h.svc.FindReviewService(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.NewErrorResponse("Review not found"))
	}

	reviewResponse := responses.ReviewResponseObject{
		ID:        review.ID,
		EventID:   review.EventID,
		UserID:    review.UserID,
		Review:    review.Review,
		Rating:    review.Rating,
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewDataResponse("Review found", reviewResponse))
}

type CustomReviewEventResponse struct {
	Event   responses.EventResponseObject    `json:"event"`
	Reviews []responses.ReviewResponseObject `json:"reviews"`
}

func (h *httpHandler) FindReviewByEventIDHandler(c *fiber.Ctx) error {
	eventID := c.Params("id")
	reviews, err := h.svc.FindReviewByEventIDService(eventID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse(err.Error()))
	}

	res := CustomReviewEventResponse{
		Event: responses.EventResponseObject{
			ID:            reviews[0].Event.ID,
			Name:          reviews[0].Event.Name,
			Location:      reviews[0].Event.Location,
			StartDate:     reviews[0].Event.StartDate,
			EndDate:       reviews[0].Event.EndDate,
			Price:         reviews[0].Event.Price,
			TotalSeat:     reviews[0].Event.TotalSeat,
			AvailableSeat: reviews[0].Event.AvailableSeat,
			Category:      reviews[0].Event.Category,
		},
	}

	for _, review := range reviews {
		res.Reviews = append(res.Reviews, responses.ReviewResponseObject{
			ID:        review.ID,
			EventID:   review.EventID,
			UserID:    review.UserID,
			Review:    review.Review,
			Rating:    review.Rating,
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewDataResponse("Reviews found", res))
}

type CustomReviewUserResponse struct {
	User    responses.UserResponseObject     `json:"user"`
	Reviews []responses.ReviewResponseObject `json:"reviews"`
}

func (h *httpHandler) FindReviewByUserIDHandler(c *fiber.Ctx) error {
	userID := c.Params("id")
	reviews, err := h.svc.FindReviewByUserIDService(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse(err.Error()))
	}

	res := CustomReviewUserResponse{
		User: responses.UserResponseObject{
			ID:    reviews[0].User.ID,
			Name:  reviews[0].User.Name,
			Email: reviews[0].User.Email,
			Role:  reviews[0].User.Role,
		},
	}

	for _, review := range reviews {
		res.Reviews = append(res.Reviews, responses.ReviewResponseObject{
			ID:        review.ID,
			EventID:   review.EventID,
			UserID:    review.UserID,
			Review:    review.Review,
			Rating:    review.Rating,
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewDataResponse("Reviews found", res))
}

func (h *httpHandler) DeleteReviewHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.svc.DeleteReviewService(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewDataResponse("Review deleted", nil))
}
