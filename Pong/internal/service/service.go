package service

import (
	"Pong/internal/rabbit"
	"fmt"
	"go.uber.org/zap"
	"log"
)

type Service struct {
	logger *zap.Logger
	rabbit *rabbit.Rabbit
	Ind    int
}

func NewService(logger *zap.Logger, rabbit *rabbit.Rabbit) *Service {
	return &Service{
		logger: logger,
		rabbit: rabbit,
		Ind:    0,
	}
}

func (s *Service) StartGame() error {

	msgs, err := s.rabbit.Get()
	if err != nil {
		return fmt.Errorf("get msgs from pong error: %w", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("recieved: '%s'\n", d.Body)

			err := s.rabbit.Send(s.Ind)
			if err != nil {
				return
			}
			log.Printf("sent: '%s %d'\n", "Pong", s.Ind)
			s.Ind++

		}

	}()

	fmt.Println("Waiting for message")
	<-forever

	return fmt.Errorf("unexpected error")
}
