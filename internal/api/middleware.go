package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Middleware func(*fiber.Ctx) error

func chainMiddleware(app *fiber.App, middlewares ...Middleware) {
	for _, middleware := range middlewares {
		app.Use(middleware)
	}
}

func logSeverity(statusCode int) zerolog.Level {
	switch {
	case statusCode >= 500:
		return zerolog.ErrorLevel
	case statusCode >= 400:
		return zerolog.ErrorLevel
	case statusCode >= 300:
		return zerolog.WarnLevel
	case statusCode >= 200:
		return zerolog.InfoLevel
	default:
		return zerolog.DebugLevel
	}
}

type logFields struct {
	RemoteIP   string
	Host       string
	UserAgent  string
	Method     string
	Path       string
	Body       string
	StatusCode int
	Latency    float64
}

func (l *logFields) MarshalZerologObject(e *zerolog.Event) {
	e.
		Str("remote_ip", l.RemoteIP).
		Str("host", l.Host).
		Str("user_agent", l.UserAgent).
		Str("method", l.Method).
		Str("path", l.Path).
		Str("body", l.Body).
		Int("status_code", l.StatusCode).
		Float64("latency", l.Latency)
}

func loggerHandler(filter func(c *fiber.Ctx) bool) Middleware {
	return func(c *fiber.Ctx) error {
		// Check filter
		if filter != nil && filter(c) {
			return c.Next()
		}

		// Start timer
		start := time.Now()

		// Read request body
		var buf []byte
		if c.Body() != nil {
			buf = c.Body()
		}

		// Process request
		err := c.Next()

		dur := float64(time.Since(start).Nanoseconds()/1e4) / 100.0

		// Create log fields
		fields := &logFields{
			RemoteIP:   c.IP(),
			Host:       c.Hostname(),
			UserAgent:  c.Get("User-Agent"),
			Method:     c.Method(),
			Path:       c.Path(),
			Body:       formatReqBody(c, buf),
			StatusCode: c.Response().StatusCode(),
			Latency:    dur,
		}

		sev := logSeverity(c.Response().StatusCode())
		logEntry := log.Ctx(c.Context()).WithLevel(sev).EmbedObject(fields)
		logEntry.Msg("http request")

		return err
	}
}

func formatReqBody(c *fiber.Ctx, data []byte) string {
	var js map[string]interface{}
	if json.Unmarshal(data, &js) != nil {
		return string(data)
	}

	result := new(bytes.Buffer)
	if err := json.Compact(result, data); err != nil {
		log.Ctx(c.Context()).Error().Msgf("error compacting body request json: %s", err.Error())
		return ""
	}

	return result.String()
}

func realIPHandler() Middleware {
	return func(c *fiber.Ctx) error {
		if rip := realIP(c); rip != "" {
			c.Request().Header.Set(fiber.HeaderXForwardedFor, rip+", "+c.IP())
		}

		return c.Next()
	}
}

func realIP(c *fiber.Ctx) string {
	trueClientIP := "True-Client-IP"
	xForwardedFor := "X-Forwarded-For"
	xRealIP := "X-Real-IP"

	var ip string
	if tcip := c.Get(trueClientIP); tcip != "" {
		ip = tcip
	} else if xrip := c.Get(xRealIP); xrip != "" {
		ip = xrip
	} else if xff := c.Get(xForwardedFor); xff != "" {
		i := strings.Index(xff, ",")
		if i == -1 {
			i = len(xff)
		}
		ip = xff[:i]
	}
	if ip == "" {
		return ""
	}
	return ip
}

func recoverHandler() Middleware {
	return func(c *fiber.Ctx) error {
		defer func() {
			if rvr := recover(); rvr != nil {
				if rvr == fiber.ErrInternalServerError {
					// we don't recover fiber.ErrAbortHandler so the response
					// to the client is aborted, this should not be logged
					panic(rvr)
				}

				err, ok := rvr.(error)
				if !ok {
					err = fmt.Errorf("%v", rvr)
				}

				log.Ctx(c.Context()).
					Error().
					Err(err).
					Bytes("stack", debug.Stack()).
					Msg("panic recover")

				c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
		}()

		return c.Next()
	}
}

func requestIDHandler() Middleware {
	return func(c *fiber.Ctx) error {
		requestIDHeader := "X-Request-Id"
		if c.Get(requestIDHeader) == "" {
			c.Set(requestIDHeader, uuid.NewString())
		}

		ctx := log.With().
			Str("request_id", c.Get(requestIDHeader)).
			Logger().
			WithContext(c.Context())

		c.SetUserContext(ctx)
		return c.Next()
	}
}

func corsHandler() Middleware {
	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")

		if c.Method() == fiber.MethodOptions {
			c.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Qiscus-App-Id, Qiscus-Secret-Key")
			c.Set("Access-Control-Allow-Methods", "GET, PUT, POST, HEAD, DELETE, OPTIONS")
			c.Set("Access-Control-Allow-Credentials", "true")
			c.Set("Access-Control-Max-Age", "3600")
			return c.SendStatus(fiber.StatusNoContent)
		}

		return c.Next()
	}
}
