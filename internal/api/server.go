package api

import (
	"context"
	"errors"
	"event-booking/internal/account"
	"event-booking/internal/auth"
	"event-booking/internal/booking"
	"event-booking/internal/config"
	"event-booking/internal/event"
	"event-booking/internal/postgres"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func NewServer() *Server {
	cfg := config.Load()

	db := postgres.NewGORM(cfg.Database)
	postgres.Migrate(db)

	jwtService := auth.NewJwtService(cfg.App.JwtSecretKey)

	// Account
	accountRepo := account.NewRepository(db)
	accountSvc := account.NewService(accountRepo)
	accountHandler := account.NewHttpHandler(accountSvc, jwtService)

	// Event
	eventRepo := event.NewRepository(db)
	eventSvc := event.NewService(eventRepo)
	eventHandler := event.NewHttpHandler(eventSvc)

	// Booking
	bookingRepo := booking.NewRepository(db)
	bookingSvc := booking.NewService(bookingRepo)
	bookingHandler := booking.NewHttpHandler(bookingSvc)

	app := fiber.New()

	// Root
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Event Booking API",
		})
	})

	// Account routes
	app.Post("/api/signup", accountHandler.SignUpUserHandler)
	app.Post("/api/signin", accountHandler.SignInUserHandler)

	// Event routes
	// there are problem on the middleware, can't get proper response
	app.Post("/api/event", jwtService.AdminOnly, eventHandler.CreateEventHandler)
	app.Get("/api/event", jwtService.AdminOnly, eventHandler.FindAllEventHandler)
	app.Get("/api/event/:id", jwtService.AdminOnly, eventHandler.FindEventHandler)
	app.Put("/api/event/:id", jwtService.AdminOnly, eventHandler.SaveEventHandler)
	app.Delete("/api/event/:id", jwtService.AdminOnly, eventHandler.DeleteEventHandler)

	// Booking routes
	app.Post("/api/booking", jwtService.AuthRequired, bookingHandler.BookEventHandler)
	app.Get("/api/booking", jwtService.AuthRequired, bookingHandler.GetBookedEventsHandler)
	app.Get("/api/booking/:id", jwtService.AuthRequired, bookingHandler.GetBookedEventByIDHandler)
	app.Delete("/api/booking/:id", jwtService.AuthRequired, bookingHandler.CancelBookedEventHandler)
	app.Put("/api/booking/:id", jwtService.AuthRequired, bookingHandler.UpdateBookedEventHandler)

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
