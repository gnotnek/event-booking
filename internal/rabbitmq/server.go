package rabbitmq

import (
	"event-booking/internal/config"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

func InitRabbitMQ(c *config.RabbitMQ) *amqp091.Connection {
	conn, err := amqp091.Dial(c.Url)
	if err != nil {
		log.Fatal().Msgf("can't connect to RabbitMQ server: %s", err.Error())
		return nil
	}

	return conn
}
