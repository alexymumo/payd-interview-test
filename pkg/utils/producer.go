package utils

import (
	"fmt"

	"github.com/streadway/amqp"
)

func PublishMessage(message []byte) error {
	rabbitmqUrl := "amqp://guest:guest@localhost:5672/"
	queueName := "test"

	conn, err := amqp.Dial(rabbitmqUrl)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, false, false, false, false, nil,
	)
	if err != nil {
		return fmt.Errorf("error occuured", err)
	}
	err = ch.Publish(
		"",
		q.Name, false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		return fmt.Errorf("error occurred", err)
	}
	return nil
}
