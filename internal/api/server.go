package api

import (
	"context"
	"errors"
	"event-booking/internal/account"
	"event-booking/internal/api/validator"
	"event-booking/internal/auth"
	"event-booking/internal/booking"
	"event-booking/internal/config"
	"event-booking/internal/email"
	"event-booking/internal/event"
	"event-booking/internal/export"
	"event-booking/internal/health"
	"event-booking/internal/postgres"
	"event-booking/internal/rabbitmq"
	"event-booking/internal/review"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog/log"
)

func NewServer() *Server {
	cfg := config.Load()

	// Database Connection
	db := postgres.NewGORM(cfg.Database)
	postgres.Migrate(db)

	// Email Service
	emailService := email.NewEmailService(&cfg.Smtp)

	// RabbitMQ
	rabbitCon := rabbitmq.InitRabbitMQ(&cfg.RabbitMQ)

	// middleware
	jwtService := auth.NewJwtService(cfg.App.JwtSecretKey)
	middleware := auth.NewMiddleware(jwtService)

	// validator
	validatorService := validator.NewValidator()

	// Health
	healthRepo := health.NewRepository(db)
	healthSvc := health.NewService(healthRepo)
	healthHandler := health.NewHttpHandler(healthSvc)

	// Account
	accountRepo := account.NewRepository(db)
	accountSvc := account.NewService(accountRepo, emailService)
	accountHandler := account.NewHttpHandler(accountSvc, jwtService, validatorService)

	// Event
	eventRepo := event.NewRepository(db)
	eventSvc := event.NewService(eventRepo)
	eventHandler := event.NewHttpHandler(eventSvc, validatorService)

	// Booking
	bookingRepo := booking.NewRepository(db)
	bookingSvc := booking.NewService(bookingRepo, eventRepo)
	bookingHandler := booking.NewHttpHandler(bookingSvc, validatorService)

	// Review
	reviewRepo := review.NewRepository(db)
	reviewSvc := review.NewService(reviewRepo)
	reviewHandler := review.NewHttpHandler(reviewSvc, validatorService)

	// Export
	exportSvc := export.NewService(rabbitCon, eventRepo, bookingRepo)
	exportHandler := export.NewHttpHandler(exportSvc)

	app := fiber.New()

	app.Use(
		logger.New(),
	)

	// Root
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Event Booking API",
		})
	})

	// Health routes
	app.Get("/health", healthHandler.HealthCheck)

	// Account routes
	app.Post("/api/signup", accountHandler.SignUpUserHandler)
	app.Post("/api/signin", accountHandler.SignInUserHandler)
	app.Post("/api/logout", accountHandler.SignOutUserHandler)
	app.Post("/api/refresh", accountHandler.RefreshTokenHandler)
	app.Put("/api/account", middleware.AdminRequired, accountHandler.UpdateUserHandler)
	app.Get("/api/account/:id", middleware.AdminRequired, accountHandler.GetUserByIDHandler)
	app.Post("/api/account/send-verification", accountHandler.RequestVerificationCodeHandler)
	app.Post("/api/account/verify", accountHandler.RequestVerificationCodeHandler)

	// Event Admin routes
	app.Post("/api/admin/event", middleware.AdminRequired, eventHandler.CreateEventHandler)
	app.Get("/api/admin/event", middleware.AdminRequired, eventHandler.FindAllEventHandler)
	app.Get("/api/admin/event/:id", middleware.AdminRequired, eventHandler.FindEventHandler)
	app.Put("/api/admin/event/:id", middleware.AdminRequired, eventHandler.SaveEventHandler)
	app.Delete("/api/admin/event/:id", middleware.AdminRequired, eventHandler.DeleteEventHandler)
	app.Get("/api/admin/event/:id/bookings", middleware.AdminRequired, eventHandler.GetEventBookingsHandler)

	// Event routes
	app.Get("/api/event", middleware.AuthRequired, eventHandler.FindAllEventHandler)
	app.Get("/api/event/:id", middleware.AuthRequired, eventHandler.FindEventHandler)
	app.Get("/api/event/filter", middleware.AuthRequired, eventHandler.FilterByCriteria)

	// Booking routes
	app.Post("/api/booking", middleware.AuthRequired, bookingHandler.BookEventHandler)
	app.Get("/api/booking", middleware.AuthRequired, bookingHandler.GetBookedEventsHandler)
	app.Get("/api/booking/:id", middleware.AuthRequired, bookingHandler.GetBookedEventByIDHandler)
	app.Put("/api/booking/:id", middleware.AuthRequired, bookingHandler.UpdateBookedEventHandler)
	app.Delete("/api/booking/:id", middleware.AuthRequired, bookingHandler.CancelBookedEventHandler)

	// Review routes
	app.Post("/api/review", middleware.AuthRequired, reviewHandler.CreateReviewHandler)
	app.Get("/api/review", middleware.AuthRequired, reviewHandler.FindAllReviewHandler)
	app.Get("/api/review/:id", middleware.AuthRequired, reviewHandler.FindReviewHandler)
	app.Get("/api/review/event/:id", middleware.AuthRequired, reviewHandler.FindReviewByEventIDHandler)
	app.Get("/api/review/user/:id", middleware.AuthRequired, reviewHandler.FindReviewByUserIDHandler)
	app.Put("/api/review/:id", middleware.AuthRequired, reviewHandler.UpdateReviewHandler)
	app.Delete("/api/review/:id", middleware.AuthRequired, reviewHandler.DeleteReviewHandler)

	// Export routes
	app.Get("/api/export/event", middleware.AdminRequired, exportHandler.ExportAllEventHandler)
	app.Get("/api/export/booking/:id", middleware.AdminRequired, exportHandler.ExportBookingHandler)

	return &Server{fiber: app}
}

type Server struct {
	fiber *fiber.App
}

// Run method of the Server struct runs the Fiber server on the specified port.
func (s *Server) Run(port int) {
	addr := fmt.Sprintf(":%d", port)

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		log.Info().Msg("server is shutting down...")

		_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := s.fiber.Shutdown(); err != nil {
			log.Fatal().Err(err).Msg("could not gracefully shutdown the server")
		}
		close(done)
	}()

	log.Info().Msgf("server serving on port %d", port)
	if err := s.fiber.Listen(addr); err != nil && !errors.Is(err, &fiber.Error{
		Code:    500,
		Message: "Server closed",
	}) {
		log.Fatal().Err(err).Msgf("could not listen on %s", addr)
	}

	<-done
	log.Info().Msg("server stopped")
}
