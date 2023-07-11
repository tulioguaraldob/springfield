package api

import (
	"log"

	"github.com/TulioGuaraldoB/springfield/internal/config/env"
	"github.com/gofiber/fiber/v2"
)

type ServerConfig struct {
	Server *fiber.App
	Port   string
}

func New() *ServerConfig {
	return &ServerConfig{
		Server: GetRoutes(),
		Port:   env.Port,
	}
}

func (s *ServerConfig) Run() {
	log.Fatal(s.Server.Listen(":" + s.Port))
}
