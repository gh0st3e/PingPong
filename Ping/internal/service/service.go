package service

import (
	"Ping/internal/rabbit"
	"fmt"
	"go.uber.org/zap"
	"log"
	"time"
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
		for {
			time.Sleep(time.Second * 3)

			err := s.rabbit.Send(s.Ind)
			if err != nil {
				return
			}
			log.Printf("sent: '%s %d'\n", "Ping", s.Ind)
			s.Ind++

			msg := <-msgs
			log.Printf("recieved: '%s'\n", msg.Body)
		}

	}()

	fmt.Println("Waiting for message")
	<-forever

	return fmt.Errorf("unexpected error")
}
