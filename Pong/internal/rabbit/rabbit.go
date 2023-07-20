package rabbit

import (
	"fmt"
	amqp "github.com/streadway/amqp"
	"log"
)

const (
	ping = "Ping"
	pong = "Pong"
)

type Rabbit struct {
	channel *amqp.Channel
}

func InitRabbit() (*Rabbit, error) {
	var rabbit Rabbit

	conn, err := amqp.Dial("amqp://guest:guest@rabbit:5672/")
	if err != nil {
		return nil, fmt.Errorf("dial rabbit error: %w", err)
	}

	log.Println("successfully connected")

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("creating channel error: %w", err)
	}

	rabbit.channel = ch

	log.Println("creating channel - success")

	ping, err := ch.QueueDeclare(ping,
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		return nil, fmt.Errorf("decalre queue error: %w", err)
	}

	pong, err := ch.QueueDeclare(pong,
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		return nil, fmt.Errorf("decalre queue error: %w", err)
	}

	log.Println(fmt.Sprintf("%s %d %d", ping.Name, ping.Messages, ping.Consumers))
	log.Println(fmt.Sprintf("%s %d %d", pong.Name, pong.Messages, pong.Consumers))

	return &rabbit, nil
}

func (r *Rabbit) Get() (<-chan amqp.Delivery, error) {
	msgs, err := r.channel.Consume(ping,
		"",
		true,
		false,
		false,
		false,
		nil)

	return msgs, err
}

func (r *Rabbit) Send(ind int) error {
	err := r.channel.Publish("",
		pong,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("Pong %d", ind)),
		})

	return err
}
