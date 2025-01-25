package echo

import (
    "context"
    "net/http"

    "github.com/gorilla/websocket"
)

type Config struct {
    serverAddr string
}

type Service struct {
    config   Config
    upgrader websocket.Upgrader
}

type Option func(*Service)

func NewEchoService(options ...Option) *Service {
    service := &Service{
        upgrader: websocket.Upgrader{},
    }

    for _, opt := range options {
        opt(service)
    }

    return service
}

func WithAddress(serverAddr string) Option {
    return func(service *Service) {
        service.config.serverAddr = serverAddr
    }
}

func (s *Service) Start(_ context.Context) error {
    handler := NewHandler(s)

    http.HandleFunc("/echo", handler.Echo)

    return http.ListenAndServe(s.config.serverAddr, nil)
}
