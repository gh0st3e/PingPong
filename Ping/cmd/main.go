package main

import (
	"Ping/internal/handler"
	"Ping/internal/rabbit"
	"Ping/internal/service"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("init logger error")
	}

	rabbit, err := rabbit.InitRabbit()
	pingServive := service.NewService(logger, rabbit)

	server := mux.NewRouter()

	pingHandler := handler.NewHandler(pingServive, server)

	go func() {
		err := pingServive.StartGame()
		if err != nil {
			logger.Fatal("Error while ping-pong", zap.Error(err))
		}
	}()

	pingHandler.Init()

	err = http.ListenAndServe(":8081", server)
	if err != nil {
		logger.Error("init server error: ", zap.Error(err))
	}
}
