package server

import (
	"log"

	"github.com/TulioGuaraldoB/q2-payer-challenge/config/env"
	"github.com/TulioGuaraldoB/q2-payer-challenge/infra/routes"
	"github.com/gofiber/fiber/v2"
	_ "github.com/TulioGuaraldoB/q2-payer-challenge/docs"
)

type ServerConfig struct {
	Port   string
	Server *fiber.App
}

func NewServer() *ServerConfig {
	return &ServerConfig{
		Port:   env.Env.Port,
		Server: routes.SetupRoutes(),
	}
}

func (s *ServerConfig) RunServer() {
	log.Fatal(s.Server.Listen(":" + s.Port))
}
