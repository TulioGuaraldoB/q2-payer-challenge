package main

import (
	"github.com/TulioGuaraldoB/q2-payer-challenge/config/env"
	"github.com/TulioGuaraldoB/q2-payer-challenge/infra/server"
)

func main() {
	env.GetEnvironmentVariables()

	server := server.NewServer()
	server.RunServer()
}
