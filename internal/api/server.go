package api

import (
	"context"
	"errors"
	"event-booking/internal/account"
	"event-booking/internal/auth"
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

	db := postgres.NewGORM(&cfg.Database)
	postgres.Migrate(db)

	jwt := auth.NewJwtService(cfg.App.SecretKey)

	// Account
	accountRepo := account.NewRepository(db)
	accountSvc := account.NewService(accountRepo)
	accountHandler := account.NewHttpHandler(accountSvc, jwt)

	// Event
	eventRepo := event.NewRepository(db)
	eventSvc := event.NewService(eventRepo)
	eventHandler := event.NewHttpHandler(eventSvc)

	app := fiber.New()

	app.Post("/api/signup", accountHandler.SignUpUserHandler)
	app.Post("/api/signin", accountHandler.SignInUserHandler)

	app.Post("/api/event", jwt.AuthRequired, eventHandler.CreateEventHandler)
	app.Get("/api/event", jwt.AuthRequired, eventHandler.FindAllEventHandler)
	app.Get("/api/event/:id", jwt.AuthRequired, eventHandler.FindEventHandler)
	app.Put("/api/event/:id", jwt.AuthRequired, eventHandler.SaveEventHandler)
	app.Delete("/api/event/:id", jwt.AuthRequired, eventHandler.DeleteEventHandler)

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
